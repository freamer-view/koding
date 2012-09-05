%% MIT License
%% ===========

%% Copyright (c) 2012 Son Tran <esente@gmail.com>

%% Permission is hereby granted, free of charge, to any person obtaining a
%% copy of this software and associated documentation files (the "Software"),
%% to deal in the Software without restriction, including without limitation
%% the rights to use, copy, modify, merge, publish, distribute, sublicense,
%% and/or sell copies of the Software, and to permit persons to whom the
%% Software is furnished to do so, subject to the following conditions:

%% The above copyright notice and this permission notice shall be included in
%% all copies or substantial portions of the Software.

%% THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
%% IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
%% FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
%% THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
%% LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
%% FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
%% DEALINGS IN THE SOFTWARE.
-module(broker_app).

-behaviour(application).

-behaviour(cowboy_http_handler). %% To handle the default route

%% Application callbacks
-export([start/0, start/2, stop/1]).

%% Cowboy callbacks
-export([init/3, handle/2, terminate/2]).
-record (client, {id, socket_id, subscriptions=dict:new()}).
-include_lib("amqp_client/include/amqp_client.hrl").

%% ===================================================================
%% Application callbacks
%% ===================================================================

start() ->
    application:start(crypto),
    application:start(public_key),
    application:start(ssl),
    application:start(sockjs),
    application:start(cowboy),

    application:start(broker).

start(_StartType, _StartArgs) ->
    NumberOfAcceptors = 100,
    Port = 8008,

    error_logger:tty(get_env(verbose, true)),

    % This will start the Broker gen_server and the subscription_sup
    broker:start_link(),
    
    SockjsState = sockjs_handler:init_state(
                    <<"/subscribe">>, fun handle_client/3, {}, []),
    VhostRoutes = [
        {
            [<<"subscribe">>, '...'], 
            sockjs_cowboy_handler, 
            SockjsState
        },
        {
            [<<"static">>, '...'], 
            cowboy_http_static,
            [
                {directory, {priv_dir, broker, [<<"www">>]}},
                {mimetypes, {fun mimetypes:path_to_mimes/2, default}}
            ]
        },
        {'_', ?MODULE, []} % The rest is handled within this module.
    ],
    Routes = [{'_',  VhostRoutes}], % any vhost

    debug_log(" [*] Running at http://localhost:~p~n", [Port]),

    cowboy:start_listener(http, 
        NumberOfAcceptors,
        cowboy_tcp_transport, [{port, Port}],
        cowboy_http_protocol, [{dispatch, Routes}]
    ),

    broker_sup:start_link().

stop(_State) ->
    ok.

%% ===================================================================
%% Cowboy callbacks
%% ===================================================================

init({_Any, http}, Req, []) ->
    {ok, Req, []}.

handle(Req, State) ->
    {Path, Req1} = cowboy_http_req:path(Req),
    {ok, Req2} = case Path of
        [<<"broker.js">>] ->
            {ok, Data} = file:read_file("./apps/broker/priv/www/js/broker.js"),
            cowboy_http_req:reply(200, [{<<"Content-Type">>, "application/javascript"}],
                               Data, Req1);

        [<<"auth">>] ->
            {Channel, Req3} = cowboy_http_req:qs_val(<<"channel">>, Req1),
            %PrivateChannel = uuid:to_string(uuid:uuid4()),
            PrivateChannel = <<"secret-", Channel/binary, ".private">>,
            cowboy_http_req:reply(200,
                [{<<"Content-Encoding">>, <<"utf-8">>}], PrivateChannel, Req3);

        [] ->
            {ok, Data} = file:read_file("./apps/broker/priv/www/index.html"),
            cowboy_http_req:reply(200, [{<<"Content-Type">>, "text/html"}],
                               Data, Req1);
        _ ->
            cowboy_http_req:reply(404, [],
                               <<"404 - Nothing here\n">>, Req1)
        end,
    {ok, Req2, State}.

terminate(_Req, _State) ->
    ok.

%% ===================================================================
%% SockJS_MQ Handlers
%% ===================================================================

handle_client(Conn, init, _State) ->
    SocketId = list_to_binary(uuid:to_string(uuid:uuid4())),
    Event = <<"connected">>,
    EventProp = {<<"event">>, Event},
    Payload = {<<"socket_id">>, SocketId},
    Conn:send(jsx:encode([EventProp, Payload])),

    SystemExchange = get_env(system_exchange, <<"private-broker">>),
    {ok, Subscription} = broker:subscribe(Conn, SystemExchange),
    broker:trigger(Subscription, Event, jsx:encode([Payload]), []),
    broker:unsubscribe(Subscription),

    {ok, #client{socket_id=SocketId}};

handle_client(Conn, {recv, Data}, 
                    State=#client{subscriptions=Subscriptions}) ->
    [Event, Exchange, _Payload, _Meta] = Decoded = decode(Data),
    Check = {Event, dict:is_key(Exchange, Subscriptions)},
    NewSubs = handle_event(Conn, Check, Decoded, Subscriptions),
    {ok, State#client{subscriptions=NewSubs}};

handle_client(Conn, closed, #client{socket_id=SocketId,
                                    subscriptions=Subscriptions}) ->
    SystemExchange = get_env(system_exchange, <<"private-broker">>),
    {ok, SystemSubscription} = broker:subscribe(Conn, SystemExchange),
    Event = <<"disconnected">>,
    Sid = {<<"socket_id">>, SocketId},
    Exchanges = {<<"exchanges">>, dict:fetch_keys(Subscriptions)},
    Payload = jsx:encode([Sid, Exchanges]),
    broker:trigger(SystemSubscription, Event, Payload, []),
    broker:unsubscribe(SystemSubscription),

    case dict:size(Subscriptions) of 
        0 -> ok;
        _ ->
            List = dict:to_list(Subscriptions),
            [broker:unsubscribe(Subscription) 
                || {_Exchange, Subscription} <- List]
    end,
    {ok, #client{}};

handle_client(_Conn, Other, _) ->
    io:format("Other data ~p~n", [Other]).

handle_event(Conn, {<<"client-subscribe">>, false}, Data, Subs) ->
    [_Event, Exchange, _Payload, _Meta] = Data,
    case broker:subscribe(Conn, Exchange) of
        {error, _Error} -> Subs;
        {ok, Subscription} ->
            dict:store(Exchange, Subscription, Subs)
    end;

handle_event(Conn, {<<"client-presence">>, false}, Data, Subs) ->
    [_Event, Where, Who, _Meta] = Data,
    % This only acts as a key so that it can be used to
    % remove the subscription later from the dictionary.
    Exchange = <<Who/bitstring,Where/bitstring,"-presence">>,
    case dict:find(Exchange, Subs) of
        {ok, Subscription} ->
            broker:unsubscribe(Subscription),
            dict:erase(Exchange, Subs);
        error ->
            case broker:presence(Conn, Where, Who) of
                {error, _Error} -> Subs;
                {ok, Subscription} ->            
                    dict:store(Exchange, Subscription, Subs)
            end
    end;

handle_event(_Conn, {<<"client-bind-event">>, true}, Data, Subs) ->
    [_Event, Exchange, Payload, _Meta] = Data,
    Subscription = dict:fetch(Exchange, Subs),
    broker:bind(Subscription, Payload),
    Subs;

handle_event(_Conn, {<<"client-unbind-event">>, true}, Data, Subs) ->
    [_Event, Exchange, Payload, _Meta] = Data,
    Subscription = dict:fetch(Exchange, Subs),
    broker:unbind(Subscription, Payload),
    Subs;

handle_event(_Conn, {<<"client-unsubscribe">>, true}, Data, Subs) ->
    [_Event, Exchange, _Payload, _Meta] = Data,
    Subscription = dict:fetch(Exchange, Subs),
    broker:unsubscribe(Subscription),
    dict:erease(Exchange, Subs);

handle_event(_Conn, {<<"client-",_EventName/binary>>, true}, Data, Subs) ->
    [Event, Exchange, Payload, Meta] = Data,
    RegExp = "^secret-",
    case re:run(Exchange, RegExp) of
        nomatch -> Subs;
        {match, _} ->
            Subscription = dict:fetch(Exchange, Subs),
            broker:trigger(Subscription, Event, Payload, Meta),
            Subs
    end;

handle_event(_Conn, _Else, _Data, Subs) ->
    Subs.

%%--------------------------------------------------------------------
%%% Internal functions
%%--------------------------------------------------------------------

debug_log(Text, Args) ->
    case application:get_env(broker, verbose) of
        {ok, Val} when Val ->
            io:format(Text, Args);
        _ -> true
    end.

%%--------------------------------------------------------------------
%% Function: decode(Data) -> [Event, Exchange, Payload, Meta]
%% Types:
%%  Data = binary()
%%  Event = binary()
%%  Exchange = binary()
%%  Payload = binary()
%%  Meta = binary()
%% Description:  Decode a binary data from the websocket connection
%% into a list of data that the handler expects
%%--------------------------------------------------------------------
decode(Data) ->
    [{<<"event">>, Event}, 
        {<<"channel">>, Exchange} | Rest] = jsx:decode(Data),

    Payload = bin_key_find(<<"payload">>, Rest, false),
    Meta = bin_key_find(<<"meta">>, Rest, true),
    [Event, Exchange, Payload, Meta].

%%--------------------------------------------------------------------
%% Function: bin_key_find(BinKey, List, ReturnsEmptyList) -> Val || false
%% Description:  A helper to find a binary key in a binary proplist.
%% The third boolean argument specifies to return an empty list or empty
%% binary when the BinKey not found.
%%--------------------------------------------------------------------
bin_key_find(BinKey, List, ReturnEmptyList) ->
    case lists:keyfind(BinKey, 1, List) of
        {_, Val} when is_integer(Val) -> list_to_binary(integer_to_list(Val));
        {_, Val} -> Val;
        false when ReturnEmptyList -> [];
        false -> <<>>
    end.

get_env(Param, DefaultValue) ->
    case application:get_env(broker, Param) of
        {ok, Val} -> Val;
        undefined -> DefaultValue
    end.
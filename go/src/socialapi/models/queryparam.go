package models

import (
	"net/url"
	"strconv"
	"time"

	"github.com/kennygrant/sanitize"
)

type Query struct {
	Skip      int
	Limit     int
	To        time.Time
	From      time.Time
	GroupName string
	Type      string
	Privacy   string
	AccountId int64
}

func NewQuery() *Query {
	return &Query{}
}

func (q *Query) MapURL(u *url.URL) *Query {
	urlQuery := u.Query()

	q.Skip, _ = strconv.Atoi(urlQuery.Get("skip"))
	q.Limit, _ = strconv.Atoi(urlQuery.Get("limit"))

	q.GroupName = u.Query().Get("groupName")
	if q.GroupName != "" {
		q.GroupName = sanitize.Name(q.GroupName)
	}

	q.Type = u.Query().Get("type")
	if q.Type != "" {
		q.Type = sanitize.Name(q.Type)
	}

	q.Privacy = u.Query().Get("privacy")
	if q.Privacy != "" {
		q.Privacy = sanitize.Name(q.Privacy)
	}

	q.AccountId, _ = strconv.ParseInt(u.Query().Get("accountId"), 10, 64)

	q.To, _ = time.Parse(time.RFC3339, urlQuery.Get("to"))
	q.From, _ = time.Parse(time.RFC3339, urlQuery.Get("from"))

	return q
}

func (q *Query) SetDefaults() *Query {
	if q.Skip == 0 {
		// no need to do something
	}

	if q.Limit == 0 || q.Limit > 25 {
		q.Limit = 25
	}

	if q.From.IsZero() {
		q.From = time.Now().UTC()
	}

	if q.GroupName == "" {
		q.GroupName = "koding"
	}

	return q
}

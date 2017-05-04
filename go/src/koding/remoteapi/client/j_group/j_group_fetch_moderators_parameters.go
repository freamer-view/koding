package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewJGroupFetchModeratorsParams creates a new JGroupFetchModeratorsParams object
// with the default values initialized.
func NewJGroupFetchModeratorsParams() *JGroupFetchModeratorsParams {
	var ()
	return &JGroupFetchModeratorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupFetchModeratorsParamsWithTimeout creates a new JGroupFetchModeratorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupFetchModeratorsParamsWithTimeout(timeout time.Duration) *JGroupFetchModeratorsParams {
	var ()
	return &JGroupFetchModeratorsParams{

		timeout: timeout,
	}
}

// NewJGroupFetchModeratorsParamsWithContext creates a new JGroupFetchModeratorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupFetchModeratorsParamsWithContext(ctx context.Context) *JGroupFetchModeratorsParams {
	var ()
	return &JGroupFetchModeratorsParams{

		Context: ctx,
	}
}

/*JGroupFetchModeratorsParams contains all the parameters to send to the API endpoint
for the j group fetch moderators operation typically these are written to a http.Request
*/
type JGroupFetchModeratorsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j group fetch moderators params
func (o *JGroupFetchModeratorsParams) WithTimeout(timeout time.Duration) *JGroupFetchModeratorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group fetch moderators params
func (o *JGroupFetchModeratorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group fetch moderators params
func (o *JGroupFetchModeratorsParams) WithContext(ctx context.Context) *JGroupFetchModeratorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group fetch moderators params
func (o *JGroupFetchModeratorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group fetch moderators params
func (o *JGroupFetchModeratorsParams) WithBody(body models.DefaultSelector) *JGroupFetchModeratorsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group fetch moderators params
func (o *JGroupFetchModeratorsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group fetch moderators params
func (o *JGroupFetchModeratorsParams) WithID(id string) *JGroupFetchModeratorsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group fetch moderators params
func (o *JGroupFetchModeratorsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupFetchModeratorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

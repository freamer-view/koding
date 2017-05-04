package j_proxy_filter

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

// NewJProxyFilterCreateParams creates a new JProxyFilterCreateParams object
// with the default values initialized.
func NewJProxyFilterCreateParams() *JProxyFilterCreateParams {
	var ()
	return &JProxyFilterCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJProxyFilterCreateParamsWithTimeout creates a new JProxyFilterCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJProxyFilterCreateParamsWithTimeout(timeout time.Duration) *JProxyFilterCreateParams {
	var ()
	return &JProxyFilterCreateParams{

		timeout: timeout,
	}
}

// NewJProxyFilterCreateParamsWithContext creates a new JProxyFilterCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJProxyFilterCreateParamsWithContext(ctx context.Context) *JProxyFilterCreateParams {
	var ()
	return &JProxyFilterCreateParams{

		Context: ctx,
	}
}

/*JProxyFilterCreateParams contains all the parameters to send to the API endpoint
for the j proxy filter create operation typically these are written to a http.Request
*/
type JProxyFilterCreateParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j proxy filter create params
func (o *JProxyFilterCreateParams) WithTimeout(timeout time.Duration) *JProxyFilterCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j proxy filter create params
func (o *JProxyFilterCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j proxy filter create params
func (o *JProxyFilterCreateParams) WithContext(ctx context.Context) *JProxyFilterCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j proxy filter create params
func (o *JProxyFilterCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j proxy filter create params
func (o *JProxyFilterCreateParams) WithBody(body models.DefaultSelector) *JProxyFilterCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j proxy filter create params
func (o *JProxyFilterCreateParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JProxyFilterCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

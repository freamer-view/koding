package j_reward

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

// NewJRewardAddCustomRewardParams creates a new JRewardAddCustomRewardParams object
// with the default values initialized.
func NewJRewardAddCustomRewardParams() *JRewardAddCustomRewardParams {
	var ()
	return &JRewardAddCustomRewardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJRewardAddCustomRewardParamsWithTimeout creates a new JRewardAddCustomRewardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJRewardAddCustomRewardParamsWithTimeout(timeout time.Duration) *JRewardAddCustomRewardParams {
	var ()
	return &JRewardAddCustomRewardParams{

		timeout: timeout,
	}
}

// NewJRewardAddCustomRewardParamsWithContext creates a new JRewardAddCustomRewardParams object
// with the default values initialized, and the ability to set a context for a request
func NewJRewardAddCustomRewardParamsWithContext(ctx context.Context) *JRewardAddCustomRewardParams {
	var ()
	return &JRewardAddCustomRewardParams{

		Context: ctx,
	}
}

/*JRewardAddCustomRewardParams contains all the parameters to send to the API endpoint
for the j reward add custom reward operation typically these are written to a http.Request
*/
type JRewardAddCustomRewardParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j reward add custom reward params
func (o *JRewardAddCustomRewardParams) WithTimeout(timeout time.Duration) *JRewardAddCustomRewardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j reward add custom reward params
func (o *JRewardAddCustomRewardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j reward add custom reward params
func (o *JRewardAddCustomRewardParams) WithContext(ctx context.Context) *JRewardAddCustomRewardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j reward add custom reward params
func (o *JRewardAddCustomRewardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j reward add custom reward params
func (o *JRewardAddCustomRewardParams) WithBody(body models.DefaultSelector) *JRewardAddCustomRewardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j reward add custom reward params
func (o *JRewardAddCustomRewardParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JRewardAddCustomRewardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

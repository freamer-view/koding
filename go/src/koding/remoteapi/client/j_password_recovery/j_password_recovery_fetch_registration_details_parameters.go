package j_password_recovery

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

// NewJPasswordRecoveryFetchRegistrationDetailsParams creates a new JPasswordRecoveryFetchRegistrationDetailsParams object
// with the default values initialized.
func NewJPasswordRecoveryFetchRegistrationDetailsParams() *JPasswordRecoveryFetchRegistrationDetailsParams {
	var ()
	return &JPasswordRecoveryFetchRegistrationDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJPasswordRecoveryFetchRegistrationDetailsParamsWithTimeout creates a new JPasswordRecoveryFetchRegistrationDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJPasswordRecoveryFetchRegistrationDetailsParamsWithTimeout(timeout time.Duration) *JPasswordRecoveryFetchRegistrationDetailsParams {
	var ()
	return &JPasswordRecoveryFetchRegistrationDetailsParams{

		timeout: timeout,
	}
}

// NewJPasswordRecoveryFetchRegistrationDetailsParamsWithContext creates a new JPasswordRecoveryFetchRegistrationDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewJPasswordRecoveryFetchRegistrationDetailsParamsWithContext(ctx context.Context) *JPasswordRecoveryFetchRegistrationDetailsParams {
	var ()
	return &JPasswordRecoveryFetchRegistrationDetailsParams{

		Context: ctx,
	}
}

/*JPasswordRecoveryFetchRegistrationDetailsParams contains all the parameters to send to the API endpoint
for the j password recovery fetch registration details operation typically these are written to a http.Request
*/
type JPasswordRecoveryFetchRegistrationDetailsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j password recovery fetch registration details params
func (o *JPasswordRecoveryFetchRegistrationDetailsParams) WithTimeout(timeout time.Duration) *JPasswordRecoveryFetchRegistrationDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j password recovery fetch registration details params
func (o *JPasswordRecoveryFetchRegistrationDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j password recovery fetch registration details params
func (o *JPasswordRecoveryFetchRegistrationDetailsParams) WithContext(ctx context.Context) *JPasswordRecoveryFetchRegistrationDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j password recovery fetch registration details params
func (o *JPasswordRecoveryFetchRegistrationDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j password recovery fetch registration details params
func (o *JPasswordRecoveryFetchRegistrationDetailsParams) WithBody(body models.DefaultSelector) *JPasswordRecoveryFetchRegistrationDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j password recovery fetch registration details params
func (o *JPasswordRecoveryFetchRegistrationDetailsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JPasswordRecoveryFetchRegistrationDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

package data_dog

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

// NewDataDogSendMetricsParams creates a new DataDogSendMetricsParams object
// with the default values initialized.
func NewDataDogSendMetricsParams() *DataDogSendMetricsParams {
	var ()
	return &DataDogSendMetricsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataDogSendMetricsParamsWithTimeout creates a new DataDogSendMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataDogSendMetricsParamsWithTimeout(timeout time.Duration) *DataDogSendMetricsParams {
	var ()
	return &DataDogSendMetricsParams{

		timeout: timeout,
	}
}

// NewDataDogSendMetricsParamsWithContext creates a new DataDogSendMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataDogSendMetricsParamsWithContext(ctx context.Context) *DataDogSendMetricsParams {
	var ()
	return &DataDogSendMetricsParams{

		Context: ctx,
	}
}

/*DataDogSendMetricsParams contains all the parameters to send to the API endpoint
for the data dog send metrics operation typically these are written to a http.Request
*/
type DataDogSendMetricsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data dog send metrics params
func (o *DataDogSendMetricsParams) WithTimeout(timeout time.Duration) *DataDogSendMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data dog send metrics params
func (o *DataDogSendMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data dog send metrics params
func (o *DataDogSendMetricsParams) WithContext(ctx context.Context) *DataDogSendMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data dog send metrics params
func (o *DataDogSendMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the data dog send metrics params
func (o *DataDogSendMetricsParams) WithBody(body models.DefaultSelector) *DataDogSendMetricsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the data dog send metrics params
func (o *DataDogSendMetricsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DataDogSendMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

package social_channel

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

// NewSocialChannelFetchPopularPostsParams creates a new SocialChannelFetchPopularPostsParams object
// with the default values initialized.
func NewSocialChannelFetchPopularPostsParams() *SocialChannelFetchPopularPostsParams {
	var ()
	return &SocialChannelFetchPopularPostsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSocialChannelFetchPopularPostsParamsWithTimeout creates a new SocialChannelFetchPopularPostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSocialChannelFetchPopularPostsParamsWithTimeout(timeout time.Duration) *SocialChannelFetchPopularPostsParams {
	var ()
	return &SocialChannelFetchPopularPostsParams{

		timeout: timeout,
	}
}

// NewSocialChannelFetchPopularPostsParamsWithContext creates a new SocialChannelFetchPopularPostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSocialChannelFetchPopularPostsParamsWithContext(ctx context.Context) *SocialChannelFetchPopularPostsParams {
	var ()
	return &SocialChannelFetchPopularPostsParams{

		Context: ctx,
	}
}

/*SocialChannelFetchPopularPostsParams contains all the parameters to send to the API endpoint
for the social channel fetch popular posts operation typically these are written to a http.Request
*/
type SocialChannelFetchPopularPostsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the social channel fetch popular posts params
func (o *SocialChannelFetchPopularPostsParams) WithTimeout(timeout time.Duration) *SocialChannelFetchPopularPostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the social channel fetch popular posts params
func (o *SocialChannelFetchPopularPostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the social channel fetch popular posts params
func (o *SocialChannelFetchPopularPostsParams) WithContext(ctx context.Context) *SocialChannelFetchPopularPostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the social channel fetch popular posts params
func (o *SocialChannelFetchPopularPostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the social channel fetch popular posts params
func (o *SocialChannelFetchPopularPostsParams) WithBody(body models.DefaultSelector) *SocialChannelFetchPopularPostsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the social channel fetch popular posts params
func (o *SocialChannelFetchPopularPostsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SocialChannelFetchPopularPostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

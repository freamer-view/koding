package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// SocialMessageFetchPrivateMessagesReader is a Reader for the SocialMessageFetchPrivateMessages structure.
type SocialMessageFetchPrivateMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SocialMessageFetchPrivateMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSocialMessageFetchPrivateMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSocialMessageFetchPrivateMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSocialMessageFetchPrivateMessagesOK creates a SocialMessageFetchPrivateMessagesOK with default headers values
func NewSocialMessageFetchPrivateMessagesOK() *SocialMessageFetchPrivateMessagesOK {
	return &SocialMessageFetchPrivateMessagesOK{}
}

/*SocialMessageFetchPrivateMessagesOK handles this case with default header values.

Request processed successfully
*/
type SocialMessageFetchPrivateMessagesOK struct {
	Payload *models.DefaultResponse
}

func (o *SocialMessageFetchPrivateMessagesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.fetchPrivateMessages][%d] socialMessageFetchPrivateMessagesOK  %+v", 200, o.Payload)
}

func (o *SocialMessageFetchPrivateMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSocialMessageFetchPrivateMessagesUnauthorized creates a SocialMessageFetchPrivateMessagesUnauthorized with default headers values
func NewSocialMessageFetchPrivateMessagesUnauthorized() *SocialMessageFetchPrivateMessagesUnauthorized {
	return &SocialMessageFetchPrivateMessagesUnauthorized{}
}

/*SocialMessageFetchPrivateMessagesUnauthorized handles this case with default header values.

Unauthorized request
*/
type SocialMessageFetchPrivateMessagesUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SocialMessageFetchPrivateMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.fetchPrivateMessages][%d] socialMessageFetchPrivateMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *SocialMessageFetchPrivateMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

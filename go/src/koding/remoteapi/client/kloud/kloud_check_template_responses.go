package kloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// KloudCheckTemplateReader is a Reader for the KloudCheckTemplate structure.
type KloudCheckTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KloudCheckTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewKloudCheckTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewKloudCheckTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKloudCheckTemplateOK creates a KloudCheckTemplateOK with default headers values
func NewKloudCheckTemplateOK() *KloudCheckTemplateOK {
	return &KloudCheckTemplateOK{}
}

/*KloudCheckTemplateOK handles this case with default header values.

Request processed successfully
*/
type KloudCheckTemplateOK struct {
	Payload *models.DefaultResponse
}

func (o *KloudCheckTemplateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Kloud.checkTemplate][%d] kloudCheckTemplateOK  %+v", 200, o.Payload)
}

func (o *KloudCheckTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKloudCheckTemplateUnauthorized creates a KloudCheckTemplateUnauthorized with default headers values
func NewKloudCheckTemplateUnauthorized() *KloudCheckTemplateUnauthorized {
	return &KloudCheckTemplateUnauthorized{}
}

/*KloudCheckTemplateUnauthorized handles this case with default header values.

Unauthorized request
*/
type KloudCheckTemplateUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *KloudCheckTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Kloud.checkTemplate][%d] kloudCheckTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *KloudCheckTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

package collaboration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// CollaborationAddReader is a Reader for the CollaborationAdd structure.
type CollaborationAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollaborationAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCollaborationAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCollaborationAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCollaborationAddOK creates a CollaborationAddOK with default headers values
func NewCollaborationAddOK() *CollaborationAddOK {
	return &CollaborationAddOK{}
}

/*CollaborationAddOK handles this case with default header values.

Request processed successfully
*/
type CollaborationAddOK struct {
	Payload *models.DefaultResponse
}

func (o *CollaborationAddOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Collaboration.add][%d] collaborationAddOK  %+v", 200, o.Payload)
}

func (o *CollaborationAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollaborationAddUnauthorized creates a CollaborationAddUnauthorized with default headers values
func NewCollaborationAddUnauthorized() *CollaborationAddUnauthorized {
	return &CollaborationAddUnauthorized{}
}

/*CollaborationAddUnauthorized handles this case with default header values.

Unauthorized request
*/
type CollaborationAddUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *CollaborationAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Collaboration.add][%d] collaborationAddUnauthorized  %+v", 401, o.Payload)
}

func (o *CollaborationAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

package j_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JCredentialCloneReader is a Reader for the JCredentialClone structure.
type JCredentialCloneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JCredentialCloneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJCredentialCloneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJCredentialCloneOK creates a JCredentialCloneOK with default headers values
func NewJCredentialCloneOK() *JCredentialCloneOK {
	return &JCredentialCloneOK{}
}

/*JCredentialCloneOK handles this case with default header values.

OK
*/
type JCredentialCloneOK struct {
	Payload JCredentialCloneOKBody
}

func (o *JCredentialCloneOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JCredential.clone/{id}][%d] jCredentialCloneOK  %+v", 200, o.Payload)
}

func (o *JCredentialCloneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JCredentialCloneOKBody j credential clone o k body
swagger:model JCredentialCloneOKBody
*/
type JCredentialCloneOKBody struct {
	models.JCredential

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JCredentialCloneOKBody) UnmarshalJSON(raw []byte) error {

	var jCredentialCloneOKBodyAO0 models.JCredential
	if err := swag.ReadJSON(raw, &jCredentialCloneOKBodyAO0); err != nil {
		return err
	}
	o.JCredential = jCredentialCloneOKBodyAO0

	var jCredentialCloneOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jCredentialCloneOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jCredentialCloneOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JCredentialCloneOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jCredentialCloneOKBodyAO0, err := swag.WriteJSON(o.JCredential)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jCredentialCloneOKBodyAO0)

	jCredentialCloneOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jCredentialCloneOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j credential clone o k body
func (o *JCredentialCloneOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JCredential.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

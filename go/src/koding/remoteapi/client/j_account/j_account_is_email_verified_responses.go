package j_account

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

// JAccountIsEmailVerifiedReader is a Reader for the JAccountIsEmailVerified structure.
type JAccountIsEmailVerifiedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JAccountIsEmailVerifiedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJAccountIsEmailVerifiedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJAccountIsEmailVerifiedOK creates a JAccountIsEmailVerifiedOK with default headers values
func NewJAccountIsEmailVerifiedOK() *JAccountIsEmailVerifiedOK {
	return &JAccountIsEmailVerifiedOK{}
}

/*JAccountIsEmailVerifiedOK handles this case with default header values.

OK
*/
type JAccountIsEmailVerifiedOK struct {
	Payload JAccountIsEmailVerifiedOKBody
}

func (o *JAccountIsEmailVerifiedOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.isEmailVerified/{id}][%d] jAccountIsEmailVerifiedOK  %+v", 200, o.Payload)
}

func (o *JAccountIsEmailVerifiedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JAccountIsEmailVerifiedOKBody j account is email verified o k body
swagger:model JAccountIsEmailVerifiedOKBody
*/
type JAccountIsEmailVerifiedOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JAccountIsEmailVerifiedOKBody) UnmarshalJSON(raw []byte) error {

	var jAccountIsEmailVerifiedOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &jAccountIsEmailVerifiedOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = jAccountIsEmailVerifiedOKBodyAO0

	var jAccountIsEmailVerifiedOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jAccountIsEmailVerifiedOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jAccountIsEmailVerifiedOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JAccountIsEmailVerifiedOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jAccountIsEmailVerifiedOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountIsEmailVerifiedOKBodyAO0)

	jAccountIsEmailVerifiedOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountIsEmailVerifiedOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j account is email verified o k body
func (o *JAccountIsEmailVerifiedOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
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

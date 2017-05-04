package j_password_recovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JPasswordRecoveryResetPasswordReader is a Reader for the JPasswordRecoveryResetPassword structure.
type JPasswordRecoveryResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JPasswordRecoveryResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJPasswordRecoveryResetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJPasswordRecoveryResetPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJPasswordRecoveryResetPasswordOK creates a JPasswordRecoveryResetPasswordOK with default headers values
func NewJPasswordRecoveryResetPasswordOK() *JPasswordRecoveryResetPasswordOK {
	return &JPasswordRecoveryResetPasswordOK{}
}

/*JPasswordRecoveryResetPasswordOK handles this case with default header values.

Request processed successfully
*/
type JPasswordRecoveryResetPasswordOK struct {
	Payload *models.DefaultResponse
}

func (o *JPasswordRecoveryResetPasswordOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JPasswordRecovery.resetPassword][%d] jPasswordRecoveryResetPasswordOK  %+v", 200, o.Payload)
}

func (o *JPasswordRecoveryResetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJPasswordRecoveryResetPasswordUnauthorized creates a JPasswordRecoveryResetPasswordUnauthorized with default headers values
func NewJPasswordRecoveryResetPasswordUnauthorized() *JPasswordRecoveryResetPasswordUnauthorized {
	return &JPasswordRecoveryResetPasswordUnauthorized{}
}

/*JPasswordRecoveryResetPasswordUnauthorized handles this case with default header values.

Unauthorized request
*/
type JPasswordRecoveryResetPasswordUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JPasswordRecoveryResetPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JPasswordRecovery.resetPassword][%d] jPasswordRecoveryResetPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *JPasswordRecoveryResetPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

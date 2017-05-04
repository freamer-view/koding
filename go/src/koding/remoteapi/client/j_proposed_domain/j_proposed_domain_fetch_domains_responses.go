package j_proposed_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JProposedDomainFetchDomainsReader is a Reader for the JProposedDomainFetchDomains structure.
type JProposedDomainFetchDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JProposedDomainFetchDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJProposedDomainFetchDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJProposedDomainFetchDomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJProposedDomainFetchDomainsOK creates a JProposedDomainFetchDomainsOK with default headers values
func NewJProposedDomainFetchDomainsOK() *JProposedDomainFetchDomainsOK {
	return &JProposedDomainFetchDomainsOK{}
}

/*JProposedDomainFetchDomainsOK handles this case with default header values.

Request processed successfully
*/
type JProposedDomainFetchDomainsOK struct {
	Payload *models.DefaultResponse
}

func (o *JProposedDomainFetchDomainsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProposedDomain.fetchDomains][%d] jProposedDomainFetchDomainsOK  %+v", 200, o.Payload)
}

func (o *JProposedDomainFetchDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJProposedDomainFetchDomainsUnauthorized creates a JProposedDomainFetchDomainsUnauthorized with default headers values
func NewJProposedDomainFetchDomainsUnauthorized() *JProposedDomainFetchDomainsUnauthorized {
	return &JProposedDomainFetchDomainsUnauthorized{}
}

/*JProposedDomainFetchDomainsUnauthorized handles this case with default header values.

Unauthorized request
*/
type JProposedDomainFetchDomainsUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JProposedDomainFetchDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProposedDomain.fetchDomains][%d] jProposedDomainFetchDomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *JProposedDomainFetchDomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

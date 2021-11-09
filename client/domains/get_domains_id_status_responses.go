// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jxsl13/plesgo/models"
)

// GetDomainsIDStatusReader is a Reader for the GetDomainsIDStatus structure.
type GetDomainsIDStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainsIDStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomainsIDStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDomainsIDStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDomainsIDStatusOK creates a GetDomainsIDStatusOK with default headers values
func NewGetDomainsIDStatusOK() *GetDomainsIDStatusOK {
	return &GetDomainsIDStatusOK{}
}

/* GetDomainsIDStatusOK describes a response with status code 200, with default header values.

OK
*/
type GetDomainsIDStatusOK struct {
	Payload *models.DomainStatus
}

func (o *GetDomainsIDStatusOK) Error() string {
	return fmt.Sprintf("[GET /domains/{id}/status][%d] getDomainsIdStatusOK  %+v", 200, o.Payload)
}
func (o *GetDomainsIDStatusOK) GetPayload() *models.DomainStatus {
	return o.Payload
}

func (o *GetDomainsIDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainsIDStatusNotFound creates a GetDomainsIDStatusNotFound with default headers values
func NewGetDomainsIDStatusNotFound() *GetDomainsIDStatusNotFound {
	return &GetDomainsIDStatusNotFound{}
}

/* GetDomainsIDStatusNotFound describes a response with status code 404, with default header values.

Domain is not found
*/
type GetDomainsIDStatusNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetDomainsIDStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /domains/{id}/status][%d] getDomainsIdStatusNotFound  %+v", 404, o.Payload)
}
func (o *GetDomainsIDStatusNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDomainsIDStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
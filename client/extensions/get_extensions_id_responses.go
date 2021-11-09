// Code generated by go-swagger; DO NOT EDIT.

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jxsl13/plesgo/models"
)

// GetExtensionsIDReader is a Reader for the GetExtensionsID structure.
type GetExtensionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExtensionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetExtensionsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExtensionsIDOK creates a GetExtensionsIDOK with default headers values
func NewGetExtensionsIDOK() *GetExtensionsIDOK {
	return &GetExtensionsIDOK{}
}

/* GetExtensionsIDOK describes a response with status code 200, with default header values.

OK
*/
type GetExtensionsIDOK struct {
	Payload *models.Extension
}

func (o *GetExtensionsIDOK) Error() string {
	return fmt.Sprintf("[GET /extensions/{id}][%d] getExtensionsIdOK  %+v", 200, o.Payload)
}
func (o *GetExtensionsIDOK) GetPayload() *models.Extension {
	return o.Payload
}

func (o *GetExtensionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Extension)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionsIDNotFound creates a GetExtensionsIDNotFound with default headers values
func NewGetExtensionsIDNotFound() *GetExtensionsIDNotFound {
	return &GetExtensionsIDNotFound{}
}

/* GetExtensionsIDNotFound describes a response with status code 404, with default header values.

Extension is not installed
*/
type GetExtensionsIDNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetExtensionsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /extensions/{id}][%d] getExtensionsIdNotFound  %+v", 404, o.Payload)
}
func (o *GetExtensionsIDNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetExtensionsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

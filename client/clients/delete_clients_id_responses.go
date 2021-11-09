// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jxsl13/plesgo/models"
)

// DeleteClientsIDReader is a Reader for the DeleteClientsID structure.
type DeleteClientsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClientsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClientsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteClientsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteClientsIDOK creates a DeleteClientsIDOK with default headers values
func NewDeleteClientsIDOK() *DeleteClientsIDOK {
	return &DeleteClientsIDOK{}
}

/* DeleteClientsIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteClientsIDOK struct {
	Payload *models.CreatedResponse
}

func (o *DeleteClientsIDOK) Error() string {
	return fmt.Sprintf("[DELETE /clients/{id}][%d] deleteClientsIdOK  %+v", 200, o.Payload)
}
func (o *DeleteClientsIDOK) GetPayload() *models.CreatedResponse {
	return o.Payload
}

func (o *DeleteClientsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreatedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClientsIDNotFound creates a DeleteClientsIDNotFound with default headers values
func NewDeleteClientsIDNotFound() *DeleteClientsIDNotFound {
	return &DeleteClientsIDNotFound{}
}

/* DeleteClientsIDNotFound describes a response with status code 404, with default header values.

Client is not found
*/
type DeleteClientsIDNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DeleteClientsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /clients/{id}][%d] deleteClientsIdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteClientsIDNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClientsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jxsl13/plesgo/models"
)

// DeleteDbusersIDReader is a Reader for the DeleteDbusersID structure.
type DeleteDbusersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDbusersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDbusersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteDbusersIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDbusersIDOK creates a DeleteDbusersIDOK with default headers values
func NewDeleteDbusersIDOK() *DeleteDbusersIDOK {
	return &DeleteDbusersIDOK{}
}

/* DeleteDbusersIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteDbusersIDOK struct {
	Payload *models.StatusResponse
}

func (o *DeleteDbusersIDOK) Error() string {
	return fmt.Sprintf("[DELETE /dbusers/{id}][%d] deleteDbusersIdOK  %+v", 200, o.Payload)
}
func (o *DeleteDbusersIDOK) GetPayload() *models.StatusResponse {
	return o.Payload
}

func (o *DeleteDbusersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDbusersIDNotFound creates a DeleteDbusersIDNotFound with default headers values
func NewDeleteDbusersIDNotFound() *DeleteDbusersIDNotFound {
	return &DeleteDbusersIDNotFound{}
}

/* DeleteDbusersIDNotFound describes a response with status code 404, with default header values.

Database user is not found
*/
type DeleteDbusersIDNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DeleteDbusersIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dbusers/{id}][%d] deleteDbusersIdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDbusersIDNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteDbusersIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

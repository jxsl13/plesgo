// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jxsl13/plesgo/models"
)

// PostServerInitReader is a Reader for the PostServerInit structure.
type PostServerInitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostServerInitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostServerInitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostServerInitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostServerInitOK creates a PostServerInitOK with default headers values
func NewPostServerInitOK() *PostServerInitOK {
	return &PostServerInitOK{}
}

/* PostServerInitOK describes a response with status code 200, with default header values.

Inital server setup was succesfully performed
*/
type PostServerInitOK struct {
	Payload *models.StatusResponse
}

func (o *PostServerInitOK) Error() string {
	return fmt.Sprintf("[POST /server/init][%d] postServerInitOK  %+v", 200, o.Payload)
}
func (o *PostServerInitOK) GetPayload() *models.StatusResponse {
	return o.Payload
}

func (o *PostServerInitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServerInitBadRequest creates a PostServerInitBadRequest with default headers values
func NewPostServerInitBadRequest() *PostServerInitBadRequest {
	return &PostServerInitBadRequest{}
}

/* PostServerInitBadRequest describes a response with status code 400, with default header values.

Invalid request data
*/
type PostServerInitBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostServerInitBadRequest) Error() string {
	return fmt.Sprintf("[POST /server/init][%d] postServerInitBadRequest  %+v", 400, o.Payload)
}
func (o *PostServerInitBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostServerInitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

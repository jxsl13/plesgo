// Code generated by go-swagger; DO NOT EDIT.

package ftp_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jxsl13/plesgo/models"
)

// PostFtpusersReader is a Reader for the PostFtpusers structure.
type PostFtpusersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFtpusersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFtpusersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostFtpusersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostFtpusersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFtpusersOK creates a PostFtpusersOK with default headers values
func NewPostFtpusersOK() *PostFtpusersOK {
	return &PostFtpusersOK{}
}

/* PostFtpusersOK describes a response with status code 200, with default header values.

OK
*/
type PostFtpusersOK struct {
	Payload *models.FtpUser
}

func (o *PostFtpusersOK) Error() string {
	return fmt.Sprintf("[POST /ftpusers][%d] postFtpusersOK  %+v", 200, o.Payload)
}
func (o *PostFtpusersOK) GetPayload() *models.FtpUser {
	return o.Payload
}

func (o *PostFtpusersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FtpUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFtpusersBadRequest creates a PostFtpusersBadRequest with default headers values
func NewPostFtpusersBadRequest() *PostFtpusersBadRequest {
	return &PostFtpusersBadRequest{}
}

/* PostFtpusersBadRequest describes a response with status code 400, with default header values.

Invalid request data
*/
type PostFtpusersBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostFtpusersBadRequest) Error() string {
	return fmt.Sprintf("[POST /ftpusers][%d] postFtpusersBadRequest  %+v", 400, o.Payload)
}
func (o *PostFtpusersBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostFtpusersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFtpusersNotFound creates a PostFtpusersNotFound with default headers values
func NewPostFtpusersNotFound() *PostFtpusersNotFound {
	return &PostFtpusersNotFound{}
}

/* PostFtpusersNotFound describes a response with status code 404, with default header values.

Domain is not found
*/
type PostFtpusersNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PostFtpusersNotFound) Error() string {
	return fmt.Sprintf("[POST /ftpusers][%d] postFtpusersNotFound  %+v", 404, o.Payload)
}
func (o *PostFtpusersNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostFtpusersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetCliCommandsReader is a Reader for the GetCliCommands structure.
type GetCliCommandsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCliCommandsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCliCommandsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCliCommandsOK creates a GetCliCommandsOK with default headers values
func NewGetCliCommandsOK() *GetCliCommandsOK {
	return &GetCliCommandsOK{}
}

/* GetCliCommandsOK describes a response with status code 200, with default header values.

OK
*/
type GetCliCommandsOK struct {
	Payload []string
}

func (o *GetCliCommandsOK) Error() string {
	return fmt.Sprintf("[GET /cli/commands][%d] getCliCommandsOK  %+v", 200, o.Payload)
}
func (o *GetCliCommandsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetCliCommandsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

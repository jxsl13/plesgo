// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CliCallResponse Command execution result
//
// swagger:model CliCallResponse
type CliCallResponse struct {

	// Command return code.
	// Example: 0
	Code int64 `json:"code,omitempty"`

	// Command stderr.
	// Example: Execution failed
	Stderr string `json:"stderr,omitempty"`

	// Command stdout.
	// Example: Done
	Stdout string `json:"stdout,omitempty"`
}

// Validate validates this cli call response
func (m *CliCallResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cli call response based on context it is used
func (m *CliCallResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CliCallResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CliCallResponse) UnmarshalBinary(b []byte) error {
	var res CliCallResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

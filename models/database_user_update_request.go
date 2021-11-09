// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatabaseUserUpdateRequest database user update request
//
// swagger:model DatabaseUserUpdateRequest
type DatabaseUserUpdateRequest struct {

	// Database user login
	// Example: exampledbuser
	Login string `json:"login,omitempty"`

	// New user password
	// Example: changeme1Q**
	Password string `json:"password,omitempty"`
}

// Validate validates this database user update request
func (m *DatabaseUserUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database user update request based on context it is used
func (m *DatabaseUserUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseUserUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseUserUpdateRequest) UnmarshalBinary(b []byte) error {
	var res DatabaseUserUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatabaseUser database user
//
// swagger:model DatabaseUser
type DatabaseUser struct {

	// ID of the database
	// Example: 12
	DatabaseID int64 `json:"database_id,omitempty"`

	// ID of the database user
	// Example: 5
	ID int64 `json:"id,omitempty"`

	// Database user login
	// Example: exampledbuser
	Login string `json:"login,omitempty"`
}

// Validate validates this database user
func (m *DatabaseUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database user based on context it is used
func (m *DatabaseUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseUser) UnmarshalBinary(b []byte) error {
	var res DatabaseUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

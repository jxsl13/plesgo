// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExtensionInstallRequest Extension installation request
//
// swagger:model ExtensionInstallRequest
type ExtensionInstallRequest struct {

	// Installation by file.
	// Example: /path/to/extension.zip
	File string `json:"file,omitempty"`

	// Installation by Identifier.
	// Example: letsencrypt
	ID string `json:"id,omitempty"`

	// Installation by URL.
	// Example: https://example.com/catalog/letsencrypt.zip
	URL string `json:"url,omitempty"`
}

// Validate validates this extension install request
func (m *ExtensionInstallRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this extension install request based on context it is used
func (m *ExtensionInstallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionInstallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionInstallRequest) UnmarshalBinary(b []byte) error {
	var res ExtensionInstallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
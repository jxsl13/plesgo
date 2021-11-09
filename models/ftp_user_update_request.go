// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FtpUserUpdateRequest ftp user update request
//
// swagger:model FtpUserUpdateRequest
type FtpUserUpdateRequest struct {

	// Subdirectory of the WWW Root that user is restricted to
	// Example: /httpdocs
	Home string `json:"home,omitempty"`

	// User name in the system
	// Example: exampleuser
	Name string `json:"name,omitempty"`

	// User password
	Password string `json:"password,omitempty"`

	// permissions
	Permissions *FtpUserUpdateRequestPermissions `json:"permissions,omitempty"`

	// Hard disk quota in bytes (if supported by platform)
	// Example: -1
	Quota int64 `json:"quota,omitempty"`
}

// Validate validates this ftp user update request
func (m *FtpUserUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FtpUserUpdateRequest) validatePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ftp user update request based on the context it is used
func (m *FtpUserUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FtpUserUpdateRequest) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	if m.Permissions != nil {
		if err := m.Permissions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FtpUserUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FtpUserUpdateRequest) UnmarshalBinary(b []byte) error {
	var res FtpUserUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FtpUserUpdateRequestPermissions Access permissions of the user (if supported by platform)
//
// swagger:model FtpUserUpdateRequestPermissions
type FtpUserUpdateRequestPermissions struct {

	// read
	// Enum: [true false]
	Read string `json:"read,omitempty"`

	// write
	// Enum: [true false]
	Write string `json:"write,omitempty"`
}

// Validate validates this ftp user update request permissions
func (m *FtpUserUpdateRequestPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWrite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ftpUserUpdateRequestPermissionsTypeReadPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["true","false"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ftpUserUpdateRequestPermissionsTypeReadPropEnum = append(ftpUserUpdateRequestPermissionsTypeReadPropEnum, v)
	}
}

const (

	// FtpUserUpdateRequestPermissionsReadTrue captures enum value "true"
	FtpUserUpdateRequestPermissionsReadTrue string = "true"

	// FtpUserUpdateRequestPermissionsReadFalse captures enum value "false"
	FtpUserUpdateRequestPermissionsReadFalse string = "false"
)

// prop value enum
func (m *FtpUserUpdateRequestPermissions) validateReadEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ftpUserUpdateRequestPermissionsTypeReadPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FtpUserUpdateRequestPermissions) validateRead(formats strfmt.Registry) error {
	if swag.IsZero(m.Read) { // not required
		return nil
	}

	// value enum
	if err := m.validateReadEnum("permissions"+"."+"read", "body", m.Read); err != nil {
		return err
	}

	return nil
}

var ftpUserUpdateRequestPermissionsTypeWritePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["true","false"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ftpUserUpdateRequestPermissionsTypeWritePropEnum = append(ftpUserUpdateRequestPermissionsTypeWritePropEnum, v)
	}
}

const (

	// FtpUserUpdateRequestPermissionsWriteTrue captures enum value "true"
	FtpUserUpdateRequestPermissionsWriteTrue string = "true"

	// FtpUserUpdateRequestPermissionsWriteFalse captures enum value "false"
	FtpUserUpdateRequestPermissionsWriteFalse string = "false"
)

// prop value enum
func (m *FtpUserUpdateRequestPermissions) validateWriteEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ftpUserUpdateRequestPermissionsTypeWritePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FtpUserUpdateRequestPermissions) validateWrite(formats strfmt.Registry) error {
	if swag.IsZero(m.Write) { // not required
		return nil
	}

	// value enum
	if err := m.validateWriteEnum("permissions"+"."+"write", "body", m.Write); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ftp user update request permissions based on context it is used
func (m *FtpUserUpdateRequestPermissions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FtpUserUpdateRequestPermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FtpUserUpdateRequestPermissions) UnmarshalBinary(b []byte) error {
	var res FtpUserUpdateRequestPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

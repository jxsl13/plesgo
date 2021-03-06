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

// DomainResponse domain response
//
// swagger:model DomainResponse
type DomainResponse struct {

	// Domain display name.
	// Example: ascii-example.com
	// Required: true
	ASCIIName *string `json:"ascii_name"`

	// Base domain ID.
	// Example: 3
	// Required: true
	BaseDomainID *int64 `json:"base_domain_id"`

	// Creation date.
	// Example: 2016-11-13
	// Required: true
	// Format: date
	Created *strfmt.Date `json:"created"`

	// Domain GUID.
	// Example: b623e93d-dc72-4102-b5f0-ded427cf0fb1
	// Required: true
	GUID *string `json:"guid"`

	// Domain hosting type.
	// Required: true
	// Enum: [virtual standard_forwarding frame_forwarding none]
	HostingType *string `json:"hosting_type"`

	// Domain ID.
	// Example: 7
	// Required: true
	ID *int64 `json:"id"`

	// Domain name.
	// Example: example.com
	// Required: true
	Name *string `json:"name"`

	// WWW Root on filesystem
	// Example: /var/www/vhosts/example.com/httpdocs
	// Required: true
	WwwRoot *string `json:"www_root"`
}

// Validate validates this domain response
func (m *DomainResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateASCIIName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseDomainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWwwRoot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainResponse) validateASCIIName(formats strfmt.Registry) error {

	if err := validate.Required("ascii_name", "body", m.ASCIIName); err != nil {
		return err
	}

	return nil
}

func (m *DomainResponse) validateBaseDomainID(formats strfmt.Registry) error {

	if err := validate.Required("base_domain_id", "body", m.BaseDomainID); err != nil {
		return err
	}

	return nil
}

func (m *DomainResponse) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainResponse) validateGUID(formats strfmt.Registry) error {

	if err := validate.Required("guid", "body", m.GUID); err != nil {
		return err
	}

	return nil
}

var domainResponseTypeHostingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","standard_forwarding","frame_forwarding","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainResponseTypeHostingTypePropEnum = append(domainResponseTypeHostingTypePropEnum, v)
	}
}

const (

	// DomainResponseHostingTypeVirtual captures enum value "virtual"
	DomainResponseHostingTypeVirtual string = "virtual"

	// DomainResponseHostingTypeStandardForwarding captures enum value "standard_forwarding"
	DomainResponseHostingTypeStandardForwarding string = "standard_forwarding"

	// DomainResponseHostingTypeFrameForwarding captures enum value "frame_forwarding"
	DomainResponseHostingTypeFrameForwarding string = "frame_forwarding"

	// DomainResponseHostingTypeNone captures enum value "none"
	DomainResponseHostingTypeNone string = "none"
)

// prop value enum
func (m *DomainResponse) validateHostingTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, domainResponseTypeHostingTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DomainResponse) validateHostingType(formats strfmt.Registry) error {

	if err := validate.Required("hosting_type", "body", m.HostingType); err != nil {
		return err
	}

	// value enum
	if err := m.validateHostingTypeEnum("hosting_type", "body", *m.HostingType); err != nil {
		return err
	}

	return nil
}

func (m *DomainResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DomainResponse) validateWwwRoot(formats strfmt.Registry) error {

	if err := validate.Required("www_root", "body", m.WwwRoot); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain response based on context it is used
func (m *DomainResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainResponse) UnmarshalBinary(b []byte) error {
	var res DomainResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

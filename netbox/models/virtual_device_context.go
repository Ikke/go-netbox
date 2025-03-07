// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualDeviceContext virtual device context
//
// swagger:model VirtualDeviceContext
type VirtualDeviceContext struct {

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Identifier
	//
	// Numeric identifier unique to the parent device
	// Maximum: 32767
	// Minimum: 0
	Identifier *int64 `json:"identifier,omitempty"`

	// Interface count
	// Read Only: true
	InterfaceCount int64 `json:"interface_count,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// primary ip
	PrimaryIP *NestedIPAddress `json:"primary_ip,omitempty"`

	// primary ip4
	PrimaryIp4 *NestedIPAddress `json:"primary_ip4,omitempty"`

	// primary ip6
	PrimaryIp6 *NestedIPAddress `json:"primary_ip6,omitempty"`

	// Status
	// Required: true
	// Enum: [active planned offline]
	Status *string `json:"status"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this virtual device context
func (m *VirtualDeviceContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualDeviceContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualDeviceContext) validateIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.Identifier) { // not required
		return nil
	}

	if err := validate.MinimumInt("identifier", "body", *m.Identifier, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("identifier", "body", *m.Identifier, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) validatePrimaryIP(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIP) { // not required
		return nil
	}

	if m.PrimaryIP != nil {
		if err := m.PrimaryIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualDeviceContext) validatePrimaryIp4(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp4) { // not required
		return nil
	}

	if m.PrimaryIp4 != nil {
		if err := m.PrimaryIp4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualDeviceContext) validatePrimaryIp6(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp6) { // not required
		return nil
	}

	if m.PrimaryIp6 != nil {
		if err := m.PrimaryIp6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

var virtualDeviceContextTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","planned","offline"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualDeviceContextTypeStatusPropEnum = append(virtualDeviceContextTypeStatusPropEnum, v)
	}
}

const (

	// VirtualDeviceContextStatusActive captures enum value "active"
	VirtualDeviceContextStatusActive string = "active"

	// VirtualDeviceContextStatusPlanned captures enum value "planned"
	VirtualDeviceContextStatusPlanned string = "planned"

	// VirtualDeviceContextStatusOffline captures enum value "offline"
	VirtualDeviceContextStatusOffline string = "offline"
)

// prop value enum
func (m *VirtualDeviceContext) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, virtualDeviceContextTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VirtualDeviceContext) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualDeviceContext) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualDeviceContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this virtual device context based on the context it is used
func (m *VirtualDeviceContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaceCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIp4(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIp6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualDeviceContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {

		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidateInterfaceCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface_count", "body", int64(m.InterfaceCount)); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidatePrimaryIP(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIP != nil {

		if swag.IsZero(m.PrimaryIP) { // not required
			return nil
		}

		if err := m.PrimaryIP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidatePrimaryIp4(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIp4 != nil {

		if swag.IsZero(m.PrimaryIp4) { // not required
			return nil
		}

		if err := m.PrimaryIp4.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidatePrimaryIp6(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIp6 != nil {

		if swag.IsZero(m.PrimaryIp6) { // not required
			return nil
		}

		if err := m.PrimaryIp6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualDeviceContext) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {

		if swag.IsZero(m.Tenant) { // not required
			return nil
		}

		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualDeviceContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualDeviceContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualDeviceContext) UnmarshalBinary(b []byte) error {
	var res VirtualDeviceContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

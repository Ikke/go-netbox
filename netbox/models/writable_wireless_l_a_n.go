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

// WritableWirelessLAN writable wireless l a n
//
// swagger:model WritableWirelessLAN
type WritableWirelessLAN struct {

	// Auth cipher
	// Enum: [auto tkip aes]
	AuthCipher string `json:"auth_cipher,omitempty"`

	// Pre-shared key
	// Max Length: 64
	AuthPsk string `json:"auth_psk,omitempty"`

	// Auth Type
	// Enum: [open wep wpa-personal wpa-enterprise]
	AuthType string `json:"auth_type,omitempty"`

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

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Group
	Group *int64 `json:"group,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// SSID
	// Required: true
	// Max Length: 32
	// Min Length: 1
	Ssid *string `json:"ssid"`

	// Status
	// Enum: [active reserved disabled deprecated]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// VLAN
	Vlan *int64 `json:"vlan,omitempty"`
}

// Validate validates this writable wireless l a n
func (m *WritableWirelessLAN) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthCipher(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthPsk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

var writableWirelessLANTypeAuthCipherPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","tkip","aes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableWirelessLANTypeAuthCipherPropEnum = append(writableWirelessLANTypeAuthCipherPropEnum, v)
	}
}

const (

	// WritableWirelessLANAuthCipherAuto captures enum value "auto"
	WritableWirelessLANAuthCipherAuto string = "auto"

	// WritableWirelessLANAuthCipherTkip captures enum value "tkip"
	WritableWirelessLANAuthCipherTkip string = "tkip"

	// WritableWirelessLANAuthCipherAes captures enum value "aes"
	WritableWirelessLANAuthCipherAes string = "aes"
)

// prop value enum
func (m *WritableWirelessLAN) validateAuthCipherEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableWirelessLANTypeAuthCipherPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableWirelessLAN) validateAuthCipher(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthCipher) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthCipherEnum("auth_cipher", "body", m.AuthCipher); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) validateAuthPsk(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthPsk) { // not required
		return nil
	}

	if err := validate.MaxLength("auth_psk", "body", m.AuthPsk, 64); err != nil {
		return err
	}

	return nil
}

var writableWirelessLANTypeAuthTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["open","wep","wpa-personal","wpa-enterprise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableWirelessLANTypeAuthTypePropEnum = append(writableWirelessLANTypeAuthTypePropEnum, v)
	}
}

const (

	// WritableWirelessLANAuthTypeOpen captures enum value "open"
	WritableWirelessLANAuthTypeOpen string = "open"

	// WritableWirelessLANAuthTypeWep captures enum value "wep"
	WritableWirelessLANAuthTypeWep string = "wep"

	// WritableWirelessLANAuthTypeWpaDashPersonal captures enum value "wpa-personal"
	WritableWirelessLANAuthTypeWpaDashPersonal string = "wpa-personal"

	// WritableWirelessLANAuthTypeWpaDashEnterprise captures enum value "wpa-enterprise"
	WritableWirelessLANAuthTypeWpaDashEnterprise string = "wpa-enterprise"
)

// prop value enum
func (m *WritableWirelessLAN) validateAuthTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableWirelessLANTypeAuthTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableWirelessLAN) validateAuthType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthTypeEnum("auth_type", "body", m.AuthType); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) validateSsid(formats strfmt.Registry) error {

	if err := validate.Required("ssid", "body", m.Ssid); err != nil {
		return err
	}

	if err := validate.MinLength("ssid", "body", *m.Ssid, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("ssid", "body", *m.Ssid, 32); err != nil {
		return err
	}

	return nil
}

var writableWirelessLANTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","reserved","disabled","deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableWirelessLANTypeStatusPropEnum = append(writableWirelessLANTypeStatusPropEnum, v)
	}
}

const (

	// WritableWirelessLANStatusActive captures enum value "active"
	WritableWirelessLANStatusActive string = "active"

	// WritableWirelessLANStatusReserved captures enum value "reserved"
	WritableWirelessLANStatusReserved string = "reserved"

	// WritableWirelessLANStatusDisabled captures enum value "disabled"
	WritableWirelessLANStatusDisabled string = "disabled"

	// WritableWirelessLANStatusDeprecated captures enum value "deprecated"
	WritableWirelessLANStatusDeprecated string = "deprecated"
)

// prop value enum
func (m *WritableWirelessLAN) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableWirelessLANTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableWirelessLAN) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) validateTags(formats strfmt.Registry) error {
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

func (m *WritableWirelessLAN) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable wireless l a n based on the context it is used
func (m *WritableWirelessLAN) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
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

func (m *WritableWirelessLAN) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLAN) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritableWirelessLAN) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableWirelessLAN) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableWirelessLAN) UnmarshalBinary(b []byte) error {
	var res WritableWirelessLAN
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

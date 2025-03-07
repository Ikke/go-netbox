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

// WritableWirelessLink writable wireless link
//
// swagger:model WritableWirelessLink
type WritableWirelessLink struct {

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

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Interface A
	// Required: true
	Interfacea *int64 `json:"interface_a"`

	// Interface B
	// Required: true
	Interfaceb *int64 `json:"interface_b"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// SSID
	// Max Length: 32
	Ssid string `json:"ssid,omitempty"`

	// Status
	// Enum: [connected planned decommissioning]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable wireless link
func (m *WritableWirelessLink) Validate(formats strfmt.Registry) error {
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

	if err := m.validateInterfacea(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaceb(formats); err != nil {
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

var writableWirelessLinkTypeAuthCipherPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","tkip","aes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableWirelessLinkTypeAuthCipherPropEnum = append(writableWirelessLinkTypeAuthCipherPropEnum, v)
	}
}

const (

	// WritableWirelessLinkAuthCipherAuto captures enum value "auto"
	WritableWirelessLinkAuthCipherAuto string = "auto"

	// WritableWirelessLinkAuthCipherTkip captures enum value "tkip"
	WritableWirelessLinkAuthCipherTkip string = "tkip"

	// WritableWirelessLinkAuthCipherAes captures enum value "aes"
	WritableWirelessLinkAuthCipherAes string = "aes"
)

// prop value enum
func (m *WritableWirelessLink) validateAuthCipherEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableWirelessLinkTypeAuthCipherPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableWirelessLink) validateAuthCipher(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthCipher) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthCipherEnum("auth_cipher", "body", m.AuthCipher); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) validateAuthPsk(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthPsk) { // not required
		return nil
	}

	if err := validate.MaxLength("auth_psk", "body", m.AuthPsk, 64); err != nil {
		return err
	}

	return nil
}

var writableWirelessLinkTypeAuthTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["open","wep","wpa-personal","wpa-enterprise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableWirelessLinkTypeAuthTypePropEnum = append(writableWirelessLinkTypeAuthTypePropEnum, v)
	}
}

const (

	// WritableWirelessLinkAuthTypeOpen captures enum value "open"
	WritableWirelessLinkAuthTypeOpen string = "open"

	// WritableWirelessLinkAuthTypeWep captures enum value "wep"
	WritableWirelessLinkAuthTypeWep string = "wep"

	// WritableWirelessLinkAuthTypeWpaDashPersonal captures enum value "wpa-personal"
	WritableWirelessLinkAuthTypeWpaDashPersonal string = "wpa-personal"

	// WritableWirelessLinkAuthTypeWpaDashEnterprise captures enum value "wpa-enterprise"
	WritableWirelessLinkAuthTypeWpaDashEnterprise string = "wpa-enterprise"
)

// prop value enum
func (m *WritableWirelessLink) validateAuthTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableWirelessLinkTypeAuthTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableWirelessLink) validateAuthType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthTypeEnum("auth_type", "body", m.AuthType); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) validateInterfacea(formats strfmt.Registry) error {

	if err := validate.Required("interface_a", "body", m.Interfacea); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) validateInterfaceb(formats strfmt.Registry) error {

	if err := validate.Required("interface_b", "body", m.Interfaceb); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) validateSsid(formats strfmt.Registry) error {
	if swag.IsZero(m.Ssid) { // not required
		return nil
	}

	if err := validate.MaxLength("ssid", "body", m.Ssid, 32); err != nil {
		return err
	}

	return nil
}

var writableWirelessLinkTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","planned","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableWirelessLinkTypeStatusPropEnum = append(writableWirelessLinkTypeStatusPropEnum, v)
	}
}

const (

	// WritableWirelessLinkStatusConnected captures enum value "connected"
	WritableWirelessLinkStatusConnected string = "connected"

	// WritableWirelessLinkStatusPlanned captures enum value "planned"
	WritableWirelessLinkStatusPlanned string = "planned"

	// WritableWirelessLinkStatusDecommissioning captures enum value "decommissioning"
	WritableWirelessLinkStatusDecommissioning string = "decommissioning"
)

// prop value enum
func (m *WritableWirelessLink) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableWirelessLinkTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableWirelessLink) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) validateTags(formats strfmt.Registry) error {
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

func (m *WritableWirelessLink) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable wireless link based on the context it is used
func (m *WritableWirelessLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *WritableWirelessLink) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritableWirelessLink) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritableWirelessLink) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableWirelessLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableWirelessLink) UnmarshalBinary(b []byte) error {
	var res WritableWirelessLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

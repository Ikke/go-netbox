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

// WirelessLink wireless link
//
// swagger:model WirelessLink
type WirelessLink struct {

	// auth cipher
	AuthCipher *WirelessLinkAuthCipher `json:"auth_cipher,omitempty"`

	// Pre-shared key
	// Max Length: 64
	AuthPsk string `json:"auth_psk,omitempty"`

	// auth type
	AuthType *WirelessLinkAuthType `json:"auth_type,omitempty"`

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

	// interface a
	// Required: true
	Interfacea *NestedInterface `json:"interface_a"`

	// interface b
	// Required: true
	Interfaceb *NestedInterface `json:"interface_b"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// SSID
	// Max Length: 32
	Ssid string `json:"ssid,omitempty"`

	// status
	Status *WirelessLinkStatus `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this wireless link
func (m *WirelessLink) Validate(formats strfmt.Registry) error {
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

func (m *WirelessLink) validateAuthCipher(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthCipher) { // not required
		return nil
	}

	if m.AuthCipher != nil {
		if err := m.AuthCipher.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_cipher")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_cipher")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) validateAuthPsk(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthPsk) { // not required
		return nil
	}

	if err := validate.MaxLength("auth_psk", "body", m.AuthPsk, 64); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLink) validateAuthType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthType) { // not required
		return nil
	}

	if m.AuthType != nil {
		if err := m.AuthType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_type")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLink) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLink) validateInterfacea(formats strfmt.Registry) error {

	if err := validate.Required("interface_a", "body", m.Interfacea); err != nil {
		return err
	}

	if m.Interfacea != nil {
		if err := m.Interfacea.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_a")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface_a")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) validateInterfaceb(formats strfmt.Registry) error {

	if err := validate.Required("interface_b", "body", m.Interfaceb); err != nil {
		return err
	}

	if m.Interfaceb != nil {
		if err := m.Interfaceb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_b")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface_b")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLink) validateSsid(formats strfmt.Registry) error {
	if swag.IsZero(m.Ssid) { // not required
		return nil
	}

	if err := validate.MaxLength("ssid", "body", m.Ssid, 32); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLink) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) validateTags(formats strfmt.Registry) error {
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

func (m *WirelessLink) validateTenant(formats strfmt.Registry) error {
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

func (m *WirelessLink) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wireless link based on the context it is used
func (m *WirelessLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthCipher(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfacea(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaceb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
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

func (m *WirelessLink) contextValidateAuthCipher(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthCipher != nil {

		if swag.IsZero(m.AuthCipher) { // not required
			return nil
		}

		if err := m.AuthCipher.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_cipher")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_cipher")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) contextValidateAuthType(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthType != nil {

		if swag.IsZero(m.AuthType) { // not required
			return nil
		}

		if err := m.AuthType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_type")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLink) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLink) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLink) contextValidateInterfacea(ctx context.Context, formats strfmt.Registry) error {

	if m.Interfacea != nil {

		if err := m.Interfacea.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_a")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface_a")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) contextValidateInterfaceb(ctx context.Context, formats strfmt.Registry) error {

	if m.Interfaceb != nil {

		if err := m.Interfaceb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_b")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interface_b")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLink) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLink) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WirelessLink) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WirelessLink) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WirelessLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WirelessLink) UnmarshalBinary(b []byte) error {
	var res WirelessLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WirelessLinkAuthCipher Auth cipher
//
// swagger:model WirelessLinkAuthCipher
type WirelessLinkAuthCipher struct {

	// label
	// Required: true
	// Enum: [Auto TKIP AES]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [auto tkip aes]
	Value *string `json:"value"`
}

// Validate validates this wireless link auth cipher
func (m *WirelessLinkAuthCipher) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var wirelessLinkAuthCipherTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Auto","TKIP","AES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLinkAuthCipherTypeLabelPropEnum = append(wirelessLinkAuthCipherTypeLabelPropEnum, v)
	}
}

const (

	// WirelessLinkAuthCipherLabelAuto captures enum value "Auto"
	WirelessLinkAuthCipherLabelAuto string = "Auto"

	// WirelessLinkAuthCipherLabelTKIP captures enum value "TKIP"
	WirelessLinkAuthCipherLabelTKIP string = "TKIP"

	// WirelessLinkAuthCipherLabelAES captures enum value "AES"
	WirelessLinkAuthCipherLabelAES string = "AES"
)

// prop value enum
func (m *WirelessLinkAuthCipher) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLinkAuthCipherTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLinkAuthCipher) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("auth_cipher"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("auth_cipher"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var wirelessLinkAuthCipherTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","tkip","aes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLinkAuthCipherTypeValuePropEnum = append(wirelessLinkAuthCipherTypeValuePropEnum, v)
	}
}

const (

	// WirelessLinkAuthCipherValueAuto captures enum value "auto"
	WirelessLinkAuthCipherValueAuto string = "auto"

	// WirelessLinkAuthCipherValueTkip captures enum value "tkip"
	WirelessLinkAuthCipherValueTkip string = "tkip"

	// WirelessLinkAuthCipherValueAes captures enum value "aes"
	WirelessLinkAuthCipherValueAes string = "aes"
)

// prop value enum
func (m *WirelessLinkAuthCipher) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLinkAuthCipherTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLinkAuthCipher) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("auth_cipher"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("auth_cipher"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wireless link auth cipher based on context it is used
func (m *WirelessLinkAuthCipher) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WirelessLinkAuthCipher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WirelessLinkAuthCipher) UnmarshalBinary(b []byte) error {
	var res WirelessLinkAuthCipher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WirelessLinkAuthType Auth type
//
// swagger:model WirelessLinkAuthType
type WirelessLinkAuthType struct {

	// label
	// Required: true
	// Enum: [Open WEP WPA Personal (PSK) WPA Enterprise]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [open wep wpa-personal wpa-enterprise]
	Value *string `json:"value"`
}

// Validate validates this wireless link auth type
func (m *WirelessLinkAuthType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var wirelessLinkAuthTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Open","WEP","WPA Personal (PSK)","WPA Enterprise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLinkAuthTypeTypeLabelPropEnum = append(wirelessLinkAuthTypeTypeLabelPropEnum, v)
	}
}

const (

	// WirelessLinkAuthTypeLabelOpen captures enum value "Open"
	WirelessLinkAuthTypeLabelOpen string = "Open"

	// WirelessLinkAuthTypeLabelWEP captures enum value "WEP"
	WirelessLinkAuthTypeLabelWEP string = "WEP"

	// WirelessLinkAuthTypeLabelWPAPersonalPSK captures enum value "WPA Personal (PSK)"
	WirelessLinkAuthTypeLabelWPAPersonalPSK string = "WPA Personal (PSK)"

	// WirelessLinkAuthTypeLabelWPAEnterprise captures enum value "WPA Enterprise"
	WirelessLinkAuthTypeLabelWPAEnterprise string = "WPA Enterprise"
)

// prop value enum
func (m *WirelessLinkAuthType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLinkAuthTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLinkAuthType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("auth_type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("auth_type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var wirelessLinkAuthTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["open","wep","wpa-personal","wpa-enterprise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLinkAuthTypeTypeValuePropEnum = append(wirelessLinkAuthTypeTypeValuePropEnum, v)
	}
}

const (

	// WirelessLinkAuthTypeValueOpen captures enum value "open"
	WirelessLinkAuthTypeValueOpen string = "open"

	// WirelessLinkAuthTypeValueWep captures enum value "wep"
	WirelessLinkAuthTypeValueWep string = "wep"

	// WirelessLinkAuthTypeValueWpaDashPersonal captures enum value "wpa-personal"
	WirelessLinkAuthTypeValueWpaDashPersonal string = "wpa-personal"

	// WirelessLinkAuthTypeValueWpaDashEnterprise captures enum value "wpa-enterprise"
	WirelessLinkAuthTypeValueWpaDashEnterprise string = "wpa-enterprise"
)

// prop value enum
func (m *WirelessLinkAuthType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLinkAuthTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLinkAuthType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("auth_type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("auth_type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wireless link auth type based on context it is used
func (m *WirelessLinkAuthType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WirelessLinkAuthType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WirelessLinkAuthType) UnmarshalBinary(b []byte) error {
	var res WirelessLinkAuthType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WirelessLinkStatus Status
//
// swagger:model WirelessLinkStatus
type WirelessLinkStatus struct {

	// label
	// Required: true
	// Enum: [Connected Planned Decommissioning]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [connected planned decommissioning]
	Value *string `json:"value"`
}

// Validate validates this wireless link status
func (m *WirelessLinkStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var wirelessLinkStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Connected","Planned","Decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLinkStatusTypeLabelPropEnum = append(wirelessLinkStatusTypeLabelPropEnum, v)
	}
}

const (

	// WirelessLinkStatusLabelConnected captures enum value "Connected"
	WirelessLinkStatusLabelConnected string = "Connected"

	// WirelessLinkStatusLabelPlanned captures enum value "Planned"
	WirelessLinkStatusLabelPlanned string = "Planned"

	// WirelessLinkStatusLabelDecommissioning captures enum value "Decommissioning"
	WirelessLinkStatusLabelDecommissioning string = "Decommissioning"
)

// prop value enum
func (m *WirelessLinkStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLinkStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLinkStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var wirelessLinkStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","planned","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLinkStatusTypeValuePropEnum = append(wirelessLinkStatusTypeValuePropEnum, v)
	}
}

const (

	// WirelessLinkStatusValueConnected captures enum value "connected"
	WirelessLinkStatusValueConnected string = "connected"

	// WirelessLinkStatusValuePlanned captures enum value "planned"
	WirelessLinkStatusValuePlanned string = "planned"

	// WirelessLinkStatusValueDecommissioning captures enum value "decommissioning"
	WirelessLinkStatusValueDecommissioning string = "decommissioning"
)

// prop value enum
func (m *WirelessLinkStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLinkStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLinkStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wireless link status based on context it is used
func (m *WirelessLinkStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WirelessLinkStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WirelessLinkStatus) UnmarshalBinary(b []byte) error {
	var res WirelessLinkStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

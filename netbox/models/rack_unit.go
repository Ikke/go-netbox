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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RackUnit rack unit
//
// swagger:model RackUnit
type RackUnit struct {

	// device
	Device *NestedDevice `json:"device,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// face
	Face *RackUnitFace `json:"face,omitempty"`

	// Id
	// Read Only: true
	ID float64 `json:"id,omitempty"`

	// Name
	// Read Only: true
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// Occupied
	// Read Only: true
	Occupied *bool `json:"occupied,omitempty"`
}

// Validate validates this rack unit
func (m *RackUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackUnit) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.Device) { // not required
		return nil
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

func (m *RackUnit) validateFace(formats strfmt.Registry) error {
	if swag.IsZero(m.Face) { // not required
		return nil
	}

	if m.Face != nil {
		if err := m.Face.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("face")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("face")
			}
			return err
		}
	}

	return nil
}

func (m *RackUnit) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rack unit based on the context it is used
func (m *RackUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackUnit) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {

		if swag.IsZero(m.Device) { // not required
			return nil
		}

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

func (m *RackUnit) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *RackUnit) contextValidateFace(ctx context.Context, formats strfmt.Registry) error {

	if m.Face != nil {

		if swag.IsZero(m.Face) { // not required
			return nil
		}

		if err := m.Face.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("face")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("face")
			}
			return err
		}
	}

	return nil
}

func (m *RackUnit) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", float64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *RackUnit) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *RackUnit) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackUnit) UnmarshalBinary(b []byte) error {
	var res RackUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackUnitFace Face
//
// swagger:model RackUnitFace
type RackUnitFace struct {

	// label
	// Required: true
	// Enum: [Front Rear]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [front rear]
	Value *string `json:"value"`
}

// Validate validates this rack unit face
func (m *RackUnitFace) Validate(formats strfmt.Registry) error {
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

var rackUnitFaceTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Front","Rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackUnitFaceTypeLabelPropEnum = append(rackUnitFaceTypeLabelPropEnum, v)
	}
}

const (

	// RackUnitFaceLabelFront captures enum value "Front"
	RackUnitFaceLabelFront string = "Front"

	// RackUnitFaceLabelRear captures enum value "Rear"
	RackUnitFaceLabelRear string = "Rear"
)

// prop value enum
func (m *RackUnitFace) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackUnitFaceTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackUnitFace) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("face"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rackUnitFaceTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["front","rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackUnitFaceTypeValuePropEnum = append(rackUnitFaceTypeValuePropEnum, v)
	}
}

const (

	// RackUnitFaceValueFront captures enum value "front"
	RackUnitFaceValueFront string = "front"

	// RackUnitFaceValueRear captures enum value "rear"
	RackUnitFaceValueRear string = "rear"
)

// prop value enum
func (m *RackUnitFace) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rackUnitFaceTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RackUnitFace) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("face"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rack unit face based on the context it is used
func (m *RackUnitFace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RackUnitFace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackUnitFace) UnmarshalBinary(b []byte) error {
	var res RackUnitFace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

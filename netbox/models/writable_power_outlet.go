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

// WritablePowerOutlet writable power outlet
//
// swagger:model WritablePowerOutlet
type WritablePowerOutlet struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable end
	// Read Only: true
	// Min Length: 1
	CableEnd string `json:"cable_end,omitempty"`

	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoints []interface{} `json:"connected_endpoints"`

	// Connected endpoints reachable
	// Read Only: true
	ConnectedEndpointsReachable *bool `json:"connected_endpoints_reachable,omitempty"`

	// Connected endpoints type
	// Read Only: true
	ConnectedEndpointsType string `json:"connected_endpoints_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Feed leg
	//
	// Phase (for three-phase feeds)
	// Enum: [A B C]
	FeedLeg string `json:"feed_leg"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	//
	// Return the appropriate serializer for the link termination model.
	//
	// Read Only: true
	LinkPeers []interface{} `json:"link_peers"`

	// Link peers type
	// Read Only: true
	LinkPeersType string `json:"link_peers_type,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected"`

	// Module
	Module *int64 `json:"module"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Power port
	PowerPort *int64 `json:"power_port"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Type
	//
	// Physical port type
	// Enum: [iec-60320-c5 iec-60320-c7 iec-60320-c13 iec-60320-c15 iec-60320-c19 iec-60320-c21 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-1-15r nema-5-15r nema-5-20r nema-5-30r nema-5-50r nema-6-15r nema-6-20r nema-6-30r nema-6-50r nema-10-30r nema-10-50r nema-14-20r nema-14-30r nema-14-50r nema-14-60r nema-15-15r nema-15-20r nema-15-30r nema-15-50r nema-15-60r nema-l1-15r nema-l5-15r nema-l5-20r nema-l5-30r nema-l5-50r nema-l6-15r nema-l6-20r nema-l6-30r nema-l6-50r nema-l10-30r nema-l14-20r nema-l14-30r nema-l14-50r nema-l14-60r nema-l15-20r nema-l15-30r nema-l15-50r nema-l15-60r nema-l21-20r nema-l21-30r nema-l22-30r CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ita-e ita-f ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o ita-multistandard usb-a usb-micro-b usb-c dc-terminal hdot-cx saf-d-grid neutrik-powercon-20a neutrik-powercon-32a neutrik-powercon-true1 neutrik-powercon-true1-top ubiquiti-smartpower hardwired other]
	Type string `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable power outlet
func (m *WritablePowerOutlet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCableEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeedLeg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *WritablePowerOutlet) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritablePowerOutlet) validateCableEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.CableEnd) { // not required
		return nil
	}

	if err := validate.MinLength("cable_end", "body", m.CableEnd, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

var writablePowerOutletTypeFeedLegPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerOutletTypeFeedLegPropEnum = append(writablePowerOutletTypeFeedLegPropEnum, v)
	}
}

const (

	// WritablePowerOutletFeedLegA captures enum value "A"
	WritablePowerOutletFeedLegA string = "A"

	// WritablePowerOutletFeedLegB captures enum value "B"
	WritablePowerOutletFeedLegB string = "B"

	// WritablePowerOutletFeedLegC captures enum value "C"
	WritablePowerOutletFeedLegC string = "C"
)

// prop value enum
func (m *WritablePowerOutlet) validateFeedLegEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerOutletTypeFeedLegPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerOutlet) validateFeedLeg(formats strfmt.Registry) error {
	if swag.IsZero(m.FeedLeg) { // not required
		return nil
	}

	// value enum
	if err := m.validateFeedLegEnum("feed_leg", "body", m.FeedLeg); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateName(formats strfmt.Registry) error {

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

func (m *WritablePowerOutlet) validateTags(formats strfmt.Registry) error {
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

var writablePowerOutletTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c5","iec-60320-c7","iec-60320-c13","iec-60320-c15","iec-60320-c19","iec-60320-c21","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15r","nema-5-15r","nema-5-20r","nema-5-30r","nema-5-50r","nema-6-15r","nema-6-20r","nema-6-30r","nema-6-50r","nema-10-30r","nema-10-50r","nema-14-20r","nema-14-30r","nema-14-50r","nema-14-60r","nema-15-15r","nema-15-20r","nema-15-30r","nema-15-50r","nema-15-60r","nema-l1-15r","nema-l5-15r","nema-l5-20r","nema-l5-30r","nema-l5-50r","nema-l6-15r","nema-l6-20r","nema-l6-30r","nema-l6-50r","nema-l10-30r","nema-l14-20r","nema-l14-30r","nema-l14-50r","nema-l14-60r","nema-l15-20r","nema-l15-30r","nema-l15-50r","nema-l15-60r","nema-l21-20r","nema-l21-30r","nema-l22-30r","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ita-e","ita-f","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","ita-multistandard","usb-a","usb-micro-b","usb-c","dc-terminal","hdot-cx","saf-d-grid","neutrik-powercon-20a","neutrik-powercon-32a","neutrik-powercon-true1","neutrik-powercon-true1-top","ubiquiti-smartpower","hardwired","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerOutletTypeTypePropEnum = append(writablePowerOutletTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerOutletTypeIecDash60320DashC5 captures enum value "iec-60320-c5"
	WritablePowerOutletTypeIecDash60320DashC5 string = "iec-60320-c5"

	// WritablePowerOutletTypeIecDash60320DashC7 captures enum value "iec-60320-c7"
	WritablePowerOutletTypeIecDash60320DashC7 string = "iec-60320-c7"

	// WritablePowerOutletTypeIecDash60320DashC13 captures enum value "iec-60320-c13"
	WritablePowerOutletTypeIecDash60320DashC13 string = "iec-60320-c13"

	// WritablePowerOutletTypeIecDash60320DashC15 captures enum value "iec-60320-c15"
	WritablePowerOutletTypeIecDash60320DashC15 string = "iec-60320-c15"

	// WritablePowerOutletTypeIecDash60320DashC19 captures enum value "iec-60320-c19"
	WritablePowerOutletTypeIecDash60320DashC19 string = "iec-60320-c19"

	// WritablePowerOutletTypeIecDash60320DashC21 captures enum value "iec-60320-c21"
	WritablePowerOutletTypeIecDash60320DashC21 string = "iec-60320-c21"

	// WritablePowerOutletTypeIecDash60309DashpDashnDasheDash4h captures enum value "iec-60309-p-n-e-4h"
	WritablePowerOutletTypeIecDash60309DashpDashnDasheDash4h string = "iec-60309-p-n-e-4h"

	// WritablePowerOutletTypeIecDash60309DashpDashnDasheDash6h captures enum value "iec-60309-p-n-e-6h"
	WritablePowerOutletTypeIecDash60309DashpDashnDasheDash6h string = "iec-60309-p-n-e-6h"

	// WritablePowerOutletTypeIecDash60309DashpDashnDasheDash9h captures enum value "iec-60309-p-n-e-9h"
	WritablePowerOutletTypeIecDash60309DashpDashnDasheDash9h string = "iec-60309-p-n-e-9h"

	// WritablePowerOutletTypeIecDash60309Dash2pDasheDash4h captures enum value "iec-60309-2p-e-4h"
	WritablePowerOutletTypeIecDash60309Dash2pDasheDash4h string = "iec-60309-2p-e-4h"

	// WritablePowerOutletTypeIecDash60309Dash2pDasheDash6h captures enum value "iec-60309-2p-e-6h"
	WritablePowerOutletTypeIecDash60309Dash2pDasheDash6h string = "iec-60309-2p-e-6h"

	// WritablePowerOutletTypeIecDash60309Dash2pDasheDash9h captures enum value "iec-60309-2p-e-9h"
	WritablePowerOutletTypeIecDash60309Dash2pDasheDash9h string = "iec-60309-2p-e-9h"

	// WritablePowerOutletTypeIecDash60309Dash3pDasheDash4h captures enum value "iec-60309-3p-e-4h"
	WritablePowerOutletTypeIecDash60309Dash3pDasheDash4h string = "iec-60309-3p-e-4h"

	// WritablePowerOutletTypeIecDash60309Dash3pDasheDash6h captures enum value "iec-60309-3p-e-6h"
	WritablePowerOutletTypeIecDash60309Dash3pDasheDash6h string = "iec-60309-3p-e-6h"

	// WritablePowerOutletTypeIecDash60309Dash3pDasheDash9h captures enum value "iec-60309-3p-e-9h"
	WritablePowerOutletTypeIecDash60309Dash3pDasheDash9h string = "iec-60309-3p-e-9h"

	// WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash4h captures enum value "iec-60309-3p-n-e-4h"
	WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash4h string = "iec-60309-3p-n-e-4h"

	// WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash6h captures enum value "iec-60309-3p-n-e-6h"
	WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash6h string = "iec-60309-3p-n-e-6h"

	// WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash9h captures enum value "iec-60309-3p-n-e-9h"
	WritablePowerOutletTypeIecDash60309Dash3pDashnDasheDash9h string = "iec-60309-3p-n-e-9h"

	// WritablePowerOutletTypeNemaDash1Dash15r captures enum value "nema-1-15r"
	WritablePowerOutletTypeNemaDash1Dash15r string = "nema-1-15r"

	// WritablePowerOutletTypeNemaDash5Dash15r captures enum value "nema-5-15r"
	WritablePowerOutletTypeNemaDash5Dash15r string = "nema-5-15r"

	// WritablePowerOutletTypeNemaDash5Dash20r captures enum value "nema-5-20r"
	WritablePowerOutletTypeNemaDash5Dash20r string = "nema-5-20r"

	// WritablePowerOutletTypeNemaDash5Dash30r captures enum value "nema-5-30r"
	WritablePowerOutletTypeNemaDash5Dash30r string = "nema-5-30r"

	// WritablePowerOutletTypeNemaDash5Dash50r captures enum value "nema-5-50r"
	WritablePowerOutletTypeNemaDash5Dash50r string = "nema-5-50r"

	// WritablePowerOutletTypeNemaDash6Dash15r captures enum value "nema-6-15r"
	WritablePowerOutletTypeNemaDash6Dash15r string = "nema-6-15r"

	// WritablePowerOutletTypeNemaDash6Dash20r captures enum value "nema-6-20r"
	WritablePowerOutletTypeNemaDash6Dash20r string = "nema-6-20r"

	// WritablePowerOutletTypeNemaDash6Dash30r captures enum value "nema-6-30r"
	WritablePowerOutletTypeNemaDash6Dash30r string = "nema-6-30r"

	// WritablePowerOutletTypeNemaDash6Dash50r captures enum value "nema-6-50r"
	WritablePowerOutletTypeNemaDash6Dash50r string = "nema-6-50r"

	// WritablePowerOutletTypeNemaDash10Dash30r captures enum value "nema-10-30r"
	WritablePowerOutletTypeNemaDash10Dash30r string = "nema-10-30r"

	// WritablePowerOutletTypeNemaDash10Dash50r captures enum value "nema-10-50r"
	WritablePowerOutletTypeNemaDash10Dash50r string = "nema-10-50r"

	// WritablePowerOutletTypeNemaDash14Dash20r captures enum value "nema-14-20r"
	WritablePowerOutletTypeNemaDash14Dash20r string = "nema-14-20r"

	// WritablePowerOutletTypeNemaDash14Dash30r captures enum value "nema-14-30r"
	WritablePowerOutletTypeNemaDash14Dash30r string = "nema-14-30r"

	// WritablePowerOutletTypeNemaDash14Dash50r captures enum value "nema-14-50r"
	WritablePowerOutletTypeNemaDash14Dash50r string = "nema-14-50r"

	// WritablePowerOutletTypeNemaDash14Dash60r captures enum value "nema-14-60r"
	WritablePowerOutletTypeNemaDash14Dash60r string = "nema-14-60r"

	// WritablePowerOutletTypeNemaDash15Dash15r captures enum value "nema-15-15r"
	WritablePowerOutletTypeNemaDash15Dash15r string = "nema-15-15r"

	// WritablePowerOutletTypeNemaDash15Dash20r captures enum value "nema-15-20r"
	WritablePowerOutletTypeNemaDash15Dash20r string = "nema-15-20r"

	// WritablePowerOutletTypeNemaDash15Dash30r captures enum value "nema-15-30r"
	WritablePowerOutletTypeNemaDash15Dash30r string = "nema-15-30r"

	// WritablePowerOutletTypeNemaDash15Dash50r captures enum value "nema-15-50r"
	WritablePowerOutletTypeNemaDash15Dash50r string = "nema-15-50r"

	// WritablePowerOutletTypeNemaDash15Dash60r captures enum value "nema-15-60r"
	WritablePowerOutletTypeNemaDash15Dash60r string = "nema-15-60r"

	// WritablePowerOutletTypeNemaDashL1Dash15r captures enum value "nema-l1-15r"
	WritablePowerOutletTypeNemaDashL1Dash15r string = "nema-l1-15r"

	// WritablePowerOutletTypeNemaDashL5Dash15r captures enum value "nema-l5-15r"
	WritablePowerOutletTypeNemaDashL5Dash15r string = "nema-l5-15r"

	// WritablePowerOutletTypeNemaDashL5Dash20r captures enum value "nema-l5-20r"
	WritablePowerOutletTypeNemaDashL5Dash20r string = "nema-l5-20r"

	// WritablePowerOutletTypeNemaDashL5Dash30r captures enum value "nema-l5-30r"
	WritablePowerOutletTypeNemaDashL5Dash30r string = "nema-l5-30r"

	// WritablePowerOutletTypeNemaDashL5Dash50r captures enum value "nema-l5-50r"
	WritablePowerOutletTypeNemaDashL5Dash50r string = "nema-l5-50r"

	// WritablePowerOutletTypeNemaDashL6Dash15r captures enum value "nema-l6-15r"
	WritablePowerOutletTypeNemaDashL6Dash15r string = "nema-l6-15r"

	// WritablePowerOutletTypeNemaDashL6Dash20r captures enum value "nema-l6-20r"
	WritablePowerOutletTypeNemaDashL6Dash20r string = "nema-l6-20r"

	// WritablePowerOutletTypeNemaDashL6Dash30r captures enum value "nema-l6-30r"
	WritablePowerOutletTypeNemaDashL6Dash30r string = "nema-l6-30r"

	// WritablePowerOutletTypeNemaDashL6Dash50r captures enum value "nema-l6-50r"
	WritablePowerOutletTypeNemaDashL6Dash50r string = "nema-l6-50r"

	// WritablePowerOutletTypeNemaDashL10Dash30r captures enum value "nema-l10-30r"
	WritablePowerOutletTypeNemaDashL10Dash30r string = "nema-l10-30r"

	// WritablePowerOutletTypeNemaDashL14Dash20r captures enum value "nema-l14-20r"
	WritablePowerOutletTypeNemaDashL14Dash20r string = "nema-l14-20r"

	// WritablePowerOutletTypeNemaDashL14Dash30r captures enum value "nema-l14-30r"
	WritablePowerOutletTypeNemaDashL14Dash30r string = "nema-l14-30r"

	// WritablePowerOutletTypeNemaDashL14Dash50r captures enum value "nema-l14-50r"
	WritablePowerOutletTypeNemaDashL14Dash50r string = "nema-l14-50r"

	// WritablePowerOutletTypeNemaDashL14Dash60r captures enum value "nema-l14-60r"
	WritablePowerOutletTypeNemaDashL14Dash60r string = "nema-l14-60r"

	// WritablePowerOutletTypeNemaDashL15Dash20r captures enum value "nema-l15-20r"
	WritablePowerOutletTypeNemaDashL15Dash20r string = "nema-l15-20r"

	// WritablePowerOutletTypeNemaDashL15Dash30r captures enum value "nema-l15-30r"
	WritablePowerOutletTypeNemaDashL15Dash30r string = "nema-l15-30r"

	// WritablePowerOutletTypeNemaDashL15Dash50r captures enum value "nema-l15-50r"
	WritablePowerOutletTypeNemaDashL15Dash50r string = "nema-l15-50r"

	// WritablePowerOutletTypeNemaDashL15Dash60r captures enum value "nema-l15-60r"
	WritablePowerOutletTypeNemaDashL15Dash60r string = "nema-l15-60r"

	// WritablePowerOutletTypeNemaDashL21Dash20r captures enum value "nema-l21-20r"
	WritablePowerOutletTypeNemaDashL21Dash20r string = "nema-l21-20r"

	// WritablePowerOutletTypeNemaDashL21Dash30r captures enum value "nema-l21-30r"
	WritablePowerOutletTypeNemaDashL21Dash30r string = "nema-l21-30r"

	// WritablePowerOutletTypeNemaDashL22Dash30r captures enum value "nema-l22-30r"
	WritablePowerOutletTypeNemaDashL22Dash30r string = "nema-l22-30r"

	// WritablePowerOutletTypeCS6360C captures enum value "CS6360C"
	WritablePowerOutletTypeCS6360C string = "CS6360C"

	// WritablePowerOutletTypeCS6364C captures enum value "CS6364C"
	WritablePowerOutletTypeCS6364C string = "CS6364C"

	// WritablePowerOutletTypeCS8164C captures enum value "CS8164C"
	WritablePowerOutletTypeCS8164C string = "CS8164C"

	// WritablePowerOutletTypeCS8264C captures enum value "CS8264C"
	WritablePowerOutletTypeCS8264C string = "CS8264C"

	// WritablePowerOutletTypeCS8364C captures enum value "CS8364C"
	WritablePowerOutletTypeCS8364C string = "CS8364C"

	// WritablePowerOutletTypeCS8464C captures enum value "CS8464C"
	WritablePowerOutletTypeCS8464C string = "CS8464C"

	// WritablePowerOutletTypeItaDashe captures enum value "ita-e"
	WritablePowerOutletTypeItaDashe string = "ita-e"

	// WritablePowerOutletTypeItaDashf captures enum value "ita-f"
	WritablePowerOutletTypeItaDashf string = "ita-f"

	// WritablePowerOutletTypeItaDashg captures enum value "ita-g"
	WritablePowerOutletTypeItaDashg string = "ita-g"

	// WritablePowerOutletTypeItaDashh captures enum value "ita-h"
	WritablePowerOutletTypeItaDashh string = "ita-h"

	// WritablePowerOutletTypeItaDashi captures enum value "ita-i"
	WritablePowerOutletTypeItaDashi string = "ita-i"

	// WritablePowerOutletTypeItaDashj captures enum value "ita-j"
	WritablePowerOutletTypeItaDashj string = "ita-j"

	// WritablePowerOutletTypeItaDashk captures enum value "ita-k"
	WritablePowerOutletTypeItaDashk string = "ita-k"

	// WritablePowerOutletTypeItaDashl captures enum value "ita-l"
	WritablePowerOutletTypeItaDashl string = "ita-l"

	// WritablePowerOutletTypeItaDashm captures enum value "ita-m"
	WritablePowerOutletTypeItaDashm string = "ita-m"

	// WritablePowerOutletTypeItaDashn captures enum value "ita-n"
	WritablePowerOutletTypeItaDashn string = "ita-n"

	// WritablePowerOutletTypeItaDasho captures enum value "ita-o"
	WritablePowerOutletTypeItaDasho string = "ita-o"

	// WritablePowerOutletTypeItaDashMultistandard captures enum value "ita-multistandard"
	WritablePowerOutletTypeItaDashMultistandard string = "ita-multistandard"

	// WritablePowerOutletTypeUsbDasha captures enum value "usb-a"
	WritablePowerOutletTypeUsbDasha string = "usb-a"

	// WritablePowerOutletTypeUsbDashMicroDashb captures enum value "usb-micro-b"
	WritablePowerOutletTypeUsbDashMicroDashb string = "usb-micro-b"

	// WritablePowerOutletTypeUsbDashc captures enum value "usb-c"
	WritablePowerOutletTypeUsbDashc string = "usb-c"

	// WritablePowerOutletTypeDcDashTerminal captures enum value "dc-terminal"
	WritablePowerOutletTypeDcDashTerminal string = "dc-terminal"

	// WritablePowerOutletTypeHdotDashCx captures enum value "hdot-cx"
	WritablePowerOutletTypeHdotDashCx string = "hdot-cx"

	// WritablePowerOutletTypeSafDashdDashGrid captures enum value "saf-d-grid"
	WritablePowerOutletTypeSafDashdDashGrid string = "saf-d-grid"

	// WritablePowerOutletTypeNeutrikDashPowerconDash20a captures enum value "neutrik-powercon-20a"
	WritablePowerOutletTypeNeutrikDashPowerconDash20a string = "neutrik-powercon-20a"

	// WritablePowerOutletTypeNeutrikDashPowerconDash32a captures enum value "neutrik-powercon-32a"
	WritablePowerOutletTypeNeutrikDashPowerconDash32a string = "neutrik-powercon-32a"

	// WritablePowerOutletTypeNeutrikDashPowerconDashTrue1 captures enum value "neutrik-powercon-true1"
	WritablePowerOutletTypeNeutrikDashPowerconDashTrue1 string = "neutrik-powercon-true1"

	// WritablePowerOutletTypeNeutrikDashPowerconDashTrue1DashTop captures enum value "neutrik-powercon-true1-top"
	WritablePowerOutletTypeNeutrikDashPowerconDashTrue1DashTop string = "neutrik-powercon-true1-top"

	// WritablePowerOutletTypeUbiquitiDashSmartpower captures enum value "ubiquiti-smartpower"
	WritablePowerOutletTypeUbiquitiDashSmartpower string = "ubiquiti-smartpower"

	// WritablePowerOutletTypeHardwired captures enum value "hardwired"
	WritablePowerOutletTypeHardwired string = "hardwired"

	// WritablePowerOutletTypeOther captures enum value "other"
	WritablePowerOutletTypeOther string = "other"
)

// prop value enum
func (m *WritablePowerOutlet) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerOutletTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerOutlet) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable power outlet based on the context it is used
func (m *WritablePowerOutlet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCableEnd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointsReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointsType(ctx, formats); err != nil {
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

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeersType(ctx, formats); err != nil {
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

func (m *WritablePowerOutlet) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {

		if swag.IsZero(m.Cable) { // not required
			return nil
		}

		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateCableEnd(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_end", "body", string(m.CableEnd)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateConnectedEndpoints(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints", "body", []interface{}(m.ConnectedEndpoints)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateConnectedEndpointsReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_reachable", "body", m.ConnectedEndpointsReachable); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateConnectedEndpointsType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_type", "body", string(m.ConnectedEndpointsType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateLinkPeers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers", "body", []interface{}(m.LinkPeers)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateLinkPeersType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers_type", "body", string(m.LinkPeersType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutlet) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritablePowerOutlet) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerOutlet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerOutlet) UnmarshalBinary(b []byte) error {
	var res WritablePowerOutlet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

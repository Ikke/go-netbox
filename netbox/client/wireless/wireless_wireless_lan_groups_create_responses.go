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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// WirelessWirelessLanGroupsCreateReader is a Reader for the WirelessWirelessLanGroupsCreate structure.
type WirelessWirelessLanGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLanGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWirelessWirelessLanGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLanGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLanGroupsCreateCreated creates a WirelessWirelessLanGroupsCreateCreated with default headers values
func NewWirelessWirelessLanGroupsCreateCreated() *WirelessWirelessLanGroupsCreateCreated {
	return &WirelessWirelessLanGroupsCreateCreated{}
}

/*
WirelessWirelessLanGroupsCreateCreated describes a response with status code 201, with default header values.

WirelessWirelessLanGroupsCreateCreated wireless wireless lan groups create created
*/
type WirelessWirelessLanGroupsCreateCreated struct {
	Payload *models.WirelessLANGroup
}

// IsSuccess returns true when this wireless wireless lan groups create created response has a 2xx status code
func (o *WirelessWirelessLanGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lan groups create created response has a 3xx status code
func (o *WirelessWirelessLanGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lan groups create created response has a 4xx status code
func (o *WirelessWirelessLanGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lan groups create created response has a 5xx status code
func (o *WirelessWirelessLanGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lan groups create created response a status code equal to that given
func (o *WirelessWirelessLanGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the wireless wireless lan groups create created response
func (o *WirelessWirelessLanGroupsCreateCreated) Code() int {
	return 201
}

func (o *WirelessWirelessLanGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /wireless/wireless-lan-groups/][%d] wirelessWirelessLanGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *WirelessWirelessLanGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /wireless/wireless-lan-groups/][%d] wirelessWirelessLanGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *WirelessWirelessLanGroupsCreateCreated) GetPayload() *models.WirelessLANGroup {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLanGroupsCreateDefault creates a WirelessWirelessLanGroupsCreateDefault with default headers values
func NewWirelessWirelessLanGroupsCreateDefault(code int) *WirelessWirelessLanGroupsCreateDefault {
	return &WirelessWirelessLanGroupsCreateDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLanGroupsCreateDefault describes a response with status code -1, with default header values.

WirelessWirelessLanGroupsCreateDefault wireless wireless lan groups create default
*/
type WirelessWirelessLanGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this wireless wireless lan groups create default response has a 2xx status code
func (o *WirelessWirelessLanGroupsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless lan groups create default response has a 3xx status code
func (o *WirelessWirelessLanGroupsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless lan groups create default response has a 4xx status code
func (o *WirelessWirelessLanGroupsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless lan groups create default response has a 5xx status code
func (o *WirelessWirelessLanGroupsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless lan groups create default response a status code equal to that given
func (o *WirelessWirelessLanGroupsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wireless wireless lan groups create default response
func (o *WirelessWirelessLanGroupsCreateDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLanGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /wireless/wireless-lan-groups/][%d] wireless_wireless-lan-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLanGroupsCreateDefault) String() string {
	return fmt.Sprintf("[POST /wireless/wireless-lan-groups/][%d] wireless_wireless-lan-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLanGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

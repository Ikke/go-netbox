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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamFhrpGroupsUpdateReader is a Reader for the IpamFhrpGroupsUpdate structure.
type IpamFhrpGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupsUpdateOK creates a IpamFhrpGroupsUpdateOK with default headers values
func NewIpamFhrpGroupsUpdateOK() *IpamFhrpGroupsUpdateOK {
	return &IpamFhrpGroupsUpdateOK{}
}

/*
IpamFhrpGroupsUpdateOK describes a response with status code 200, with default header values.

IpamFhrpGroupsUpdateOK ipam fhrp groups update o k
*/
type IpamFhrpGroupsUpdateOK struct {
	Payload *models.FHRPGroup
}

// IsSuccess returns true when this ipam fhrp groups update o k response has a 2xx status code
func (o *IpamFhrpGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups update o k response has a 3xx status code
func (o *IpamFhrpGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups update o k response has a 4xx status code
func (o *IpamFhrpGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups update o k response has a 5xx status code
func (o *IpamFhrpGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups update o k response a status code equal to that given
func (o *IpamFhrpGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam fhrp groups update o k response
func (o *IpamFhrpGroupsUpdateOK) Code() int {
	return 200
}

func (o *IpamFhrpGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupsUpdateOK) GetPayload() *models.FHRPGroup {
	return o.Payload
}

func (o *IpamFhrpGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupsUpdateDefault creates a IpamFhrpGroupsUpdateDefault with default headers values
func NewIpamFhrpGroupsUpdateDefault(code int) *IpamFhrpGroupsUpdateDefault {
	return &IpamFhrpGroupsUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupsUpdateDefault describes a response with status code -1, with default header values.

IpamFhrpGroupsUpdateDefault ipam fhrp groups update default
*/
type IpamFhrpGroupsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam fhrp groups update default response has a 2xx status code
func (o *IpamFhrpGroupsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp groups update default response has a 3xx status code
func (o *IpamFhrpGroupsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp groups update default response has a 4xx status code
func (o *IpamFhrpGroupsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp groups update default response has a 5xx status code
func (o *IpamFhrpGroupsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp groups update default response a status code equal to that given
func (o *IpamFhrpGroupsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam fhrp groups update default response
func (o *IpamFhrpGroupsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamFhrpGroupsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/{id}/][%d] ipam_fhrp-groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/{id}/][%d] ipam_fhrp-groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// IpamRolesPartialUpdateReader is a Reader for the IpamRolesPartialUpdate structure.
type IpamRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRolesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRolesPartialUpdateOK creates a IpamRolesPartialUpdateOK with default headers values
func NewIpamRolesPartialUpdateOK() *IpamRolesPartialUpdateOK {
	return &IpamRolesPartialUpdateOK{}
}

/*
IpamRolesPartialUpdateOK describes a response with status code 200, with default header values.

IpamRolesPartialUpdateOK ipam roles partial update o k
*/
type IpamRolesPartialUpdateOK struct {
	Payload *models.Role
}

// IsSuccess returns true when this ipam roles partial update o k response has a 2xx status code
func (o *IpamRolesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam roles partial update o k response has a 3xx status code
func (o *IpamRolesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam roles partial update o k response has a 4xx status code
func (o *IpamRolesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam roles partial update o k response has a 5xx status code
func (o *IpamRolesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam roles partial update o k response a status code equal to that given
func (o *IpamRolesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam roles partial update o k response
func (o *IpamRolesPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamRolesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/roles/{id}/][%d] ipamRolesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRolesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/roles/{id}/][%d] ipamRolesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRolesPartialUpdateOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRolesPartialUpdateDefault creates a IpamRolesPartialUpdateDefault with default headers values
func NewIpamRolesPartialUpdateDefault(code int) *IpamRolesPartialUpdateDefault {
	return &IpamRolesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamRolesPartialUpdateDefault describes a response with status code -1, with default header values.

IpamRolesPartialUpdateDefault ipam roles partial update default
*/
type IpamRolesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam roles partial update default response has a 2xx status code
func (o *IpamRolesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam roles partial update default response has a 3xx status code
func (o *IpamRolesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam roles partial update default response has a 4xx status code
func (o *IpamRolesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam roles partial update default response has a 5xx status code
func (o *IpamRolesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam roles partial update default response a status code equal to that given
func (o *IpamRolesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam roles partial update default response
func (o *IpamRolesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamRolesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/roles/{id}/][%d] ipam_roles_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRolesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/roles/{id}/][%d] ipam_roles_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRolesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRolesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

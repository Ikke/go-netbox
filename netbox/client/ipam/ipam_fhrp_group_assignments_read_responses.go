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

// IpamFhrpGroupAssignmentsReadReader is a Reader for the IpamFhrpGroupAssignmentsRead structure.
type IpamFhrpGroupAssignmentsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupAssignmentsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupAssignmentsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupAssignmentsReadOK creates a IpamFhrpGroupAssignmentsReadOK with default headers values
func NewIpamFhrpGroupAssignmentsReadOK() *IpamFhrpGroupAssignmentsReadOK {
	return &IpamFhrpGroupAssignmentsReadOK{}
}

/*
IpamFhrpGroupAssignmentsReadOK describes a response with status code 200, with default header values.

IpamFhrpGroupAssignmentsReadOK ipam fhrp group assignments read o k
*/
type IpamFhrpGroupAssignmentsReadOK struct {
	Payload *models.FHRPGroupAssignment
}

// IsSuccess returns true when this ipam fhrp group assignments read o k response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp group assignments read o k response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments read o k response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp group assignments read o k response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments read o k response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam fhrp group assignments read o k response
func (o *IpamFhrpGroupAssignmentsReadOK) Code() int {
	return 200
}

func (o *IpamFhrpGroupAssignmentsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsReadOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsReadOK) String() string {
	return fmt.Sprintf("[GET /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsReadOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsReadOK) GetPayload() *models.FHRPGroupAssignment {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroupAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupAssignmentsReadDefault creates a IpamFhrpGroupAssignmentsReadDefault with default headers values
func NewIpamFhrpGroupAssignmentsReadDefault(code int) *IpamFhrpGroupAssignmentsReadDefault {
	return &IpamFhrpGroupAssignmentsReadDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupAssignmentsReadDefault describes a response with status code -1, with default header values.

IpamFhrpGroupAssignmentsReadDefault ipam fhrp group assignments read default
*/
type IpamFhrpGroupAssignmentsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam fhrp group assignments read default response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp group assignments read default response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp group assignments read default response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp group assignments read default response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp group assignments read default response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam fhrp group assignments read default response
func (o *IpamFhrpGroupAssignmentsReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamFhrpGroupAssignmentsReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/fhrp-group-assignments/{id}/][%d] ipam_fhrp-group-assignments_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsReadDefault) String() string {
	return fmt.Sprintf("[GET /ipam/fhrp-group-assignments/{id}/][%d] ipam_fhrp-group-assignments_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

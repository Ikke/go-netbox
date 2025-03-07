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

// IpamVrfsCreateReader is a Reader for the IpamVrfsCreate structure.
type IpamVrfsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVrfsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsCreateCreated creates a IpamVrfsCreateCreated with default headers values
func NewIpamVrfsCreateCreated() *IpamVrfsCreateCreated {
	return &IpamVrfsCreateCreated{}
}

/*
IpamVrfsCreateCreated describes a response with status code 201, with default header values.

IpamVrfsCreateCreated ipam vrfs create created
*/
type IpamVrfsCreateCreated struct {
	Payload *models.VRF
}

// IsSuccess returns true when this ipam vrfs create created response has a 2xx status code
func (o *IpamVrfsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs create created response has a 3xx status code
func (o *IpamVrfsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs create created response has a 4xx status code
func (o *IpamVrfsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs create created response has a 5xx status code
func (o *IpamVrfsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs create created response a status code equal to that given
func (o *IpamVrfsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam vrfs create created response
func (o *IpamVrfsCreateCreated) Code() int {
	return 201
}

func (o *IpamVrfsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vrfs/][%d] ipamVrfsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVrfsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/vrfs/][%d] ipamVrfsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVrfsCreateCreated) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVrfsCreateDefault creates a IpamVrfsCreateDefault with default headers values
func NewIpamVrfsCreateDefault(code int) *IpamVrfsCreateDefault {
	return &IpamVrfsCreateDefault{
		_statusCode: code,
	}
}

/*
IpamVrfsCreateDefault describes a response with status code -1, with default header values.

IpamVrfsCreateDefault ipam vrfs create default
*/
type IpamVrfsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam vrfs create default response has a 2xx status code
func (o *IpamVrfsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vrfs create default response has a 3xx status code
func (o *IpamVrfsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vrfs create default response has a 4xx status code
func (o *IpamVrfsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vrfs create default response has a 5xx status code
func (o *IpamVrfsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vrfs create default response a status code equal to that given
func (o *IpamVrfsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam vrfs create default response
func (o *IpamVrfsCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVrfsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/vrfs/][%d] ipam_vrfs_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/vrfs/][%d] ipam_vrfs_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

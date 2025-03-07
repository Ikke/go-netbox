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

// IpamPrefixesAvailableIpsCreateReader is a Reader for the IpamPrefixesAvailableIpsCreate structure.
type IpamPrefixesAvailableIpsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailableIpsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamPrefixesAvailableIpsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesAvailableIpsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesAvailableIpsCreateCreated creates a IpamPrefixesAvailableIpsCreateCreated with default headers values
func NewIpamPrefixesAvailableIpsCreateCreated() *IpamPrefixesAvailableIpsCreateCreated {
	return &IpamPrefixesAvailableIpsCreateCreated{}
}

/*
IpamPrefixesAvailableIpsCreateCreated describes a response with status code 201, with default header values.

IpamPrefixesAvailableIpsCreateCreated ipam prefixes available ips create created
*/
type IpamPrefixesAvailableIpsCreateCreated struct {
	Payload []*models.IPAddress
}

// IsSuccess returns true when this ipam prefixes available ips create created response has a 2xx status code
func (o *IpamPrefixesAvailableIpsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes available ips create created response has a 3xx status code
func (o *IpamPrefixesAvailableIpsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes available ips create created response has a 4xx status code
func (o *IpamPrefixesAvailableIpsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes available ips create created response has a 5xx status code
func (o *IpamPrefixesAvailableIpsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes available ips create created response a status code equal to that given
func (o *IpamPrefixesAvailableIpsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam prefixes available ips create created response
func (o *IpamPrefixesAvailableIpsCreateCreated) Code() int {
	return 201
}

func (o *IpamPrefixesAvailableIpsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamPrefixesAvailableIpsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamPrefixesAvailableIpsCreateCreated) GetPayload() []*models.IPAddress {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesAvailableIpsCreateDefault creates a IpamPrefixesAvailableIpsCreateDefault with default headers values
func NewIpamPrefixesAvailableIpsCreateDefault(code int) *IpamPrefixesAvailableIpsCreateDefault {
	return &IpamPrefixesAvailableIpsCreateDefault{
		_statusCode: code,
	}
}

/*
IpamPrefixesAvailableIpsCreateDefault describes a response with status code -1, with default header values.

IpamPrefixesAvailableIpsCreateDefault ipam prefixes available ips create default
*/
type IpamPrefixesAvailableIpsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam prefixes available ips create default response has a 2xx status code
func (o *IpamPrefixesAvailableIpsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam prefixes available ips create default response has a 3xx status code
func (o *IpamPrefixesAvailableIpsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam prefixes available ips create default response has a 4xx status code
func (o *IpamPrefixesAvailableIpsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam prefixes available ips create default response has a 5xx status code
func (o *IpamPrefixesAvailableIpsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam prefixes available ips create default response a status code equal to that given
func (o *IpamPrefixesAvailableIpsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam prefixes available ips create default response
func (o *IpamPrefixesAvailableIpsCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamPrefixesAvailableIpsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-ips/][%d] ipam_prefixes_available-ips_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesAvailableIpsCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-ips/][%d] ipam_prefixes_available-ips_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesAvailableIpsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

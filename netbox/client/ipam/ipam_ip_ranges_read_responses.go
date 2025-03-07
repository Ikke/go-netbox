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

// IpamIPRangesReadReader is a Reader for the IpamIPRangesRead structure.
type IpamIPRangesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPRangesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPRangesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPRangesReadOK creates a IpamIPRangesReadOK with default headers values
func NewIpamIPRangesReadOK() *IpamIPRangesReadOK {
	return &IpamIPRangesReadOK{}
}

/*
IpamIPRangesReadOK describes a response with status code 200, with default header values.

IpamIPRangesReadOK ipam Ip ranges read o k
*/
type IpamIPRangesReadOK struct {
	Payload *models.IPRange
}

// IsSuccess returns true when this ipam Ip ranges read o k response has a 2xx status code
func (o *IpamIPRangesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges read o k response has a 3xx status code
func (o *IpamIPRangesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges read o k response has a 4xx status code
func (o *IpamIPRangesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges read o k response has a 5xx status code
func (o *IpamIPRangesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges read o k response a status code equal to that given
func (o *IpamIPRangesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam Ip ranges read o k response
func (o *IpamIPRangesReadOK) Code() int {
	return 200
}

func (o *IpamIPRangesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/ip-ranges/{id}/][%d] ipamIpRangesReadOK  %+v", 200, o.Payload)
}

func (o *IpamIPRangesReadOK) String() string {
	return fmt.Sprintf("[GET /ipam/ip-ranges/{id}/][%d] ipamIpRangesReadOK  %+v", 200, o.Payload)
}

func (o *IpamIPRangesReadOK) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPRangesReadDefault creates a IpamIPRangesReadDefault with default headers values
func NewIpamIPRangesReadDefault(code int) *IpamIPRangesReadDefault {
	return &IpamIPRangesReadDefault{
		_statusCode: code,
	}
}

/*
IpamIPRangesReadDefault describes a response with status code -1, with default header values.

IpamIPRangesReadDefault ipam ip ranges read default
*/
type IpamIPRangesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam ip ranges read default response has a 2xx status code
func (o *IpamIPRangesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip ranges read default response has a 3xx status code
func (o *IpamIPRangesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip ranges read default response has a 4xx status code
func (o *IpamIPRangesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip ranges read default response has a 5xx status code
func (o *IpamIPRangesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip ranges read default response a status code equal to that given
func (o *IpamIPRangesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam ip ranges read default response
func (o *IpamIPRangesReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamIPRangesReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/ip-ranges/{id}/][%d] ipam_ip-ranges_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesReadDefault) String() string {
	return fmt.Sprintf("[GET /ipam/ip-ranges/{id}/][%d] ipam_ip-ranges_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// IpamServiceTemplatesCreateReader is a Reader for the IpamServiceTemplatesCreate structure.
type IpamServiceTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServiceTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamServiceTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServiceTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServiceTemplatesCreateCreated creates a IpamServiceTemplatesCreateCreated with default headers values
func NewIpamServiceTemplatesCreateCreated() *IpamServiceTemplatesCreateCreated {
	return &IpamServiceTemplatesCreateCreated{}
}

/*
IpamServiceTemplatesCreateCreated describes a response with status code 201, with default header values.

IpamServiceTemplatesCreateCreated ipam service templates create created
*/
type IpamServiceTemplatesCreateCreated struct {
	Payload *models.ServiceTemplate
}

// IsSuccess returns true when this ipam service templates create created response has a 2xx status code
func (o *IpamServiceTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam service templates create created response has a 3xx status code
func (o *IpamServiceTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam service templates create created response has a 4xx status code
func (o *IpamServiceTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam service templates create created response has a 5xx status code
func (o *IpamServiceTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam service templates create created response a status code equal to that given
func (o *IpamServiceTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam service templates create created response
func (o *IpamServiceTemplatesCreateCreated) Code() int {
	return 201
}

func (o *IpamServiceTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/service-templates/][%d] ipamServiceTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamServiceTemplatesCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/service-templates/][%d] ipamServiceTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamServiceTemplatesCreateCreated) GetPayload() *models.ServiceTemplate {
	return o.Payload
}

func (o *IpamServiceTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServiceTemplatesCreateDefault creates a IpamServiceTemplatesCreateDefault with default headers values
func NewIpamServiceTemplatesCreateDefault(code int) *IpamServiceTemplatesCreateDefault {
	return &IpamServiceTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*
IpamServiceTemplatesCreateDefault describes a response with status code -1, with default header values.

IpamServiceTemplatesCreateDefault ipam service templates create default
*/
type IpamServiceTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam service templates create default response has a 2xx status code
func (o *IpamServiceTemplatesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam service templates create default response has a 3xx status code
func (o *IpamServiceTemplatesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam service templates create default response has a 4xx status code
func (o *IpamServiceTemplatesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam service templates create default response has a 5xx status code
func (o *IpamServiceTemplatesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam service templates create default response a status code equal to that given
func (o *IpamServiceTemplatesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam service templates create default response
func (o *IpamServiceTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamServiceTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/service-templates/][%d] ipam_service-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServiceTemplatesCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/service-templates/][%d] ipam_service-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServiceTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServiceTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

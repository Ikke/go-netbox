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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// TenancyContactRolesCreateReader is a Reader for the TenancyContactRolesCreate structure.
type TenancyContactRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyContactRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactRolesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactRolesCreateCreated creates a TenancyContactRolesCreateCreated with default headers values
func NewTenancyContactRolesCreateCreated() *TenancyContactRolesCreateCreated {
	return &TenancyContactRolesCreateCreated{}
}

/*
TenancyContactRolesCreateCreated describes a response with status code 201, with default header values.

TenancyContactRolesCreateCreated tenancy contact roles create created
*/
type TenancyContactRolesCreateCreated struct {
	Payload *models.ContactRole
}

// IsSuccess returns true when this tenancy contact roles create created response has a 2xx status code
func (o *TenancyContactRolesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact roles create created response has a 3xx status code
func (o *TenancyContactRolesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact roles create created response has a 4xx status code
func (o *TenancyContactRolesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact roles create created response has a 5xx status code
func (o *TenancyContactRolesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact roles create created response a status code equal to that given
func (o *TenancyContactRolesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the tenancy contact roles create created response
func (o *TenancyContactRolesCreateCreated) Code() int {
	return 201
}

func (o *TenancyContactRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /tenancy/contact-roles/][%d] tenancyContactRolesCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyContactRolesCreateCreated) String() string {
	return fmt.Sprintf("[POST /tenancy/contact-roles/][%d] tenancyContactRolesCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyContactRolesCreateCreated) GetPayload() *models.ContactRole {
	return o.Payload
}

func (o *TenancyContactRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactRolesCreateDefault creates a TenancyContactRolesCreateDefault with default headers values
func NewTenancyContactRolesCreateDefault(code int) *TenancyContactRolesCreateDefault {
	return &TenancyContactRolesCreateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactRolesCreateDefault describes a response with status code -1, with default header values.

TenancyContactRolesCreateDefault tenancy contact roles create default
*/
type TenancyContactRolesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact roles create default response has a 2xx status code
func (o *TenancyContactRolesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact roles create default response has a 3xx status code
func (o *TenancyContactRolesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact roles create default response has a 4xx status code
func (o *TenancyContactRolesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact roles create default response has a 5xx status code
func (o *TenancyContactRolesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact roles create default response a status code equal to that given
func (o *TenancyContactRolesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact roles create default response
func (o *TenancyContactRolesCreateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactRolesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /tenancy/contact-roles/][%d] tenancy_contact-roles_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesCreateDefault) String() string {
	return fmt.Sprintf("[POST /tenancy/contact-roles/][%d] tenancy_contact-roles_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactRolesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

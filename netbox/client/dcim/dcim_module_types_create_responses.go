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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimModuleTypesCreateReader is a Reader for the DcimModuleTypesCreate structure.
type DcimModuleTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimModuleTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleTypesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleTypesCreateCreated creates a DcimModuleTypesCreateCreated with default headers values
func NewDcimModuleTypesCreateCreated() *DcimModuleTypesCreateCreated {
	return &DcimModuleTypesCreateCreated{}
}

/*
DcimModuleTypesCreateCreated describes a response with status code 201, with default header values.

DcimModuleTypesCreateCreated dcim module types create created
*/
type DcimModuleTypesCreateCreated struct {
	Payload *models.ModuleType
}

// IsSuccess returns true when this dcim module types create created response has a 2xx status code
func (o *DcimModuleTypesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module types create created response has a 3xx status code
func (o *DcimModuleTypesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module types create created response has a 4xx status code
func (o *DcimModuleTypesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module types create created response has a 5xx status code
func (o *DcimModuleTypesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module types create created response a status code equal to that given
func (o *DcimModuleTypesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim module types create created response
func (o *DcimModuleTypesCreateCreated) Code() int {
	return 201
}

func (o *DcimModuleTypesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/module-types/][%d] dcimModuleTypesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimModuleTypesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/module-types/][%d] dcimModuleTypesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimModuleTypesCreateCreated) GetPayload() *models.ModuleType {
	return o.Payload
}

func (o *DcimModuleTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleTypesCreateDefault creates a DcimModuleTypesCreateDefault with default headers values
func NewDcimModuleTypesCreateDefault(code int) *DcimModuleTypesCreateDefault {
	return &DcimModuleTypesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleTypesCreateDefault describes a response with status code -1, with default header values.

DcimModuleTypesCreateDefault dcim module types create default
*/
type DcimModuleTypesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim module types create default response has a 2xx status code
func (o *DcimModuleTypesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module types create default response has a 3xx status code
func (o *DcimModuleTypesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module types create default response has a 4xx status code
func (o *DcimModuleTypesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module types create default response has a 5xx status code
func (o *DcimModuleTypesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module types create default response a status code equal to that given
func (o *DcimModuleTypesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim module types create default response
func (o *DcimModuleTypesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimModuleTypesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/module-types/][%d] dcim_module-types_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/module-types/][%d] dcim_module-types_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleTypesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

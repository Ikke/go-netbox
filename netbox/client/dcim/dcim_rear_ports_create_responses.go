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

// DcimRearPortsCreateReader is a Reader for the DcimRearPortsCreate structure.
type DcimRearPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRearPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortsCreateCreated creates a DcimRearPortsCreateCreated with default headers values
func NewDcimRearPortsCreateCreated() *DcimRearPortsCreateCreated {
	return &DcimRearPortsCreateCreated{}
}

/*
DcimRearPortsCreateCreated describes a response with status code 201, with default header values.

DcimRearPortsCreateCreated dcim rear ports create created
*/
type DcimRearPortsCreateCreated struct {
	Payload *models.RearPort
}

// IsSuccess returns true when this dcim rear ports create created response has a 2xx status code
func (o *DcimRearPortsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear ports create created response has a 3xx status code
func (o *DcimRearPortsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear ports create created response has a 4xx status code
func (o *DcimRearPortsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear ports create created response has a 5xx status code
func (o *DcimRearPortsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear ports create created response a status code equal to that given
func (o *DcimRearPortsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim rear ports create created response
func (o *DcimRearPortsCreateCreated) Code() int {
	return 201
}

func (o *DcimRearPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rear-ports/][%d] dcimRearPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRearPortsCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/rear-ports/][%d] dcimRearPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRearPortsCreateCreated) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortsCreateDefault creates a DcimRearPortsCreateDefault with default headers values
func NewDcimRearPortsCreateDefault(code int) *DcimRearPortsCreateDefault {
	return &DcimRearPortsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortsCreateDefault describes a response with status code -1, with default header values.

DcimRearPortsCreateDefault dcim rear ports create default
*/
type DcimRearPortsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rear ports create default response has a 2xx status code
func (o *DcimRearPortsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear ports create default response has a 3xx status code
func (o *DcimRearPortsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear ports create default response has a 4xx status code
func (o *DcimRearPortsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear ports create default response has a 5xx status code
func (o *DcimRearPortsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear ports create default response a status code equal to that given
func (o *DcimRearPortsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rear ports create default response
func (o *DcimRearPortsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/rear-ports/][%d] dcim_rear-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/rear-ports/][%d] dcim_rear-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

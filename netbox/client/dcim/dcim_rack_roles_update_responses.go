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

// DcimRackRolesUpdateReader is a Reader for the DcimRackRolesUpdate structure.
type DcimRackRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackRolesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackRolesUpdateOK creates a DcimRackRolesUpdateOK with default headers values
func NewDcimRackRolesUpdateOK() *DcimRackRolesUpdateOK {
	return &DcimRackRolesUpdateOK{}
}

/*
DcimRackRolesUpdateOK describes a response with status code 200, with default header values.

DcimRackRolesUpdateOK dcim rack roles update o k
*/
type DcimRackRolesUpdateOK struct {
	Payload *models.RackRole
}

// IsSuccess returns true when this dcim rack roles update o k response has a 2xx status code
func (o *DcimRackRolesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack roles update o k response has a 3xx status code
func (o *DcimRackRolesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles update o k response has a 4xx status code
func (o *DcimRackRolesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack roles update o k response has a 5xx status code
func (o *DcimRackRolesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles update o k response a status code equal to that given
func (o *DcimRackRolesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rack roles update o k response
func (o *DcimRackRolesUpdateOK) Code() int {
	return 200
}

func (o *DcimRackRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/{id}/][%d] dcimRackRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/{id}/][%d] dcimRackRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesUpdateOK) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesUpdateDefault creates a DcimRackRolesUpdateDefault with default headers values
func NewDcimRackRolesUpdateDefault(code int) *DcimRackRolesUpdateDefault {
	return &DcimRackRolesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRackRolesUpdateDefault describes a response with status code -1, with default header values.

DcimRackRolesUpdateDefault dcim rack roles update default
*/
type DcimRackRolesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rack roles update default response has a 2xx status code
func (o *DcimRackRolesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rack roles update default response has a 3xx status code
func (o *DcimRackRolesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rack roles update default response has a 4xx status code
func (o *DcimRackRolesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rack roles update default response has a 5xx status code
func (o *DcimRackRolesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rack roles update default response a status code equal to that given
func (o *DcimRackRolesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rack roles update default response
func (o *DcimRackRolesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackRolesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/{id}/][%d] dcim_rack-roles_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/{id}/][%d] dcim_rack-roles_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

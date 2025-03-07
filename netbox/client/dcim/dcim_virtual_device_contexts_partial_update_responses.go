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

// DcimVirtualDeviceContextsPartialUpdateReader is a Reader for the DcimVirtualDeviceContextsPartialUpdate structure.
type DcimVirtualDeviceContextsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualDeviceContextsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualDeviceContextsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualDeviceContextsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualDeviceContextsPartialUpdateOK creates a DcimVirtualDeviceContextsPartialUpdateOK with default headers values
func NewDcimVirtualDeviceContextsPartialUpdateOK() *DcimVirtualDeviceContextsPartialUpdateOK {
	return &DcimVirtualDeviceContextsPartialUpdateOK{}
}

/*
DcimVirtualDeviceContextsPartialUpdateOK describes a response with status code 200, with default header values.

DcimVirtualDeviceContextsPartialUpdateOK dcim virtual device contexts partial update o k
*/
type DcimVirtualDeviceContextsPartialUpdateOK struct {
	Payload *models.VirtualDeviceContext
}

// IsSuccess returns true when this dcim virtual device contexts partial update o k response has a 2xx status code
func (o *DcimVirtualDeviceContextsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual device contexts partial update o k response has a 3xx status code
func (o *DcimVirtualDeviceContextsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual device contexts partial update o k response has a 4xx status code
func (o *DcimVirtualDeviceContextsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual device contexts partial update o k response has a 5xx status code
func (o *DcimVirtualDeviceContextsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual device contexts partial update o k response a status code equal to that given
func (o *DcimVirtualDeviceContextsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim virtual device contexts partial update o k response
func (o *DcimVirtualDeviceContextsPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimVirtualDeviceContextsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualDeviceContextsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualDeviceContextsPartialUpdateOK) GetPayload() *models.VirtualDeviceContext {
	return o.Payload
}

func (o *DcimVirtualDeviceContextsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualDeviceContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualDeviceContextsPartialUpdateDefault creates a DcimVirtualDeviceContextsPartialUpdateDefault with default headers values
func NewDcimVirtualDeviceContextsPartialUpdateDefault(code int) *DcimVirtualDeviceContextsPartialUpdateDefault {
	return &DcimVirtualDeviceContextsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimVirtualDeviceContextsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimVirtualDeviceContextsPartialUpdateDefault dcim virtual device contexts partial update default
*/
type DcimVirtualDeviceContextsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim virtual device contexts partial update default response has a 2xx status code
func (o *DcimVirtualDeviceContextsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim virtual device contexts partial update default response has a 3xx status code
func (o *DcimVirtualDeviceContextsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim virtual device contexts partial update default response has a 4xx status code
func (o *DcimVirtualDeviceContextsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim virtual device contexts partial update default response has a 5xx status code
func (o *DcimVirtualDeviceContextsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim virtual device contexts partial update default response a status code equal to that given
func (o *DcimVirtualDeviceContextsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim virtual device contexts partial update default response
func (o *DcimVirtualDeviceContextsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimVirtualDeviceContextsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-device-contexts/{id}/][%d] dcim_virtual-device-contexts_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualDeviceContextsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-device-contexts/{id}/][%d] dcim_virtual-device-contexts_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualDeviceContextsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualDeviceContextsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

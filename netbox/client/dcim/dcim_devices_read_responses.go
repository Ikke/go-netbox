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

// DcimDevicesReadReader is a Reader for the DcimDevicesRead structure.
type DcimDevicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDevicesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesReadOK creates a DcimDevicesReadOK with default headers values
func NewDcimDevicesReadOK() *DcimDevicesReadOK {
	return &DcimDevicesReadOK{}
}

/*
DcimDevicesReadOK describes a response with status code 200, with default header values.

DcimDevicesReadOK dcim devices read o k
*/
type DcimDevicesReadOK struct {
	Payload *models.DeviceWithConfigContext
}

// IsSuccess returns true when this dcim devices read o k response has a 2xx status code
func (o *DcimDevicesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim devices read o k response has a 3xx status code
func (o *DcimDevicesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices read o k response has a 4xx status code
func (o *DcimDevicesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim devices read o k response has a 5xx status code
func (o *DcimDevicesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices read o k response a status code equal to that given
func (o *DcimDevicesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim devices read o k response
func (o *DcimDevicesReadOK) Code() int {
	return 200
}

func (o *DcimDevicesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/][%d] dcimDevicesReadOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/][%d] dcimDevicesReadOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesReadOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDevicesReadDefault creates a DcimDevicesReadDefault with default headers values
func NewDcimDevicesReadDefault(code int) *DcimDevicesReadDefault {
	return &DcimDevicesReadDefault{
		_statusCode: code,
	}
}

/*
DcimDevicesReadDefault describes a response with status code -1, with default header values.

DcimDevicesReadDefault dcim devices read default
*/
type DcimDevicesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim devices read default response has a 2xx status code
func (o *DcimDevicesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim devices read default response has a 3xx status code
func (o *DcimDevicesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim devices read default response has a 4xx status code
func (o *DcimDevicesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim devices read default response has a 5xx status code
func (o *DcimDevicesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim devices read default response a status code equal to that given
func (o *DcimDevicesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim devices read default response
func (o *DcimDevicesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimDevicesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/][%d] dcim_devices_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/][%d] dcim_devices_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// DcimDeviceBayTemplatesUpdateReader is a Reader for the DcimDeviceBayTemplatesUpdate structure.
type DcimDeviceBayTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBayTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBayTemplatesUpdateOK creates a DcimDeviceBayTemplatesUpdateOK with default headers values
func NewDcimDeviceBayTemplatesUpdateOK() *DcimDeviceBayTemplatesUpdateOK {
	return &DcimDeviceBayTemplatesUpdateOK{}
}

/*
DcimDeviceBayTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesUpdateOK dcim device bay templates update o k
*/
type DcimDeviceBayTemplatesUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

// IsSuccess returns true when this dcim device bay templates update o k response has a 2xx status code
func (o *DcimDeviceBayTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates update o k response has a 3xx status code
func (o *DcimDeviceBayTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates update o k response has a 4xx status code
func (o *DcimDeviceBayTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates update o k response has a 5xx status code
func (o *DcimDeviceBayTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates update o k response a status code equal to that given
func (o *DcimDeviceBayTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim device bay templates update o k response
func (o *DcimDeviceBayTemplatesUpdateOK) Code() int {
	return 200
}

func (o *DcimDeviceBayTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesUpdateOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBayTemplatesUpdateDefault creates a DcimDeviceBayTemplatesUpdateDefault with default headers values
func NewDcimDeviceBayTemplatesUpdateDefault(code int) *DcimDeviceBayTemplatesUpdateDefault {
	return &DcimDeviceBayTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceBayTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceBayTemplatesUpdateDefault dcim device bay templates update default
*/
type DcimDeviceBayTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim device bay templates update default response has a 2xx status code
func (o *DcimDeviceBayTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device bay templates update default response has a 3xx status code
func (o *DcimDeviceBayTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device bay templates update default response has a 4xx status code
func (o *DcimDeviceBayTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device bay templates update default response has a 5xx status code
func (o *DcimDeviceBayTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device bay templates update default response a status code equal to that given
func (o *DcimDeviceBayTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim device bay templates update default response
func (o *DcimDeviceBayTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceBayTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/{id}/][%d] dcim_device-bay-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/{id}/][%d] dcim_device-bay-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

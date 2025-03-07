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

// DcimPowerOutletsTraceReader is a Reader for the DcimPowerOutletsTrace structure.
type DcimPowerOutletsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletsTraceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletsTraceOK creates a DcimPowerOutletsTraceOK with default headers values
func NewDcimPowerOutletsTraceOK() *DcimPowerOutletsTraceOK {
	return &DcimPowerOutletsTraceOK{}
}

/*
DcimPowerOutletsTraceOK describes a response with status code 200, with default header values.

DcimPowerOutletsTraceOK dcim power outlets trace o k
*/
type DcimPowerOutletsTraceOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets trace o k response has a 2xx status code
func (o *DcimPowerOutletsTraceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets trace o k response has a 3xx status code
func (o *DcimPowerOutletsTraceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets trace o k response has a 4xx status code
func (o *DcimPowerOutletsTraceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets trace o k response has a 5xx status code
func (o *DcimPowerOutletsTraceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets trace o k response a status code equal to that given
func (o *DcimPowerOutletsTraceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlets trace o k response
func (o *DcimPowerOutletsTraceOK) Code() int {
	return 200
}

func (o *DcimPowerOutletsTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/trace/][%d] dcimPowerOutletsTraceOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsTraceOK) String() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/trace/][%d] dcimPowerOutletsTraceOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsTraceOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletsTraceDefault creates a DcimPowerOutletsTraceDefault with default headers values
func NewDcimPowerOutletsTraceDefault(code int) *DcimPowerOutletsTraceDefault {
	return &DcimPowerOutletsTraceDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletsTraceDefault describes a response with status code -1, with default header values.

DcimPowerOutletsTraceDefault dcim power outlets trace default
*/
type DcimPowerOutletsTraceDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power outlets trace default response has a 2xx status code
func (o *DcimPowerOutletsTraceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlets trace default response has a 3xx status code
func (o *DcimPowerOutletsTraceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlets trace default response has a 4xx status code
func (o *DcimPowerOutletsTraceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlets trace default response has a 5xx status code
func (o *DcimPowerOutletsTraceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlets trace default response a status code equal to that given
func (o *DcimPowerOutletsTraceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power outlets trace default response
func (o *DcimPowerOutletsTraceDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletsTraceDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/trace/][%d] dcim_power-outlets_trace default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsTraceDefault) String() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/trace/][%d] dcim_power-outlets_trace default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsTraceDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletsTraceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

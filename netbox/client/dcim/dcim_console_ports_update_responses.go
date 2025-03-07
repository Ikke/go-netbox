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

// DcimConsolePortsUpdateReader is a Reader for the DcimConsolePortsUpdate structure.
type DcimConsolePortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortsUpdateOK creates a DcimConsolePortsUpdateOK with default headers values
func NewDcimConsolePortsUpdateOK() *DcimConsolePortsUpdateOK {
	return &DcimConsolePortsUpdateOK{}
}

/*
DcimConsolePortsUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortsUpdateOK dcim console ports update o k
*/
type DcimConsolePortsUpdateOK struct {
	Payload *models.ConsolePort
}

// IsSuccess returns true when this dcim console ports update o k response has a 2xx status code
func (o *DcimConsolePortsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console ports update o k response has a 3xx status code
func (o *DcimConsolePortsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console ports update o k response has a 4xx status code
func (o *DcimConsolePortsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console ports update o k response has a 5xx status code
func (o *DcimConsolePortsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console ports update o k response a status code equal to that given
func (o *DcimConsolePortsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console ports update o k response
func (o *DcimConsolePortsUpdateOK) Code() int {
	return 200
}

func (o *DcimConsolePortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/{id}/][%d] dcimConsolePortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/{id}/][%d] dcimConsolePortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortsUpdateOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortsUpdateDefault creates a DcimConsolePortsUpdateDefault with default headers values
func NewDcimConsolePortsUpdateDefault(code int) *DcimConsolePortsUpdateDefault {
	return &DcimConsolePortsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortsUpdateDefault describes a response with status code -1, with default header values.

DcimConsolePortsUpdateDefault dcim console ports update default
*/
type DcimConsolePortsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim console ports update default response has a 2xx status code
func (o *DcimConsolePortsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console ports update default response has a 3xx status code
func (o *DcimConsolePortsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console ports update default response has a 4xx status code
func (o *DcimConsolePortsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console ports update default response has a 5xx status code
func (o *DcimConsolePortsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console ports update default response a status code equal to that given
func (o *DcimConsolePortsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim console ports update default response
func (o *DcimConsolePortsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsolePortsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/{id}/][%d] dcim_console-ports_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/console-ports/{id}/][%d] dcim_console-ports_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

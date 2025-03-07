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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasCustomFieldsPartialUpdateReader is a Reader for the ExtrasCustomFieldsPartialUpdate structure.
type ExtrasCustomFieldsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldsPartialUpdateOK creates a ExtrasCustomFieldsPartialUpdateOK with default headers values
func NewExtrasCustomFieldsPartialUpdateOK() *ExtrasCustomFieldsPartialUpdateOK {
	return &ExtrasCustomFieldsPartialUpdateOK{}
}

/*
ExtrasCustomFieldsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsPartialUpdateOK extras custom fields partial update o k
*/
type ExtrasCustomFieldsPartialUpdateOK struct {
	Payload *models.CustomField
}

// IsSuccess returns true when this extras custom fields partial update o k response has a 2xx status code
func (o *ExtrasCustomFieldsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom fields partial update o k response has a 3xx status code
func (o *ExtrasCustomFieldsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom fields partial update o k response has a 4xx status code
func (o *ExtrasCustomFieldsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom fields partial update o k response has a 5xx status code
func (o *ExtrasCustomFieldsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom fields partial update o k response a status code equal to that given
func (o *ExtrasCustomFieldsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras custom fields partial update o k response
func (o *ExtrasCustomFieldsPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasCustomFieldsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-fields/{id}/][%d] extrasCustomFieldsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/custom-fields/{id}/][%d] extrasCustomFieldsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldsPartialUpdateOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldsPartialUpdateDefault creates a ExtrasCustomFieldsPartialUpdateDefault with default headers values
func NewExtrasCustomFieldsPartialUpdateDefault(code int) *ExtrasCustomFieldsPartialUpdateDefault {
	return &ExtrasCustomFieldsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldsPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasCustomFieldsPartialUpdateDefault extras custom fields partial update default
*/
type ExtrasCustomFieldsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras custom fields partial update default response has a 2xx status code
func (o *ExtrasCustomFieldsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom fields partial update default response has a 3xx status code
func (o *ExtrasCustomFieldsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom fields partial update default response has a 4xx status code
func (o *ExtrasCustomFieldsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom fields partial update default response has a 5xx status code
func (o *ExtrasCustomFieldsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom fields partial update default response a status code equal to that given
func (o *ExtrasCustomFieldsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras custom fields partial update default response
func (o *ExtrasCustomFieldsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomFieldsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-fields/{id}/][%d] extras_custom-fields_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/custom-fields/{id}/][%d] extras_custom-fields_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// ExtrasCustomFieldChoiceSetsBulkUpdateReader is a Reader for the ExtrasCustomFieldChoiceSetsBulkUpdate structure.
type ExtrasCustomFieldChoiceSetsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldChoiceSetsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldChoiceSetsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldChoiceSetsBulkUpdateOK creates a ExtrasCustomFieldChoiceSetsBulkUpdateOK with default headers values
func NewExtrasCustomFieldChoiceSetsBulkUpdateOK() *ExtrasCustomFieldChoiceSetsBulkUpdateOK {
	return &ExtrasCustomFieldChoiceSetsBulkUpdateOK{}
}

/*
ExtrasCustomFieldChoiceSetsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldChoiceSetsBulkUpdateOK extras custom field choice sets bulk update o k
*/
type ExtrasCustomFieldChoiceSetsBulkUpdateOK struct {
	Payload *models.CustomFieldChoiceSet
}

// IsSuccess returns true when this extras custom field choice sets bulk update o k response has a 2xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom field choice sets bulk update o k response has a 3xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom field choice sets bulk update o k response has a 4xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom field choice sets bulk update o k response has a 5xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom field choice sets bulk update o k response a status code equal to that given
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras custom field choice sets bulk update o k response
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) Code() int {
	return 200
}

func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choice-sets/][%d] extrasCustomFieldChoiceSetsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choice-sets/][%d] extrasCustomFieldChoiceSetsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) GetPayload() *models.CustomFieldChoiceSet {
	return o.Payload
}

func (o *ExtrasCustomFieldChoiceSetsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomFieldChoiceSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldChoiceSetsBulkUpdateDefault creates a ExtrasCustomFieldChoiceSetsBulkUpdateDefault with default headers values
func NewExtrasCustomFieldChoiceSetsBulkUpdateDefault(code int) *ExtrasCustomFieldChoiceSetsBulkUpdateDefault {
	return &ExtrasCustomFieldChoiceSetsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldChoiceSetsBulkUpdateDefault describes a response with status code -1, with default header values.

ExtrasCustomFieldChoiceSetsBulkUpdateDefault extras custom field choice sets bulk update default
*/
type ExtrasCustomFieldChoiceSetsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras custom field choice sets bulk update default response has a 2xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom field choice sets bulk update default response has a 3xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom field choice sets bulk update default response has a 4xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom field choice sets bulk update default response has a 5xx status code
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom field choice sets bulk update default response a status code equal to that given
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras custom field choice sets bulk update default response
func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choice-sets/][%d] extras_custom-field-choice-sets_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choice-sets/][%d] extras_custom-field-choice-sets_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldChoiceSetsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

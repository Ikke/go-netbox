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
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimRackRolesListReader is a Reader for the DcimRackRolesList structure.
type DcimRackRolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackRolesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackRolesListOK creates a DcimRackRolesListOK with default headers values
func NewDcimRackRolesListOK() *DcimRackRolesListOK {
	return &DcimRackRolesListOK{}
}

/*
DcimRackRolesListOK describes a response with status code 200, with default header values.

DcimRackRolesListOK dcim rack roles list o k
*/
type DcimRackRolesListOK struct {
	Payload *DcimRackRolesListOKBody
}

// IsSuccess returns true when this dcim rack roles list o k response has a 2xx status code
func (o *DcimRackRolesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack roles list o k response has a 3xx status code
func (o *DcimRackRolesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles list o k response has a 4xx status code
func (o *DcimRackRolesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack roles list o k response has a 5xx status code
func (o *DcimRackRolesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles list o k response a status code equal to that given
func (o *DcimRackRolesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rack roles list o k response
func (o *DcimRackRolesListOK) Code() int {
	return 200
}

func (o *DcimRackRolesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/][%d] dcimRackRolesListOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesListOK) String() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/][%d] dcimRackRolesListOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesListOK) GetPayload() *DcimRackRolesListOKBody {
	return o.Payload
}

func (o *DcimRackRolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimRackRolesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesListDefault creates a DcimRackRolesListDefault with default headers values
func NewDcimRackRolesListDefault(code int) *DcimRackRolesListDefault {
	return &DcimRackRolesListDefault{
		_statusCode: code,
	}
}

/*
DcimRackRolesListDefault describes a response with status code -1, with default header values.

DcimRackRolesListDefault dcim rack roles list default
*/
type DcimRackRolesListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rack roles list default response has a 2xx status code
func (o *DcimRackRolesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rack roles list default response has a 3xx status code
func (o *DcimRackRolesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rack roles list default response has a 4xx status code
func (o *DcimRackRolesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rack roles list default response has a 5xx status code
func (o *DcimRackRolesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rack roles list default response a status code equal to that given
func (o *DcimRackRolesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rack roles list default response
func (o *DcimRackRolesListDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackRolesListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/][%d] dcim_rack-roles_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/rack-roles/][%d] dcim_rack-roles_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimRackRolesListOKBody dcim rack roles list o k body
swagger:model DcimRackRolesListOKBody
*/
type DcimRackRolesListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.RackRole `json:"results"`
}

// Validate validates this dcim rack roles list o k body
func (o *DcimRackRolesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimRackRolesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimRackRolesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimRackRolesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimRackRolesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimRackRolesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimRackRolesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimRackRolesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimRackRolesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimRackRolesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimRackRolesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim rack roles list o k body based on the context it is used
func (o *DcimRackRolesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimRackRolesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimRackRolesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimRackRolesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimRackRolesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimRackRolesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimRackRolesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

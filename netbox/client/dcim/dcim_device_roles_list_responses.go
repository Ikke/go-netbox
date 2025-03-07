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

// DcimDeviceRolesListReader is a Reader for the DcimDeviceRolesList structure.
type DcimDeviceRolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceRolesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesListOK creates a DcimDeviceRolesListOK with default headers values
func NewDcimDeviceRolesListOK() *DcimDeviceRolesListOK {
	return &DcimDeviceRolesListOK{}
}

/*
DcimDeviceRolesListOK describes a response with status code 200, with default header values.

DcimDeviceRolesListOK dcim device roles list o k
*/
type DcimDeviceRolesListOK struct {
	Payload *DcimDeviceRolesListOKBody
}

// IsSuccess returns true when this dcim device roles list o k response has a 2xx status code
func (o *DcimDeviceRolesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device roles list o k response has a 3xx status code
func (o *DcimDeviceRolesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device roles list o k response has a 4xx status code
func (o *DcimDeviceRolesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device roles list o k response has a 5xx status code
func (o *DcimDeviceRolesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device roles list o k response a status code equal to that given
func (o *DcimDeviceRolesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim device roles list o k response
func (o *DcimDeviceRolesListOK) Code() int {
	return 200
}

func (o *DcimDeviceRolesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-roles/][%d] dcimDeviceRolesListOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceRolesListOK) String() string {
	return fmt.Sprintf("[GET /dcim/device-roles/][%d] dcimDeviceRolesListOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceRolesListOK) GetPayload() *DcimDeviceRolesListOKBody {
	return o.Payload
}

func (o *DcimDeviceRolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimDeviceRolesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceRolesListDefault creates a DcimDeviceRolesListDefault with default headers values
func NewDcimDeviceRolesListDefault(code int) *DcimDeviceRolesListDefault {
	return &DcimDeviceRolesListDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceRolesListDefault describes a response with status code -1, with default header values.

DcimDeviceRolesListDefault dcim device roles list default
*/
type DcimDeviceRolesListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim device roles list default response has a 2xx status code
func (o *DcimDeviceRolesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device roles list default response has a 3xx status code
func (o *DcimDeviceRolesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device roles list default response has a 4xx status code
func (o *DcimDeviceRolesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device roles list default response has a 5xx status code
func (o *DcimDeviceRolesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device roles list default response a status code equal to that given
func (o *DcimDeviceRolesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim device roles list default response
func (o *DcimDeviceRolesListDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceRolesListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/device-roles/][%d] dcim_device-roles_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/device-roles/][%d] dcim_device-roles_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimDeviceRolesListOKBody dcim device roles list o k body
swagger:model DcimDeviceRolesListOKBody
*/
type DcimDeviceRolesListOKBody struct {

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
	Results []*models.DeviceRole `json:"results"`
}

// Validate validates this dcim device roles list o k body
func (o *DcimDeviceRolesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimDeviceRolesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimDeviceRolesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimDeviceRolesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimDeviceRolesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimDeviceRolesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimDeviceRolesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimDeviceRolesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimDeviceRolesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimDeviceRolesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimDeviceRolesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim device roles list o k body based on the context it is used
func (o *DcimDeviceRolesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimDeviceRolesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimDeviceRolesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimDeviceRolesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimDeviceRolesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimDeviceRolesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimDeviceRolesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

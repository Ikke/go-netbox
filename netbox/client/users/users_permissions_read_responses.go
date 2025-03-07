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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// UsersPermissionsReadReader is a Reader for the UsersPermissionsRead structure.
type UsersPermissionsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersPermissionsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersPermissionsReadOK creates a UsersPermissionsReadOK with default headers values
func NewUsersPermissionsReadOK() *UsersPermissionsReadOK {
	return &UsersPermissionsReadOK{}
}

/*
UsersPermissionsReadOK describes a response with status code 200, with default header values.

UsersPermissionsReadOK users permissions read o k
*/
type UsersPermissionsReadOK struct {
	Payload *models.ObjectPermission
}

// IsSuccess returns true when this users permissions read o k response has a 2xx status code
func (o *UsersPermissionsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users permissions read o k response has a 3xx status code
func (o *UsersPermissionsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users permissions read o k response has a 4xx status code
func (o *UsersPermissionsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users permissions read o k response has a 5xx status code
func (o *UsersPermissionsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users permissions read o k response a status code equal to that given
func (o *UsersPermissionsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users permissions read o k response
func (o *UsersPermissionsReadOK) Code() int {
	return 200
}

func (o *UsersPermissionsReadOK) Error() string {
	return fmt.Sprintf("[GET /users/permissions/{id}/][%d] usersPermissionsReadOK  %+v", 200, o.Payload)
}

func (o *UsersPermissionsReadOK) String() string {
	return fmt.Sprintf("[GET /users/permissions/{id}/][%d] usersPermissionsReadOK  %+v", 200, o.Payload)
}

func (o *UsersPermissionsReadOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersPermissionsReadDefault creates a UsersPermissionsReadDefault with default headers values
func NewUsersPermissionsReadDefault(code int) *UsersPermissionsReadDefault {
	return &UsersPermissionsReadDefault{
		_statusCode: code,
	}
}

/*
UsersPermissionsReadDefault describes a response with status code -1, with default header values.

UsersPermissionsReadDefault users permissions read default
*/
type UsersPermissionsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this users permissions read default response has a 2xx status code
func (o *UsersPermissionsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users permissions read default response has a 3xx status code
func (o *UsersPermissionsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users permissions read default response has a 4xx status code
func (o *UsersPermissionsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users permissions read default response has a 5xx status code
func (o *UsersPermissionsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users permissions read default response a status code equal to that given
func (o *UsersPermissionsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the users permissions read default response
func (o *UsersPermissionsReadDefault) Code() int {
	return o._statusCode
}

func (o *UsersPermissionsReadDefault) Error() string {
	return fmt.Sprintf("[GET /users/permissions/{id}/][%d] users_permissions_read default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsReadDefault) String() string {
	return fmt.Sprintf("[GET /users/permissions/{id}/][%d] users_permissions_read default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersPermissionsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

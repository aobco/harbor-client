// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/goharbor-client/v5/apiv2/model"
)

// SetUserSysAdminReader is a Reader for the SetUserSysAdmin structure.
type SetUserSysAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetUserSysAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetUserSysAdminOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetUserSysAdminUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetUserSysAdminForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetUserSysAdminNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetUserSysAdminInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetUserSysAdminOK creates a SetUserSysAdminOK with default headers values
func NewSetUserSysAdminOK() *SetUserSysAdminOK {
	return &SetUserSysAdminOK{}
}

/*SetUserSysAdminOK handles this case with default header values.

Success
*/
type SetUserSysAdminOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *SetUserSysAdminOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/sysadmin][%d] setUserSysAdminOK ", 200)
}

func (o *SetUserSysAdminOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewSetUserSysAdminUnauthorized creates a SetUserSysAdminUnauthorized with default headers values
func NewSetUserSysAdminUnauthorized() *SetUserSysAdminUnauthorized {
	return &SetUserSysAdminUnauthorized{}
}

/*SetUserSysAdminUnauthorized handles this case with default header values.

Unauthorized
*/
type SetUserSysAdminUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetUserSysAdminUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/sysadmin][%d] setUserSysAdminUnauthorized  %+v", 401, o.Payload)
}

func (o *SetUserSysAdminUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetUserSysAdminUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetUserSysAdminForbidden creates a SetUserSysAdminForbidden with default headers values
func NewSetUserSysAdminForbidden() *SetUserSysAdminForbidden {
	return &SetUserSysAdminForbidden{}
}

/*SetUserSysAdminForbidden handles this case with default header values.

Forbidden
*/
type SetUserSysAdminForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetUserSysAdminForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/sysadmin][%d] setUserSysAdminForbidden  %+v", 403, o.Payload)
}

func (o *SetUserSysAdminForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetUserSysAdminForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetUserSysAdminNotFound creates a SetUserSysAdminNotFound with default headers values
func NewSetUserSysAdminNotFound() *SetUserSysAdminNotFound {
	return &SetUserSysAdminNotFound{}
}

/*SetUserSysAdminNotFound handles this case with default header values.

Not found
*/
type SetUserSysAdminNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetUserSysAdminNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/sysadmin][%d] setUserSysAdminNotFound  %+v", 404, o.Payload)
}

func (o *SetUserSysAdminNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetUserSysAdminNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetUserSysAdminInternalServerError creates a SetUserSysAdminInternalServerError with default headers values
func NewSetUserSysAdminInternalServerError() *SetUserSysAdminInternalServerError {
	return &SetUserSysAdminInternalServerError{}
}

/*SetUserSysAdminInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type SetUserSysAdminInternalServerError struct {
}

func (o *SetUserSysAdminInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/sysadmin][%d] setUserSysAdminInternalServerError ", 500)
}

func (o *SetUserSysAdminInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

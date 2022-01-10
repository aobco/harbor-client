// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/harbor-client/apiv1/model"
)

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*GetUsersOK handles this case with default header values.

Searched for users of Harbor successfully.
*/
type GetUsersOK struct {
	Payload []*model.User
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) GetPayload() []*model.User {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersBadRequest creates a GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {
	return &GetUsersBadRequest{}
}

/*GetUsersBadRequest handles this case with default header values.

Invalid user ID.
*/
type GetUsersBadRequest struct {
}

func (o *GetUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersBadRequest ", 400)
}

func (o *GetUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersUnauthorized creates a GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

/*GetUsersUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetUsersUnauthorized struct {
}

func (o *GetUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersUnauthorized ", 401)
}

func (o *GetUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForbidden creates a GetUsersForbidden with default headers values
func NewGetUsersForbidden() *GetUsersForbidden {
	return &GetUsersForbidden{}
}

/*GetUsersForbidden handles this case with default header values.

User does not have permission of admin role.
*/
type GetUsersForbidden struct {
}

func (o *GetUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersForbidden ", 403)
}

func (o *GetUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersInternalServerError creates a GetUsersInternalServerError with default headers values
func NewGetUsersInternalServerError() *GetUsersInternalServerError {
	return &GetUsersInternalServerError{}
}

/*GetUsersInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetUsersInternalServerError struct {
}

func (o *GetUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersInternalServerError ", 500)
}

func (o *GetUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

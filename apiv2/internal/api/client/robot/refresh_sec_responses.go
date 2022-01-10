// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/harbor-client/v5/apiv2/model"
)

// RefreshSecReader is a Reader for the RefreshSec structure.
type RefreshSecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshSecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshSecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRefreshSecBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRefreshSecUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRefreshSecForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRefreshSecNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRefreshSecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRefreshSecOK creates a RefreshSecOK with default headers values
func NewRefreshSecOK() *RefreshSecOK {
	return &RefreshSecOK{}
}

/*RefreshSecOK handles this case with default header values.

Return refreshed robot sec.
*/
type RefreshSecOK struct {
	Payload *model.RobotSec
}

func (o *RefreshSecOK) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecOK  %+v", 200, o.Payload)
}

func (o *RefreshSecOK) GetPayload() *model.RobotSec {
	return o.Payload
}

func (o *RefreshSecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RobotSec)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshSecBadRequest creates a RefreshSecBadRequest with default headers values
func NewRefreshSecBadRequest() *RefreshSecBadRequest {
	return &RefreshSecBadRequest{}
}

/*RefreshSecBadRequest handles this case with default header values.

Bad request
*/
type RefreshSecBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *RefreshSecBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecBadRequest  %+v", 400, o.Payload)
}

func (o *RefreshSecBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *RefreshSecBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshSecUnauthorized creates a RefreshSecUnauthorized with default headers values
func NewRefreshSecUnauthorized() *RefreshSecUnauthorized {
	return &RefreshSecUnauthorized{}
}

/*RefreshSecUnauthorized handles this case with default header values.

Unauthorized
*/
type RefreshSecUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *RefreshSecUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecUnauthorized  %+v", 401, o.Payload)
}

func (o *RefreshSecUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *RefreshSecUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshSecForbidden creates a RefreshSecForbidden with default headers values
func NewRefreshSecForbidden() *RefreshSecForbidden {
	return &RefreshSecForbidden{}
}

/*RefreshSecForbidden handles this case with default header values.

Forbidden
*/
type RefreshSecForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *RefreshSecForbidden) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecForbidden  %+v", 403, o.Payload)
}

func (o *RefreshSecForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *RefreshSecForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshSecNotFound creates a RefreshSecNotFound with default headers values
func NewRefreshSecNotFound() *RefreshSecNotFound {
	return &RefreshSecNotFound{}
}

/*RefreshSecNotFound handles this case with default header values.

Not found
*/
type RefreshSecNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *RefreshSecNotFound) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecNotFound  %+v", 404, o.Payload)
}

func (o *RefreshSecNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *RefreshSecNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshSecInternalServerError creates a RefreshSecInternalServerError with default headers values
func NewRefreshSecInternalServerError() *RefreshSecInternalServerError {
	return &RefreshSecInternalServerError{}
}

/*RefreshSecInternalServerError handles this case with default header values.

Internal server error
*/
type RefreshSecInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *RefreshSecInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecInternalServerError  %+v", 500, o.Payload)
}

func (o *RefreshSecInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *RefreshSecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

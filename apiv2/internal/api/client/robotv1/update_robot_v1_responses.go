// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/harbor-client/v5/apiv2/model"
)

// UpdateRobotV1Reader is a Reader for the UpdateRobotV1 structure.
type UpdateRobotV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRobotV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRobotV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRobotV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRobotV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRobotV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRobotV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRobotV1Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRobotV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRobotV1OK creates a UpdateRobotV1OK with default headers values
func NewUpdateRobotV1OK() *UpdateRobotV1OK {
	return &UpdateRobotV1OK{}
}

/*UpdateRobotV1OK handles this case with default header values.

Success
*/
type UpdateRobotV1OK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateRobotV1OK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/robots/{robot_id}][%d] updateRobotV1OK ", 200)
}

func (o *UpdateRobotV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateRobotV1BadRequest creates a UpdateRobotV1BadRequest with default headers values
func NewUpdateRobotV1BadRequest() *UpdateRobotV1BadRequest {
	return &UpdateRobotV1BadRequest{}
}

/*UpdateRobotV1BadRequest handles this case with default header values.

Bad request
*/
type UpdateRobotV1BadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateRobotV1BadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/robots/{robot_id}][%d] updateRobotV1BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRobotV1BadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateRobotV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRobotV1Unauthorized creates a UpdateRobotV1Unauthorized with default headers values
func NewUpdateRobotV1Unauthorized() *UpdateRobotV1Unauthorized {
	return &UpdateRobotV1Unauthorized{}
}

/*UpdateRobotV1Unauthorized handles this case with default header values.

Unauthorized
*/
type UpdateRobotV1Unauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateRobotV1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/robots/{robot_id}][%d] updateRobotV1Unauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRobotV1Unauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateRobotV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRobotV1Forbidden creates a UpdateRobotV1Forbidden with default headers values
func NewUpdateRobotV1Forbidden() *UpdateRobotV1Forbidden {
	return &UpdateRobotV1Forbidden{}
}

/*UpdateRobotV1Forbidden handles this case with default header values.

Forbidden
*/
type UpdateRobotV1Forbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateRobotV1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/robots/{robot_id}][%d] updateRobotV1Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateRobotV1Forbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateRobotV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRobotV1NotFound creates a UpdateRobotV1NotFound with default headers values
func NewUpdateRobotV1NotFound() *UpdateRobotV1NotFound {
	return &UpdateRobotV1NotFound{}
}

/*UpdateRobotV1NotFound handles this case with default header values.

Not found
*/
type UpdateRobotV1NotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateRobotV1NotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/robots/{robot_id}][%d] updateRobotV1NotFound  %+v", 404, o.Payload)
}

func (o *UpdateRobotV1NotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateRobotV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRobotV1Conflict creates a UpdateRobotV1Conflict with default headers values
func NewUpdateRobotV1Conflict() *UpdateRobotV1Conflict {
	return &UpdateRobotV1Conflict{}
}

/*UpdateRobotV1Conflict handles this case with default header values.

Conflict
*/
type UpdateRobotV1Conflict struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateRobotV1Conflict) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/robots/{robot_id}][%d] updateRobotV1Conflict  %+v", 409, o.Payload)
}

func (o *UpdateRobotV1Conflict) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateRobotV1Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRobotV1InternalServerError creates a UpdateRobotV1InternalServerError with default headers values
func NewUpdateRobotV1InternalServerError() *UpdateRobotV1InternalServerError {
	return &UpdateRobotV1InternalServerError{}
}

/*UpdateRobotV1InternalServerError handles this case with default header values.

Internal server error
*/
type UpdateRobotV1InternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateRobotV1InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/robots/{robot_id}][%d] updateRobotV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRobotV1InternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateRobotV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

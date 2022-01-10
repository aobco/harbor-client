// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/harbor-client/apiv2/model"
)

// CreateTagReader is a Reader for the CreateTag structure.
type CreateTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTagCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateTagMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateTagConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTagCreated creates a CreateTagCreated with default headers values
func NewCreateTagCreated() *CreateTagCreated {
	return &CreateTagCreated{}
}

/*CreateTagCreated handles this case with default header values.

Created
*/
type CreateTagCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreateTagCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagCreated ", 201)
}

func (o *CreateTagCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewCreateTagBadRequest creates a CreateTagBadRequest with default headers values
func NewCreateTagBadRequest() *CreateTagBadRequest {
	return &CreateTagBadRequest{}
}

/*CreateTagBadRequest handles this case with default header values.

Bad request
*/
type CreateTagBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateTagBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTagBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagUnauthorized creates a CreateTagUnauthorized with default headers values
func NewCreateTagUnauthorized() *CreateTagUnauthorized {
	return &CreateTagUnauthorized{}
}

/*CreateTagUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateTagUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateTagUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateTagUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagForbidden creates a CreateTagForbidden with default headers values
func NewCreateTagForbidden() *CreateTagForbidden {
	return &CreateTagForbidden{}
}

/*CreateTagForbidden handles this case with default header values.

Forbidden
*/
type CreateTagForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateTagForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagForbidden  %+v", 403, o.Payload)
}

func (o *CreateTagForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagNotFound creates a CreateTagNotFound with default headers values
func NewCreateTagNotFound() *CreateTagNotFound {
	return &CreateTagNotFound{}
}

/*CreateTagNotFound handles this case with default header values.

Not found
*/
type CreateTagNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateTagNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagNotFound  %+v", 404, o.Payload)
}

func (o *CreateTagNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagMethodNotAllowed creates a CreateTagMethodNotAllowed with default headers values
func NewCreateTagMethodNotAllowed() *CreateTagMethodNotAllowed {
	return &CreateTagMethodNotAllowed{}
}

/*CreateTagMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type CreateTagMethodNotAllowed struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateTagMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CreateTagMethodNotAllowed) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateTagMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagConflict creates a CreateTagConflict with default headers values
func NewCreateTagConflict() *CreateTagConflict {
	return &CreateTagConflict{}
}

/*CreateTagConflict handles this case with default header values.

Conflict
*/
type CreateTagConflict struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateTagConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagConflict  %+v", 409, o.Payload)
}

func (o *CreateTagConflict) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateTagConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagInternalServerError creates a CreateTagInternalServerError with default headers values
func NewCreateTagInternalServerError() *CreateTagInternalServerError {
	return &CreateTagInternalServerError{}
}

/*CreateTagInternalServerError handles this case with default header values.

Internal server error
*/
type CreateTagInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateTagInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] createTagInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTagInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

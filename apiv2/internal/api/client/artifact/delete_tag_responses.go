// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/goharbor-client/v5/apiv2/model"
)

// DeleteTagReader is a Reader for the DeleteTag structure.
type DeleteTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTagOK creates a DeleteTagOK with default headers values
func NewDeleteTagOK() *DeleteTagOK {
	return &DeleteTagOK{}
}

/*DeleteTagOK handles this case with default header values.

Success
*/
type DeleteTagOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteTagOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagOK ", 200)
}

func (o *DeleteTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewDeleteTagUnauthorized creates a DeleteTagUnauthorized with default headers values
func NewDeleteTagUnauthorized() *DeleteTagUnauthorized {
	return &DeleteTagUnauthorized{}
}

/*DeleteTagUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteTagUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteTagUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTagUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagForbidden creates a DeleteTagForbidden with default headers values
func NewDeleteTagForbidden() *DeleteTagForbidden {
	return &DeleteTagForbidden{}
}

/*DeleteTagForbidden handles this case with default header values.

Forbidden
*/
type DeleteTagForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteTagForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTagForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagNotFound creates a DeleteTagNotFound with default headers values
func NewDeleteTagNotFound() *DeleteTagNotFound {
	return &DeleteTagNotFound{}
}

/*DeleteTagNotFound handles this case with default header values.

Not found
*/
type DeleteTagNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteTagNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTagNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTagInternalServerError creates a DeleteTagInternalServerError with default headers values
func NewDeleteTagInternalServerError() *DeleteTagInternalServerError {
	return &DeleteTagInternalServerError{}
}

/*DeleteTagInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteTagInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteTagInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags/{tag_name}][%d] deleteTagInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTagInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

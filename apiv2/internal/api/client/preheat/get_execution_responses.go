// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/harbor-client/apiv2/model"
)

// GetExecutionReader is a Reader for the GetExecution structure.
type GetExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExecutionOK creates a GetExecutionOK with default headers values
func NewGetExecutionOK() *GetExecutionOK {
	return &GetExecutionOK{}
}

/*GetExecutionOK handles this case with default header values.

Get execution success
*/
type GetExecutionOK struct {
	Payload *model.Execution
}

func (o *GetExecutionOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] getExecutionOK  %+v", 200, o.Payload)
}

func (o *GetExecutionOK) GetPayload() *model.Execution {
	return o.Payload
}

func (o *GetExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Execution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutionBadRequest creates a GetExecutionBadRequest with default headers values
func NewGetExecutionBadRequest() *GetExecutionBadRequest {
	return &GetExecutionBadRequest{}
}

/*GetExecutionBadRequest handles this case with default header values.

Bad request
*/
type GetExecutionBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetExecutionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] getExecutionBadRequest  %+v", 400, o.Payload)
}

func (o *GetExecutionBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutionUnauthorized creates a GetExecutionUnauthorized with default headers values
func NewGetExecutionUnauthorized() *GetExecutionUnauthorized {
	return &GetExecutionUnauthorized{}
}

/*GetExecutionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetExecutionUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] getExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExecutionUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutionForbidden creates a GetExecutionForbidden with default headers values
func NewGetExecutionForbidden() *GetExecutionForbidden {
	return &GetExecutionForbidden{}
}

/*GetExecutionForbidden handles this case with default header values.

Forbidden
*/
type GetExecutionForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetExecutionForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] getExecutionForbidden  %+v", 403, o.Payload)
}

func (o *GetExecutionForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutionNotFound creates a GetExecutionNotFound with default headers values
func NewGetExecutionNotFound() *GetExecutionNotFound {
	return &GetExecutionNotFound{}
}

/*GetExecutionNotFound handles this case with default header values.

Not found
*/
type GetExecutionNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetExecutionNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] getExecutionNotFound  %+v", 404, o.Payload)
}

func (o *GetExecutionNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutionInternalServerError creates a GetExecutionInternalServerError with default headers values
func NewGetExecutionInternalServerError() *GetExecutionInternalServerError {
	return &GetExecutionInternalServerError{}
}

/*GetExecutionInternalServerError handles this case with default header values.

Internal server error
*/
type GetExecutionInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] getExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExecutionInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

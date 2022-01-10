// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/goharbor-client/v5/apiv2/model"
)

// GetProjectMetadataReader is a Reader for the GetProjectMetadata structure.
type GetProjectMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectMetadataOK creates a GetProjectMetadataOK with default headers values
func NewGetProjectMetadataOK() *GetProjectMetadataOK {
	return &GetProjectMetadataOK{}
}

/*GetProjectMetadataOK handles this case with default header values.

Success
*/
type GetProjectMetadataOK struct {
	Payload map[string]string
}

func (o *GetProjectMetadataOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/{meta_name}][%d] getProjectMetadataOK  %+v", 200, o.Payload)
}

func (o *GetProjectMetadataOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *GetProjectMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectMetadataBadRequest creates a GetProjectMetadataBadRequest with default headers values
func NewGetProjectMetadataBadRequest() *GetProjectMetadataBadRequest {
	return &GetProjectMetadataBadRequest{}
}

/*GetProjectMetadataBadRequest handles this case with default header values.

Bad request
*/
type GetProjectMetadataBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectMetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/{meta_name}][%d] getProjectMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectMetadataBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectMetadataUnauthorized creates a GetProjectMetadataUnauthorized with default headers values
func NewGetProjectMetadataUnauthorized() *GetProjectMetadataUnauthorized {
	return &GetProjectMetadataUnauthorized{}
}

/*GetProjectMetadataUnauthorized handles this case with default header values.

Unauthorized
*/
type GetProjectMetadataUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/{meta_name}][%d] getProjectMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProjectMetadataUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectMetadataForbidden creates a GetProjectMetadataForbidden with default headers values
func NewGetProjectMetadataForbidden() *GetProjectMetadataForbidden {
	return &GetProjectMetadataForbidden{}
}

/*GetProjectMetadataForbidden handles this case with default header values.

Forbidden
*/
type GetProjectMetadataForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/{meta_name}][%d] getProjectMetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectMetadataForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectMetadataNotFound creates a GetProjectMetadataNotFound with default headers values
func NewGetProjectMetadataNotFound() *GetProjectMetadataNotFound {
	return &GetProjectMetadataNotFound{}
}

/*GetProjectMetadataNotFound handles this case with default header values.

Not found
*/
type GetProjectMetadataNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/{meta_name}][%d] getProjectMetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectMetadataNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectMetadataInternalServerError creates a GetProjectMetadataInternalServerError with default headers values
func NewGetProjectMetadataInternalServerError() *GetProjectMetadataInternalServerError {
	return &GetProjectMetadataInternalServerError{}
}

/*GetProjectMetadataInternalServerError handles this case with default header values.

Internal server error
*/
type GetProjectMetadataInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/{meta_name}][%d] getProjectMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectMetadataInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

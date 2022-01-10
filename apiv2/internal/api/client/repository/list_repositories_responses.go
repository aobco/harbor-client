// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/aobco/goharbor-client/v5/apiv2/model"
)

// ListRepositoriesReader is a Reader for the ListRepositories structure.
type ListRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListRepositoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListRepositoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRepositoriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRepositoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRepositoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRepositoriesOK creates a ListRepositoriesOK with default headers values
func NewListRepositoriesOK() *ListRepositoriesOK {
	return &ListRepositoriesOK{}
}

/*ListRepositoriesOK handles this case with default header values.

Success
*/
type ListRepositoriesOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of repositories
	 */
	XTotalCount int64

	Payload []*model.Repository
}

func (o *ListRepositoriesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesOK  %+v", 200, o.Payload)
}

func (o *ListRepositoriesOK) GetPayload() []*model.Repository {
	return o.Payload
}

func (o *ListRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepositoriesBadRequest creates a ListRepositoriesBadRequest with default headers values
func NewListRepositoriesBadRequest() *ListRepositoriesBadRequest {
	return &ListRepositoriesBadRequest{}
}

/*ListRepositoriesBadRequest handles this case with default header values.

Bad request
*/
type ListRepositoriesBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRepositoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesBadRequest  %+v", 400, o.Payload)
}

func (o *ListRepositoriesBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRepositoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepositoriesUnauthorized creates a ListRepositoriesUnauthorized with default headers values
func NewListRepositoriesUnauthorized() *ListRepositoriesUnauthorized {
	return &ListRepositoriesUnauthorized{}
}

/*ListRepositoriesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListRepositoriesUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRepositoriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRepositoriesUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRepositoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepositoriesForbidden creates a ListRepositoriesForbidden with default headers values
func NewListRepositoriesForbidden() *ListRepositoriesForbidden {
	return &ListRepositoriesForbidden{}
}

/*ListRepositoriesForbidden handles this case with default header values.

Forbidden
*/
type ListRepositoriesForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRepositoriesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesForbidden  %+v", 403, o.Payload)
}

func (o *ListRepositoriesForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRepositoriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepositoriesNotFound creates a ListRepositoriesNotFound with default headers values
func NewListRepositoriesNotFound() *ListRepositoriesNotFound {
	return &ListRepositoriesNotFound{}
}

/*ListRepositoriesNotFound handles this case with default header values.

Not found
*/
type ListRepositoriesNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRepositoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesNotFound  %+v", 404, o.Payload)
}

func (o *ListRepositoriesNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRepositoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepositoriesInternalServerError creates a ListRepositoriesInternalServerError with default headers values
func NewListRepositoriesInternalServerError() *ListRepositoriesInternalServerError {
	return &ListRepositoriesInternalServerError{}
}

/*ListRepositoriesInternalServerError handles this case with default header values.

Internal server error
*/
type ListRepositoriesInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRepositoriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRepositoriesInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRepositoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

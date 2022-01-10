// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/harbor-client/v5/apiv2/model"
)

// HeadProjectReader is a Reader for the HeadProject structure.
type HeadProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewHeadProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHeadProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHeadProjectOK creates a HeadProjectOK with default headers values
func NewHeadProjectOK() *HeadProjectOK {
	return &HeadProjectOK{}
}

/*HeadProjectOK handles this case with default header values.

Success
*/
type HeadProjectOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *HeadProjectOK) Error() string {
	return fmt.Sprintf("[HEAD /projects][%d] headProjectOK ", 200)
}

func (o *HeadProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewHeadProjectNotFound creates a HeadProjectNotFound with default headers values
func NewHeadProjectNotFound() *HeadProjectNotFound {
	return &HeadProjectNotFound{}
}

/*HeadProjectNotFound handles this case with default header values.

Not found
*/
type HeadProjectNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *HeadProjectNotFound) Error() string {
	return fmt.Sprintf("[HEAD /projects][%d] headProjectNotFound  %+v", 404, o.Payload)
}

func (o *HeadProjectNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *HeadProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadProjectInternalServerError creates a HeadProjectInternalServerError with default headers values
func NewHeadProjectInternalServerError() *HeadProjectInternalServerError {
	return &HeadProjectInternalServerError{}
}

/*HeadProjectInternalServerError handles this case with default header values.

Internal server error
*/
type HeadProjectInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *HeadProjectInternalServerError) Error() string {
	return fmt.Sprintf("[HEAD /projects][%d] headProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *HeadProjectInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *HeadProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

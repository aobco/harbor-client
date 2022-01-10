// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/aobco/harbor-client/apiv2/model"
)

// GetGCHistoryReader is a Reader for the GetGCHistory structure.
type GetGCHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGCHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGCHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGCHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGCHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGCHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGCHistoryOK creates a GetGCHistoryOK with default headers values
func NewGetGCHistoryOK() *GetGCHistoryOK {
	return &GetGCHistoryOK{}
}

/*GetGCHistoryOK handles this case with default header values.

Get gc results successfully.
*/
type GetGCHistoryOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of history
	 */
	XTotalCount int64

	Payload []*model.GCHistory
}

func (o *GetGCHistoryOK) Error() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryOK  %+v", 200, o.Payload)
}

func (o *GetGCHistoryOK) GetPayload() []*model.GCHistory {
	return o.Payload
}

func (o *GetGCHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetGCHistoryUnauthorized creates a GetGCHistoryUnauthorized with default headers values
func NewGetGCHistoryUnauthorized() *GetGCHistoryUnauthorized {
	return &GetGCHistoryUnauthorized{}
}

/*GetGCHistoryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetGCHistoryUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGCHistoryUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCHistoryForbidden creates a GetGCHistoryForbidden with default headers values
func NewGetGCHistoryForbidden() *GetGCHistoryForbidden {
	return &GetGCHistoryForbidden{}
}

/*GetGCHistoryForbidden handles this case with default header values.

Forbidden
*/
type GetGCHistoryForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryForbidden  %+v", 403, o.Payload)
}

func (o *GetGCHistoryForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCHistoryInternalServerError creates a GetGCHistoryInternalServerError with default headers values
func NewGetGCHistoryInternalServerError() *GetGCHistoryInternalServerError {
	return &GetGCHistoryInternalServerError{}
}

/*GetGCHistoryInternalServerError handles this case with default header values.

Internal server error
*/
type GetGCHistoryInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGCHistoryInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

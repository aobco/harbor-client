// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/harbor-client/apiv2/model"
)

// StopScanAllReader is a Reader for the StopScanAll structure.
type StopScanAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopScanAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStopScanAllAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopScanAllBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStopScanAllUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopScanAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopScanAllInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopScanAllAccepted creates a StopScanAllAccepted with default headers values
func NewStopScanAllAccepted() *StopScanAllAccepted {
	return &StopScanAllAccepted{}
}

/*StopScanAllAccepted handles this case with default header values.

Accepted
*/
type StopScanAllAccepted struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *StopScanAllAccepted) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/stop][%d] stopScanAllAccepted ", 202)
}

func (o *StopScanAllAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewStopScanAllBadRequest creates a StopScanAllBadRequest with default headers values
func NewStopScanAllBadRequest() *StopScanAllBadRequest {
	return &StopScanAllBadRequest{}
}

/*StopScanAllBadRequest handles this case with default header values.

Bad request
*/
type StopScanAllBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopScanAllBadRequest) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/stop][%d] stopScanAllBadRequest  %+v", 400, o.Payload)
}

func (o *StopScanAllBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopScanAllBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopScanAllUnauthorized creates a StopScanAllUnauthorized with default headers values
func NewStopScanAllUnauthorized() *StopScanAllUnauthorized {
	return &StopScanAllUnauthorized{}
}

/*StopScanAllUnauthorized handles this case with default header values.

Unauthorized
*/
type StopScanAllUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopScanAllUnauthorized) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/stop][%d] stopScanAllUnauthorized  %+v", 401, o.Payload)
}

func (o *StopScanAllUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopScanAllUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopScanAllForbidden creates a StopScanAllForbidden with default headers values
func NewStopScanAllForbidden() *StopScanAllForbidden {
	return &StopScanAllForbidden{}
}

/*StopScanAllForbidden handles this case with default header values.

Forbidden
*/
type StopScanAllForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopScanAllForbidden) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/stop][%d] stopScanAllForbidden  %+v", 403, o.Payload)
}

func (o *StopScanAllForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopScanAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopScanAllInternalServerError creates a StopScanAllInternalServerError with default headers values
func NewStopScanAllInternalServerError() *StopScanAllInternalServerError {
	return &StopScanAllInternalServerError{}
}

/*StopScanAllInternalServerError handles this case with default header values.

Internal server error
*/
type StopScanAllInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopScanAllInternalServerError) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/stop][%d] stopScanAllInternalServerError  %+v", 500, o.Payload)
}

func (o *StopScanAllInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopScanAllInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

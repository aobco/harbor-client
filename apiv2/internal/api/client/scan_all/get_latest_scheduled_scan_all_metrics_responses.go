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

// GetLatestScheduledScanAllMetricsReader is a Reader for the GetLatestScheduledScanAllMetrics structure.
type GetLatestScheduledScanAllMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestScheduledScanAllMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestScheduledScanAllMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLatestScheduledScanAllMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLatestScheduledScanAllMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewGetLatestScheduledScanAllMetricsPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLatestScheduledScanAllMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLatestScheduledScanAllMetricsOK creates a GetLatestScheduledScanAllMetricsOK with default headers values
func NewGetLatestScheduledScanAllMetricsOK() *GetLatestScheduledScanAllMetricsOK {
	return &GetLatestScheduledScanAllMetricsOK{}
}

/*GetLatestScheduledScanAllMetricsOK handles this case with default header values.

OK
*/
type GetLatestScheduledScanAllMetricsOK struct {
	Payload *model.Stats
}

func (o *GetLatestScheduledScanAllMetricsOK) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsOK  %+v", 200, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsOK) GetPayload() *model.Stats {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Stats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScheduledScanAllMetricsUnauthorized creates a GetLatestScheduledScanAllMetricsUnauthorized with default headers values
func NewGetLatestScheduledScanAllMetricsUnauthorized() *GetLatestScheduledScanAllMetricsUnauthorized {
	return &GetLatestScheduledScanAllMetricsUnauthorized{}
}

/*GetLatestScheduledScanAllMetricsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetLatestScheduledScanAllMetricsUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLatestScheduledScanAllMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScheduledScanAllMetricsForbidden creates a GetLatestScheduledScanAllMetricsForbidden with default headers values
func NewGetLatestScheduledScanAllMetricsForbidden() *GetLatestScheduledScanAllMetricsForbidden {
	return &GetLatestScheduledScanAllMetricsForbidden{}
}

/*GetLatestScheduledScanAllMetricsForbidden handles this case with default header values.

Forbidden
*/
type GetLatestScheduledScanAllMetricsForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLatestScheduledScanAllMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsForbidden  %+v", 403, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScheduledScanAllMetricsPreconditionFailed creates a GetLatestScheduledScanAllMetricsPreconditionFailed with default headers values
func NewGetLatestScheduledScanAllMetricsPreconditionFailed() *GetLatestScheduledScanAllMetricsPreconditionFailed {
	return &GetLatestScheduledScanAllMetricsPreconditionFailed{}
}

/*GetLatestScheduledScanAllMetricsPreconditionFailed handles this case with default header values.

Precondition failed
*/
type GetLatestScheduledScanAllMetricsPreconditionFailed struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsPreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScheduledScanAllMetricsInternalServerError creates a GetLatestScheduledScanAllMetricsInternalServerError with default headers values
func NewGetLatestScheduledScanAllMetricsInternalServerError() *GetLatestScheduledScanAllMetricsInternalServerError {
	return &GetLatestScheduledScanAllMetricsInternalServerError{}
}

/*GetLatestScheduledScanAllMetricsInternalServerError handles this case with default header values.

Internal server error
*/
type GetLatestScheduledScanAllMetricsInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLatestScheduledScanAllMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

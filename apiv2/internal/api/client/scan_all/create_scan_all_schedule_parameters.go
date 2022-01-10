// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/aobco/harbor-client/v5/apiv2/model"
)

// NewCreateScanAllScheduleParams creates a new CreateScanAllScheduleParams object
// with the default values initialized.
func NewCreateScanAllScheduleParams() *CreateScanAllScheduleParams {
	var ()
	return &CreateScanAllScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateScanAllScheduleParamsWithTimeout creates a new CreateScanAllScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateScanAllScheduleParamsWithTimeout(timeout time.Duration) *CreateScanAllScheduleParams {
	var ()
	return &CreateScanAllScheduleParams{

		timeout: timeout,
	}
}

// NewCreateScanAllScheduleParamsWithContext creates a new CreateScanAllScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateScanAllScheduleParamsWithContext(ctx context.Context) *CreateScanAllScheduleParams {
	var ()
	return &CreateScanAllScheduleParams{

		Context: ctx,
	}
}

// NewCreateScanAllScheduleParamsWithHTTPClient creates a new CreateScanAllScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateScanAllScheduleParamsWithHTTPClient(client *http.Client) *CreateScanAllScheduleParams {
	var ()
	return &CreateScanAllScheduleParams{
		HTTPClient: client,
	}
}

/*CreateScanAllScheduleParams contains all the parameters to send to the API endpoint
for the create scan all schedule operation typically these are written to a http.Request
*/
type CreateScanAllScheduleParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Schedule
	  Create a schedule or a manual trigger for the scan all job.

	*/
	Schedule *model.Schedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create scan all schedule params
func (o *CreateScanAllScheduleParams) WithTimeout(timeout time.Duration) *CreateScanAllScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create scan all schedule params
func (o *CreateScanAllScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create scan all schedule params
func (o *CreateScanAllScheduleParams) WithContext(ctx context.Context) *CreateScanAllScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create scan all schedule params
func (o *CreateScanAllScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create scan all schedule params
func (o *CreateScanAllScheduleParams) WithHTTPClient(client *http.Client) *CreateScanAllScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create scan all schedule params
func (o *CreateScanAllScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create scan all schedule params
func (o *CreateScanAllScheduleParams) WithXRequestID(xRequestID *string) *CreateScanAllScheduleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create scan all schedule params
func (o *CreateScanAllScheduleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSchedule adds the schedule to the create scan all schedule params
func (o *CreateScanAllScheduleParams) WithSchedule(schedule *model.Schedule) *CreateScanAllScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the create scan all schedule params
func (o *CreateScanAllScheduleParams) SetSchedule(schedule *model.Schedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *CreateScanAllScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

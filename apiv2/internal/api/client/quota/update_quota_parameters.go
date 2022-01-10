// Code generated by go-swagger; DO NOT EDIT.

package quota

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
	"github.com/go-openapi/swag"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// NewUpdateQuotaParams creates a new UpdateQuotaParams object
// with the default values initialized.
func NewUpdateQuotaParams() *UpdateQuotaParams {
	var ()
	return &UpdateQuotaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateQuotaParamsWithTimeout creates a new UpdateQuotaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateQuotaParamsWithTimeout(timeout time.Duration) *UpdateQuotaParams {
	var ()
	return &UpdateQuotaParams{

		timeout: timeout,
	}
}

// NewUpdateQuotaParamsWithContext creates a new UpdateQuotaParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateQuotaParamsWithContext(ctx context.Context) *UpdateQuotaParams {
	var ()
	return &UpdateQuotaParams{

		Context: ctx,
	}
}

// NewUpdateQuotaParamsWithHTTPClient creates a new UpdateQuotaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateQuotaParamsWithHTTPClient(client *http.Client) *UpdateQuotaParams {
	var ()
	return &UpdateQuotaParams{
		HTTPClient: client,
	}
}

/*UpdateQuotaParams contains all the parameters to send to the API endpoint
for the update quota operation typically these are written to a http.Request
*/
type UpdateQuotaParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Hard
	  The new hard limits for the quota

	*/
	Hard *model.QuotaUpdateReq
	/*ID
	  Quota ID

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update quota params
func (o *UpdateQuotaParams) WithTimeout(timeout time.Duration) *UpdateQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update quota params
func (o *UpdateQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update quota params
func (o *UpdateQuotaParams) WithContext(ctx context.Context) *UpdateQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update quota params
func (o *UpdateQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update quota params
func (o *UpdateQuotaParams) WithHTTPClient(client *http.Client) *UpdateQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update quota params
func (o *UpdateQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update quota params
func (o *UpdateQuotaParams) WithXRequestID(xRequestID *string) *UpdateQuotaParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update quota params
func (o *UpdateQuotaParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithHard adds the hard to the update quota params
func (o *UpdateQuotaParams) WithHard(hard *model.QuotaUpdateReq) *UpdateQuotaParams {
	o.SetHard(hard)
	return o
}

// SetHard adds the hard to the update quota params
func (o *UpdateQuotaParams) SetHard(hard *model.QuotaUpdateReq) {
	o.Hard = hard
}

// WithID adds the id to the update quota params
func (o *UpdateQuotaParams) WithID(id int64) *UpdateQuotaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update quota params
func (o *UpdateQuotaParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Hard != nil {
		if err := r.SetBodyParam(o.Hard); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

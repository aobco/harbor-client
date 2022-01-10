// Code generated by go-swagger; DO NOT EDIT.

package configure

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
)

// NewGetInternalconfigParams creates a new GetInternalconfigParams object
// with the default values initialized.
func NewGetInternalconfigParams() *GetInternalconfigParams {
	var ()
	return &GetInternalconfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInternalconfigParamsWithTimeout creates a new GetInternalconfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInternalconfigParamsWithTimeout(timeout time.Duration) *GetInternalconfigParams {
	var ()
	return &GetInternalconfigParams{

		timeout: timeout,
	}
}

// NewGetInternalconfigParamsWithContext creates a new GetInternalconfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInternalconfigParamsWithContext(ctx context.Context) *GetInternalconfigParams {
	var ()
	return &GetInternalconfigParams{

		Context: ctx,
	}
}

// NewGetInternalconfigParamsWithHTTPClient creates a new GetInternalconfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInternalconfigParamsWithHTTPClient(client *http.Client) *GetInternalconfigParams {
	var ()
	return &GetInternalconfigParams{
		HTTPClient: client,
	}
}

/*GetInternalconfigParams contains all the parameters to send to the API endpoint
for the get internalconfig operation typically these are written to a http.Request
*/
type GetInternalconfigParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get internalconfig params
func (o *GetInternalconfigParams) WithTimeout(timeout time.Duration) *GetInternalconfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get internalconfig params
func (o *GetInternalconfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get internalconfig params
func (o *GetInternalconfigParams) WithContext(ctx context.Context) *GetInternalconfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get internalconfig params
func (o *GetInternalconfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get internalconfig params
func (o *GetInternalconfigParams) WithHTTPClient(client *http.Client) *GetInternalconfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get internalconfig params
func (o *GetInternalconfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get internalconfig params
func (o *GetInternalconfigParams) WithXRequestID(xRequestID *string) *GetInternalconfigParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get internalconfig params
func (o *GetInternalconfigParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInternalconfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

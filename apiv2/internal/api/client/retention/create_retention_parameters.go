// Code generated by go-swagger; DO NOT EDIT.

package retention

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

// NewCreateRetentionParams creates a new CreateRetentionParams object
// with the default values initialized.
func NewCreateRetentionParams() *CreateRetentionParams {
	var ()
	return &CreateRetentionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRetentionParamsWithTimeout creates a new CreateRetentionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRetentionParamsWithTimeout(timeout time.Duration) *CreateRetentionParams {
	var ()
	return &CreateRetentionParams{

		timeout: timeout,
	}
}

// NewCreateRetentionParamsWithContext creates a new CreateRetentionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRetentionParamsWithContext(ctx context.Context) *CreateRetentionParams {
	var ()
	return &CreateRetentionParams{

		Context: ctx,
	}
}

// NewCreateRetentionParamsWithHTTPClient creates a new CreateRetentionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRetentionParamsWithHTTPClient(client *http.Client) *CreateRetentionParams {
	var ()
	return &CreateRetentionParams{
		HTTPClient: client,
	}
}

/*CreateRetentionParams contains all the parameters to send to the API endpoint
for the create retention operation typically these are written to a http.Request
*/
type CreateRetentionParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Policy
	  Create Retention Policy successfully.

	*/
	Policy *model.RetentionPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create retention params
func (o *CreateRetentionParams) WithTimeout(timeout time.Duration) *CreateRetentionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create retention params
func (o *CreateRetentionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create retention params
func (o *CreateRetentionParams) WithContext(ctx context.Context) *CreateRetentionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create retention params
func (o *CreateRetentionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create retention params
func (o *CreateRetentionParams) WithHTTPClient(client *http.Client) *CreateRetentionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create retention params
func (o *CreateRetentionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create retention params
func (o *CreateRetentionParams) WithXRequestID(xRequestID *string) *CreateRetentionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create retention params
func (o *CreateRetentionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPolicy adds the policy to the create retention params
func (o *CreateRetentionParams) WithPolicy(policy *model.RetentionPolicy) *CreateRetentionParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the create retention params
func (o *CreateRetentionParams) SetPolicy(policy *model.RetentionPolicy) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRetentionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

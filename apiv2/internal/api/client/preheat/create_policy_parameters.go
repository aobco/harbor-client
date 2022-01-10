// Code generated by go-swagger; DO NOT EDIT.

package preheat

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

// NewCreatePolicyParams creates a new CreatePolicyParams object
// with the default values initialized.
func NewCreatePolicyParams() *CreatePolicyParams {
	var ()
	return &CreatePolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePolicyParamsWithTimeout creates a new CreatePolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePolicyParamsWithTimeout(timeout time.Duration) *CreatePolicyParams {
	var ()
	return &CreatePolicyParams{

		timeout: timeout,
	}
}

// NewCreatePolicyParamsWithContext creates a new CreatePolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePolicyParamsWithContext(ctx context.Context) *CreatePolicyParams {
	var ()
	return &CreatePolicyParams{

		Context: ctx,
	}
}

// NewCreatePolicyParamsWithHTTPClient creates a new CreatePolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePolicyParamsWithHTTPClient(client *http.Client) *CreatePolicyParams {
	var ()
	return &CreatePolicyParams{
		HTTPClient: client,
	}
}

/*CreatePolicyParams contains all the parameters to send to the API endpoint
for the create policy operation typically these are written to a http.Request
*/
type CreatePolicyParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Policy
	  The policy schema info

	*/
	Policy *model.PreheatPolicy
	/*ProjectName
	  The name of the project

	*/
	ProjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create policy params
func (o *CreatePolicyParams) WithTimeout(timeout time.Duration) *CreatePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policy params
func (o *CreatePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policy params
func (o *CreatePolicyParams) WithContext(ctx context.Context) *CreatePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policy params
func (o *CreatePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policy params
func (o *CreatePolicyParams) WithHTTPClient(client *http.Client) *CreatePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policy params
func (o *CreatePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create policy params
func (o *CreatePolicyParams) WithXRequestID(xRequestID *string) *CreatePolicyParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create policy params
func (o *CreatePolicyParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPolicy adds the policy to the create policy params
func (o *CreatePolicyParams) WithPolicy(policy *model.PreheatPolicy) *CreatePolicyParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the create policy params
func (o *CreatePolicyParams) SetPolicy(policy *model.PreheatPolicy) {
	o.Policy = policy
}

// WithProjectName adds the projectName to the create policy params
func (o *CreatePolicyParams) WithProjectName(projectName string) *CreatePolicyParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the create policy params
func (o *CreatePolicyParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

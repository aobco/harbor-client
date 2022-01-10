// Code generated by go-swagger; DO NOT EDIT.

package products

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

	"github.com/aobco/harbor-client/v5/apiv1/model"
)

// NewPutProjectsProjectIDMembersMidParams creates a new PutProjectsProjectIDMembersMidParams object
// with the default values initialized.
func NewPutProjectsProjectIDMembersMidParams() *PutProjectsProjectIDMembersMidParams {
	var ()
	return &PutProjectsProjectIDMembersMidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectsProjectIDMembersMidParamsWithTimeout creates a new PutProjectsProjectIDMembersMidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutProjectsProjectIDMembersMidParamsWithTimeout(timeout time.Duration) *PutProjectsProjectIDMembersMidParams {
	var ()
	return &PutProjectsProjectIDMembersMidParams{

		timeout: timeout,
	}
}

// NewPutProjectsProjectIDMembersMidParamsWithContext creates a new PutProjectsProjectIDMembersMidParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutProjectsProjectIDMembersMidParamsWithContext(ctx context.Context) *PutProjectsProjectIDMembersMidParams {
	var ()
	return &PutProjectsProjectIDMembersMidParams{

		Context: ctx,
	}
}

// NewPutProjectsProjectIDMembersMidParamsWithHTTPClient creates a new PutProjectsProjectIDMembersMidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutProjectsProjectIDMembersMidParamsWithHTTPClient(client *http.Client) *PutProjectsProjectIDMembersMidParams {
	var ()
	return &PutProjectsProjectIDMembersMidParams{
		HTTPClient: client,
	}
}

/*PutProjectsProjectIDMembersMidParams contains all the parameters to send to the API endpoint
for the put projects project ID members mid operation typically these are written to a http.Request
*/
type PutProjectsProjectIDMembersMidParams struct {

	/*Mid
	  Member ID.

	*/
	Mid int64
	/*ProjectID
	  Relevant project ID.

	*/
	ProjectID int64
	/*Role*/
	Role *model.RoleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) WithTimeout(timeout time.Duration) *PutProjectsProjectIDMembersMidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) WithContext(ctx context.Context) *PutProjectsProjectIDMembersMidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) WithHTTPClient(client *http.Client) *PutProjectsProjectIDMembersMidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMid adds the mid to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) WithMid(mid int64) *PutProjectsProjectIDMembersMidParams {
	o.SetMid(mid)
	return o
}

// SetMid adds the mid to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) SetMid(mid int64) {
	o.Mid = mid
}

// WithProjectID adds the projectID to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) WithProjectID(projectID int64) *PutProjectsProjectIDMembersMidParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithRole adds the role to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) WithRole(role *model.RoleRequest) *PutProjectsProjectIDMembersMidParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the put projects project ID members mid params
func (o *PutProjectsProjectIDMembersMidParams) SetRole(role *model.RoleRequest) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectsProjectIDMembersMidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mid
	if err := r.SetPathParam("mid", swag.FormatInt64(o.Mid)); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if o.Role != nil {
		if err := r.SetBodyParam(o.Role); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package usergroup

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

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// NewCreateUserGroupParams creates a new CreateUserGroupParams object
// with the default values initialized.
func NewCreateUserGroupParams() *CreateUserGroupParams {
	var ()
	return &CreateUserGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserGroupParamsWithTimeout creates a new CreateUserGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserGroupParamsWithTimeout(timeout time.Duration) *CreateUserGroupParams {
	var ()
	return &CreateUserGroupParams{

		timeout: timeout,
	}
}

// NewCreateUserGroupParamsWithContext creates a new CreateUserGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUserGroupParamsWithContext(ctx context.Context) *CreateUserGroupParams {
	var ()
	return &CreateUserGroupParams{

		Context: ctx,
	}
}

// NewCreateUserGroupParamsWithHTTPClient creates a new CreateUserGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUserGroupParamsWithHTTPClient(client *http.Client) *CreateUserGroupParams {
	var ()
	return &CreateUserGroupParams{
		HTTPClient: client,
	}
}

/*CreateUserGroupParams contains all the parameters to send to the API endpoint
for the create user group operation typically these are written to a http.Request
*/
type CreateUserGroupParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Usergroup*/
	Usergroup *model.UserGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create user group params
func (o *CreateUserGroupParams) WithTimeout(timeout time.Duration) *CreateUserGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user group params
func (o *CreateUserGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user group params
func (o *CreateUserGroupParams) WithContext(ctx context.Context) *CreateUserGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user group params
func (o *CreateUserGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user group params
func (o *CreateUserGroupParams) WithHTTPClient(client *http.Client) *CreateUserGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user group params
func (o *CreateUserGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create user group params
func (o *CreateUserGroupParams) WithXRequestID(xRequestID *string) *CreateUserGroupParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create user group params
func (o *CreateUserGroupParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithUsergroup adds the usergroup to the create user group params
func (o *CreateUserGroupParams) WithUsergroup(usergroup *model.UserGroup) *CreateUserGroupParams {
	o.SetUsergroup(usergroup)
	return o
}

// SetUsergroup adds the usergroup to the create user group params
func (o *CreateUserGroupParams) SetUsergroup(usergroup *model.UserGroup) {
	o.Usergroup = usergroup
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Usergroup != nil {
		if err := r.SetBodyParam(o.Usergroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package replication

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

	"github.com/aobco/harbor-client/v5/apiv2/model"
)

// NewUpdateReplicationPolicyParams creates a new UpdateReplicationPolicyParams object
// with the default values initialized.
func NewUpdateReplicationPolicyParams() *UpdateReplicationPolicyParams {
	var ()
	return &UpdateReplicationPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReplicationPolicyParamsWithTimeout creates a new UpdateReplicationPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateReplicationPolicyParamsWithTimeout(timeout time.Duration) *UpdateReplicationPolicyParams {
	var ()
	return &UpdateReplicationPolicyParams{

		timeout: timeout,
	}
}

// NewUpdateReplicationPolicyParamsWithContext creates a new UpdateReplicationPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateReplicationPolicyParamsWithContext(ctx context.Context) *UpdateReplicationPolicyParams {
	var ()
	return &UpdateReplicationPolicyParams{

		Context: ctx,
	}
}

// NewUpdateReplicationPolicyParamsWithHTTPClient creates a new UpdateReplicationPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateReplicationPolicyParamsWithHTTPClient(client *http.Client) *UpdateReplicationPolicyParams {
	var ()
	return &UpdateReplicationPolicyParams{
		HTTPClient: client,
	}
}

/*UpdateReplicationPolicyParams contains all the parameters to send to the API endpoint
for the update replication policy operation typically these are written to a http.Request
*/
type UpdateReplicationPolicyParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ID
	  The policy ID

	*/
	ID int64
	/*Policy
	  The replication policy

	*/
	Policy *model.ReplicationPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update replication policy params
func (o *UpdateReplicationPolicyParams) WithTimeout(timeout time.Duration) *UpdateReplicationPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update replication policy params
func (o *UpdateReplicationPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update replication policy params
func (o *UpdateReplicationPolicyParams) WithContext(ctx context.Context) *UpdateReplicationPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update replication policy params
func (o *UpdateReplicationPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update replication policy params
func (o *UpdateReplicationPolicyParams) WithHTTPClient(client *http.Client) *UpdateReplicationPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update replication policy params
func (o *UpdateReplicationPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update replication policy params
func (o *UpdateReplicationPolicyParams) WithXRequestID(xRequestID *string) *UpdateReplicationPolicyParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update replication policy params
func (o *UpdateReplicationPolicyParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the update replication policy params
func (o *UpdateReplicationPolicyParams) WithID(id int64) *UpdateReplicationPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update replication policy params
func (o *UpdateReplicationPolicyParams) SetID(id int64) {
	o.ID = id
}

// WithPolicy adds the policy to the update replication policy params
func (o *UpdateReplicationPolicyParams) WithPolicy(policy *model.ReplicationPolicy) *UpdateReplicationPolicyParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the update replication policy params
func (o *UpdateReplicationPolicyParams) SetPolicy(policy *model.ReplicationPolicy) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReplicationPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
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

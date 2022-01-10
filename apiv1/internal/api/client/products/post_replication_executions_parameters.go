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

	"github.com/aobco/harbor-client/apiv1/model"
)

// NewPostReplicationExecutionsParams creates a new PostReplicationExecutionsParams object
// with the default values initialized.
func NewPostReplicationExecutionsParams() *PostReplicationExecutionsParams {
	var ()
	return &PostReplicationExecutionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostReplicationExecutionsParamsWithTimeout creates a new PostReplicationExecutionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostReplicationExecutionsParamsWithTimeout(timeout time.Duration) *PostReplicationExecutionsParams {
	var ()
	return &PostReplicationExecutionsParams{

		timeout: timeout,
	}
}

// NewPostReplicationExecutionsParamsWithContext creates a new PostReplicationExecutionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostReplicationExecutionsParamsWithContext(ctx context.Context) *PostReplicationExecutionsParams {
	var ()
	return &PostReplicationExecutionsParams{

		Context: ctx,
	}
}

// NewPostReplicationExecutionsParamsWithHTTPClient creates a new PostReplicationExecutionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostReplicationExecutionsParamsWithHTTPClient(client *http.Client) *PostReplicationExecutionsParams {
	var ()
	return &PostReplicationExecutionsParams{
		HTTPClient: client,
	}
}

/*PostReplicationExecutionsParams contains all the parameters to send to the API endpoint
for the post replication executions operation typically these are written to a http.Request
*/
type PostReplicationExecutionsParams struct {

	/*Execution
	  The execution that needs to be started, only the property "policy_id" is needed.

	*/
	Execution *model.ReplicationExecution

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post replication executions params
func (o *PostReplicationExecutionsParams) WithTimeout(timeout time.Duration) *PostReplicationExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post replication executions params
func (o *PostReplicationExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post replication executions params
func (o *PostReplicationExecutionsParams) WithContext(ctx context.Context) *PostReplicationExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post replication executions params
func (o *PostReplicationExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post replication executions params
func (o *PostReplicationExecutionsParams) WithHTTPClient(client *http.Client) *PostReplicationExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post replication executions params
func (o *PostReplicationExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecution adds the execution to the post replication executions params
func (o *PostReplicationExecutionsParams) WithExecution(execution *model.ReplicationExecution) *PostReplicationExecutionsParams {
	o.SetExecution(execution)
	return o
}

// SetExecution adds the execution to the post replication executions params
func (o *PostReplicationExecutionsParams) SetExecution(execution *model.ReplicationExecution) {
	o.Execution = execution
}

// WriteToRequest writes these params to a swagger request
func (o *PostReplicationExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Execution != nil {
		if err := r.SetBodyParam(o.Execution); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new robotv1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for robotv1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRobotV1(params *CreateRobotV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateRobotV1Created, error)

	DeleteRobotV1(params *DeleteRobotV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteRobotV1OK, error)

	GetRobotByIDV1(params *GetRobotByIDV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetRobotByIDV1OK, error)

	ListRobotV1(params *ListRobotV1Params, authInfo runtime.ClientAuthInfoWriter) (*ListRobotV1OK, error)

	UpdateRobotV1(params *UpdateRobotV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateRobotV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRobotV1 creates a robot account

  Create a robot account
*/
func (a *Client) CreateRobotV1(params *CreateRobotV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateRobotV1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRobotV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRobotV1",
		Method:             "POST",
		PathPattern:        "/projects/{project_name_or_id}/robots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRobotV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRobotV1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateRobotV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRobotV1 deletes a robot account

  This endpoint deletes specific robot account information by robot ID.
*/
func (a *Client) DeleteRobotV1(params *DeleteRobotV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteRobotV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRobotV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRobotV1",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name_or_id}/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRobotV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRobotV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteRobotV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRobotByIDV1 gets a robot account

  This endpoint returns specific robot account information by robot ID.
*/
func (a *Client) GetRobotByIDV1(params *GetRobotByIDV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetRobotByIDV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRobotByIDV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRobotByIDV1",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRobotByIDV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRobotByIDV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRobotByIDV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRobotV1 gets all robot accounts of specified project

  Get all robot accounts of specified project
*/
func (a *Client) ListRobotV1(params *ListRobotV1Params, authInfo runtime.ClientAuthInfoWriter) (*ListRobotV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRobotV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListRobotV1",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/robots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRobotV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRobotV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListRobotV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRobotV1 updates status of robot account

  Used to disable/enable a specified robot account.
*/
func (a *Client) UpdateRobotV1(params *UpdateRobotV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateRobotV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRobotV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateRobotV1",
		Method:             "PUT",
		PathPattern:        "/projects/{project_name_or_id}/robots/{robot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateRobotV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRobotV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateRobotV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

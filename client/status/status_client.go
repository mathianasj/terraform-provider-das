// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetStatuses(params *GetStatusesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStatusesOK, error)

	UpdateStatus(params *UpdateStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateStatusOK, error)

	UpdateStatusWithPartition(params *UpdateStatusWithPartitionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateStatusWithPartitionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetStatuses gets current o p a statuses
*/
func (a *Client) GetStatuses(params *GetStatusesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStatusesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatusesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStatuses",
		Method:             "GET",
		PathPattern:        "/v1/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStatusesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStatusesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStatuses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateStatus updates current o p a status
*/
func (a *Client) UpdateStatus(params *UpdateStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateStatus",
		Method:             "POST",
		PathPattern:        "/v1/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateStatusWithPartition updates current o p a status
*/
func (a *Client) UpdateStatusWithPartition(params *UpdateStatusWithPartitionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateStatusWithPartitionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStatusWithPartitionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateStatusWithPartition",
		Method:             "POST",
		PathPattern:        "/v1/status/{partition}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateStatusWithPartitionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateStatusWithPartitionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateStatusWithPartition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
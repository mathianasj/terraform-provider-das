// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tokens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tokens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateToken(params *CreateTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTokenOK, error)

	GetToken(params *GetTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTokenOK, error)

	ListTokens(params *ListTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTokensOK, error)

	RevokeToken(params *RevokeTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RevokeTokenOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateToken creates or update a token

  If If-None-Match header is set to *, tries to create a token, otherwise will try to either update or create depending on whether an unexpired token with that ID already exists. Token creation errors with a 409 code if an unexpired one already exists, on success returns the token secret (valid for the TTL whose default value is ~10 years). Token updates return nothing unless `regenerate` is true, in which case it returns the new secret. WARNING: If allow_path_patterns is unset or an empty list, all paths are allowed.
*/
func (a *Client) CreateToken(params *CreateTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateToken",
		Method:             "PUT",
		PathPattern:        "/v1/tokens/{tokenId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTokenReader{formats: a.formats},
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
	success, ok := result.(*CreateTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetToken gets token
*/
func (a *Client) GetToken(params *GetTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetToken",
		Method:             "GET",
		PathPattern:        "/v1/tokens/{tokenId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTokenReader{formats: a.formats},
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
	success, ok := result.(*GetTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListTokens lists tokens
*/
func (a *Client) ListTokens(params *ListTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTokensParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListTokens",
		Method:             "GET",
		PathPattern:        "/v1/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListTokensReader{formats: a.formats},
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
	success, ok := result.(*ListTokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListTokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RevokeToken revokes token
*/
func (a *Client) RevokeToken(params *RevokeTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RevokeTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RevokeToken",
		Method:             "DELETE",
		PathPattern:        "/v1/tokens/{tokenId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RevokeTokenReader{formats: a.formats},
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
	success, ok := result.(*RevokeTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RevokeToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
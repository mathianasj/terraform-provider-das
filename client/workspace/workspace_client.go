// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workspace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workspace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CommitFilesToSourceControlWorkspace(params *CommitFilesToSourceControlWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommitFilesToSourceControlWorkspaceOK, error)

	DeleteUserBranchWorkspace(params *DeleteUserBranchWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserBranchWorkspaceOK, error)

	GetRegions(params *GetRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRegionsOK, error)

	GetS3Config(params *GetS3ConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetS3ConfigOK, error)

	GetSourceControlFilesBranchWorkspace(params *GetSourceControlFilesBranchWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSourceControlFilesBranchWorkspaceOK, error)

	GetSourceControlFilesMasterWorkspace(params *GetSourceControlFilesMasterWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSourceControlFilesMasterWorkspaceOK, error)

	GetWorkspace(params *GetWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkspaceOK, error)

	KafkaVerifyConfig(params *KafkaVerifyConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KafkaVerifyConfigOK, error)

	S3VerifyConfig(params *S3VerifyConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*S3VerifyConfigOK, error)

	SourceControlVerifyConfigWorkspace(params *SourceControlVerifyConfigWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SourceControlVerifyConfigWorkspaceOK, error)

	UpdateS3Config(params *UpdateS3ConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateS3ConfigOK, error)

	UpdateWorkspace(params *UpdateWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateWorkspaceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CommitFilesToSourceControlWorkspace commits files to workspace source control

  Commit files to source control associated with a workspace
*/
func (a *Client) CommitFilesToSourceControlWorkspace(params *CommitFilesToSourceControlWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommitFilesToSourceControlWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitFilesToSourceControlWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CommitFilesToSourceControlWorkspace",
		Method:             "POST",
		PathPattern:        "/v1/workspace/{id}/commits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CommitFilesToSourceControlWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*CommitFilesToSourceControlWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CommitFilesToSourceControlWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUserBranchWorkspace deletes a user owned branch
*/
func (a *Client) DeleteUserBranchWorkspace(params *DeleteUserBranchWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserBranchWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserBranchWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteUserBranchWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v1/workspace/{id}/branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteUserBranchWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*DeleteUserBranchWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUserBranchWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRegions gets s3 regions list

  Get list of valid regions for S3 integration type
*/
func (a *Client) GetRegions(params *GetRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRegionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRegions",
		Method:             "GET",
		PathPattern:        "/v1/workspace/regions/{storagesvc}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRegionsReader{formats: a.formats},
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
	success, ok := result.(*GetRegionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRegions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetS3Config gets s3 decision configuration
*/
func (a *Client) GetS3Config(params *GetS3ConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetS3ConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetS3ConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetS3Config",
		Method:             "GET",
		PathPattern:        "/v1/workspace/s3-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetS3ConfigReader{formats: a.formats},
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
	success, ok := result.(*GetS3ConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetS3Config: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSourceControlFilesBranchWorkspace lists files in styra d a s created branch

  Gets the list of files for the branch that the Styra DAS creates when modifying rego in the Styra DAS UI and pushing the changes to GitHub in a branch for review.
*/
func (a *Client) GetSourceControlFilesBranchWorkspace(params *GetSourceControlFilesBranchWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSourceControlFilesBranchWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSourceControlFilesBranchWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSourceControlFilesBranchWorkspace",
		Method:             "GET",
		PathPattern:        "/v1/workspace/{id}/branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSourceControlFilesBranchWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*GetSourceControlFilesBranchWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSourceControlFilesBranchWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSourceControlFilesMasterWorkspace lists files in current branch

  Gets the list of files in the currently chosen branch.
*/
func (a *Client) GetSourceControlFilesMasterWorkspace(params *GetSourceControlFilesMasterWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSourceControlFilesMasterWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSourceControlFilesMasterWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSourceControlFilesMasterWorkspace",
		Method:             "GET",
		PathPattern:        "/v1/workspace/{id}/master",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSourceControlFilesMasterWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*GetSourceControlFilesMasterWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSourceControlFilesMasterWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWorkspace gets workspace
*/
func (a *Client) GetWorkspace(params *GetWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWorkspace",
		Method:             "GET",
		PathPattern:        "/v1/workspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*GetWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KafkaVerifyConfig kafkas connectivity test

  Verifies that the Kafka topic can be accessed with the provided credentials.
*/
func (a *Client) KafkaVerifyConfig(params *KafkaVerifyConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KafkaVerifyConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKafkaVerifyConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KafkaVerifyConfig",
		Method:             "POST",
		PathPattern:        "/v1/workspace/kafka/verify-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KafkaVerifyConfigReader{formats: a.formats},
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
	success, ok := result.(*KafkaVerifyConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for KafkaVerifyConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  S3VerifyConfig s3s connectivity test

  Verifies that the S3 bucket can be accessed with the provided credentials. Creates styra_test.json file
*/
func (a *Client) S3VerifyConfig(params *S3VerifyConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*S3VerifyConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewS3VerifyConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "S3VerifyConfig",
		Method:             "POST",
		PathPattern:        "/v1/workspace/s3/verify-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &S3VerifyConfigReader{formats: a.formats},
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
	success, ok := result.(*S3VerifyConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for S3VerifyConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SourceControlVerifyConfigWorkspace verifies git access

  Verifies that the repository can be accessed with the provided credentials
*/
func (a *Client) SourceControlVerifyConfigWorkspace(params *SourceControlVerifyConfigWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SourceControlVerifyConfigWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSourceControlVerifyConfigWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SourceControlVerifyConfigWorkspace",
		Method:             "POST",
		PathPattern:        "/v1/workspace/source-control/verify-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SourceControlVerifyConfigWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*SourceControlVerifyConfigWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SourceControlVerifyConfigWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateS3Config updates s3 decision configuration
*/
func (a *Client) UpdateS3Config(params *UpdateS3ConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateS3ConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateS3ConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateS3Config",
		Method:             "PUT",
		PathPattern:        "/v1/workspace/s3-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateS3ConfigReader{formats: a.formats},
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
	success, ok := result.(*UpdateS3ConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateS3Config: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateWorkspace updates workspace

  Updates workspace configuration
*/
func (a *Client) UpdateWorkspace(params *UpdateWorkspaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkspaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateWorkspace",
		Method:             "PUT",
		PathPattern:        "/v1/workspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateWorkspaceReader{formats: a.formats},
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
	success, ok := result.(*UpdateWorkspaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateWorkspace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

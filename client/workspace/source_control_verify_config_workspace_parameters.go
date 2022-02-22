// Code generated by go-swagger; DO NOT EDIT.

package workspace

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

	"github.com/mathianasj/terraform-provider-das/models"
)

// NewSourceControlVerifyConfigWorkspaceParams creates a new SourceControlVerifyConfigWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSourceControlVerifyConfigWorkspaceParams() *SourceControlVerifyConfigWorkspaceParams {
	return &SourceControlVerifyConfigWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSourceControlVerifyConfigWorkspaceParamsWithTimeout creates a new SourceControlVerifyConfigWorkspaceParams object
// with the ability to set a timeout on a request.
func NewSourceControlVerifyConfigWorkspaceParamsWithTimeout(timeout time.Duration) *SourceControlVerifyConfigWorkspaceParams {
	return &SourceControlVerifyConfigWorkspaceParams{
		timeout: timeout,
	}
}

// NewSourceControlVerifyConfigWorkspaceParamsWithContext creates a new SourceControlVerifyConfigWorkspaceParams object
// with the ability to set a context for a request.
func NewSourceControlVerifyConfigWorkspaceParamsWithContext(ctx context.Context) *SourceControlVerifyConfigWorkspaceParams {
	return &SourceControlVerifyConfigWorkspaceParams{
		Context: ctx,
	}
}

// NewSourceControlVerifyConfigWorkspaceParamsWithHTTPClient creates a new SourceControlVerifyConfigWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewSourceControlVerifyConfigWorkspaceParamsWithHTTPClient(client *http.Client) *SourceControlVerifyConfigWorkspaceParams {
	return &SourceControlVerifyConfigWorkspaceParams{
		HTTPClient: client,
	}
}

/* SourceControlVerifyConfigWorkspaceParams contains all the parameters to send to the API endpoint
   for the source control verify config workspace operation.

   Typically these are written to a http.Request.
*/
type SourceControlVerifyConfigWorkspaceParams struct {

	// Body.
	Body *models.GitV1VerifyConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the source control verify config workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SourceControlVerifyConfigWorkspaceParams) WithDefaults() *SourceControlVerifyConfigWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the source control verify config workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SourceControlVerifyConfigWorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the source control verify config workspace params
func (o *SourceControlVerifyConfigWorkspaceParams) WithTimeout(timeout time.Duration) *SourceControlVerifyConfigWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the source control verify config workspace params
func (o *SourceControlVerifyConfigWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the source control verify config workspace params
func (o *SourceControlVerifyConfigWorkspaceParams) WithContext(ctx context.Context) *SourceControlVerifyConfigWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the source control verify config workspace params
func (o *SourceControlVerifyConfigWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the source control verify config workspace params
func (o *SourceControlVerifyConfigWorkspaceParams) WithHTTPClient(client *http.Client) *SourceControlVerifyConfigWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the source control verify config workspace params
func (o *SourceControlVerifyConfigWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the source control verify config workspace params
func (o *SourceControlVerifyConfigWorkspaceParams) WithBody(body *models.GitV1VerifyConfigRequest) *SourceControlVerifyConfigWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the source control verify config workspace params
func (o *SourceControlVerifyConfigWorkspaceParams) SetBody(body *models.GitV1VerifyConfigRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SourceControlVerifyConfigWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
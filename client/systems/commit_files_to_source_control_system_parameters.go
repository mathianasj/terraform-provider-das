// Code generated by go-swagger; DO NOT EDIT.

package systems

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

// NewCommitFilesToSourceControlSystemParams creates a new CommitFilesToSourceControlSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCommitFilesToSourceControlSystemParams() *CommitFilesToSourceControlSystemParams {
	return &CommitFilesToSourceControlSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCommitFilesToSourceControlSystemParamsWithTimeout creates a new CommitFilesToSourceControlSystemParams object
// with the ability to set a timeout on a request.
func NewCommitFilesToSourceControlSystemParamsWithTimeout(timeout time.Duration) *CommitFilesToSourceControlSystemParams {
	return &CommitFilesToSourceControlSystemParams{
		timeout: timeout,
	}
}

// NewCommitFilesToSourceControlSystemParamsWithContext creates a new CommitFilesToSourceControlSystemParams object
// with the ability to set a context for a request.
func NewCommitFilesToSourceControlSystemParamsWithContext(ctx context.Context) *CommitFilesToSourceControlSystemParams {
	return &CommitFilesToSourceControlSystemParams{
		Context: ctx,
	}
}

// NewCommitFilesToSourceControlSystemParamsWithHTTPClient creates a new CommitFilesToSourceControlSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewCommitFilesToSourceControlSystemParamsWithHTTPClient(client *http.Client) *CommitFilesToSourceControlSystemParams {
	return &CommitFilesToSourceControlSystemParams{
		HTTPClient: client,
	}
}

/* CommitFilesToSourceControlSystemParams contains all the parameters to send to the API endpoint
   for the commit files to source control system operation.

   Typically these are written to a http.Request.
*/
type CommitFilesToSourceControlSystemParams struct {

	// Body.
	Body *models.GitV1CommitInput

	/* ID.

	   system id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the commit files to source control system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommitFilesToSourceControlSystemParams) WithDefaults() *CommitFilesToSourceControlSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the commit files to source control system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommitFilesToSourceControlSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) WithTimeout(timeout time.Duration) *CommitFilesToSourceControlSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) WithContext(ctx context.Context) *CommitFilesToSourceControlSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) WithHTTPClient(client *http.Client) *CommitFilesToSourceControlSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) WithBody(body *models.GitV1CommitInput) *CommitFilesToSourceControlSystemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) SetBody(body *models.GitV1CommitInput) {
	o.Body = body
}

// WithID adds the id to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) WithID(id string) *CommitFilesToSourceControlSystemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the commit files to source control system params
func (o *CommitFilesToSourceControlSystemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CommitFilesToSourceControlSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
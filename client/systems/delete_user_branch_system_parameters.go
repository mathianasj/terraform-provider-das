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
)

// NewDeleteUserBranchSystemParams creates a new DeleteUserBranchSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserBranchSystemParams() *DeleteUserBranchSystemParams {
	return &DeleteUserBranchSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserBranchSystemParamsWithTimeout creates a new DeleteUserBranchSystemParams object
// with the ability to set a timeout on a request.
func NewDeleteUserBranchSystemParamsWithTimeout(timeout time.Duration) *DeleteUserBranchSystemParams {
	return &DeleteUserBranchSystemParams{
		timeout: timeout,
	}
}

// NewDeleteUserBranchSystemParamsWithContext creates a new DeleteUserBranchSystemParams object
// with the ability to set a context for a request.
func NewDeleteUserBranchSystemParamsWithContext(ctx context.Context) *DeleteUserBranchSystemParams {
	return &DeleteUserBranchSystemParams{
		Context: ctx,
	}
}

// NewDeleteUserBranchSystemParamsWithHTTPClient creates a new DeleteUserBranchSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserBranchSystemParamsWithHTTPClient(client *http.Client) *DeleteUserBranchSystemParams {
	return &DeleteUserBranchSystemParams{
		HTTPClient: client,
	}
}

/* DeleteUserBranchSystemParams contains all the parameters to send to the API endpoint
   for the delete user branch system operation.

   Typically these are written to a http.Request.
*/
type DeleteUserBranchSystemParams struct {

	/* ID.

	   system id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user branch system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserBranchSystemParams) WithDefaults() *DeleteUserBranchSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user branch system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserBranchSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user branch system params
func (o *DeleteUserBranchSystemParams) WithTimeout(timeout time.Duration) *DeleteUserBranchSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user branch system params
func (o *DeleteUserBranchSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user branch system params
func (o *DeleteUserBranchSystemParams) WithContext(ctx context.Context) *DeleteUserBranchSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user branch system params
func (o *DeleteUserBranchSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user branch system params
func (o *DeleteUserBranchSystemParams) WithHTTPClient(client *http.Client) *DeleteUserBranchSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user branch system params
func (o *DeleteUserBranchSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete user branch system params
func (o *DeleteUserBranchSystemParams) WithID(id string) *DeleteUserBranchSystemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete user branch system params
func (o *DeleteUserBranchSystemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserBranchSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

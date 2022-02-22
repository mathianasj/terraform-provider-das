// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewDeleteAgentInfoParams creates a new DeleteAgentInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAgentInfoParams() *DeleteAgentInfoParams {
	return &DeleteAgentInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAgentInfoParamsWithTimeout creates a new DeleteAgentInfoParams object
// with the ability to set a timeout on a request.
func NewDeleteAgentInfoParamsWithTimeout(timeout time.Duration) *DeleteAgentInfoParams {
	return &DeleteAgentInfoParams{
		timeout: timeout,
	}
}

// NewDeleteAgentInfoParamsWithContext creates a new DeleteAgentInfoParams object
// with the ability to set a context for a request.
func NewDeleteAgentInfoParamsWithContext(ctx context.Context) *DeleteAgentInfoParams {
	return &DeleteAgentInfoParams{
		Context: ctx,
	}
}

// NewDeleteAgentInfoParamsWithHTTPClient creates a new DeleteAgentInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAgentInfoParamsWithHTTPClient(client *http.Client) *DeleteAgentInfoParams {
	return &DeleteAgentInfoParams{
		HTTPClient: client,
	}
}

/* DeleteAgentInfoParams contains all the parameters to send to the API endpoint
   for the delete agent info operation.

   Typically these are written to a http.Request.
*/
type DeleteAgentInfoParams struct {

	/* ID.

	   agent id
	*/
	ID string

	/* Kind.

	   agent kind such as "agents", "datasources", "slps", "exporters"
	*/
	Kind string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete agent info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentInfoParams) WithDefaults() *DeleteAgentInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete agent info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete agent info params
func (o *DeleteAgentInfoParams) WithTimeout(timeout time.Duration) *DeleteAgentInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete agent info params
func (o *DeleteAgentInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete agent info params
func (o *DeleteAgentInfoParams) WithContext(ctx context.Context) *DeleteAgentInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete agent info params
func (o *DeleteAgentInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete agent info params
func (o *DeleteAgentInfoParams) WithHTTPClient(client *http.Client) *DeleteAgentInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete agent info params
func (o *DeleteAgentInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete agent info params
func (o *DeleteAgentInfoParams) WithID(id string) *DeleteAgentInfoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete agent info params
func (o *DeleteAgentInfoParams) SetID(id string) {
	o.ID = id
}

// WithKind adds the kind to the delete agent info params
func (o *DeleteAgentInfoParams) WithKind(kind string) *DeleteAgentInfoParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the delete agent info params
func (o *DeleteAgentInfoParams) SetKind(kind string) {
	o.Kind = kind
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAgentInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

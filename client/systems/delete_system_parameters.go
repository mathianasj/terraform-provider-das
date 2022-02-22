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

// NewDeleteSystemParams creates a new DeleteSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSystemParams() *DeleteSystemParams {
	return &DeleteSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSystemParamsWithTimeout creates a new DeleteSystemParams object
// with the ability to set a timeout on a request.
func NewDeleteSystemParamsWithTimeout(timeout time.Duration) *DeleteSystemParams {
	return &DeleteSystemParams{
		timeout: timeout,
	}
}

// NewDeleteSystemParamsWithContext creates a new DeleteSystemParams object
// with the ability to set a context for a request.
func NewDeleteSystemParamsWithContext(ctx context.Context) *DeleteSystemParams {
	return &DeleteSystemParams{
		Context: ctx,
	}
}

// NewDeleteSystemParamsWithHTTPClient creates a new DeleteSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSystemParamsWithHTTPClient(client *http.Client) *DeleteSystemParams {
	return &DeleteSystemParams{
		HTTPClient: client,
	}
}

/* DeleteSystemParams contains all the parameters to send to the API endpoint
   for the delete system operation.

   Typically these are written to a http.Request.
*/
type DeleteSystemParams struct {

	/* Recursive.

	   if set to 'false', only deletes the system configuration and does not delete associated objects
	*/
	Recursive *string

	/* System.

	   system ID
	*/
	System string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSystemParams) WithDefaults() *DeleteSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete system params
func (o *DeleteSystemParams) WithTimeout(timeout time.Duration) *DeleteSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete system params
func (o *DeleteSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete system params
func (o *DeleteSystemParams) WithContext(ctx context.Context) *DeleteSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete system params
func (o *DeleteSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete system params
func (o *DeleteSystemParams) WithHTTPClient(client *http.Client) *DeleteSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete system params
func (o *DeleteSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecursive adds the recursive to the delete system params
func (o *DeleteSystemParams) WithRecursive(recursive *string) *DeleteSystemParams {
	o.SetRecursive(recursive)
	return o
}

// SetRecursive adds the recursive to the delete system params
func (o *DeleteSystemParams) SetRecursive(recursive *string) {
	o.Recursive = recursive
}

// WithSystem adds the system to the delete system params
func (o *DeleteSystemParams) WithSystem(system string) *DeleteSystemParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the delete system params
func (o *DeleteSystemParams) SetSystem(system string) {
	o.System = system
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Recursive != nil {

		// query param recursive
		var qrRecursive string

		if o.Recursive != nil {
			qrRecursive = *o.Recursive
		}
		qRecursive := qrRecursive
		if qRecursive != "" {

			if err := r.SetQueryParam("recursive", qRecursive); err != nil {
				return err
			}
		}
	}

	// path param system
	if err := r.SetPathParam("system", o.System); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
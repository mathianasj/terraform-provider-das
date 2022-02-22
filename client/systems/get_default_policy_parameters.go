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

// NewGetDefaultPolicyParams creates a new GetDefaultPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDefaultPolicyParams() *GetDefaultPolicyParams {
	return &GetDefaultPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDefaultPolicyParamsWithTimeout creates a new GetDefaultPolicyParams object
// with the ability to set a timeout on a request.
func NewGetDefaultPolicyParamsWithTimeout(timeout time.Duration) *GetDefaultPolicyParams {
	return &GetDefaultPolicyParams{
		timeout: timeout,
	}
}

// NewGetDefaultPolicyParamsWithContext creates a new GetDefaultPolicyParams object
// with the ability to set a context for a request.
func NewGetDefaultPolicyParamsWithContext(ctx context.Context) *GetDefaultPolicyParams {
	return &GetDefaultPolicyParams{
		Context: ctx,
	}
}

// NewGetDefaultPolicyParamsWithHTTPClient creates a new GetDefaultPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDefaultPolicyParamsWithHTTPClient(client *http.Client) *GetDefaultPolicyParams {
	return &GetDefaultPolicyParams{
		HTTPClient: client,
	}
}

/* GetDefaultPolicyParams contains all the parameters to send to the API endpoint
   for the get default policy operation.

   Typically these are written to a http.Request.
*/
type GetDefaultPolicyParams struct {

	/* Path.

	   policy path
	*/
	Path string

	/* System.

	   system ID
	*/
	System string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get default policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDefaultPolicyParams) WithDefaults() *GetDefaultPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get default policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDefaultPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get default policy params
func (o *GetDefaultPolicyParams) WithTimeout(timeout time.Duration) *GetDefaultPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get default policy params
func (o *GetDefaultPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get default policy params
func (o *GetDefaultPolicyParams) WithContext(ctx context.Context) *GetDefaultPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get default policy params
func (o *GetDefaultPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get default policy params
func (o *GetDefaultPolicyParams) WithHTTPClient(client *http.Client) *GetDefaultPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get default policy params
func (o *GetDefaultPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the get default policy params
func (o *GetDefaultPolicyParams) WithPath(path string) *GetDefaultPolicyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get default policy params
func (o *GetDefaultPolicyParams) SetPath(path string) {
	o.Path = path
}

// WithSystem adds the system to the get default policy params
func (o *GetDefaultPolicyParams) WithSystem(system string) *GetDefaultPolicyParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the get default policy params
func (o *GetDefaultPolicyParams) SetSystem(system string) {
	o.System = system
}

// WriteToRequest writes these params to a swagger request
func (o *GetDefaultPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
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

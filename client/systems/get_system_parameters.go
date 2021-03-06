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
	"github.com/go-openapi/swag"
)

// NewGetSystemParams creates a new GetSystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemParams() *GetSystemParams {
	return &GetSystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemParamsWithTimeout creates a new GetSystemParams object
// with the ability to set a timeout on a request.
func NewGetSystemParamsWithTimeout(timeout time.Duration) *GetSystemParams {
	return &GetSystemParams{
		timeout: timeout,
	}
}

// NewGetSystemParamsWithContext creates a new GetSystemParams object
// with the ability to set a context for a request.
func NewGetSystemParamsWithContext(ctx context.Context) *GetSystemParams {
	return &GetSystemParams{
		Context: ctx,
	}
}

// NewGetSystemParamsWithHTTPClient creates a new GetSystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemParamsWithHTTPClient(client *http.Client) *GetSystemParams {
	return &GetSystemParams{
		HTTPClient: client,
	}
}

/* GetSystemParams contains all the parameters to send to the API endpoint
   for the get system operation.

   Typically these are written to a http.Request.
*/
type GetSystemParams struct {

	/* Datasources.

	   set to 'false' to omit datasources from the output
	*/
	Datasources *bool

	/* Errors.

	   set to 'false' to omit errors/warnings from the output
	*/
	Errors *bool

	/* Modules.

	   set to 'false' to omit modules from the output
	*/
	Modules *bool

	/* Policies.

	   set to 'false' to omit policies from the output
	*/
	Policies *bool

	/* Stacks.

	   set to 'false' to omit matching_stacks from the output
	*/
	Stacks *bool

	/* System.

	   system ID
	*/
	System string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemParams) WithDefaults() *GetSystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system params
func (o *GetSystemParams) WithTimeout(timeout time.Duration) *GetSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system params
func (o *GetSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system params
func (o *GetSystemParams) WithContext(ctx context.Context) *GetSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system params
func (o *GetSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system params
func (o *GetSystemParams) WithHTTPClient(client *http.Client) *GetSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system params
func (o *GetSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasources adds the datasources to the get system params
func (o *GetSystemParams) WithDatasources(datasources *bool) *GetSystemParams {
	o.SetDatasources(datasources)
	return o
}

// SetDatasources adds the datasources to the get system params
func (o *GetSystemParams) SetDatasources(datasources *bool) {
	o.Datasources = datasources
}

// WithErrors adds the errors to the get system params
func (o *GetSystemParams) WithErrors(errors *bool) *GetSystemParams {
	o.SetErrors(errors)
	return o
}

// SetErrors adds the errors to the get system params
func (o *GetSystemParams) SetErrors(errors *bool) {
	o.Errors = errors
}

// WithModules adds the modules to the get system params
func (o *GetSystemParams) WithModules(modules *bool) *GetSystemParams {
	o.SetModules(modules)
	return o
}

// SetModules adds the modules to the get system params
func (o *GetSystemParams) SetModules(modules *bool) {
	o.Modules = modules
}

// WithPolicies adds the policies to the get system params
func (o *GetSystemParams) WithPolicies(policies *bool) *GetSystemParams {
	o.SetPolicies(policies)
	return o
}

// SetPolicies adds the policies to the get system params
func (o *GetSystemParams) SetPolicies(policies *bool) {
	o.Policies = policies
}

// WithStacks adds the stacks to the get system params
func (o *GetSystemParams) WithStacks(stacks *bool) *GetSystemParams {
	o.SetStacks(stacks)
	return o
}

// SetStacks adds the stacks to the get system params
func (o *GetSystemParams) SetStacks(stacks *bool) {
	o.Stacks = stacks
}

// WithSystem adds the system to the get system params
func (o *GetSystemParams) WithSystem(system string) *GetSystemParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the get system params
func (o *GetSystemParams) SetSystem(system string) {
	o.System = system
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Datasources != nil {

		// query param datasources
		var qrDatasources bool

		if o.Datasources != nil {
			qrDatasources = *o.Datasources
		}
		qDatasources := swag.FormatBool(qrDatasources)
		if qDatasources != "" {

			if err := r.SetQueryParam("datasources", qDatasources); err != nil {
				return err
			}
		}
	}

	if o.Errors != nil {

		// query param errors
		var qrErrors bool

		if o.Errors != nil {
			qrErrors = *o.Errors
		}
		qErrors := swag.FormatBool(qrErrors)
		if qErrors != "" {

			if err := r.SetQueryParam("errors", qErrors); err != nil {
				return err
			}
		}
	}

	if o.Modules != nil {

		// query param modules
		var qrModules bool

		if o.Modules != nil {
			qrModules = *o.Modules
		}
		qModules := swag.FormatBool(qrModules)
		if qModules != "" {

			if err := r.SetQueryParam("modules", qModules); err != nil {
				return err
			}
		}
	}

	if o.Policies != nil {

		// query param policies
		var qrPolicies bool

		if o.Policies != nil {
			qrPolicies = *o.Policies
		}
		qPolicies := swag.FormatBool(qrPolicies)
		if qPolicies != "" {

			if err := r.SetQueryParam("policies", qPolicies); err != nil {
				return err
			}
		}
	}

	if o.Stacks != nil {

		// query param stacks
		var qrStacks bool

		if o.Stacks != nil {
			qrStacks = *o.Stacks
		}
		qStacks := swag.FormatBool(qrStacks)
		if qStacks != "" {

			if err := r.SetQueryParam("stacks", qStacks); err != nil {
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

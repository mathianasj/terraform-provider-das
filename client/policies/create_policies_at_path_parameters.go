// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewCreatePoliciesAtPathParams creates a new CreatePoliciesAtPathParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePoliciesAtPathParams() *CreatePoliciesAtPathParams {
	return &CreatePoliciesAtPathParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePoliciesAtPathParamsWithTimeout creates a new CreatePoliciesAtPathParams object
// with the ability to set a timeout on a request.
func NewCreatePoliciesAtPathParamsWithTimeout(timeout time.Duration) *CreatePoliciesAtPathParams {
	return &CreatePoliciesAtPathParams{
		timeout: timeout,
	}
}

// NewCreatePoliciesAtPathParamsWithContext creates a new CreatePoliciesAtPathParams object
// with the ability to set a context for a request.
func NewCreatePoliciesAtPathParamsWithContext(ctx context.Context) *CreatePoliciesAtPathParams {
	return &CreatePoliciesAtPathParams{
		Context: ctx,
	}
}

// NewCreatePoliciesAtPathParamsWithHTTPClient creates a new CreatePoliciesAtPathParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePoliciesAtPathParamsWithHTTPClient(client *http.Client) *CreatePoliciesAtPathParams {
	return &CreatePoliciesAtPathParams{
		HTTPClient: client,
	}
}

/* CreatePoliciesAtPathParams contains all the parameters to send to the API endpoint
   for the create policies at path operation.

   Typically these are written to a http.Request.
*/
type CreatePoliciesAtPathParams struct {

	// Body.
	Body *models.PoliciesV2PoliciesPostRequest

	/* Path.

	   base path
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create policies at path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePoliciesAtPathParams) WithDefaults() *CreatePoliciesAtPathParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create policies at path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePoliciesAtPathParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create policies at path params
func (o *CreatePoliciesAtPathParams) WithTimeout(timeout time.Duration) *CreatePoliciesAtPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policies at path params
func (o *CreatePoliciesAtPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policies at path params
func (o *CreatePoliciesAtPathParams) WithContext(ctx context.Context) *CreatePoliciesAtPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policies at path params
func (o *CreatePoliciesAtPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policies at path params
func (o *CreatePoliciesAtPathParams) WithHTTPClient(client *http.Client) *CreatePoliciesAtPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policies at path params
func (o *CreatePoliciesAtPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create policies at path params
func (o *CreatePoliciesAtPathParams) WithBody(body *models.PoliciesV2PoliciesPostRequest) *CreatePoliciesAtPathParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create policies at path params
func (o *CreatePoliciesAtPathParams) SetBody(body *models.PoliciesV2PoliciesPostRequest) {
	o.Body = body
}

// WithPath adds the path to the create policies at path params
func (o *CreatePoliciesAtPathParams) WithPath(path string) *CreatePoliciesAtPathParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the create policies at path params
func (o *CreatePoliciesAtPathParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePoliciesAtPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

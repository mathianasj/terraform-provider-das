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

// NewPromotePoliciesParams creates a new PromotePoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPromotePoliciesParams() *PromotePoliciesParams {
	return &PromotePoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPromotePoliciesParamsWithTimeout creates a new PromotePoliciesParams object
// with the ability to set a timeout on a request.
func NewPromotePoliciesParamsWithTimeout(timeout time.Duration) *PromotePoliciesParams {
	return &PromotePoliciesParams{
		timeout: timeout,
	}
}

// NewPromotePoliciesParamsWithContext creates a new PromotePoliciesParams object
// with the ability to set a context for a request.
func NewPromotePoliciesParamsWithContext(ctx context.Context) *PromotePoliciesParams {
	return &PromotePoliciesParams{
		Context: ctx,
	}
}

// NewPromotePoliciesParamsWithHTTPClient creates a new PromotePoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPromotePoliciesParamsWithHTTPClient(client *http.Client) *PromotePoliciesParams {
	return &PromotePoliciesParams{
		HTTPClient: client,
	}
}

/* PromotePoliciesParams contains all the parameters to send to the API endpoint
   for the promote policies operation.

   Typically these are written to a http.Request.
*/
type PromotePoliciesParams struct {

	// Body.
	Body *models.PoliciesV2PoliciesPutRequest

	/* Path.

	   base path
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the promote policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PromotePoliciesParams) WithDefaults() *PromotePoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the promote policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PromotePoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the promote policies params
func (o *PromotePoliciesParams) WithTimeout(timeout time.Duration) *PromotePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the promote policies params
func (o *PromotePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the promote policies params
func (o *PromotePoliciesParams) WithContext(ctx context.Context) *PromotePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the promote policies params
func (o *PromotePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the promote policies params
func (o *PromotePoliciesParams) WithHTTPClient(client *http.Client) *PromotePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the promote policies params
func (o *PromotePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the promote policies params
func (o *PromotePoliciesParams) WithBody(body *models.PoliciesV2PoliciesPutRequest) *PromotePoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the promote policies params
func (o *PromotePoliciesParams) SetBody(body *models.PoliciesV2PoliciesPutRequest) {
	o.Body = body
}

// WithPath adds the path to the promote policies params
func (o *PromotePoliciesParams) WithPath(path string) *PromotePoliciesParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the promote policies params
func (o *PromotePoliciesParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *PromotePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

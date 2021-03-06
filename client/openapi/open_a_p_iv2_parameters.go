// Code generated by go-swagger; DO NOT EDIT.

package openapi

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

// NewOpenAPIv2Params creates a new OpenAPIv2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpenAPIv2Params() *OpenAPIv2Params {
	return &OpenAPIv2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpenAPIv2ParamsWithTimeout creates a new OpenAPIv2Params object
// with the ability to set a timeout on a request.
func NewOpenAPIv2ParamsWithTimeout(timeout time.Duration) *OpenAPIv2Params {
	return &OpenAPIv2Params{
		timeout: timeout,
	}
}

// NewOpenAPIv2ParamsWithContext creates a new OpenAPIv2Params object
// with the ability to set a context for a request.
func NewOpenAPIv2ParamsWithContext(ctx context.Context) *OpenAPIv2Params {
	return &OpenAPIv2Params{
		Context: ctx,
	}
}

// NewOpenAPIv2ParamsWithHTTPClient creates a new OpenAPIv2Params object
// with the ability to set a custom HTTPClient for a request.
func NewOpenAPIv2ParamsWithHTTPClient(client *http.Client) *OpenAPIv2Params {
	return &OpenAPIv2Params{
		HTTPClient: client,
	}
}

/* OpenAPIv2Params contains all the parameters to send to the API endpoint
   for the open a p iv2 operation.

   Typically these are written to a http.Request.
*/
type OpenAPIv2Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the open a p iv2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAPIv2Params) WithDefaults() *OpenAPIv2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the open a p iv2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAPIv2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the open a p iv2 params
func (o *OpenAPIv2Params) WithTimeout(timeout time.Duration) *OpenAPIv2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the open a p iv2 params
func (o *OpenAPIv2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the open a p iv2 params
func (o *OpenAPIv2Params) WithContext(ctx context.Context) *OpenAPIv2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the open a p iv2 params
func (o *OpenAPIv2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the open a p iv2 params
func (o *OpenAPIv2Params) WithHTTPClient(client *http.Client) *OpenAPIv2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the open a p iv2 params
func (o *OpenAPIv2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OpenAPIv2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewBulkUploadPlaygroundPoliciesParams creates a new BulkUploadPlaygroundPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkUploadPlaygroundPoliciesParams() *BulkUploadPlaygroundPoliciesParams {
	return &BulkUploadPlaygroundPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkUploadPlaygroundPoliciesParamsWithTimeout creates a new BulkUploadPlaygroundPoliciesParams object
// with the ability to set a timeout on a request.
func NewBulkUploadPlaygroundPoliciesParamsWithTimeout(timeout time.Duration) *BulkUploadPlaygroundPoliciesParams {
	return &BulkUploadPlaygroundPoliciesParams{
		timeout: timeout,
	}
}

// NewBulkUploadPlaygroundPoliciesParamsWithContext creates a new BulkUploadPlaygroundPoliciesParams object
// with the ability to set a context for a request.
func NewBulkUploadPlaygroundPoliciesParamsWithContext(ctx context.Context) *BulkUploadPlaygroundPoliciesParams {
	return &BulkUploadPlaygroundPoliciesParams{
		Context: ctx,
	}
}

// NewBulkUploadPlaygroundPoliciesParamsWithHTTPClient creates a new BulkUploadPlaygroundPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkUploadPlaygroundPoliciesParamsWithHTTPClient(client *http.Client) *BulkUploadPlaygroundPoliciesParams {
	return &BulkUploadPlaygroundPoliciesParams{
		HTTPClient: client,
	}
}

/* BulkUploadPlaygroundPoliciesParams contains all the parameters to send to the API endpoint
   for the bulk upload playground policies operation.

   Typically these are written to a http.Request.
*/
type BulkUploadPlaygroundPoliciesParams struct {

	/* Body.

	   Policy bundle

	   Format: binary
	*/
	Body io.ReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk upload playground policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkUploadPlaygroundPoliciesParams) WithDefaults() *BulkUploadPlaygroundPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk upload playground policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkUploadPlaygroundPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk upload playground policies params
func (o *BulkUploadPlaygroundPoliciesParams) WithTimeout(timeout time.Duration) *BulkUploadPlaygroundPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk upload playground policies params
func (o *BulkUploadPlaygroundPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk upload playground policies params
func (o *BulkUploadPlaygroundPoliciesParams) WithContext(ctx context.Context) *BulkUploadPlaygroundPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk upload playground policies params
func (o *BulkUploadPlaygroundPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk upload playground policies params
func (o *BulkUploadPlaygroundPoliciesParams) WithHTTPClient(client *http.Client) *BulkUploadPlaygroundPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk upload playground policies params
func (o *BulkUploadPlaygroundPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bulk upload playground policies params
func (o *BulkUploadPlaygroundPoliciesParams) WithBody(body io.ReadCloser) *BulkUploadPlaygroundPoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk upload playground policies params
func (o *BulkUploadPlaygroundPoliciesParams) SetBody(body io.ReadCloser) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BulkUploadPlaygroundPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

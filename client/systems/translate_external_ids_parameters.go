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

// NewTranslateExternalIdsParams creates a new TranslateExternalIdsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTranslateExternalIdsParams() *TranslateExternalIdsParams {
	return &TranslateExternalIdsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTranslateExternalIdsParamsWithTimeout creates a new TranslateExternalIdsParams object
// with the ability to set a timeout on a request.
func NewTranslateExternalIdsParamsWithTimeout(timeout time.Duration) *TranslateExternalIdsParams {
	return &TranslateExternalIdsParams{
		timeout: timeout,
	}
}

// NewTranslateExternalIdsParamsWithContext creates a new TranslateExternalIdsParams object
// with the ability to set a context for a request.
func NewTranslateExternalIdsParamsWithContext(ctx context.Context) *TranslateExternalIdsParams {
	return &TranslateExternalIdsParams{
		Context: ctx,
	}
}

// NewTranslateExternalIdsParamsWithHTTPClient creates a new TranslateExternalIdsParams object
// with the ability to set a custom HTTPClient for a request.
func NewTranslateExternalIdsParamsWithHTTPClient(client *http.Client) *TranslateExternalIdsParams {
	return &TranslateExternalIdsParams{
		HTTPClient: client,
	}
}

/* TranslateExternalIdsParams contains all the parameters to send to the API endpoint
   for the translate external ids operation.

   Typically these are written to a http.Request.
*/
type TranslateExternalIdsParams struct {

	// Body.
	Body *models.SystemsV1SystemsTranslateExternalIdsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the translate external ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TranslateExternalIdsParams) WithDefaults() *TranslateExternalIdsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the translate external ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TranslateExternalIdsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the translate external ids params
func (o *TranslateExternalIdsParams) WithTimeout(timeout time.Duration) *TranslateExternalIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the translate external ids params
func (o *TranslateExternalIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the translate external ids params
func (o *TranslateExternalIdsParams) WithContext(ctx context.Context) *TranslateExternalIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the translate external ids params
func (o *TranslateExternalIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the translate external ids params
func (o *TranslateExternalIdsParams) WithHTTPClient(client *http.Client) *TranslateExternalIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the translate external ids params
func (o *TranslateExternalIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the translate external ids params
func (o *TranslateExternalIdsParams) WithBody(body *models.SystemsV1SystemsTranslateExternalIdsRequest) *TranslateExternalIdsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the translate external ids params
func (o *TranslateExternalIdsParams) SetBody(body *models.SystemsV1SystemsTranslateExternalIdsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TranslateExternalIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

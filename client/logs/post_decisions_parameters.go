// Code generated by go-swagger; DO NOT EDIT.

package logs

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

// NewPostDecisionsParams creates a new PostDecisionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDecisionsParams() *PostDecisionsParams {
	return &PostDecisionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDecisionsParamsWithTimeout creates a new PostDecisionsParams object
// with the ability to set a timeout on a request.
func NewPostDecisionsParamsWithTimeout(timeout time.Duration) *PostDecisionsParams {
	return &PostDecisionsParams{
		timeout: timeout,
	}
}

// NewPostDecisionsParamsWithContext creates a new PostDecisionsParams object
// with the ability to set a context for a request.
func NewPostDecisionsParamsWithContext(ctx context.Context) *PostDecisionsParams {
	return &PostDecisionsParams{
		Context: ctx,
	}
}

// NewPostDecisionsParamsWithHTTPClient creates a new PostDecisionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDecisionsParamsWithHTTPClient(client *http.Client) *PostDecisionsParams {
	return &PostDecisionsParams{
		HTTPClient: client,
	}
}

/* PostDecisionsParams contains all the parameters to send to the API endpoint
   for the post decisions operation.

   Typically these are written to a http.Request.
*/
type PostDecisionsParams struct {

	// Body.
	Body []models.MetaV1RequestObject

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post decisions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDecisionsParams) WithDefaults() *PostDecisionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post decisions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDecisionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post decisions params
func (o *PostDecisionsParams) WithTimeout(timeout time.Duration) *PostDecisionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post decisions params
func (o *PostDecisionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post decisions params
func (o *PostDecisionsParams) WithContext(ctx context.Context) *PostDecisionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post decisions params
func (o *PostDecisionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post decisions params
func (o *PostDecisionsParams) WithHTTPClient(client *http.Client) *PostDecisionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post decisions params
func (o *PostDecisionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post decisions params
func (o *PostDecisionsParams) WithBody(body []models.MetaV1RequestObject) *PostDecisionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post decisions params
func (o *PostDecisionsParams) SetBody(body []models.MetaV1RequestObject) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDecisionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

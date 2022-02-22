// Code generated by go-swagger; DO NOT EDIT.

package workspace

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

// NewKafkaVerifyConfigParams creates a new KafkaVerifyConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKafkaVerifyConfigParams() *KafkaVerifyConfigParams {
	return &KafkaVerifyConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKafkaVerifyConfigParamsWithTimeout creates a new KafkaVerifyConfigParams object
// with the ability to set a timeout on a request.
func NewKafkaVerifyConfigParamsWithTimeout(timeout time.Duration) *KafkaVerifyConfigParams {
	return &KafkaVerifyConfigParams{
		timeout: timeout,
	}
}

// NewKafkaVerifyConfigParamsWithContext creates a new KafkaVerifyConfigParams object
// with the ability to set a context for a request.
func NewKafkaVerifyConfigParamsWithContext(ctx context.Context) *KafkaVerifyConfigParams {
	return &KafkaVerifyConfigParams{
		Context: ctx,
	}
}

// NewKafkaVerifyConfigParamsWithHTTPClient creates a new KafkaVerifyConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewKafkaVerifyConfigParamsWithHTTPClient(client *http.Client) *KafkaVerifyConfigParams {
	return &KafkaVerifyConfigParams{
		HTTPClient: client,
	}
}

/* KafkaVerifyConfigParams contains all the parameters to send to the API endpoint
   for the kafka verify config operation.

   Typically these are written to a http.Request.
*/
type KafkaVerifyConfigParams struct {

	// Body.
	Body *models.WorkspaceV1KafkaConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kafka verify config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafkaVerifyConfigParams) WithDefaults() *KafkaVerifyConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kafka verify config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KafkaVerifyConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kafka verify config params
func (o *KafkaVerifyConfigParams) WithTimeout(timeout time.Duration) *KafkaVerifyConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kafka verify config params
func (o *KafkaVerifyConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kafka verify config params
func (o *KafkaVerifyConfigParams) WithContext(ctx context.Context) *KafkaVerifyConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kafka verify config params
func (o *KafkaVerifyConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kafka verify config params
func (o *KafkaVerifyConfigParams) WithHTTPClient(client *http.Client) *KafkaVerifyConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kafka verify config params
func (o *KafkaVerifyConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kafka verify config params
func (o *KafkaVerifyConfigParams) WithBody(body *models.WorkspaceV1KafkaConfig) *KafkaVerifyConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kafka verify config params
func (o *KafkaVerifyConfigParams) SetBody(body *models.WorkspaceV1KafkaConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *KafkaVerifyConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

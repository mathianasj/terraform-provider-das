// Code generated by go-swagger; DO NOT EDIT.

package logreplay

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

// NewRunLogReplayParams creates a new RunLogReplayParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRunLogReplayParams() *RunLogReplayParams {
	return &RunLogReplayParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRunLogReplayParamsWithTimeout creates a new RunLogReplayParams object
// with the ability to set a timeout on a request.
func NewRunLogReplayParamsWithTimeout(timeout time.Duration) *RunLogReplayParams {
	return &RunLogReplayParams{
		timeout: timeout,
	}
}

// NewRunLogReplayParamsWithContext creates a new RunLogReplayParams object
// with the ability to set a context for a request.
func NewRunLogReplayParamsWithContext(ctx context.Context) *RunLogReplayParams {
	return &RunLogReplayParams{
		Context: ctx,
	}
}

// NewRunLogReplayParamsWithHTTPClient creates a new RunLogReplayParams object
// with the ability to set a custom HTTPClient for a request.
func NewRunLogReplayParamsWithHTTPClient(client *http.Client) *RunLogReplayParams {
	return &RunLogReplayParams{
		HTTPClient: client,
	}
}

/* RunLogReplayParams contains all the parameters to send to the API endpoint
   for the run log replay operation.

   Typically these are written to a http.Request.
*/
type RunLogReplayParams struct {

	// Body.
	Body *models.LogreplayV1ReplayRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the run log replay params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunLogReplayParams) WithDefaults() *RunLogReplayParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the run log replay params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunLogReplayParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the run log replay params
func (o *RunLogReplayParams) WithTimeout(timeout time.Duration) *RunLogReplayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run log replay params
func (o *RunLogReplayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run log replay params
func (o *RunLogReplayParams) WithContext(ctx context.Context) *RunLogReplayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run log replay params
func (o *RunLogReplayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run log replay params
func (o *RunLogReplayParams) WithHTTPClient(client *http.Client) *RunLogReplayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run log replay params
func (o *RunLogReplayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the run log replay params
func (o *RunLogReplayParams) WithBody(body *models.LogreplayV1ReplayRequest) *RunLogReplayParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the run log replay params
func (o *RunLogReplayParams) SetBody(body *models.LogreplayV1ReplayRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RunLogReplayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

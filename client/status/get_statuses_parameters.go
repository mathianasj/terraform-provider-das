// Code generated by go-swagger; DO NOT EDIT.

package status

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

// NewGetStatusesParams creates a new GetStatusesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStatusesParams() *GetStatusesParams {
	return &GetStatusesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatusesParamsWithTimeout creates a new GetStatusesParams object
// with the ability to set a timeout on a request.
func NewGetStatusesParamsWithTimeout(timeout time.Duration) *GetStatusesParams {
	return &GetStatusesParams{
		timeout: timeout,
	}
}

// NewGetStatusesParamsWithContext creates a new GetStatusesParams object
// with the ability to set a context for a request.
func NewGetStatusesParamsWithContext(ctx context.Context) *GetStatusesParams {
	return &GetStatusesParams{
		Context: ctx,
	}
}

// NewGetStatusesParamsWithHTTPClient creates a new GetStatusesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStatusesParamsWithHTTPClient(client *http.Client) *GetStatusesParams {
	return &GetStatusesParams{
		HTTPClient: client,
	}
}

/* GetStatusesParams contains all the parameters to send to the API endpoint
   for the get statuses operation.

   Typically these are written to a http.Request.
*/
type GetStatusesParams struct {

	/* System.

	   return only statuses for one or more system ID
	*/
	System *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatusesParams) WithDefaults() *GetStatusesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatusesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get statuses params
func (o *GetStatusesParams) WithTimeout(timeout time.Duration) *GetStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get statuses params
func (o *GetStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get statuses params
func (o *GetStatusesParams) WithContext(ctx context.Context) *GetStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get statuses params
func (o *GetStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get statuses params
func (o *GetStatusesParams) WithHTTPClient(client *http.Client) *GetStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get statuses params
func (o *GetStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSystem adds the system to the get statuses params
func (o *GetStatusesParams) WithSystem(system *string) *GetStatusesParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the get statuses params
func (o *GetStatusesParams) SetSystem(system *string) {
	o.System = system
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.System != nil {

		// query param system
		var qrSystem string

		if o.System != nil {
			qrSystem = *o.System
		}
		qSystem := qrSystem
		if qSystem != "" {

			if err := r.SetQueryParam("system", qSystem); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package datasources

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

	"github.com/mathianasj/terraform-provider-das/models"
)

// NewExecuteDatasourceParams creates a new ExecuteDatasourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecuteDatasourceParams() *ExecuteDatasourceParams {
	return &ExecuteDatasourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteDatasourceParamsWithTimeout creates a new ExecuteDatasourceParams object
// with the ability to set a timeout on a request.
func NewExecuteDatasourceParamsWithTimeout(timeout time.Duration) *ExecuteDatasourceParams {
	return &ExecuteDatasourceParams{
		timeout: timeout,
	}
}

// NewExecuteDatasourceParamsWithContext creates a new ExecuteDatasourceParams object
// with the ability to set a context for a request.
func NewExecuteDatasourceParamsWithContext(ctx context.Context) *ExecuteDatasourceParams {
	return &ExecuteDatasourceParams{
		Context: ctx,
	}
}

// NewExecuteDatasourceParamsWithHTTPClient creates a new ExecuteDatasourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecuteDatasourceParamsWithHTTPClient(client *http.Client) *ExecuteDatasourceParams {
	return &ExecuteDatasourceParams{
		HTTPClient: client,
	}
}

/* ExecuteDatasourceParams contains all the parameters to send to the API endpoint
   for the execute datasource operation.

   Typically these are written to a http.Request.
*/
type ExecuteDatasourceParams struct {

	// Body.
	Body models.DatasourcesV1DatasourcesPutRequest

	/* Datasource.

	   Data source ID
	*/
	Datasource string

	/* Execute.

	   Execute data source
	*/
	Execute *bool

	/* Preview.

	   Preview data source
	*/
	Preview *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute datasource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteDatasourceParams) WithDefaults() *ExecuteDatasourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute datasource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteDatasourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execute datasource params
func (o *ExecuteDatasourceParams) WithTimeout(timeout time.Duration) *ExecuteDatasourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute datasource params
func (o *ExecuteDatasourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute datasource params
func (o *ExecuteDatasourceParams) WithContext(ctx context.Context) *ExecuteDatasourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute datasource params
func (o *ExecuteDatasourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute datasource params
func (o *ExecuteDatasourceParams) WithHTTPClient(client *http.Client) *ExecuteDatasourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute datasource params
func (o *ExecuteDatasourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the execute datasource params
func (o *ExecuteDatasourceParams) WithBody(body models.DatasourcesV1DatasourcesPutRequest) *ExecuteDatasourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the execute datasource params
func (o *ExecuteDatasourceParams) SetBody(body models.DatasourcesV1DatasourcesPutRequest) {
	o.Body = body
}

// WithDatasource adds the datasource to the execute datasource params
func (o *ExecuteDatasourceParams) WithDatasource(datasource string) *ExecuteDatasourceParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the execute datasource params
func (o *ExecuteDatasourceParams) SetDatasource(datasource string) {
	o.Datasource = datasource
}

// WithExecute adds the execute to the execute datasource params
func (o *ExecuteDatasourceParams) WithExecute(execute *bool) *ExecuteDatasourceParams {
	o.SetExecute(execute)
	return o
}

// SetExecute adds the execute to the execute datasource params
func (o *ExecuteDatasourceParams) SetExecute(execute *bool) {
	o.Execute = execute
}

// WithPreview adds the preview to the execute datasource params
func (o *ExecuteDatasourceParams) WithPreview(preview *bool) *ExecuteDatasourceParams {
	o.SetPreview(preview)
	return o
}

// SetPreview adds the preview to the execute datasource params
func (o *ExecuteDatasourceParams) SetPreview(preview *bool) {
	o.Preview = preview
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteDatasourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param datasource
	if err := r.SetPathParam("datasource", o.Datasource); err != nil {
		return err
	}

	if o.Execute != nil {

		// query param execute
		var qrExecute bool

		if o.Execute != nil {
			qrExecute = *o.Execute
		}
		qExecute := swag.FormatBool(qrExecute)
		if qExecute != "" {

			if err := r.SetQueryParam("execute", qExecute); err != nil {
				return err
			}
		}
	}

	if o.Preview != nil {

		// query param preview
		var qrPreview bool

		if o.Preview != nil {
			qrPreview = *o.Preview
		}
		qPreview := swag.FormatBool(qrPreview)
		if qPreview != "" {

			if err := r.SetQueryParam("preview", qPreview); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
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

	"github.com/mathianasj/terraform-provider-das/models"
)

// NewUpdateStatusWithPartitionParams creates a new UpdateStatusWithPartitionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateStatusWithPartitionParams() *UpdateStatusWithPartitionParams {
	return &UpdateStatusWithPartitionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateStatusWithPartitionParamsWithTimeout creates a new UpdateStatusWithPartitionParams object
// with the ability to set a timeout on a request.
func NewUpdateStatusWithPartitionParamsWithTimeout(timeout time.Duration) *UpdateStatusWithPartitionParams {
	return &UpdateStatusWithPartitionParams{
		timeout: timeout,
	}
}

// NewUpdateStatusWithPartitionParamsWithContext creates a new UpdateStatusWithPartitionParams object
// with the ability to set a context for a request.
func NewUpdateStatusWithPartitionParamsWithContext(ctx context.Context) *UpdateStatusWithPartitionParams {
	return &UpdateStatusWithPartitionParams{
		Context: ctx,
	}
}

// NewUpdateStatusWithPartitionParamsWithHTTPClient creates a new UpdateStatusWithPartitionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateStatusWithPartitionParamsWithHTTPClient(client *http.Client) *UpdateStatusWithPartitionParams {
	return &UpdateStatusWithPartitionParams{
		HTTPClient: client,
	}
}

/* UpdateStatusWithPartitionParams contains all the parameters to send to the API endpoint
   for the update status with partition operation.

   Typically these are written to a http.Request.
*/
type UpdateStatusWithPartitionParams struct {

	// Body.
	Body models.StatusV1AgentStatus

	/* Partition.

	   partition name. Currently not used
	*/
	Partition string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update status with partition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateStatusWithPartitionParams) WithDefaults() *UpdateStatusWithPartitionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update status with partition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateStatusWithPartitionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update status with partition params
func (o *UpdateStatusWithPartitionParams) WithTimeout(timeout time.Duration) *UpdateStatusWithPartitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update status with partition params
func (o *UpdateStatusWithPartitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update status with partition params
func (o *UpdateStatusWithPartitionParams) WithContext(ctx context.Context) *UpdateStatusWithPartitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update status with partition params
func (o *UpdateStatusWithPartitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update status with partition params
func (o *UpdateStatusWithPartitionParams) WithHTTPClient(client *http.Client) *UpdateStatusWithPartitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update status with partition params
func (o *UpdateStatusWithPartitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update status with partition params
func (o *UpdateStatusWithPartitionParams) WithBody(body models.StatusV1AgentStatus) *UpdateStatusWithPartitionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update status with partition params
func (o *UpdateStatusWithPartitionParams) SetBody(body models.StatusV1AgentStatus) {
	o.Body = body
}

// WithPartition adds the partition to the update status with partition params
func (o *UpdateStatusWithPartitionParams) WithPartition(partition string) *UpdateStatusWithPartitionParams {
	o.SetPartition(partition)
	return o
}

// SetPartition adds the partition to the update status with partition params
func (o *UpdateStatusWithPartitionParams) SetPartition(partition string) {
	o.Partition = partition
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateStatusWithPartitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param partition
	if err := r.SetPathParam("partition", o.Partition); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
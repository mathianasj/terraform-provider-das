// Code generated by go-swagger; DO NOT EDIT.

package stacks

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

// NewValidateStackTestsParams creates a new ValidateStackTestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateStackTestsParams() *ValidateStackTestsParams {
	return &ValidateStackTestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateStackTestsParamsWithTimeout creates a new ValidateStackTestsParams object
// with the ability to set a timeout on a request.
func NewValidateStackTestsParamsWithTimeout(timeout time.Duration) *ValidateStackTestsParams {
	return &ValidateStackTestsParams{
		timeout: timeout,
	}
}

// NewValidateStackTestsParamsWithContext creates a new ValidateStackTestsParams object
// with the ability to set a context for a request.
func NewValidateStackTestsParamsWithContext(ctx context.Context) *ValidateStackTestsParams {
	return &ValidateStackTestsParams{
		Context: ctx,
	}
}

// NewValidateStackTestsParamsWithHTTPClient creates a new ValidateStackTestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateStackTestsParamsWithHTTPClient(client *http.Client) *ValidateStackTestsParams {
	return &ValidateStackTestsParams{
		HTTPClient: client,
	}
}

/* ValidateStackTestsParams contains all the parameters to send to the API endpoint
   for the validate stack tests operation.

   Typically these are written to a http.Request.
*/
type ValidateStackTestsParams struct {

	// Body.
	Body *models.StacksV1StacksTestsRequest

	/* Stack.

	   stack id
	*/
	Stack string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate stack tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateStackTestsParams) WithDefaults() *ValidateStackTestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate stack tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateStackTestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate stack tests params
func (o *ValidateStackTestsParams) WithTimeout(timeout time.Duration) *ValidateStackTestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate stack tests params
func (o *ValidateStackTestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate stack tests params
func (o *ValidateStackTestsParams) WithContext(ctx context.Context) *ValidateStackTestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate stack tests params
func (o *ValidateStackTestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate stack tests params
func (o *ValidateStackTestsParams) WithHTTPClient(client *http.Client) *ValidateStackTestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate stack tests params
func (o *ValidateStackTestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate stack tests params
func (o *ValidateStackTestsParams) WithBody(body *models.StacksV1StacksTestsRequest) *ValidateStackTestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate stack tests params
func (o *ValidateStackTestsParams) SetBody(body *models.StacksV1StacksTestsRequest) {
	o.Body = body
}

// WithStack adds the stack to the validate stack tests params
func (o *ValidateStackTestsParams) WithStack(stack string) *ValidateStackTestsParams {
	o.SetStack(stack)
	return o
}

// SetStack adds the stack to the validate stack tests params
func (o *ValidateStackTestsParams) SetStack(stack string) {
	o.Stack = stack
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateStackTestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param stack
	if err := r.SetPathParam("stack", o.Stack); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

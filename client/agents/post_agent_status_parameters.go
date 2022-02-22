// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewPostAgentStatusParams creates a new PostAgentStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAgentStatusParams() *PostAgentStatusParams {
	return &PostAgentStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAgentStatusParamsWithTimeout creates a new PostAgentStatusParams object
// with the ability to set a timeout on a request.
func NewPostAgentStatusParamsWithTimeout(timeout time.Duration) *PostAgentStatusParams {
	return &PostAgentStatusParams{
		timeout: timeout,
	}
}

// NewPostAgentStatusParamsWithContext creates a new PostAgentStatusParams object
// with the ability to set a context for a request.
func NewPostAgentStatusParamsWithContext(ctx context.Context) *PostAgentStatusParams {
	return &PostAgentStatusParams{
		Context: ctx,
	}
}

// NewPostAgentStatusParamsWithHTTPClient creates a new PostAgentStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAgentStatusParamsWithHTTPClient(client *http.Client) *PostAgentStatusParams {
	return &PostAgentStatusParams{
		HTTPClient: client,
	}
}

/* PostAgentStatusParams contains all the parameters to send to the API endpoint
   for the post agent status operation.

   Typically these are written to a http.Request.
*/
type PostAgentStatusParams struct {

	// Body.
	Body models.StatusV1AgentStatus

	/* Kind.

	   agent kind such as "agents", "datasources", "slps", "exporters"
	*/
	Kind string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post agent status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAgentStatusParams) WithDefaults() *PostAgentStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post agent status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAgentStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post agent status params
func (o *PostAgentStatusParams) WithTimeout(timeout time.Duration) *PostAgentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post agent status params
func (o *PostAgentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post agent status params
func (o *PostAgentStatusParams) WithContext(ctx context.Context) *PostAgentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post agent status params
func (o *PostAgentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post agent status params
func (o *PostAgentStatusParams) WithHTTPClient(client *http.Client) *PostAgentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post agent status params
func (o *PostAgentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post agent status params
func (o *PostAgentStatusParams) WithBody(body models.StatusV1AgentStatus) *PostAgentStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post agent status params
func (o *PostAgentStatusParams) SetBody(body models.StatusV1AgentStatus) {
	o.Body = body
}

// WithKind adds the kind to the post agent status params
func (o *PostAgentStatusParams) WithKind(kind string) *PostAgentStatusParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the post agent status params
func (o *PostAgentStatusParams) SetKind(kind string) {
	o.Kind = kind
}

// WriteToRequest writes these params to a swagger request
func (o *PostAgentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

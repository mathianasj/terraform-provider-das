// Code generated by go-swagger; DO NOT EDIT.

package bundles

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

// NewGetPolicyBundle2Params creates a new GetPolicyBundle2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPolicyBundle2Params() *GetPolicyBundle2Params {
	return &GetPolicyBundle2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyBundle2ParamsWithTimeout creates a new GetPolicyBundle2Params object
// with the ability to set a timeout on a request.
func NewGetPolicyBundle2ParamsWithTimeout(timeout time.Duration) *GetPolicyBundle2Params {
	return &GetPolicyBundle2Params{
		timeout: timeout,
	}
}

// NewGetPolicyBundle2ParamsWithContext creates a new GetPolicyBundle2Params object
// with the ability to set a context for a request.
func NewGetPolicyBundle2ParamsWithContext(ctx context.Context) *GetPolicyBundle2Params {
	return &GetPolicyBundle2Params{
		Context: ctx,
	}
}

// NewGetPolicyBundle2ParamsWithHTTPClient creates a new GetPolicyBundle2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetPolicyBundle2ParamsWithHTTPClient(client *http.Client) *GetPolicyBundle2Params {
	return &GetPolicyBundle2Params{
		HTTPClient: client,
	}
}

/* GetPolicyBundle2Params contains all the parameters to send to the API endpoint
   for the get policy bundle2 operation.

   Typically these are written to a http.Request.
*/
type GetPolicyBundle2Params struct {

	/* IfNoneMatch.

	   etag
	*/
	IfNoneMatch *string

	/* EvalPath.

	   path to partial evaluation
	*/
	EvalPath *string

	/* Policy.

	   policy name
	*/
	Policy string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get policy bundle2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyBundle2Params) WithDefaults() *GetPolicyBundle2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get policy bundle2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPolicyBundle2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get policy bundle2 params
func (o *GetPolicyBundle2Params) WithTimeout(timeout time.Duration) *GetPolicyBundle2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy bundle2 params
func (o *GetPolicyBundle2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy bundle2 params
func (o *GetPolicyBundle2Params) WithContext(ctx context.Context) *GetPolicyBundle2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy bundle2 params
func (o *GetPolicyBundle2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy bundle2 params
func (o *GetPolicyBundle2Params) WithHTTPClient(client *http.Client) *GetPolicyBundle2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy bundle2 params
func (o *GetPolicyBundle2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get policy bundle2 params
func (o *GetPolicyBundle2Params) WithIfNoneMatch(ifNoneMatch *string) *GetPolicyBundle2Params {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get policy bundle2 params
func (o *GetPolicyBundle2Params) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithEvalPath adds the evalPath to the get policy bundle2 params
func (o *GetPolicyBundle2Params) WithEvalPath(evalPath *string) *GetPolicyBundle2Params {
	o.SetEvalPath(evalPath)
	return o
}

// SetEvalPath adds the evalPath to the get policy bundle2 params
func (o *GetPolicyBundle2Params) SetEvalPath(evalPath *string) {
	o.EvalPath = evalPath
}

// WithPolicy adds the policy to the get policy bundle2 params
func (o *GetPolicyBundle2Params) WithPolicy(policy string) *GetPolicyBundle2Params {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the get policy bundle2 params
func (o *GetPolicyBundle2Params) SetPolicy(policy string) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyBundle2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}
	}

	if o.EvalPath != nil {

		// query param eval_path
		var qrEvalPath string

		if o.EvalPath != nil {
			qrEvalPath = *o.EvalPath
		}
		qEvalPath := qrEvalPath
		if qEvalPath != "" {

			if err := r.SetQueryParam("eval_path", qEvalPath); err != nil {
				return err
			}
		}
	}

	// path param policy
	if err := r.SetPathParam("policy", o.Policy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

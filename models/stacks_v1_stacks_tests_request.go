// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StacksV1StacksTestsRequest stacks v1 stacks tests request
//
// swagger:model stacks.v1.StacksTestsRequest
type StacksV1StacksTestsRequest struct {

	// draft policies to be used for 'new' violations computation (path => rego)
	Drafts map[string]string `json:"drafts,omitempty"`

	// validation mode. One of (delta, all, delta-count, all-count)
	Mode *string `json:"mode,omitempty"`

	// policy type to narrow the monitor policy search (e.g. validating, mutating). Default (empty string or missing) is to run all monitoring policies
	PolicyType string `json:"policy_type,omitempty"`
}

// Validate validates this stacks v1 stacks tests request
func (m *StacksV1StacksTestsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stacks v1 stacks tests request based on context it is used
func (m *StacksV1StacksTestsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StacksV1StacksTestsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksV1StacksTestsRequest) UnmarshalBinary(b []byte) error {
	var res StacksV1StacksTestsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

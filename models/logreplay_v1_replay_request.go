// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogreplayV1ReplayRequest logreplay v1 replay request
//
// swagger:model logreplay.v1.ReplayRequest
type LogreplayV1ReplayRequest struct {

	// do not compare decisions by system-type-dependent significant fields
	CompareFullResults *bool `json:"compare_full_results,omitempty"`

	// list of JSON Patches to apply to the data namespace
	DataPatches []JSONJSONPatchSpec `json:"data_patches"`

	// list of JSON Patches to apply to the decisions before they evaluated
	DecisionPatches []JSONJSONPatchSpec `json:"decision_patches"`

	// signals that decisions having the same inputs, data and revision always evaluate to the same result and therefore can be cached
	DeterministicPolicies *bool `json:"deterministic_policies,omitempty"`

	// maximum replay duration (e.g. "20s")
	Duration string `json:"duration,omitempty"`

	// maximum number of samples to return
	MaxSamples int32 `json:"max_samples,omitempty"`

	// mock http.send and opa.runtime builtins
	Mocks *SystemsV1BuiltinMocks `json:"mocks,omitempty"`

	// modified rego policies (path => rego content)
	Policies map[string]string `json:"policies,omitempty"`

	// list of scopes to narrow the decision search
	Scope []*LogreplayV1ReplayScope `json:"scope"`

	// list of batch IDs to skip
	SkipBatches []string `json:"skip_batches"`
}

// Validate validates this logreplay v1 replay request
func (m *LogreplayV1ReplayRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogreplayV1ReplayRequest) validateMocks(formats strfmt.Registry) error {
	if swag.IsZero(m.Mocks) { // not required
		return nil
	}

	if m.Mocks != nil {
		if err := m.Mocks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mocks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mocks")
			}
			return err
		}
	}

	return nil
}

func (m *LogreplayV1ReplayRequest) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	for i := 0; i < len(m.Scope); i++ {
		if swag.IsZero(m.Scope[i]) { // not required
			continue
		}

		if m.Scope[i] != nil {
			if err := m.Scope[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scope" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scope" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this logreplay v1 replay request based on the context it is used
func (m *LogreplayV1ReplayRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogreplayV1ReplayRequest) contextValidateMocks(ctx context.Context, formats strfmt.Registry) error {

	if m.Mocks != nil {
		if err := m.Mocks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mocks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mocks")
			}
			return err
		}
	}

	return nil
}

func (m *LogreplayV1ReplayRequest) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Scope); i++ {

		if m.Scope[i] != nil {
			if err := m.Scope[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scope" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scope" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogreplayV1ReplayRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogreplayV1ReplayRequest) UnmarshalBinary(b []byte) error {
	var res LogreplayV1ReplayRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PoliciesV2PoliciesGetResponse policies v2 policies get response
//
// swagger:model policies.v2.PoliciesGetResponse
type PoliciesV2PoliciesGetResponse struct {

	// policies
	// Required: true
	Policies map[string]PoliciesV2PoliciesUnderRoot `json:"policies"`

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this policies v2 policies get response
func (m *PoliciesV2PoliciesGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoliciesV2PoliciesGetResponse) validatePolicies(formats strfmt.Registry) error {

	if err := validate.Required("policies", "body", m.Policies); err != nil {
		return err
	}

	for k := range m.Policies {

		if err := validate.Required("policies"+"."+k, "body", m.Policies[k]); err != nil {
			return err
		}
		if val, ok := m.Policies[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this policies v2 policies get response based on the context it is used
func (m *PoliciesV2PoliciesGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoliciesV2PoliciesGetResponse) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("policies", "body", m.Policies); err != nil {
		return err
	}

	for k := range m.Policies {

		if val, ok := m.Policies[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoliciesV2PoliciesGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoliciesV2PoliciesGetResponse) UnmarshalBinary(b []byte) error {
	var res PoliciesV2PoliciesGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

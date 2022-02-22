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

// AuthzV2RoleBindingsPostSubjectsResponse authz v2 role bindings post subjects response
//
// swagger:model authz.v2.RoleBindingsPostSubjectsResponse
type AuthzV2RoleBindingsPostSubjectsResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`

	// rolebinding
	// Required: true
	Rolebinding *AuthzV2RoleBindingConfig `json:"rolebinding"`
}

// Validate validates this authz v2 role bindings post subjects response
func (m *AuthzV2RoleBindingsPostSubjectsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRolebinding(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthzV2RoleBindingsPostSubjectsResponse) validateRolebinding(formats strfmt.Registry) error {

	if err := validate.Required("rolebinding", "body", m.Rolebinding); err != nil {
		return err
	}

	if m.Rolebinding != nil {
		if err := m.Rolebinding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rolebinding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rolebinding")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this authz v2 role bindings post subjects response based on the context it is used
func (m *AuthzV2RoleBindingsPostSubjectsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRolebinding(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthzV2RoleBindingsPostSubjectsResponse) contextValidateRolebinding(ctx context.Context, formats strfmt.Registry) error {

	if m.Rolebinding != nil {
		if err := m.Rolebinding.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rolebinding")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rolebinding")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthzV2RoleBindingsPostSubjectsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthzV2RoleBindingsPostSubjectsResponse) UnmarshalBinary(b []byte) error {
	var res AuthzV2RoleBindingsPostSubjectsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

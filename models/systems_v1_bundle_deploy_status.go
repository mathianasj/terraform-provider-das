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

// SystemsV1BundleDeployStatus systems v1 bundle deploy status
//
// swagger:model systems.v1.BundleDeployStatus
type SystemsV1BundleDeployStatus struct {

	// build errors (per bundle id) when building the latest bundle, if any
	// Required: true
	BuildErrors map[string]string `json:"build_errors"`

	// primary bundle to activate
	Primary *SystemsV1BundleActivation `json:"primary,omitempty"`
}

// Validate validates this systems v1 bundle deploy status
func (m *SystemsV1BundleDeployStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1BundleDeployStatus) validateBuildErrors(formats strfmt.Registry) error {

	if err := validate.Required("build_errors", "body", m.BuildErrors); err != nil {
		return err
	}

	return nil
}

func (m *SystemsV1BundleDeployStatus) validatePrimary(formats strfmt.Registry) error {
	if swag.IsZero(m.Primary) { // not required
		return nil
	}

	if m.Primary != nil {
		if err := m.Primary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this systems v1 bundle deploy status based on the context it is used
func (m *SystemsV1BundleDeployStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrimary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1BundleDeployStatus) contextValidatePrimary(ctx context.Context, formats strfmt.Registry) error {

	if m.Primary != nil {
		if err := m.Primary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1BundleDeployStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1BundleDeployStatus) UnmarshalBinary(b []byte) error {
	var res SystemsV1BundleDeployStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

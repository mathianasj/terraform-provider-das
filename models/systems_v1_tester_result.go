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

// SystemsV1TesterResult systems v1 tester result
//
// swagger:model systems.v1.TesterResult
type SystemsV1TesterResult struct {

	// duration
	// Required: true
	Duration *int64 `json:"duration"`

	// error
	Error string `json:"error,omitempty"`

	// fail
	// Required: true
	Fail *bool `json:"fail"`

	// failed at
	FailedAt *AstExpr `json:"failed_at,omitempty"`

	// location
	// Required: true
	Location *LocationLocation `json:"location"`

	// name
	// Required: true
	Name *string `json:"name"`

	// package
	// Required: true
	Package *string `json:"package"`

	// presence changed
	PresenceChanged bool `json:"presence_changed,omitempty"`
}

// Validate validates this systems v1 tester result
func (m *SystemsV1TesterResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1TesterResult) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *SystemsV1TesterResult) validateFail(formats strfmt.Registry) error {

	if err := validate.Required("fail", "body", m.Fail); err != nil {
		return err
	}

	return nil
}

func (m *SystemsV1TesterResult) validateFailedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.FailedAt) { // not required
		return nil
	}

	if m.FailedAt != nil {
		if err := m.FailedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failed_at")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("failed_at")
			}
			return err
		}
	}

	return nil
}

func (m *SystemsV1TesterResult) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *SystemsV1TesterResult) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SystemsV1TesterResult) validatePackage(formats strfmt.Registry) error {

	if err := validate.Required("package", "body", m.Package); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this systems v1 tester result based on the context it is used
func (m *SystemsV1TesterResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFailedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1TesterResult) contextValidateFailedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.FailedAt != nil {
		if err := m.FailedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failed_at")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("failed_at")
			}
			return err
		}
	}

	return nil
}

func (m *SystemsV1TesterResult) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1TesterResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1TesterResult) UnmarshalBinary(b []byte) error {
	var res SystemsV1TesterResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

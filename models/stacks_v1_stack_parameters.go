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

// StacksV1StackParameters stacks v1 stack parameters
//
// swagger:model stacks.v1.StackParameters
type StacksV1StackParameters struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`

	// read only
	// Required: true
	ReadOnly *bool `json:"read_only"`

	// source control
	SourceControl *StacksV1SourceControlConfig `json:"source_control,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`

	// stack type parameter values (for template.* types)
	TypeParameters interface{} `json:"type_parameters,omitempty"`
}

// Validate validates this stacks v1 stack parameters
func (m *StacksV1StackParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceControl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1StackParameters) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *StacksV1StackParameters) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StacksV1StackParameters) validateReadOnly(formats strfmt.Registry) error {

	if err := validate.Required("read_only", "body", m.ReadOnly); err != nil {
		return err
	}

	return nil
}

func (m *StacksV1StackParameters) validateSourceControl(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceControl) { // not required
		return nil
	}

	if m.SourceControl != nil {
		if err := m.SourceControl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_control")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_control")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StackParameters) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this stacks v1 stack parameters based on the context it is used
func (m *StacksV1StackParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceControl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1StackParameters) contextValidateSourceControl(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceControl != nil {
		if err := m.SourceControl.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_control")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_control")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StacksV1StackParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksV1StackParameters) UnmarshalBinary(b []byte) error {
	var res StacksV1StackParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
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

// StacksV1SourceControlConfig stacks v1 source control config
//
// swagger:model stacks.v1.SourceControlConfig
type StacksV1SourceControlConfig struct {

	// origin
	Origin *GitV1GitRepoConfig `json:"origin,omitempty"`

	// stack origin
	StackOrigin *GitV1GitRepoConfig `json:"stack_origin,omitempty"`

	// use workspace settings
	// Required: true
	UseWorkspaceSettings *bool `json:"use_workspace_settings"`
}

// Validate validates this stacks v1 source control config
func (m *StacksV1SourceControlConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseWorkspaceSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1SourceControlConfig) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1SourceControlConfig) validateStackOrigin(formats strfmt.Registry) error {
	if swag.IsZero(m.StackOrigin) { // not required
		return nil
	}

	if m.StackOrigin != nil {
		if err := m.StackOrigin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack_origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stack_origin")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1SourceControlConfig) validateUseWorkspaceSettings(formats strfmt.Registry) error {

	if err := validate.Required("use_workspace_settings", "body", m.UseWorkspaceSettings); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this stacks v1 source control config based on the context it is used
func (m *StacksV1SourceControlConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStackOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1SourceControlConfig) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if m.Origin != nil {
		if err := m.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1SourceControlConfig) contextValidateStackOrigin(ctx context.Context, formats strfmt.Registry) error {

	if m.StackOrigin != nil {
		if err := m.StackOrigin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stack_origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stack_origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StacksV1SourceControlConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksV1SourceControlConfig) UnmarshalBinary(b []byte) error {
	var res StacksV1SourceControlConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

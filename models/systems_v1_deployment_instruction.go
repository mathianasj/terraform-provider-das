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
	"github.com/go-openapi/validate"
)

// SystemsV1DeploymentInstruction systems v1 deployment instruction
//
// swagger:model systems.v1.DeploymentInstruction
type SystemsV1DeploymentInstruction struct {

	// category/tool name
	// Required: true
	Category *string `json:"category"`

	// commands
	// Required: true
	Commands []*SystemsV1DeploymentCommand `json:"commands"`
}

// Validate validates this systems v1 deployment instruction
func (m *SystemsV1DeploymentInstruction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1DeploymentInstruction) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *SystemsV1DeploymentInstruction) validateCommands(formats strfmt.Registry) error {

	if err := validate.Required("commands", "body", m.Commands); err != nil {
		return err
	}

	for i := 0; i < len(m.Commands); i++ {
		if swag.IsZero(m.Commands[i]) { // not required
			continue
		}

		if m.Commands[i] != nil {
			if err := m.Commands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this systems v1 deployment instruction based on the context it is used
func (m *SystemsV1DeploymentInstruction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommands(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1DeploymentInstruction) contextValidateCommands(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Commands); i++ {

		if m.Commands[i] != nil {
			if err := m.Commands[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1DeploymentInstruction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1DeploymentInstruction) UnmarshalBinary(b []byte) error {
	var res SystemsV1DeploymentInstruction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

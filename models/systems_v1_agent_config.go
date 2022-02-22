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

// SystemsV1AgentConfig systems v1 agent config
//
// swagger:model systems.v1.AgentConfig
type SystemsV1AgentConfig struct {

	// agent ID
	// Required: true
	ID *string `json:"id"`

	// agent status
	// Read Only: true
	Status interface{} `json:"status,omitempty"`
}

// Validate validates this systems v1 agent config
func (m *SystemsV1AgentConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1AgentConfig) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this systems v1 agent config based on context it is used
func (m *SystemsV1AgentConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1AgentConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1AgentConfig) UnmarshalBinary(b []byte) error {
	var res SystemsV1AgentConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

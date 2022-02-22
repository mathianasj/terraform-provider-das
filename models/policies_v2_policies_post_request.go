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

// PoliciesV2PoliciesPostRequest policies v2 policies post request
//
// swagger:model policies.v2.PoliciesPostRequest
type PoliciesV2PoliciesPostRequest struct {

	// context
	// Required: true
	Context *string `json:"context"`

	// message
	// Required: true
	Message *string `json:"message"`

	// payload
	// Required: true
	Payload map[string]PoliciesV2PolicyFile `json:"payload"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this policies v2 policies post request
func (m *PoliciesV2PoliciesPostRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
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

func (m *PoliciesV2PoliciesPostRequest) validateContext(formats strfmt.Registry) error {

	if err := validate.Required("context", "body", m.Context); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV2PoliciesPostRequest) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV2PoliciesPostRequest) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	for k := range m.Payload {

		if err := validate.Required("payload"+"."+k, "body", m.Payload[k]); err != nil {
			return err
		}
		if val, ok := m.Payload[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payload" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("payload" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *PoliciesV2PoliciesPostRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this policies v2 policies post request based on the context it is used
func (m *PoliciesV2PoliciesPostRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePayload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoliciesV2PoliciesPostRequest) contextValidatePayload(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	for k := range m.Payload {

		if val, ok := m.Payload[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoliciesV2PoliciesPostRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoliciesV2PoliciesPostRequest) UnmarshalBinary(b []byte) error {
	var res PoliciesV2PoliciesPostRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
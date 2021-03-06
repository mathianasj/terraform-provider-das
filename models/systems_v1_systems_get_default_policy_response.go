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

// SystemsV1SystemsGetDefaultPolicyResponse systems v1 systems get default policy response
//
// swagger:model systems.v1.SystemsGetDefaultPolicyResponse
type SystemsV1SystemsGetDefaultPolicyResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`

	// result
	// Required: true
	Result *string `json:"result"`
}

// Validate validates this systems v1 systems get default policy response
func (m *SystemsV1SystemsGetDefaultPolicyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1SystemsGetDefaultPolicyResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this systems v1 systems get default policy response based on context it is used
func (m *SystemsV1SystemsGetDefaultPolicyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1SystemsGetDefaultPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1SystemsGetDefaultPolicyResponse) UnmarshalBinary(b []byte) error {
	var res SystemsV1SystemsGetDefaultPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

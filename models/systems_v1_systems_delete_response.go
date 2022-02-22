// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemsV1SystemsDeleteResponse systems v1 systems delete response
//
// swagger:model systems.v1.SystemsDeleteResponse
type SystemsV1SystemsDeleteResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this systems v1 systems delete response
func (m *SystemsV1SystemsDeleteResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this systems v1 systems delete response based on context it is used
func (m *SystemsV1SystemsDeleteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1SystemsDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1SystemsDeleteResponse) UnmarshalBinary(b []byte) error {
	var res SystemsV1SystemsDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

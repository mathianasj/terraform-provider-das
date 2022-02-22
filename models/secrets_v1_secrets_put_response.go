// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SecretsV1SecretsPutResponse secrets v1 secrets put response
//
// swagger:model secrets.v1.SecretsPutResponse
type SecretsV1SecretsPutResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this secrets v1 secrets put response
func (m *SecretsV1SecretsPutResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secrets v1 secrets put response based on context it is used
func (m *SecretsV1SecretsPutResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecretsV1SecretsPutResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretsV1SecretsPutResponse) UnmarshalBinary(b []byte) error {
	var res SecretsV1SecretsPutResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
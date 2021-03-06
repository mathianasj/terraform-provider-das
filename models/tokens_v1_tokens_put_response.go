// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokensV1TokensPutResponse tokens v1 tokens put response
//
// swagger:model tokens.v1.TokensPutResponse
type TokensV1TokensPutResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`

	// result
	Result string `json:"result,omitempty"`
}

// Validate validates this tokens v1 tokens put response
func (m *TokensV1TokensPutResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tokens v1 tokens put response based on context it is used
func (m *TokensV1TokensPutResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokensV1TokensPutResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokensV1TokensPutResponse) UnmarshalBinary(b []byte) error {
	var res TokensV1TokensPutResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

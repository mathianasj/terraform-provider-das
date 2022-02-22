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

// TokensV1TokensPutRequest tokens v1 tokens put request
//
// swagger:model tokens.v1.TokensPutRequest
type TokensV1TokensPutRequest struct {

	// allow path patterns
	// Required: true
	AllowPathPatterns []string `json:"allow_path_patterns"`

	// description
	// Required: true
	Description *string `json:"description"`

	// regenerate
	// Required: true
	Regenerate *bool `json:"regenerate"`

	// ttl
	TTL string `json:"ttl,omitempty"`
}

// Validate validates this tokens v1 tokens put request
func (m *TokensV1TokensPutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowPathPatterns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegenerate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokensV1TokensPutRequest) validateAllowPathPatterns(formats strfmt.Registry) error {

	if err := validate.Required("allow_path_patterns", "body", m.AllowPathPatterns); err != nil {
		return err
	}

	return nil
}

func (m *TokensV1TokensPutRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *TokensV1TokensPutRequest) validateRegenerate(formats strfmt.Registry) error {

	if err := validate.Required("regenerate", "body", m.Regenerate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tokens v1 tokens put request based on context it is used
func (m *TokensV1TokensPutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokensV1TokensPutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokensV1TokensPutRequest) UnmarshalBinary(b []byte) error {
	var res TokensV1TokensPutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

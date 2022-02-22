// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkspaceV1KafkaPlain workspace v1 kafka plain
//
// swagger:model workspace.v1.KafkaPlain
type WorkspaceV1KafkaPlain struct {

	// Authenticated Username/Password, stored in /v1/secrets/${user}
	User string `json:"user,omitempty"`
}

// Validate validates this workspace v1 kafka plain
func (m *WorkspaceV1KafkaPlain) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this workspace v1 kafka plain based on context it is used
func (m *WorkspaceV1KafkaPlain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceV1KafkaPlain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceV1KafkaPlain) UnmarshalBinary(b []byte) error {
	var res WorkspaceV1KafkaPlain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
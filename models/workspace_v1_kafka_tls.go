// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkspaceV1KafkaTLS workspace v1 kafka Tls
//
// swagger:model workspace.v1.KafkaTls
type WorkspaceV1KafkaTLS struct {

	// Optional client public/private certificate PEM, stored in /v1/secrets/${client_cert}
	ClientCert string `json:"client_cert,omitempty"`

	// Skip server certificate chain and host verification
	InsecureSkipVerify bool `json:"insecure_skip_verify,omitempty"`

	// Trusted Root CA(s) PEM
	Rootca string `json:"rootca,omitempty"`
}

// Validate validates this workspace v1 kafka Tls
func (m *WorkspaceV1KafkaTLS) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this workspace v1 kafka Tls based on context it is used
func (m *WorkspaceV1KafkaTLS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceV1KafkaTLS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceV1KafkaTLS) UnmarshalBinary(b []byte) error {
	var res WorkspaceV1KafkaTLS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

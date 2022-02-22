// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RelayProtocolClient relay protocol client
//
// swagger:model relay.protocol.Client
type RelayProtocolClient struct {

	// client key
	ClientKey string `json:"client_key,omitempty"`

	// remote address
	RemoteAddress string `json:"remote_address,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this relay protocol client
func (m *RelayProtocolClient) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this relay protocol client based on context it is used
func (m *RelayProtocolClient) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RelayProtocolClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelayProtocolClient) UnmarshalBinary(b []byte) error {
	var res RelayProtocolClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PoliciesV2PoliciesPutResponse policies v2 policies put response
//
// swagger:model policies.v2.PoliciesPutResponse
type PoliciesV2PoliciesPutResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this policies v2 policies put response
func (m *PoliciesV2PoliciesPutResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policies v2 policies put response based on context it is used
func (m *PoliciesV2PoliciesPutResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PoliciesV2PoliciesPutResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoliciesV2PoliciesPutResponse) UnmarshalBinary(b []byte) error {
	var res PoliciesV2PoliciesPutResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

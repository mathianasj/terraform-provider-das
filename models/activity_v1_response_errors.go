// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActivityV1ResponseErrors activity v1 response errors
//
// swagger:model activity.v1.ResponseErrors
type ActivityV1ResponseErrors struct {

	// processing
	Processing string `json:"processing,omitempty"`
}

// Validate validates this activity v1 response errors
func (m *ActivityV1ResponseErrors) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this activity v1 response errors based on context it is used
func (m *ActivityV1ResponseErrors) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActivityV1ResponseErrors) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityV1ResponseErrors) UnmarshalBinary(b []byte) error {
	var res ActivityV1ResponseErrors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

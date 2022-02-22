// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogreplayV1ReplayScope logreplay v1 replay scope
//
// swagger:model logreplay.v1.ReplayScope
type LogreplayV1ReplayScope struct {

	// maximum decision age (either in relative duration or in absolute time string)
	MaxAge string `json:"max_age,omitempty"`

	// maximum number of revisions to analyze back from the latest revision
	MaxRevisions *int64 `json:"max_revisions,omitempty"`

	// minimum decision age (either in relative duration or in absolute time string)
	MinAge string `json:"min_age,omitempty"`

	// decision path must overlap with this one (can be subpath of the decision path to compare portion of the result)
	Path string `json:"path,omitempty"`
}

// Validate validates this logreplay v1 replay scope
func (m *LogreplayV1ReplayScope) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this logreplay v1 replay scope based on context it is used
func (m *LogreplayV1ReplayScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogreplayV1ReplayScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogreplayV1ReplayScope) UnmarshalBinary(b []byte) error {
	var res LogreplayV1ReplayScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
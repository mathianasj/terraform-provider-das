// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogsV1PostDecisionsResponse logs v1 post decisions response
//
// swagger:model logs.v1.PostDecisionsResponse
type LogsV1PostDecisionsResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this logs v1 post decisions response
func (m *LogsV1PostDecisionsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this logs v1 post decisions response based on context it is used
func (m *LogsV1PostDecisionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogsV1PostDecisionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogsV1PostDecisionsResponse) UnmarshalBinary(b []byte) error {
	var res LogsV1PostDecisionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

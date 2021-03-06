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

// StatusV1StatusGetResponse status v1 status get response
//
// swagger:model status.v1.StatusGetResponse
type StatusV1StatusGetResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`

	// result
	// Required: true
	Result map[string]StatusV1AgentStatus `json:"result"`
}

// Validate validates this status v1 status get response
func (m *StatusV1StatusGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusV1StatusGetResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	for k := range m.Result {

		if err := validate.Required("result"+"."+k, "body", m.Result[k]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this status v1 status get response based on context it is used
func (m *StatusV1StatusGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatusV1StatusGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusV1StatusGetResponse) UnmarshalBinary(b []byte) error {
	var res StatusV1StatusGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

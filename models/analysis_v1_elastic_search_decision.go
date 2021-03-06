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

// AnalysisV1ElasticSearchDecision analysis v1 elastic search decision
//
// swagger:model analysis.v1.ElasticSearchDecision
type AnalysisV1ElasticSearchDecision struct {

	// decision
	// Required: true
	Decision *string `json:"Decision"`

	// processed
	// Format: date-time
	Processed strfmt.DateTime `json:"processed,omitempty"`
}

// Validate validates this analysis v1 elastic search decision
func (m *AnalysisV1ElasticSearchDecision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalysisV1ElasticSearchDecision) validateDecision(formats strfmt.Registry) error {

	if err := validate.Required("Decision", "body", m.Decision); err != nil {
		return err
	}

	return nil
}

func (m *AnalysisV1ElasticSearchDecision) validateProcessed(formats strfmt.Registry) error {
	if swag.IsZero(m.Processed) { // not required
		return nil
	}

	if err := validate.FormatOf("processed", "body", "date-time", m.Processed.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this analysis v1 elastic search decision based on context it is used
func (m *AnalysisV1ElasticSearchDecision) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnalysisV1ElasticSearchDecision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalysisV1ElasticSearchDecision) UnmarshalBinary(b []byte) error {
	var res AnalysisV1ElasticSearchDecision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

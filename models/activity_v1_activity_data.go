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

// ActivityV1ActivityData activity v1 activity data
//
// swagger:model activity.v1.ActivityData
type ActivityV1ActivityData struct {

	// decision
	Decision *ServicesAuthzDecisionDecision `json:"decision,omitempty"`

	// duration
	// Required: true
	Duration *int64 `json:"duration"`

	// request
	// Required: true
	Request *ActivityV1RequestData `json:"request"`

	// response
	// Required: true
	Response *ActivityV1ResponseData `json:"response"`
}

// Validate validates this activity v1 activity data
func (m *ActivityV1ActivityData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityV1ActivityData) validateDecision(formats strfmt.Registry) error {
	if swag.IsZero(m.Decision) { // not required
		return nil
	}

	if m.Decision != nil {
		if err := m.Decision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("decision")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityV1ActivityData) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *ActivityV1ActivityData) validateRequest(formats strfmt.Registry) error {

	if err := validate.Required("request", "body", m.Request); err != nil {
		return err
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityV1ActivityData) validateResponse(formats strfmt.Registry) error {

	if err := validate.Required("response", "body", m.Response); err != nil {
		return err
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this activity v1 activity data based on the context it is used
func (m *ActivityV1ActivityData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDecision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityV1ActivityData) contextValidateDecision(ctx context.Context, formats strfmt.Registry) error {

	if m.Decision != nil {
		if err := m.Decision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("decision")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityV1ActivityData) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {
		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityV1ActivityData) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.Response != nil {
		if err := m.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivityV1ActivityData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityV1ActivityData) UnmarshalBinary(b []byte) error {
	var res ActivityV1ActivityData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// ActivityV1RequestData activity v1 request data
//
// swagger:model activity.v1.RequestData
type ActivityV1RequestData struct {

	// body
	Body string `json:"body,omitempty"`

	// errors
	Errors *ActivityV1RequestErrors `json:"errors,omitempty"`

	// host
	// Required: true
	Host *string `json:"host"`

	// id
	// Required: true
	ID *string `json:"id"`

	// method
	// Required: true
	Method *string `json:"method"`

	// path
	// Required: true
	Path *string `json:"path"`

	// request body
	RequestBody string `json:"request_body,omitempty"`

	// requested by
	RequestedBy string `json:"requested_by,omitempty"`

	// requested through
	RequestedThrough string `json:"requested_through,omitempty"`

	// tenant
	// Required: true
	Tenant *string `json:"tenant"`

	// timestamp
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this activity v1 request data
func (m *ActivityV1RequestData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityV1RequestData) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if m.Errors != nil {
		if err := m.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errors")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityV1RequestData) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *ActivityV1RequestData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ActivityV1RequestData) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

func (m *ActivityV1RequestData) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *ActivityV1RequestData) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *ActivityV1RequestData) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this activity v1 request data based on the context it is used
func (m *ActivityV1RequestData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityV1RequestData) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if m.Errors != nil {
		if err := m.Errors.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errors")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivityV1RequestData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityV1RequestData) UnmarshalBinary(b []byte) error {
	var res ActivityV1RequestData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// MetaV2ObjectMeta meta v2 object meta
//
// swagger:model meta.v2.ObjectMeta
type MetaV2ObjectMeta struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// created through
	CreatedThrough string `json:"created_through,omitempty"`

	// last modified at
	// Format: date-time
	LastModifiedAt strfmt.DateTime `json:"last_modified_at,omitempty"`

	// last modified by
	LastModifiedBy string `json:"last_modified_by,omitempty"`

	// last modified through
	LastModifiedThrough string `json:"last_modified_through,omitempty"`
}

// Validate validates this meta v2 object meta
func (m *MetaV2ObjectMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetaV2ObjectMeta) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MetaV2ObjectMeta) validateLastModifiedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModifiedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_modified_at", "body", "date-time", m.LastModifiedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this meta v2 object meta based on context it is used
func (m *MetaV2ObjectMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetaV2ObjectMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetaV2ObjectMeta) UnmarshalBinary(b []byte) error {
	var res MetaV2ObjectMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemsV1MocksHTTPSendResponse systems v1 mocks Http send response
//
// swagger:model systems.v1.MocksHttpSendResponse
type SystemsV1MocksHTTPSendResponse struct {

	// mock http.send response
	Mocked []*SystemsV1MocksHTTPResponse `json:"mocked"`

	// mock http.send response
	Unmocked []*SystemsV1MocksHTTPResponse `json:"unmocked"`
}

// Validate validates this systems v1 mocks Http send response
func (m *SystemsV1MocksHTTPSendResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnmocked(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1MocksHTTPSendResponse) validateMocked(formats strfmt.Registry) error {
	if swag.IsZero(m.Mocked) { // not required
		return nil
	}

	for i := 0; i < len(m.Mocked); i++ {
		if swag.IsZero(m.Mocked[i]) { // not required
			continue
		}

		if m.Mocked[i] != nil {
			if err := m.Mocked[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mocked" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mocked" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SystemsV1MocksHTTPSendResponse) validateUnmocked(formats strfmt.Registry) error {
	if swag.IsZero(m.Unmocked) { // not required
		return nil
	}

	for i := 0; i < len(m.Unmocked); i++ {
		if swag.IsZero(m.Unmocked[i]) { // not required
			continue
		}

		if m.Unmocked[i] != nil {
			if err := m.Unmocked[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unmocked" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("unmocked" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this systems v1 mocks Http send response based on the context it is used
func (m *SystemsV1MocksHTTPSendResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMocked(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnmocked(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1MocksHTTPSendResponse) contextValidateMocked(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mocked); i++ {

		if m.Mocked[i] != nil {
			if err := m.Mocked[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mocked" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mocked" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SystemsV1MocksHTTPSendResponse) contextValidateUnmocked(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Unmocked); i++ {

		if m.Unmocked[i] != nil {
			if err := m.Unmocked[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unmocked" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("unmocked" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1MocksHTTPSendResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1MocksHTTPSendResponse) UnmarshalBinary(b []byte) error {
	var res SystemsV1MocksHTTPSendResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

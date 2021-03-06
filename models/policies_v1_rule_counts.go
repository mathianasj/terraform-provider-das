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

// PoliciesV1RuleCounts policies v1 rule counts
//
// swagger:model policies.v1.RuleCounts
type PoliciesV1RuleCounts struct {

	// number of allow rules
	// Required: true
	Allow *int32 `json:"allow"`

	// number of deny rules
	// Required: true
	Deny *int32 `json:"deny"`

	// number of enforce rules
	// Required: true
	Enforce *int32 `json:"enforce"`

	// number of ignore rules
	// Required: true
	Ignore *int32 `json:"ignore"`

	// number of monitor rules
	// Required: true
	Monitor *int32 `json:"monitor"`

	// number of notify rules
	// Required: true
	Notify *int32 `json:"notify"`

	// number of unclassified rules
	// Required: true
	Other *int32 `json:"other"`

	// number of test rules
	// Required: true
	Test *int32 `json:"test"`

	// total number of rules
	// Required: true
	Total *int32 `json:"total"`
}

// Validate validates this policies v1 rule counts
func (m *PoliciesV1RuleCounts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeny(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnforce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgnore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotify(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOther(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoliciesV1RuleCounts) validateAllow(formats strfmt.Registry) error {

	if err := validate.Required("allow", "body", m.Allow); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV1RuleCounts) validateDeny(formats strfmt.Registry) error {

	if err := validate.Required("deny", "body", m.Deny); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV1RuleCounts) validateEnforce(formats strfmt.Registry) error {

	if err := validate.Required("enforce", "body", m.Enforce); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV1RuleCounts) validateIgnore(formats strfmt.Registry) error {

	if err := validate.Required("ignore", "body", m.Ignore); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV1RuleCounts) validateMonitor(formats strfmt.Registry) error {

	if err := validate.Required("monitor", "body", m.Monitor); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV1RuleCounts) validateNotify(formats strfmt.Registry) error {

	if err := validate.Required("notify", "body", m.Notify); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV1RuleCounts) validateOther(formats strfmt.Registry) error {

	if err := validate.Required("other", "body", m.Other); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV1RuleCounts) validateTest(formats strfmt.Registry) error {

	if err := validate.Required("test", "body", m.Test); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV1RuleCounts) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this policies v1 rule counts based on context it is used
func (m *PoliciesV1RuleCounts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PoliciesV1RuleCounts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoliciesV1RuleCounts) UnmarshalBinary(b []byte) error {
	var res PoliciesV1RuleCounts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

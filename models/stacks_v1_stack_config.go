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
	"github.com/go-openapi/validate"
)

// StacksV1StackConfig stacks v1 stack config
//
// swagger:model stacks.v1.StackConfig
type StacksV1StackConfig struct {

	// authorization config
	// Read Only: true
	Authz *SystemsV1AuthzConfig `json:"authz,omitempty"`

	// datasources created for the stack
	// Read Only: true
	Datasources []*SystemsV1DatasourceConfig `json:"datasources"`

	// description
	// Required: true
	Description *string `json:"description"`

	// id
	// Required: true
	ID *string `json:"id"`

	// IDs of systems matching the stack
	MatchingSystems []string `json:"matching_systems"`

	// metadata
	// Required: true
	Metadata *MetaV1ObjectMeta `json:"metadata"`

	// minimum running OPA version for any of the matching systems
	MinimumOpaVersion string `json:"minimum_opa_version,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// policies
	// Required: true
	Policies []*SystemsV1PolicyConfig `json:"policies"`

	// read only
	// Required: true
	ReadOnly *bool `json:"read_only"`

	// source control
	SourceControl *StacksV1SourceControlConfig `json:"source_control,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// type
	// Required: true
	Type *string `json:"type"`

	// stack type parameter values (for template.* types)
	TypeParameters interface{} `json:"type_parameters,omitempty"`
}

// Validate validates this stacks v1 stack config
func (m *StacksV1StackConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatasources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceControl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1StackConfig) validateAuthz(formats strfmt.Registry) error {
	if swag.IsZero(m.Authz) { // not required
		return nil
	}

	if m.Authz != nil {
		if err := m.Authz.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authz")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authz")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StackConfig) validateDatasources(formats strfmt.Registry) error {
	if swag.IsZero(m.Datasources) { // not required
		return nil
	}

	for i := 0; i < len(m.Datasources); i++ {
		if swag.IsZero(m.Datasources[i]) { // not required
			continue
		}

		if m.Datasources[i] != nil {
			if err := m.Datasources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datasources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datasources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksV1StackConfig) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *StacksV1StackConfig) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *StacksV1StackConfig) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StackConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StacksV1StackConfig) validatePolicies(formats strfmt.Registry) error {

	if err := validate.Required("policies", "body", m.Policies); err != nil {
		return err
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksV1StackConfig) validateReadOnly(formats strfmt.Registry) error {

	if err := validate.Required("read_only", "body", m.ReadOnly); err != nil {
		return err
	}

	return nil
}

func (m *StacksV1StackConfig) validateSourceControl(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceControl) { // not required
		return nil
	}

	if m.SourceControl != nil {
		if err := m.SourceControl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_control")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_control")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StackConfig) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *StacksV1StackConfig) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this stacks v1 stack config based on the context it is used
func (m *StacksV1StackConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthz(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatasources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceControl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1StackConfig) contextValidateAuthz(ctx context.Context, formats strfmt.Registry) error {

	if m.Authz != nil {
		if err := m.Authz.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authz")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authz")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StackConfig) contextValidateDatasources(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "datasources", "body", []*SystemsV1DatasourceConfig(m.Datasources)); err != nil {
		return err
	}

	for i := 0; i < len(m.Datasources); i++ {

		if m.Datasources[i] != nil {
			if err := m.Datasources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datasources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datasources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksV1StackConfig) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StackConfig) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Policies); i++ {

		if m.Policies[i] != nil {
			if err := m.Policies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksV1StackConfig) contextValidateSourceControl(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceControl != nil {
		if err := m.SourceControl.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_control")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_control")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StacksV1StackConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksV1StackConfig) UnmarshalBinary(b []byte) error {
	var res StacksV1StackConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

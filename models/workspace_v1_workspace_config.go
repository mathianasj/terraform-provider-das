// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkspaceV1WorkspaceConfig workspace v1 workspace config
//
// swagger:model workspace.v1.WorkspaceConfig
type WorkspaceV1WorkspaceConfig struct {

	// activity exporter
	ActivityExporter *WorkspaceV1ActivityExporterConfig `json:"activity_exporter,omitempty"`

	// decisions exporter
	DecisionsExporter *WorkspaceV1DecisionExporterConfig `json:"decisions_exporter,omitempty"`

	// github
	Github *WorkspaceV1GithubConfiguration `json:"github,omitempty"`

	// metadata
	Metadata *MetaV1ObjectMeta `json:"metadata,omitempty"`

	// metrics exporter
	MetricsExporter *WorkspaceV1MetricsExporterConfig `json:"metrics_exporter,omitempty"`

	// source control
	SourceControl *GitV1SourceControlConfig `json:"source_control,omitempty"`

	// status
	Status *WorkspaceV1Status `json:"status,omitempty"`
}

// Validate validates this workspace v1 workspace config
func (m *WorkspaceV1WorkspaceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDecisionsExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGithub(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricsExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceControl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspaceV1WorkspaceConfig) validateActivityExporter(formats strfmt.Registry) error {
	if swag.IsZero(m.ActivityExporter) { // not required
		return nil
	}

	if m.ActivityExporter != nil {
		if err := m.ActivityExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity_exporter")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceV1WorkspaceConfig) validateDecisionsExporter(formats strfmt.Registry) error {
	if swag.IsZero(m.DecisionsExporter) { // not required
		return nil
	}

	if m.DecisionsExporter != nil {
		if err := m.DecisionsExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decisions_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("decisions_exporter")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceV1WorkspaceConfig) validateGithub(formats strfmt.Registry) error {
	if swag.IsZero(m.Github) { // not required
		return nil
	}

	if m.Github != nil {
		if err := m.Github.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("github")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("github")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceV1WorkspaceConfig) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
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

func (m *WorkspaceV1WorkspaceConfig) validateMetricsExporter(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricsExporter) { // not required
		return nil
	}

	if m.MetricsExporter != nil {
		if err := m.MetricsExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics_exporter")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceV1WorkspaceConfig) validateSourceControl(formats strfmt.Registry) error {
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

func (m *WorkspaceV1WorkspaceConfig) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this workspace v1 workspace config based on the context it is used
func (m *WorkspaceV1WorkspaceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActivityExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDecisionsExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGithub(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetricsExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceControl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspaceV1WorkspaceConfig) contextValidateActivityExporter(ctx context.Context, formats strfmt.Registry) error {

	if m.ActivityExporter != nil {
		if err := m.ActivityExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity_exporter")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceV1WorkspaceConfig) contextValidateDecisionsExporter(ctx context.Context, formats strfmt.Registry) error {

	if m.DecisionsExporter != nil {
		if err := m.DecisionsExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decisions_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("decisions_exporter")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceV1WorkspaceConfig) contextValidateGithub(ctx context.Context, formats strfmt.Registry) error {

	if m.Github != nil {
		if err := m.Github.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("github")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("github")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceV1WorkspaceConfig) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WorkspaceV1WorkspaceConfig) contextValidateMetricsExporter(ctx context.Context, formats strfmt.Registry) error {

	if m.MetricsExporter != nil {
		if err := m.MetricsExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics_exporter")
			}
			return err
		}
	}

	return nil
}

func (m *WorkspaceV1WorkspaceConfig) contextValidateSourceControl(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WorkspaceV1WorkspaceConfig) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceV1WorkspaceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceV1WorkspaceConfig) UnmarshalBinary(b []byte) error {
	var res WorkspaceV1WorkspaceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

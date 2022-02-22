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

// SystemsV1BundleRegistryConfig systems v1 bundle registry config
//
// swagger:model systems.v1.BundleRegistryConfig
type SystemsV1BundleRegistryConfig struct {

	// configuration for external S3 bucket to use for bundle distribution
	DistributionS3 *SystemsV1BundleDistributionS3Config `json:"distribution_s3,omitempty"`

	// manual deployment mode to prevent automatic deployment of new bundles
	ManualDeployment bool `json:"manual_deployment,omitempty"`

	// maximum number of all bundles to store
	MaxBundles int64 `json:"max_bundles,omitempty"`

	// maximum number of previously deployed bundles to store
	MaxDeployedBundles int64 `json:"max_deployed_bundles,omitempty"`
}

// Validate validates this systems v1 bundle registry config
func (m *SystemsV1BundleRegistryConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistributionS3(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1BundleRegistryConfig) validateDistributionS3(formats strfmt.Registry) error {
	if swag.IsZero(m.DistributionS3) { // not required
		return nil
	}

	if m.DistributionS3 != nil {
		if err := m.DistributionS3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distribution_s3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("distribution_s3")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this systems v1 bundle registry config based on the context it is used
func (m *SystemsV1BundleRegistryConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDistributionS3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1BundleRegistryConfig) contextValidateDistributionS3(ctx context.Context, formats strfmt.Registry) error {

	if m.DistributionS3 != nil {
		if err := m.DistributionS3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distribution_s3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("distribution_s3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1BundleRegistryConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1BundleRegistryConfig) UnmarshalBinary(b []byte) error {
	var res SystemsV1BundleRegistryConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

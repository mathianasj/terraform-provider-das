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

// SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials systems v1 bundle distribution s3 config opa credentials web identity credentials
//
// swagger:model systems.v1.BundleDistributionS3Config.opa_credentials.web_identity_credentials
type SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials struct {

	// aws region
	// Required: true
	AwsRegion *string `json:"aws_region"`

	// session name
	// Required: true
	SessionName *string `json:"session_name"`
}

// Validate validates this systems v1 bundle distribution s3 config opa credentials web identity credentials
func (m *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials) validateAwsRegion(formats strfmt.Registry) error {

	if err := validate.Required("aws_region", "body", m.AwsRegion); err != nil {
		return err
	}

	return nil
}

func (m *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials) validateSessionName(formats strfmt.Registry) error {

	if err := validate.Required("session_name", "body", m.SessionName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this systems v1 bundle distribution s3 config opa credentials web identity credentials based on context it is used
func (m *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials) UnmarshalBinary(b []byte) error {
	var res SystemsV1BundleDistributionS3ConfigOpaCredentialsWebIdentityCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

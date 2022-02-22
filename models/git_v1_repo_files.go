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

// GitV1RepoFiles git v1 repo files
//
// swagger:model git.v1.RepoFiles
type GitV1RepoFiles struct {

	// deleted files
	DeletedFiles []string `json:"deleted_files"`

	// files
	// Required: true
	Files map[string]string `json:"files"`
}

// Validate validates this git v1 repo files
func (m *GitV1RepoFiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitV1RepoFiles) validateFiles(formats strfmt.Registry) error {

	if err := validate.Required("files", "body", m.Files); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this git v1 repo files based on context it is used
func (m *GitV1RepoFiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitV1RepoFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitV1RepoFiles) UnmarshalBinary(b []byte) error {
	var res GitV1RepoFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

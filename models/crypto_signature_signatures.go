// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// CryptoSignatureSignatures crypto signature signatures
//
// swagger:model crypto.Signature.signatures
type CryptoSignatureSignatures map[string]string

// Validate validates this crypto signature signatures
func (m CryptoSignatureSignatures) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this crypto signature signatures based on context it is used
func (m CryptoSignatureSignatures) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
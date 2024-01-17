// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TestbankAccountIdentifier testbank account identifier
//
// swagger:model testbank.AccountIdentifier
type TestbankAccountIdentifier struct {

	// number
	Number string `json:"number,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this testbank account identifier
func (m *TestbankAccountIdentifier) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this testbank account identifier based on context it is used
func (m *TestbankAccountIdentifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestbankAccountIdentifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestbankAccountIdentifier) UnmarshalBinary(b []byte) error {
	var res TestbankAccountIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

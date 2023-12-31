// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CounterpartyIdentifier counterparty identifier
//
// swagger:model CounterpartyIdentifier
type CounterpartyIdentifier struct {

	// Type of the identifier. For country specific basic bank account number with bank identifier (e.g. routing or clearing number) choose BBAN. For country specific bank account number without bank identifier (e.g. routing or clearing number) choose ACCOUNT_NUMBER. For further information please refer to Account Identifiers section.
	IdentifierType string `json:"identifierType,omitempty"`

	// markets
	Markets []string `json:"markets"`
}

// Validate validates this counterparty identifier
func (m *CounterpartyIdentifier) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this counterparty identifier based on context it is used
func (m *CounterpartyIdentifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CounterpartyIdentifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CounterpartyIdentifier) UnmarshalBinary(b []byte) error {
	var res CounterpartyIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

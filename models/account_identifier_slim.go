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

// AccountIdentifierSlim Slim variant of AccountIdentifier, a strict subset.
//
// swagger:model AccountIdentifierSlim
type AccountIdentifierSlim struct {

	// holder name
	// Example: John Smith
	// Required: true
	HolderName *string `json:"holderName"`

	// market
	// Example: DE
	Market string `json:"market,omitempty"`

	// number
	// Example: DE47500105175557488443
	Number string `json:"number,omitempty"`

	// Type of the identifier. For country specific basic bank account number with bank identifier (e.g. routing or clearing number) choose BBAN. For country specific bank account number without bank identifier (e.g. routing or clearing number) choose ACCOUNT_NUMBER. For further information please refer to Account Identifiers section.
	// Example: IBAN
	Type string `json:"type,omitempty"`
}

// Validate validates this account identifier slim
func (m *AccountIdentifierSlim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHolderName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountIdentifierSlim) validateHolderName(formats strfmt.Registry) error {

	if err := validate.Required("holderName", "body", m.HolderName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this account identifier slim based on context it is used
func (m *AccountIdentifierSlim) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountIdentifierSlim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountIdentifierSlim) UnmarshalBinary(b []byte) error {
	var res AccountIdentifierSlim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
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

// AccountTrx a subset of an Account set on each Transaction
//
// swagger:model AccountTrx
type AccountTrx struct {

	// bank
	Bank *BankSlim `json:"bank,omitempty"`

	// id
	// Example: 31d593d7-fff9-4783-aa1d-92acb7b21a19
	// Required: true
	ID *string `json:"id"`

	// Name of the account
	// Example: Company Account
	Name string `json:"name,omitempty"`
}

// Validate validates this account trx
func (m *AccountTrx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountTrx) validateBank(formats strfmt.Registry) error {
	if swag.IsZero(m.Bank) { // not required
		return nil
	}

	if m.Bank != nil {
		if err := m.Bank.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bank")
			}
			return err
		}
	}

	return nil
}

func (m *AccountTrx) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this account trx based on the context it is used
func (m *AccountTrx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBank(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountTrx) contextValidateBank(ctx context.Context, formats strfmt.Registry) error {

	if m.Bank != nil {

		if swag.IsZero(m.Bank) { // not required
			return nil
		}

		if err := m.Bank.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bank")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountTrx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountTrx) UnmarshalBinary(b []byte) error {
	var res AccountTrx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

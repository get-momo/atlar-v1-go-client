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

// CreateTestbankAccountRequest create testbank account request
//
// swagger:model CreateTestbankAccountRequest
type CreateTestbankAccountRequest struct {

	// The BIC code of the bank for the account. Always formatted as 11 characters, if your BIC is only 8 chars, add "XXX" as the last 3 characters.
	Bic string `json:"bic,omitempty"`

	// booked balance
	BookedBalance *AmountInput `json:"bookedBalance,omitempty"`

	// iban
	Iban string `json:"iban,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner name
	OwnerName string `json:"ownerName,omitempty"`
}

// Validate validates this create testbank account request
func (m *CreateTestbankAccountRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBookedBalance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTestbankAccountRequest) validateBookedBalance(formats strfmt.Registry) error {
	if swag.IsZero(m.BookedBalance) { // not required
		return nil
	}

	if m.BookedBalance != nil {
		if err := m.BookedBalance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bookedBalance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bookedBalance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create testbank account request based on the context it is used
func (m *CreateTestbankAccountRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBookedBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTestbankAccountRequest) contextValidateBookedBalance(ctx context.Context, formats strfmt.Registry) error {

	if m.BookedBalance != nil {

		if swag.IsZero(m.BookedBalance) { // not required
			return nil
		}

		if err := m.BookedBalance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bookedBalance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bookedBalance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTestbankAccountRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTestbankAccountRequest) UnmarshalBinary(b []byte) error {
	var res CreateTestbankAccountRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

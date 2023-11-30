// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TestbankCurrencyExchange testbank currency exchange
//
// swagger:model testbank.CurrencyExchange
type TestbankCurrencyExchange struct {

	// exchange rate
	ExchangeRate string `json:"exchangeRate,omitempty"`
}

// Validate validates this testbank currency exchange
func (m *TestbankCurrencyExchange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this testbank currency exchange based on context it is used
func (m *TestbankCurrencyExchange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestbankCurrencyExchange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestbankCurrencyExchange) UnmarshalBinary(b []byte) error {
	var res TestbankCurrencyExchange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

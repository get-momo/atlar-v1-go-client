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

// Affiliation affiliation
//
// swagger:model Affiliation
type Affiliation struct {

	// DEPRECATED
	Bank *BankSlim `json:"bank,omitempty"`

	// created
	Created string `json:"created,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID string `json:"organizationId,omitempty"`

	// third party
	ThirdParty *ThirdParty `json:"thirdParty,omitempty"`

	// updated
	Updated string `json:"updated,omitempty"`
}

// Validate validates this affiliation
func (m *Affiliation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThirdParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Affiliation) validateBank(formats strfmt.Registry) error {
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

func (m *Affiliation) validateThirdParty(formats strfmt.Registry) error {
	if swag.IsZero(m.ThirdParty) { // not required
		return nil
	}

	if m.ThirdParty != nil {
		if err := m.ThirdParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thirdParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("thirdParty")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this affiliation based on the context it is used
func (m *Affiliation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBank(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThirdParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Affiliation) contextValidateBank(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Affiliation) contextValidateThirdParty(ctx context.Context, formats strfmt.Registry) error {

	if m.ThirdParty != nil {

		if swag.IsZero(m.ThirdParty) { // not required
			return nil
		}

		if err := m.ThirdParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thirdParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("thirdParty")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Affiliation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Affiliation) UnmarshalBinary(b []byte) error {
	var res Affiliation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

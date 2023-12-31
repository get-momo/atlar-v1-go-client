// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransactionCounterpartyDetails transaction counterparty details
//
// swagger:model TransactionCounterpartyDetails
type TransactionCounterpartyDetails struct {

	// contact details
	ContactDetails *ContactDetails `json:"contactDetails,omitempty"`

	// external account
	ExternalAccount *TransactionExternalAccount `json:"externalAccount,omitempty"`

	// populated if Mandate reference was given on the transaction by the bank. Typically, if a transaction stems from a DirectDebit payment, the mandate reference that was used to pull the funds will be given
	// Example: 123456
	MandateReference string `json:"mandateReference,omitempty"`

	// Name of the Counterparty
	// Example: Customer #312
	Name string `json:"name,omitempty"`

	// national identifier
	NationalIdentifier *NationalIdentifier `json:"nationalIdentifier,omitempty"`

	// The legal type of the Counterparty
	// Example: INDIVIDUAL
	// Enum: [INDIVIDUAL COMPANY]
	PartyType string `json:"partyType,omitempty"`
}

// Validate validates this transaction counterparty details
func (m *TransactionCounterpartyDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNationalIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartyType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionCounterpartyDetails) validateContactDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ContactDetails) { // not required
		return nil
	}

	if m.ContactDetails != nil {
		if err := m.ContactDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contactDetails")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionCounterpartyDetails) validateExternalAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalAccount) { // not required
		return nil
	}

	if m.ExternalAccount != nil {
		if err := m.ExternalAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalAccount")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionCounterpartyDetails) validateNationalIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.NationalIdentifier) { // not required
		return nil
	}

	if m.NationalIdentifier != nil {
		if err := m.NationalIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nationalIdentifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nationalIdentifier")
			}
			return err
		}
	}

	return nil
}

var transactionCounterpartyDetailsTypePartyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INDIVIDUAL","COMPANY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionCounterpartyDetailsTypePartyTypePropEnum = append(transactionCounterpartyDetailsTypePartyTypePropEnum, v)
	}
}

const (

	// TransactionCounterpartyDetailsPartyTypeINDIVIDUAL captures enum value "INDIVIDUAL"
	TransactionCounterpartyDetailsPartyTypeINDIVIDUAL string = "INDIVIDUAL"

	// TransactionCounterpartyDetailsPartyTypeCOMPANY captures enum value "COMPANY"
	TransactionCounterpartyDetailsPartyTypeCOMPANY string = "COMPANY"
)

// prop value enum
func (m *TransactionCounterpartyDetails) validatePartyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transactionCounterpartyDetailsTypePartyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransactionCounterpartyDetails) validatePartyType(formats strfmt.Registry) error {
	if swag.IsZero(m.PartyType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePartyTypeEnum("partyType", "body", m.PartyType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transaction counterparty details based on the context it is used
func (m *TransactionCounterpartyDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContactDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNationalIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionCounterpartyDetails) contextValidateContactDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ContactDetails != nil {

		if swag.IsZero(m.ContactDetails) { // not required
			return nil
		}

		if err := m.ContactDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contactDetails")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionCounterpartyDetails) contextValidateExternalAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalAccount != nil {

		if swag.IsZero(m.ExternalAccount) { // not required
			return nil
		}

		if err := m.ExternalAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalAccount")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionCounterpartyDetails) contextValidateNationalIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if m.NationalIdentifier != nil {

		if swag.IsZero(m.NationalIdentifier) { // not required
			return nil
		}

		if err := m.NationalIdentifier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nationalIdentifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nationalIdentifier")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionCounterpartyDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionCounterpartyDetails) UnmarshalBinary(b []byte) error {
	var res TransactionCounterpartyDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

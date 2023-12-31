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

// CreateDirectDebitRequest create direct debit request
//
// swagger:model CreateDirectDebitRequest
type CreateDirectDebitRequest struct {

	// The amount that should be collected.
	// Required: true
	Amount *AmountInput `json:"amount"`

	// ISO 20022 Category Purpose Code (ExternalCategoryPurpose1Code). Whether the category purpose code is
	// propagated and/or used throughout the complete chain of payment processing intermediaries and made available to the
	// debtor bank may be subject to bank and scheme specifics. Specifying category purpose is supported for
	// the following schemes: SDD_CORE.
	// Example: SUPP
	CategoryPurpose string `json:"categoryPurpose,omitempty"`

	// The requested execution date with format: `yyyy-mm-dd`. It is a local date from the destination account's point-of-view, without timezone. Cannot be set to before the current date, and not more than 60 days into the future.
	// Example: 2022-05-07
	// Required: true
	Date *string `json:"date"`

	// The ID of your internal account that the money should go to.
	// Example: 75c5b34c-e2b4-416b-9f23-8e797a620768
	// Required: true
	DestinationAccountID *string `json:"destinationAccountId"`

	// The type of this direct debit.
	// Example: SDD_CORE
	// Required: true
	// Enum: [SDD_CORE SDD_B2B AUTOGIRO]
	DirectDebitSchemeType *string `json:"directDebitSchemeType"`

	// ExternalId is optional to use, but if used, the Atlar platform will persist it, index it, as well as require it to be unique across all DirectDebits. It is also possible to retrieve the identified DirectDebit using the ExternalId.
	// Example: walVNuin6X5Mvte4xhg1ibTAVSACfN4Q9hl
	ExternalID string `json:"externalId,omitempty"`

	// Any external metadata you want to attach, such as your own IDs.
	ExternalMetadata ExternalMetadata `json:"externalMetadata,omitempty"`

	// Remittance information that might show up on the debtor's account statement. Whether or not it will show up depends on the direct debit scheme. Should not be set when scheme type is AUTOGIRO.
	RemittanceInformation *RemittanceInformation `json:"remittanceInformation,omitempty"`

	// The ID of the external account to collect money from. External accounts can be created via the API.
	// Example: ba0ad236-2904-4b01-8e85-a4ef3a73c353
	// Required: true
	SourceExternalAccountID *string `json:"sourceExternalAccountId"`
}

// Validate validates this create direct debit request
func (m *CreateDirectDebitRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitSchemeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemittanceInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceExternalAccountID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDirectDebitRequest) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

func (m *CreateDirectDebitRequest) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

func (m *CreateDirectDebitRequest) validateDestinationAccountID(formats strfmt.Registry) error {

	if err := validate.Required("destinationAccountId", "body", m.DestinationAccountID); err != nil {
		return err
	}

	return nil
}

var createDirectDebitRequestTypeDirectDebitSchemeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SDD_CORE","SDD_B2B","AUTOGIRO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createDirectDebitRequestTypeDirectDebitSchemeTypePropEnum = append(createDirectDebitRequestTypeDirectDebitSchemeTypePropEnum, v)
	}
}

const (

	// CreateDirectDebitRequestDirectDebitSchemeTypeSDDCORE captures enum value "SDD_CORE"
	CreateDirectDebitRequestDirectDebitSchemeTypeSDDCORE string = "SDD_CORE"

	// CreateDirectDebitRequestDirectDebitSchemeTypeSDDB2B captures enum value "SDD_B2B"
	CreateDirectDebitRequestDirectDebitSchemeTypeSDDB2B string = "SDD_B2B"

	// CreateDirectDebitRequestDirectDebitSchemeTypeAUTOGIRO captures enum value "AUTOGIRO"
	CreateDirectDebitRequestDirectDebitSchemeTypeAUTOGIRO string = "AUTOGIRO"
)

// prop value enum
func (m *CreateDirectDebitRequest) validateDirectDebitSchemeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createDirectDebitRequestTypeDirectDebitSchemeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateDirectDebitRequest) validateDirectDebitSchemeType(formats strfmt.Registry) error {

	if err := validate.Required("directDebitSchemeType", "body", m.DirectDebitSchemeType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectDebitSchemeTypeEnum("directDebitSchemeType", "body", *m.DirectDebitSchemeType); err != nil {
		return err
	}

	return nil
}

func (m *CreateDirectDebitRequest) validateExternalMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalMetadata) { // not required
		return nil
	}

	if m.ExternalMetadata != nil {
		if err := m.ExternalMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *CreateDirectDebitRequest) validateRemittanceInformation(formats strfmt.Registry) error {
	if swag.IsZero(m.RemittanceInformation) { // not required
		return nil
	}

	if m.RemittanceInformation != nil {
		if err := m.RemittanceInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remittanceInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remittanceInformation")
			}
			return err
		}
	}

	return nil
}

func (m *CreateDirectDebitRequest) validateSourceExternalAccountID(formats strfmt.Registry) error {

	if err := validate.Required("sourceExternalAccountId", "body", m.SourceExternalAccountID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create direct debit request based on the context it is used
func (m *CreateDirectDebitRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemittanceInformation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDirectDebitRequest) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.Amount != nil {

		if err := m.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

func (m *CreateDirectDebitRequest) contextValidateExternalMetadata(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalMetadata) { // not required
		return nil
	}

	if err := m.ExternalMetadata.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("externalMetadata")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("externalMetadata")
		}
		return err
	}

	return nil
}

func (m *CreateDirectDebitRequest) contextValidateRemittanceInformation(ctx context.Context, formats strfmt.Registry) error {

	if m.RemittanceInformation != nil {

		if swag.IsZero(m.RemittanceInformation) { // not required
			return nil
		}

		if err := m.RemittanceInformation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remittanceInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remittanceInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateDirectDebitRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDirectDebitRequest) UnmarshalBinary(b []byte) error {
	var res CreateDirectDebitRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

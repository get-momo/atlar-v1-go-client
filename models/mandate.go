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

// Mandate mandate
//
// swagger:model Mandate
type Mandate struct {

	// The ID of the instruction that was sent to the bank where this mandate was included when it was created. Subsequent instructions, such as mandate cancellations, won't update this ID. It will always refer to the instruction which created the resource.
	// Example: c1d6502c-07c7-4c77-bee5-963f9464c7d8
	ConnectionInstructionID string `json:"connectionInstructionId,omitempty"`

	// The ID of the counterparty which owns this mandate.
	// Example: f3efbb73-4e5b-4b22-adeb-918bbf1dfbd8
	// Required: true
	CounterpartyID *string `json:"counterpartyId"`

	// A timestamp showing when this mandate was created.
	// Example: 2022-07-20T18:31:12.889104898Z
	// Required: true
	Created *string `json:"created"`

	// The customer number that you have received from the bank.
	CreditorReference string `json:"creditorReference,omitempty"`

	// The ID of the destination account where money will be collected for this mandate. This is only needed for `AUTOGIRO` mandates.
	// Example: 31d593d7-fff9-4783-aa1d-92acb7b21a19
	DestinationAccountID string `json:"destinationAccountId,omitempty"`

	// The scheme type of the direct debits for which this mandate will be used.
	// Example: SDD_CORE
	// Required: true
	DirectDebitSchemeType *string `json:"directDebitSchemeType"`

	// The ETag based on version. To be passed along in `If-Match` header when updating a Mandate
	// Example: version:1
	Etag string `json:"etag,omitempty"`

	// The ID of the external account to which this mandate belongs.
	// Example: ba0ad236-2904-4b01-8e85-a4ef3a73c353
	// Required: true
	ExternalAccountID *string `json:"externalAccountId"`

	// ExternalId is optional to use, but if used, the Atlar platform will persist it, index it, as well as require it to be unique. It is also possible to retrieve the identified resource using the ExternalId.
	// Example: walVNuin6X5Mvte4xhg1ibTAVSACfN4Q9hl
	ExternalID string `json:"externalId,omitempty"`

	// Any external metadata you want to attach, such as your own internal IDs.
	ExternalMetadata ExternalMetadata `json:"externalMetadata,omitempty"`

	// The unique identifier of this mandate.
	// Example: be7e04fe-3fa0-48f5-a993-2008409255f2
	// Required: true
	ID *string `json:"id"`

	// A unique reference to the mandate that you have signed with the debtor.
	MandateReference string `json:"mandateReference,omitempty"`

	// The national identifier of the debtor. This cannot be changed after the mandate has been created.
	NationalIdentifier *NationalIdentifier `json:"nationalIdentifier,omitempty"`

	// Your organization ID.
	// Example: 605e26fc-4fce-495a-a92f-2c3592d7287e
	// Required: true
	OrganizationID *string `json:"organizationId"`

	// The date that this mandate was signed by the end user.
	// Example: 2022-08-01
	SignatureDate string `json:"signatureDate,omitempty"`

	// The current status of the mandate.
	// Example: SENT
	// Required: true
	// Enum: [CREATED PENDING_CANCELLATION PENDING_CANCELLATION_SUBMISSION CANCELLATION_SENT PENDING_SUBMISSION SENT ACCEPTED ACTIVE FAILED CANCELLED]
	Status *string `json:"status"`

	// A timestamp showing when this mandate was last updated.
	// Example: 2022-05-20T18:31:12.889104898Z
	// Required: true
	Updated *string `json:"updated"`

	// starts at 1 when the mandate is created and increases by one in each successive update
	Version int64 `json:"version,omitempty"`
}

// Validate validates this mandate
func (m *Mandate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounterpartyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitSchemeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNationalIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mandate) validateCounterpartyID(formats strfmt.Registry) error {

	if err := validate.Required("counterpartyId", "body", m.CounterpartyID); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateDirectDebitSchemeType(formats strfmt.Registry) error {

	if err := validate.Required("directDebitSchemeType", "body", m.DirectDebitSchemeType); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateExternalAccountID(formats strfmt.Registry) error {

	if err := validate.Required("externalAccountId", "body", m.ExternalAccountID); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateExternalMetadata(formats strfmt.Registry) error {
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

func (m *Mandate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateNationalIdentifier(formats strfmt.Registry) error {
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

func (m *Mandate) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

var mandateTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATED","PENDING_CANCELLATION","PENDING_CANCELLATION_SUBMISSION","CANCELLATION_SENT","PENDING_SUBMISSION","SENT","ACCEPTED","ACTIVE","FAILED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mandateTypeStatusPropEnum = append(mandateTypeStatusPropEnum, v)
	}
}

const (

	// MandateStatusCREATED captures enum value "CREATED"
	MandateStatusCREATED string = "CREATED"

	// MandateStatusPENDINGCANCELLATION captures enum value "PENDING_CANCELLATION"
	MandateStatusPENDINGCANCELLATION string = "PENDING_CANCELLATION"

	// MandateStatusPENDINGCANCELLATIONSUBMISSION captures enum value "PENDING_CANCELLATION_SUBMISSION"
	MandateStatusPENDINGCANCELLATIONSUBMISSION string = "PENDING_CANCELLATION_SUBMISSION"

	// MandateStatusCANCELLATIONSENT captures enum value "CANCELLATION_SENT"
	MandateStatusCANCELLATIONSENT string = "CANCELLATION_SENT"

	// MandateStatusPENDINGSUBMISSION captures enum value "PENDING_SUBMISSION"
	MandateStatusPENDINGSUBMISSION string = "PENDING_SUBMISSION"

	// MandateStatusSENT captures enum value "SENT"
	MandateStatusSENT string = "SENT"

	// MandateStatusACCEPTED captures enum value "ACCEPTED"
	MandateStatusACCEPTED string = "ACCEPTED"

	// MandateStatusACTIVE captures enum value "ACTIVE"
	MandateStatusACTIVE string = "ACTIVE"

	// MandateStatusFAILED captures enum value "FAILED"
	MandateStatusFAILED string = "FAILED"

	// MandateStatusCANCELLED captures enum value "CANCELLED"
	MandateStatusCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *Mandate) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mandateTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Mandate) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Mandate) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("updated", "body", m.Updated); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mandate based on the context it is used
func (m *Mandate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalMetadata(ctx, formats); err != nil {
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

func (m *Mandate) contextValidateExternalMetadata(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Mandate) contextValidateNationalIdentifier(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Mandate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mandate) UnmarshalBinary(b []byte) error {
	var res Mandate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

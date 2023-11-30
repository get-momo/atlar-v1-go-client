// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateAffiliationRequest create affiliation request
//
// swagger:model CreateAffiliationRequest
type CreateAffiliationRequest struct {

	// DEPRECATED
	BankID string `json:"bankId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// third party Id
	ThirdPartyID string `json:"thirdPartyId,omitempty"`
}

// Validate validates this create affiliation request
func (m *CreateAffiliationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create affiliation request based on context it is used
func (m *CreateAffiliationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAffiliationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAffiliationRequest) UnmarshalBinary(b []byte) error {
	var res CreateAffiliationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
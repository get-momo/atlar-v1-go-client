// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MfaSetting mfa setting
//
// swagger:model MfaSetting
type MfaSetting struct {

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this mfa setting
func (m *MfaSetting) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mfa setting based on context it is used
func (m *MfaSetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MfaSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MfaSetting) UnmarshalBinary(b []byte) error {
	var res MfaSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

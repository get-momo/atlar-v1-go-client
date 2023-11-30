// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebhookSigningKey webhook signing key
//
// swagger:model WebhookSigningKey
type WebhookSigningKey struct {

	// The time this key was created.
	// Example: 2022-07-20T18:31:12.889104898Z
	Created string `json:"created,omitempty"`

	// The unique identifier of this webhook secret.
	// Example: be7e04fe-3fa0-48f5-a993-2008409255f2
	ID string `json:"id,omitempty"`

	// A key used to sign the webhook payload.
	// Example: agj+xWKk3gqkP+SsCsljkjbDth7bxguqVMRd4K3wm1I=
	SigningKey string `json:"signingKey,omitempty"`
}

// Validate validates this webhook signing key
func (m *WebhookSigningKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this webhook signing key based on context it is used
func (m *WebhookSigningKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookSigningKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookSigningKey) UnmarshalBinary(b []byte) error {
	var res WebhookSigningKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

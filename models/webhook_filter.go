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

// WebhookFilter webhook filter
//
// swagger:model WebhookFilter
type WebhookFilter struct {

	// The API version to serialize entities as. If you are consuming events and do want to deserialize into `/v1/credit-transfers`, set `apiVersion: 1`
	// Example: 1
	// Required: true
	APIVersion *int64 `json:"apiVersion"`

	// A list of events for which to receive webhooks. See events for each resource in the [Webhooks and events](doc:api#webhooks-and-events) section.
	// If a list of events is not specified you will receive all events for the given resource. Note that new events may
	// be added to the Atlar API over time which implies you will also receiving such events as they are rolled out.
	// Example: ["SENT_TO_BANK","RECONCILED"]
	Events []string `json:"events"`

	// The API resource for which to receive webhooks. See available resources in the [Webhooks and events](doc:api#webhooks-and-events) section.
	// Example: direct_debits
	// Required: true
	Resource *string `json:"resource"`
}

// Validate validates this webhook filter
func (m *WebhookFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookFilter) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("apiVersion", "body", m.APIVersion); err != nil {
		return err
	}

	return nil
}

func (m *WebhookFilter) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this webhook filter based on context it is used
func (m *WebhookFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookFilter) UnmarshalBinary(b []byte) error {
	var res WebhookFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

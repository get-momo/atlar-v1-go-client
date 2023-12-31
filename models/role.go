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

// Role role
//
// swagger:model Role
type Role struct {

	// Created timestamp
	Created string `json:"created,omitempty"`

	// id
	// Example: 54c71c82-82f0-439c-996f-c8bd62c3fe0a
	ID string `json:"id,omitempty"`

	// name
	// Example: FINANCE
	Name string `json:"name,omitempty"`

	// organization Id
	// Example: 605e26fc-4fce-495a-a92f-2c3592d7287e
	OrganizationID string `json:"organizationId,omitempty"`

	// owner
	// Example: false
	Owner bool `json:"owner,omitempty"`

	// comma-separated scopes (which is on form resource:method)
	// Example: accounts:read,payments:create
	Scopes string `json:"scopes,omitempty"`

	// type
	// Example: REGULAR_ROLE
	// Enum: [SERVICE_ROLE OPERATOR_ROLE REGULAR_ROLE]
	Type string `json:"type,omitempty"`

	// Updated timestamp
	Updated string `json:"updated,omitempty"`
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var roleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SERVICE_ROLE","OPERATOR_ROLE","REGULAR_ROLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roleTypeTypePropEnum = append(roleTypeTypePropEnum, v)
	}
}

const (

	// RoleTypeSERVICEROLE captures enum value "SERVICE_ROLE"
	RoleTypeSERVICEROLE string = "SERVICE_ROLE"

	// RoleTypeOPERATORROLE captures enum value "OPERATOR_ROLE"
	RoleTypeOPERATORROLE string = "OPERATOR_ROLE"

	// RoleTypeREGULARROLE captures enum value "REGULAR_ROLE"
	RoleTypeREGULARROLE string = "REGULAR_ROLE"
)

// prop value enum
func (m *Role) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, roleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Role) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this role based on context it is used
func (m *Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

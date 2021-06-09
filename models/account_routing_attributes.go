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

// AccountRoutingAttributes account routing attributes
//
// swagger:model AccountRoutingAttributes
type AccountRoutingAttributes struct {

	// Service name used to generate the account number
	// Example: accountapi
	// Required: true
	// Enum: [accountapi lhv-gateway]
	AccountGenerator *string `json:"account_generator"`

	// Service name used to provision the account
	// Example: accountapi
	// Required: true
	// Enum: [accountapi lhv-gateway starlinggateway]
	AccountProvisioner *string `json:"account_provisioner"`

	// match
	// Required: true
	Match *string `json:"match"`

	// highest that match is selected first. Must be greater or equal than 0
	// Example: 0
	// Required: true
	Priority *int64 `json:"priority"`
}

// Validate validates this account routing attributes
func (m *AccountRoutingAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountGenerator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountProvisioner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountRoutingAttributesTypeAccountGeneratorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["accountapi","lhv-gateway"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountRoutingAttributesTypeAccountGeneratorPropEnum = append(accountRoutingAttributesTypeAccountGeneratorPropEnum, v)
	}
}

const (

	// AccountRoutingAttributesAccountGeneratorAccountapi captures enum value "accountapi"
	AccountRoutingAttributesAccountGeneratorAccountapi string = "accountapi"

	// AccountRoutingAttributesAccountGeneratorLhvDashGateway captures enum value "lhv-gateway"
	AccountRoutingAttributesAccountGeneratorLhvDashGateway string = "lhv-gateway"
)

// prop value enum
func (m *AccountRoutingAttributes) validateAccountGeneratorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountRoutingAttributesTypeAccountGeneratorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountRoutingAttributes) validateAccountGenerator(formats strfmt.Registry) error {

	if err := validate.Required("account_generator", "body", m.AccountGenerator); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccountGeneratorEnum("account_generator", "body", *m.AccountGenerator); err != nil {
		return err
	}

	return nil
}

var accountRoutingAttributesTypeAccountProvisionerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["accountapi","lhv-gateway","starlinggateway"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountRoutingAttributesTypeAccountProvisionerPropEnum = append(accountRoutingAttributesTypeAccountProvisionerPropEnum, v)
	}
}

const (

	// AccountRoutingAttributesAccountProvisionerAccountapi captures enum value "accountapi"
	AccountRoutingAttributesAccountProvisionerAccountapi string = "accountapi"

	// AccountRoutingAttributesAccountProvisionerLhvDashGateway captures enum value "lhv-gateway"
	AccountRoutingAttributesAccountProvisionerLhvDashGateway string = "lhv-gateway"

	// AccountRoutingAttributesAccountProvisionerStarlinggateway captures enum value "starlinggateway"
	AccountRoutingAttributesAccountProvisionerStarlinggateway string = "starlinggateway"
)

// prop value enum
func (m *AccountRoutingAttributes) validateAccountProvisionerEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountRoutingAttributesTypeAccountProvisionerPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountRoutingAttributes) validateAccountProvisioner(formats strfmt.Registry) error {

	if err := validate.Required("account_provisioner", "body", m.AccountProvisioner); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccountProvisionerEnum("account_provisioner", "body", *m.AccountProvisioner); err != nil {
		return err
	}

	return nil
}

func (m *AccountRoutingAttributes) validateMatch(formats strfmt.Registry) error {

	if err := validate.Required("match", "body", m.Match); err != nil {
		return err
	}

	return nil
}

func (m *AccountRoutingAttributes) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this account routing attributes based on context it is used
func (m *AccountRoutingAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountRoutingAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRoutingAttributes) UnmarshalBinary(b []byte) error {
	var res AccountRoutingAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

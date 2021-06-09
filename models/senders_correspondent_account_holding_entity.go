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

// SendersCorrespondentAccountHoldingEntity senders correspondent account holding entity
//
// swagger:model SendersCorrespondentAccountHoldingEntity
type SendersCorrespondentAccountHoldingEntity struct {

	// Sender's correspondent's address
	// Example: ["Liverpool Customer Service Centre","Stevenson Way","Wavertree","L13 1NW"]
	BankAddress []string `json:"bank_address,omitempty"`

	// SWIFT BIC for sender's correspondent
	// Example: 333333
	BankID string `json:"bank_id,omitempty"`

	// bank id code
	BankIDCode BankIDCode `json:"bank_id_code,omitempty"`

	// Sender's correspondent's name
	// Example: XYZ BANK PLC
	BankName string `json:"bank_name,omitempty"`

	// Identifier of the financial institution which services the account
	// Example: //AT12345
	BankPartyID string `json:"bank_party_id,omitempty"`
}

// Validate validates this senders correspondent account holding entity
func (m *SendersCorrespondentAccountHoldingEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendersCorrespondentAccountHoldingEntity) validateBankIDCode(formats strfmt.Registry) error {
	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := m.BankIDCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bank_id_code")
		}
		return err
	}

	return nil
}

// ContextValidate validate this senders correspondent account holding entity based on the context it is used
func (m *SendersCorrespondentAccountHoldingEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBankIDCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendersCorrespondentAccountHoldingEntity) contextValidateBankIDCode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BankIDCode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bank_id_code")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendersCorrespondentAccountHoldingEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendersCorrespondentAccountHoldingEntity) UnmarshalBinary(b []byte) error {
	var res SendersCorrespondentAccountHoldingEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

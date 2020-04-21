// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EburyAssociationAttributes ebury association attributes
//
// swagger:model EburyAssociationAttributes
type EburyAssociationAttributes struct {

	// funding currency
	FundingCurrency string `json:"funding_currency,omitempty"`
}

// Validate validates this ebury association attributes
func (m *EburyAssociationAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EburyAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EburyAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res EburyAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

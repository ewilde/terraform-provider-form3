// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LhvAgencySynchronisationAttributes lhv agency synchronisation attributes
// swagger:model LhvAgencySynchronisationAttributes
type LhvAgencySynchronisationAttributes struct {

	// bank id
	// Required: true
	// Min Items: 1
	// Unique: true
	BankID []string `json:"bank_id"`

	// SWIFT BIC in either 8 character format
	// Required: true
	// Pattern: ^[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]$
	Bic string `json:"bic"`

	// country
	// Required: true
	// Pattern: ^[A-Z]{2}$
	Country string `json:"country"`
}

// Validate validates this lhv agency synchronisation attributes
func (m *LhvAgencySynchronisationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LhvAgencySynchronisationAttributes) validateBankID(formats strfmt.Registry) error {

	if err := validate.Required("bank_id", "body", m.BankID); err != nil {
		return err
	}

	iBankIDSize := int64(len(m.BankID))

	if err := validate.MinItems("bank_id", "body", iBankIDSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("bank_id", "body", m.BankID); err != nil {
		return err
	}

	for i := 0; i < len(m.BankID); i++ {

		if err := validate.Pattern("bank_id"+"."+strconv.Itoa(i), "body", string(m.BankID[i]), `^[0-9]{4}([0-9]{2})?$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *LhvAgencySynchronisationAttributes) validateBic(formats strfmt.Registry) error {

	if err := validate.RequiredString("bic", "body", string(m.Bic)); err != nil {
		return err
	}

	if err := validate.Pattern("bic", "body", string(m.Bic), `^[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *LhvAgencySynchronisationAttributes) validateCountry(formats strfmt.Registry) error {

	if err := validate.RequiredString("country", "body", string(m.Country)); err != nil {
		return err
	}

	if err := validate.Pattern("country", "body", string(m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LhvAgencySynchronisationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LhvAgencySynchronisationAttributes) UnmarshalBinary(b []byte) error {
	var res LhvAgencySynchronisationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

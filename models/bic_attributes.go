// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BicAttributes bic attributes
//
// swagger:model BicAttributes
type BicAttributes struct {

	// bic
	// Pattern: ^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$
	Bic string `json:"bic,omitempty"`
}

// Validate validates this bic attributes
func (m *BicAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BicAttributes) validateBic(formats strfmt.Registry) error {

	if swag.IsZero(m.Bic) { // not required
		return nil
	}

	if err := validate.Pattern("bic", "body", string(m.Bic), `^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BicAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BicAttributes) UnmarshalBinary(b []byte) error {
	var res BicAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

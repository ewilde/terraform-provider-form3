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

// AssociationAttributes association attributes
//
// swagger:model AssociationAttributes
type AssociationAttributes struct {

	// starling account name
	StarlingAccountName string `json:"starling_account_name,omitempty"`

	// starling account uid
	// Format: uuid
	StarlingAccountUID strfmt.UUID `json:"starling_account_uid,omitempty"`
}

// Validate validates this association attributes
func (m *AssociationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStarlingAccountUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssociationAttributes) validateStarlingAccountUID(formats strfmt.Registry) error {

	if swag.IsZero(m.StarlingAccountUID) { // not required
		return nil
	}

	if err := validate.FormatOf("starling_account_uid", "body", "uuid", m.StarlingAccountUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssociationAttributes) UnmarshalBinary(b []byte) error {
	var res AssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

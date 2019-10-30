// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SepaInstantAssociationRelationships sepa instant association relationships
// swagger:model SepaInstantAssociationRelationships
type SepaInstantAssociationRelationships struct {

	// sponsor
	Sponsor SepaInstantAssociationRelationshipsSponsor `json:"sponsor,omitempty"`
}

// Validate validates this sepa instant association relationships
func (m *SepaInstantAssociationRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSponsor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaInstantAssociationRelationships) validateSponsor(formats strfmt.Registry) error {

	if swag.IsZero(m.Sponsor) { // not required
		return nil
	}

	if err := m.Sponsor.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sponsor")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepaInstantAssociationRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaInstantAssociationRelationships) UnmarshalBinary(b []byte) error {
	var res SepaInstantAssociationRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SepaInstantAssociationRelationshipsSponsor sepa instant association relationships sponsor
// swagger:model SepaInstantAssociationRelationshipsSponsor
type SepaInstantAssociationRelationshipsSponsor struct {

	// data
	Data SepaInstantAssociationReference `json:"data,omitempty"`
}

// Validate validates this sepa instant association relationships sponsor
func (m *SepaInstantAssociationRelationshipsSponsor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaInstantAssociationRelationshipsSponsor) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if err := m.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sponsor" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepaInstantAssociationRelationshipsSponsor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaInstantAssociationRelationshipsSponsor) UnmarshalBinary(b []byte) error {
	var res SepaInstantAssociationRelationshipsSponsor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
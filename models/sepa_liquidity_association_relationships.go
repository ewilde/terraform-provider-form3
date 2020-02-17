// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SepaLiquidityAssociationRelationships sepa liquidity association relationships
// swagger:model SepaLiquidityAssociationRelationships
type SepaLiquidityAssociationRelationships struct {

	// sponsor
	Sponsor SepaLiquidityAssociationRelationshipsSponsor `json:"sponsor,omitempty"`
}

// Validate validates this sepa liquidity association relationships
func (m *SepaLiquidityAssociationRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSponsor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaLiquidityAssociationRelationships) validateSponsor(formats strfmt.Registry) error {

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
func (m *SepaLiquidityAssociationRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaLiquidityAssociationRelationships) UnmarshalBinary(b []byte) error {
	var res SepaLiquidityAssociationRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SepaLiquidityAssociationRelationshipsSponsor sepa liquidity association relationships sponsor
// swagger:model SepaLiquidityAssociationRelationshipsSponsor
type SepaLiquidityAssociationRelationshipsSponsor struct {

	// data
	// Required: true
	Data SepaLiquidityRelationshipData `json:"data"`
}

// Validate validates this sepa liquidity association relationships sponsor
func (m *SepaLiquidityAssociationRelationshipsSponsor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepaLiquidityAssociationRelationshipsSponsor) validateData(formats strfmt.Registry) error {

	if err := m.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sponsor" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepaLiquidityAssociationRelationshipsSponsor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaLiquidityAssociationRelationshipsSponsor) UnmarshalBinary(b []byte) error {
	var res SepaLiquidityAssociationRelationshipsSponsor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

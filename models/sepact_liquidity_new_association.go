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

// SepactLiquidityNewAssociation sepact liquidity new association
//
// swagger:model SepactLiquidityNewAssociation
type SepactLiquidityNewAssociation struct {

	// attributes
	// Required: true
	Attributes *SepactLiquidityAssociationAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *SepactLiquidityAssociationRelationships `json:"relationships,omitempty"`

	// type
	// Required: true
	// Enum: [reconciliation_associations]
	Type string `json:"type"`
}

// Validate validates this sepact liquidity new association
func (m *SepactLiquidityNewAssociation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepactLiquidityNewAssociation) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *SepactLiquidityNewAssociation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SepactLiquidityNewAssociation) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SepactLiquidityNewAssociation) validateRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

var sepactLiquidityNewAssociationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["reconciliation_associations"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sepactLiquidityNewAssociationTypeTypePropEnum = append(sepactLiquidityNewAssociationTypeTypePropEnum, v)
	}
}

const (

	// SepactLiquidityNewAssociationTypeReconciliationAssociations captures enum value "reconciliation_associations"
	SepactLiquidityNewAssociationTypeReconciliationAssociations string = "reconciliation_associations"
)

// prop value enum
func (m *SepactLiquidityNewAssociation) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sepactLiquidityNewAssociationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SepactLiquidityNewAssociation) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sepact liquidity new association based on the context it is used
func (m *SepactLiquidityNewAssociation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SepactLiquidityNewAssociation) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.Attributes != nil {
		if err := m.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *SepactLiquidityNewAssociation) contextValidateRelationships(ctx context.Context, formats strfmt.Registry) error {

	if m.Relationships != nil {
		if err := m.Relationships.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SepactLiquidityNewAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepactLiquidityNewAssociation) UnmarshalBinary(b []byte) error {
	var res SepactLiquidityNewAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

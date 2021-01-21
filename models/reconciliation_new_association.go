// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReconciliationNewAssociation reconciliation new association
// swagger:model ReconciliationNewAssociation
type ReconciliationNewAssociation struct {

	// attributes
	// Required: true
	Attributes *ReconciliationAssociationAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// type
	// Required: true
	// Enum: [reconciliation_associations]
	Type string `json:"type"`
}

// Validate validates this reconciliation new association
func (m *ReconciliationNewAssociation) Validate(formats strfmt.Registry) error {
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

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReconciliationNewAssociation) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReconciliationNewAssociation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReconciliationNewAssociation) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var reconciliationNewAssociationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["reconciliation_associations"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reconciliationNewAssociationTypeTypePropEnum = append(reconciliationNewAssociationTypeTypePropEnum, v)
	}
}

const (

	// ReconciliationNewAssociationTypeReconciliationAssociations captures enum value "reconciliation_associations"
	ReconciliationNewAssociationTypeReconciliationAssociations string = "reconciliation_associations"
)

// prop value enum
func (m *ReconciliationNewAssociation) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, reconciliationNewAssociationTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReconciliationNewAssociation) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReconciliationNewAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReconciliationNewAssociation) UnmarshalBinary(b []byte) error {
	var res ReconciliationNewAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataPublisherNewAssociation data publisher new association
//
// swagger:model DataPublisherNewAssociation
type DataPublisherNewAssociation struct {

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
	// Enum: [data_publisher_associations]
	Type string `json:"type"`
}

// Validate validates this data publisher new association
func (m *DataPublisherNewAssociation) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *DataPublisherNewAssociation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DataPublisherNewAssociation) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var dataPublisherNewAssociationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["data_publisher_associations"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataPublisherNewAssociationTypeTypePropEnum = append(dataPublisherNewAssociationTypeTypePropEnum, v)
	}
}

const (

	// DataPublisherNewAssociationTypeDataPublisherAssociations captures enum value "data_publisher_associations"
	DataPublisherNewAssociationTypeDataPublisherAssociations string = "data_publisher_associations"
)

// prop value enum
func (m *DataPublisherNewAssociation) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dataPublisherNewAssociationTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DataPublisherNewAssociation) validateType(formats strfmt.Registry) error {

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
func (m *DataPublisherNewAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataPublisherNewAssociation) UnmarshalBinary(b []byte) error {
	var res DataPublisherNewAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VocalinkReportAssociationCertificateRelationship vocalink report association certificate relationship
// swagger:model VocalinkReportAssociationCertificateRelationship
type VocalinkReportAssociationCertificateRelationship struct {

	// data
	Data *VocalinkReportAssociationCertificateRelationshipData `json:"data,omitempty"`
}

// Validate validates this vocalink report association certificate relationship
func (m *VocalinkReportAssociationCertificateRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VocalinkReportAssociationCertificateRelationship) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VocalinkReportAssociationCertificateRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VocalinkReportAssociationCertificateRelationship) UnmarshalBinary(b []byte) error {
	var res VocalinkReportAssociationCertificateRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VocalinkReportAssociationCertificateRelationshipData vocalink report association certificate relationship data
// swagger:model VocalinkReportAssociationCertificateRelationshipData
type VocalinkReportAssociationCertificateRelationshipData struct {

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this vocalink report association certificate relationship data
func (m *VocalinkReportAssociationCertificateRelationshipData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VocalinkReportAssociationCertificateRelationshipData) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("data"+"."+"id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VocalinkReportAssociationCertificateRelationshipData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VocalinkReportAssociationCertificateRelationshipData) UnmarshalBinary(b []byte) error {
	var res VocalinkReportAssociationCertificateRelationshipData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
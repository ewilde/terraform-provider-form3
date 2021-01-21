// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MandateRelationships mandate relationships
//
// swagger:model MandateRelationships
type MandateRelationships struct {

	// mandate admission
	MandateAdmission *MandateRelationshipsMandateAdmission `json:"mandate_admission,omitempty"`
}

// Validate validates this mandate relationships
func (m *MandateRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMandateAdmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateRelationships) validateMandateAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.MandateAdmission) { // not required
		return nil
	}

	if m.MandateAdmission != nil {
		if err := m.MandateAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mandate_admission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateRelationships) UnmarshalBinary(b []byte) error {
	var res MandateRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MandateRelationshipsMandateAdmission mandate relationships mandate admission
//
// swagger:model MandateRelationshipsMandateAdmission
type MandateRelationshipsMandateAdmission struct {

	// data
	Data []*MandateAdmission `json:"data"`
}

// Validate validates this mandate relationships mandate admission
func (m *MandateRelationshipsMandateAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateRelationshipsMandateAdmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mandate_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateRelationshipsMandateAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateRelationshipsMandateAdmission) UnmarshalBinary(b []byte) error {
	var res MandateRelationshipsMandateAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// DirectDebitAdmission direct debit admission
//
// swagger:model DirectDebitAdmission
type DirectDebitAdmission struct {

	// attributes
	Attributes *DirectDebitAdmissionAttributes `json:"attributes,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// organisation id
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// Validate validates this direct debit admission
func (m *DirectDebitAdmission) Validate(formats strfmt.Registry) error {
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

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitAdmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
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

func (m *DirectDebitAdmission) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DirectDebitAdmissionAttributes direct debit admission attributes
//
// swagger:model DirectDebitAdmissionAttributes
type DirectDebitAdmissionAttributes struct {

	// admission datetime
	// Read Only: true
	// Format: date-time
	AdmissionDatetime strfmt.DateTime `json:"admission_datetime,omitempty"`

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// scheme status code description
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// status
	Status DirectDebitAdmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason DirectDebitAdmissionStatusReason `json:"status_reason,omitempty"`
}

// Validate validates this direct debit admission attributes
func (m *DirectDebitAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status")
		}
		return err
	}

	return nil
}

func (m *DirectDebitAdmissionAttributes) validateStatusReason(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusReason) { // not required
		return nil
	}

	if err := m.StatusReason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status_reason")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SepaSctAssociationAttributes sepa sct association attributes
//
// swagger:model SepaSctAssociationAttributes
type SepaSctAssociationAttributes struct {

	// bic
	Bic string `json:"bic,omitempty"`

	// business user
	BusinessUser string `json:"businessUser,omitempty"`

	// can sponsor
	CanSponsor bool `json:"can_sponsor,omitempty"`

	// reachable bics
	ReachableBics []string `json:"reachable_bics"`

	// receiver business user
	ReceiverBusinessUser string `json:"receiverBusinessUser,omitempty"`
}

// Validate validates this sepa sct association attributes
func (m *SepaSctAssociationAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SepaSctAssociationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SepaSctAssociationAttributes) UnmarshalBinary(b []byte) error {
	var res SepaSctAssociationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

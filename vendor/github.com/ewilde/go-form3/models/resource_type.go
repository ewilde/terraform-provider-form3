// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ResourceType resource type
// swagger:model ResourceType
type ResourceType string

const (
	// ResourceTypePayportAssociations captures enum value "payport_associations"
	ResourceTypePayportAssociations ResourceType = "payport_associations"
	// ResourceTypeLimits captures enum value "limits"
	ResourceTypeLimits ResourceType = "limits"
)

// for schema
var resourceTypeEnum []interface{}

func init() {
	var res []ResourceType
	if err := json.Unmarshal([]byte(`["payport_associations","limits"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeEnum = append(resourceTypeEnum, v)
	}
}

func (m ResourceType) validateResourceTypeEnum(path, location string, value ResourceType) error {
	if err := validate.Enum(path, location, value, resourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this resource type
func (m ResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

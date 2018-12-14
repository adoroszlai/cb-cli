// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PlatformDisksJSON platform disks Json
// swagger:model PlatformDisksJson

type PlatformDisksJSON struct {

	// default disks
	DefaultDisks map[string]string `json:"defaultDisks,omitempty"`

	// disk mappings
	DiskMappings map[string]map[string]string `json:"diskMappings,omitempty"`

	// disk types
	DiskTypes map[string][]string `json:"diskTypes,omitempty"`

	// disk displayNames
	DisplayNames map[string]map[string]string `json:"displayNames,omitempty"`
}

/* polymorph PlatformDisksJson defaultDisks false */

/* polymorph PlatformDisksJson diskMappings false */

/* polymorph PlatformDisksJson diskTypes false */

/* polymorph PlatformDisksJson displayNames false */

// Validate validates this platform disks Json
func (m *PlatformDisksJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiskTypes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformDisksJSON) validateDiskTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.DiskTypes) { // not required
		return nil
	}

	if swag.IsZero(m.DiskTypes) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlatformDisksJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformDisksJSON) UnmarshalBinary(b []byte) error {
	var res PlatformDisksJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
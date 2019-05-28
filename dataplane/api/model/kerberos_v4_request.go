// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KerberosV4Request kerberos v4 request
// swagger:model KerberosV4Request
type KerberosV4Request struct {

	// active directory
	ActiveDirectory *ActiveDirectoryKerberosDescriptor `json:"activeDirectory,omitempty"`

	// ambari descriptor
	AmbariDescriptor *AmbariKerberosDescriptor `json:"ambariDescriptor,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// free ipa
	FreeIpa *FreeIPAKerberosDescriptor `json:"freeIpa,omitempty"`

	// mit
	Mit *MITKerberosDescriptor `json:"mit,omitempty"`

	// the name of the kerberos configuration
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this kerberos v4 request
func (m *KerberosV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveDirectory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmbariDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeIpa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KerberosV4Request) validateActiveDirectory(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveDirectory) { // not required
		return nil
	}

	if m.ActiveDirectory != nil {
		if err := m.ActiveDirectory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activeDirectory")
			}
			return err
		}
	}

	return nil
}

func (m *KerberosV4Request) validateAmbariDescriptor(formats strfmt.Registry) error {

	if swag.IsZero(m.AmbariDescriptor) { // not required
		return nil
	}

	if m.AmbariDescriptor != nil {
		if err := m.AmbariDescriptor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambariDescriptor")
			}
			return err
		}
	}

	return nil
}

func (m *KerberosV4Request) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *KerberosV4Request) validateFreeIpa(formats strfmt.Registry) error {

	if swag.IsZero(m.FreeIpa) { // not required
		return nil
	}

	if m.FreeIpa != nil {
		if err := m.FreeIpa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *KerberosV4Request) validateMit(formats strfmt.Registry) error {

	if swag.IsZero(m.Mit) { // not required
		return nil
	}

	if m.Mit != nil {
		if err := m.Mit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mit")
			}
			return err
		}
	}

	return nil
}

func (m *KerberosV4Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KerberosV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KerberosV4Request) UnmarshalBinary(b []byte) error {
	var res KerberosV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
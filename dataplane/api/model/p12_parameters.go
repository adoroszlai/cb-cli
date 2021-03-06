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

// P12Parameters p12 parameters
// swagger:model P12Parameters
type P12Parameters struct {

	// project Id
	// Required: true
	ProjectID *string `json:"projectId"`

	// service account Id
	// Required: true
	ServiceAccountID *string `json:"serviceAccountId"`

	// service account private key
	// Required: true
	ServiceAccountPrivateKey *string `json:"serviceAccountPrivateKey"`
}

// Validate validates this p12 parameters
func (m *P12Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccountPrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *P12Parameters) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *P12Parameters) validateServiceAccountID(formats strfmt.Registry) error {

	if err := validate.Required("serviceAccountId", "body", m.ServiceAccountID); err != nil {
		return err
	}

	return nil
}

func (m *P12Parameters) validateServiceAccountPrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("serviceAccountPrivateKey", "body", m.ServiceAccountPrivateKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *P12Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *P12Parameters) UnmarshalBinary(b []byte) error {
	var res P12Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

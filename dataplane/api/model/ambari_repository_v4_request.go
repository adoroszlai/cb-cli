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

// AmbariRepositoryV4Request ambari repository v4 request
// swagger:model AmbariRepositoryV4Request
type AmbariRepositoryV4Request struct {

	// url of the cluster manager repository
	BaseURL string `json:"baseUrl,omitempty"`

	// gpg key of the cluster manager repository
	GpgKeyURL string `json:"gpgKeyUrl,omitempty"`

	// version of the cluster manager
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this ambari repository v4 request
func (m *AmbariRepositoryV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AmbariRepositoryV4Request) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AmbariRepositoryV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmbariRepositoryV4Request) UnmarshalBinary(b []byte) error {
	var res AmbariRepositoryV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
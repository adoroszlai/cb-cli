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

// GcsCloudStorageV4Parameters gcs cloud storage v4 parameters
// swagger:model GcsCloudStorageV4Parameters
type GcsCloudStorageV4Parameters struct {

	// service account email
	// Required: true
	ServiceAccountEmail *string `json:"serviceAccountEmail"`
}

// Validate validates this gcs cloud storage v4 parameters
func (m *GcsCloudStorageV4Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceAccountEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcsCloudStorageV4Parameters) validateServiceAccountEmail(formats strfmt.Registry) error {

	if err := validate.Required("serviceAccountEmail", "body", m.ServiceAccountEmail); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GcsCloudStorageV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcsCloudStorageV4Parameters) UnmarshalBinary(b []byte) error {
	var res GcsCloudStorageV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

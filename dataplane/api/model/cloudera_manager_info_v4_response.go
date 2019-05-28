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

// ClouderaManagerInfoV4Response cloudera manager info v4 response
// swagger:model ClouderaManagerInfoV4Response
type ClouderaManagerInfoV4Response struct {

	// repository
	Repository map[string]ClouderaManagerRepositoryV4Response `json:"repository,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this cloudera manager info v4 response
func (m *ClouderaManagerInfoV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClouderaManagerInfoV4Response) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	for k := range m.Repository {

		if err := validate.Required("repository"+"."+k, "body", m.Repository[k]); err != nil {
			return err
		}
		if val, ok := m.Repository[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClouderaManagerInfoV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClouderaManagerInfoV4Response) UnmarshalBinary(b []byte) error {
	var res ClouderaManagerInfoV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CloudbreakDetailsJSON cloudbreak details Json
// swagger:model CloudbreakDetailsJson

type CloudbreakDetailsJSON struct {

	// version of the Cloudbreak that provisioned the stack
	Version string `json:"version,omitempty"`
}

/* polymorph CloudbreakDetailsJson version false */

// Validate validates this cloudbreak details Json
func (m *CloudbreakDetailsJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CloudbreakDetailsJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudbreakDetailsJSON) UnmarshalBinary(b []byte) error {
	var res CloudbreakDetailsJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
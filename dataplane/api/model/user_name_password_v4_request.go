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

// UserNamePasswordV4Request user name password v4 request
// swagger:model UserNamePasswordV4Request
type UserNamePasswordV4Request struct {

	// old password in ambari
	// Required: true
	// Max Length: 2147483647
	// Min Length: 1
	OldPassword *string `json:"oldPassword"`

	// new password in ambari
	// Required: true
	// Max Length: 2147483647
	// Min Length: 1
	Password *string `json:"password"`

	// new user name in ambari
	// Required: true
	// Max Length: 2147483647
	// Min Length: 1
	UserName *string `json:"userName"`
}

// Validate validates this user name password v4 request
func (m *UserNamePasswordV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOldPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserNamePasswordV4Request) validateOldPassword(formats strfmt.Registry) error {

	if err := validate.Required("oldPassword", "body", m.OldPassword); err != nil {
		return err
	}

	if err := validate.MinLength("oldPassword", "body", string(*m.OldPassword), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("oldPassword", "body", string(*m.OldPassword), 2147483647); err != nil {
		return err
	}

	return nil
}

func (m *UserNamePasswordV4Request) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(*m.Password), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(*m.Password), 2147483647); err != nil {
		return err
	}

	return nil
}

func (m *UserNamePasswordV4Request) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	if err := validate.MinLength("userName", "body", string(*m.UserName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("userName", "body", string(*m.UserName), 2147483647); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserNamePasswordV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserNamePasswordV4Request) UnmarshalBinary(b []byte) error {
	var res UserNamePasswordV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

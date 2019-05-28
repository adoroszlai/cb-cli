// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserV4Response user v4 response
// swagger:model UserV4Response
type UserV4Response struct {

	// id
	ID int64 `json:"id,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this user v4 response
func (m *UserV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserV4Response) UnmarshalBinary(b []byte) error {
	var res UserV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
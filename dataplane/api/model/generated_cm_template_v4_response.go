// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GeneratedCmTemplateV4Response generated cm template v4 response
// swagger:model GeneratedCmTemplateV4Response
type GeneratedCmTemplateV4Response struct {

	// template
	Template string `json:"template,omitempty"`
}

// Validate validates this generated cm template v4 response
func (m *GeneratedCmTemplateV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GeneratedCmTemplateV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneratedCmTemplateV4Response) UnmarshalBinary(b []byte) error {
	var res GeneratedCmTemplateV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
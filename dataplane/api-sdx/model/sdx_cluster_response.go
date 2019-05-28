// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SdxClusterResponse sdx cluster response
// swagger:model SdxClusterResponse
type SdxClusterResponse struct {

	// sdx crn
	SdxCrn string `json:"sdxCrn,omitempty"`

	// sdx name
	SdxName string `json:"sdxName,omitempty"`

	// status
	// Enum: [REQUESTED REQUESTED_FROM_CLOUDBREAK RUNNING PROVISIONING_FAILED DELETE_REQUESTED DELETED DELETE_FAILED]
	Status string `json:"status,omitempty"`
}

// Validate validates this sdx cluster response
func (m *SdxClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sdxClusterResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","REQUESTED_FROM_CLOUDBREAK","RUNNING","PROVISIONING_FAILED","DELETE_REQUESTED","DELETED","DELETE_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterResponseTypeStatusPropEnum = append(sdxClusterResponseTypeStatusPropEnum, v)
	}
}

const (

	// SdxClusterResponseStatusREQUESTED captures enum value "REQUESTED"
	SdxClusterResponseStatusREQUESTED string = "REQUESTED"

	// SdxClusterResponseStatusREQUESTEDFROMCLOUDBREAK captures enum value "REQUESTED_FROM_CLOUDBREAK"
	SdxClusterResponseStatusREQUESTEDFROMCLOUDBREAK string = "REQUESTED_FROM_CLOUDBREAK"

	// SdxClusterResponseStatusRUNNING captures enum value "RUNNING"
	SdxClusterResponseStatusRUNNING string = "RUNNING"

	// SdxClusterResponseStatusPROVISIONINGFAILED captures enum value "PROVISIONING_FAILED"
	SdxClusterResponseStatusPROVISIONINGFAILED string = "PROVISIONING_FAILED"

	// SdxClusterResponseStatusDELETEREQUESTED captures enum value "DELETE_REQUESTED"
	SdxClusterResponseStatusDELETEREQUESTED string = "DELETE_REQUESTED"

	// SdxClusterResponseStatusDELETED captures enum value "DELETED"
	SdxClusterResponseStatusDELETED string = "DELETED"

	// SdxClusterResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	SdxClusterResponseStatusDELETEFAILED string = "DELETE_FAILED"
)

// prop value enum
func (m *SdxClusterResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SdxClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxClusterResponse) UnmarshalBinary(b []byte) error {
	var res SdxClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
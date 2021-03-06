// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateClusterV4Request update cluster v4 request
// swagger:model UpdateClusterV4Request
type UpdateClusterV4Request struct {

	// blueprint id for the cluster
	BlueprintName string `json:"blueprintName,omitempty"`

	// host group adjustment
	HostGroupAdjustment *HostGroupAdjustmentV4Request `json:"hostGroupAdjustment,omitempty"`

	// collection of hostgroups
	// Unique: true
	Hostgroups []*HostGroupV4Request `json:"hostgroups"`

	// kerberos admin password
	// Max Length: 50
	// Min Length: 5
	KerberosPassword string `json:"kerberosPassword,omitempty"`

	// kerberos principal
	KerberosPrincipal string `json:"kerberosPrincipal,omitempty"`

	// details of the Ambari stack
	StackRepository *StackRepositoryV4Request `json:"stackRepository,omitempty"`

	// request status
	// Enum: [SYNC FULL_SYNC REPAIR_FAILED_NODES STOPPED STARTED]
	Status string `json:"status,omitempty"`

	// user details
	UserNamePassword *UserNamePasswordV4Request `json:"userNamePassword,omitempty"`

	// blueprint validation
	ValidateBlueprint *bool `json:"validateBlueprint,omitempty"`
}

// Validate validates this update cluster v4 request
func (m *UpdateClusterV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostGroupAdjustment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostgroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKerberosPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserNamePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateClusterV4Request) validateHostGroupAdjustment(formats strfmt.Registry) error {

	if swag.IsZero(m.HostGroupAdjustment) { // not required
		return nil
	}

	if m.HostGroupAdjustment != nil {
		if err := m.HostGroupAdjustment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostGroupAdjustment")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateClusterV4Request) validateHostgroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Hostgroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("hostgroups", "body", m.Hostgroups); err != nil {
		return err
	}

	for i := 0; i < len(m.Hostgroups); i++ {
		if swag.IsZero(m.Hostgroups[i]) { // not required
			continue
		}

		if m.Hostgroups[i] != nil {
			if err := m.Hostgroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostgroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateClusterV4Request) validateKerberosPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.KerberosPassword) { // not required
		return nil
	}

	if err := validate.MinLength("kerberosPassword", "body", string(m.KerberosPassword), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("kerberosPassword", "body", string(m.KerberosPassword), 50); err != nil {
		return err
	}

	return nil
}

func (m *UpdateClusterV4Request) validateStackRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.StackRepository) { // not required
		return nil
	}

	if m.StackRepository != nil {
		if err := m.StackRepository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackRepository")
			}
			return err
		}
	}

	return nil
}

var updateClusterV4RequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SYNC","FULL_SYNC","REPAIR_FAILED_NODES","STOPPED","STARTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClusterV4RequestTypeStatusPropEnum = append(updateClusterV4RequestTypeStatusPropEnum, v)
	}
}

const (

	// UpdateClusterV4RequestStatusSYNC captures enum value "SYNC"
	UpdateClusterV4RequestStatusSYNC string = "SYNC"

	// UpdateClusterV4RequestStatusFULLSYNC captures enum value "FULL_SYNC"
	UpdateClusterV4RequestStatusFULLSYNC string = "FULL_SYNC"

	// UpdateClusterV4RequestStatusREPAIRFAILEDNODES captures enum value "REPAIR_FAILED_NODES"
	UpdateClusterV4RequestStatusREPAIRFAILEDNODES string = "REPAIR_FAILED_NODES"

	// UpdateClusterV4RequestStatusSTOPPED captures enum value "STOPPED"
	UpdateClusterV4RequestStatusSTOPPED string = "STOPPED"

	// UpdateClusterV4RequestStatusSTARTED captures enum value "STARTED"
	UpdateClusterV4RequestStatusSTARTED string = "STARTED"
)

// prop value enum
func (m *UpdateClusterV4Request) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateClusterV4RequestTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClusterV4Request) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *UpdateClusterV4Request) validateUserNamePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.UserNamePassword) { // not required
		return nil
	}

	if m.UserNamePassword != nil {
		if err := m.UserNamePassword.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userNamePassword")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateClusterV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateClusterV4Request) UnmarshalBinary(b []byte) error {
	var res UpdateClusterV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

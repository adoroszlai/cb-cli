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

// ClusterTemplateViewV4Response cluster template view v4 response
// swagger:model ClusterTemplateViewV4Response
type ClusterTemplateViewV4Response struct {

	// cloudplatform which this template is compatible with
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// datalake required which this template is compatible with
	// Enum: [NONE OPTIONAL REQUIRED]
	DatalakeRequired string `json:"datalakeRequired,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// environment name
	EnvironmentName string `json:"environmentName,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// node count of the cluster template
	NodeCount int32 `json:"nodeCount,omitempty"`

	// stack type of cluster template
	StackType string `json:"stackType,omitempty"`

	// stack version of cluster template
	StackVersion string `json:"stackVersion,omitempty"`

	// status
	// Enum: [DEFAULT DEFAULT_DELETED USER_MANAGED OUTDATED]
	Status string `json:"status,omitempty"`

	// type of the cluster template
	// Enum: [SPARK HIVE DATASCIENCE EDW ETL OTHER]
	Type string `json:"type,omitempty"`
}

// Validate validates this cluster template view v4 response
func (m *ClusterTemplateViewV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clusterTemplateViewV4ResponseTypeDatalakeRequiredPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","OPTIONAL","REQUIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateViewV4ResponseTypeDatalakeRequiredPropEnum = append(clusterTemplateViewV4ResponseTypeDatalakeRequiredPropEnum, v)
	}
}

const (

	// ClusterTemplateViewV4ResponseDatalakeRequiredNONE captures enum value "NONE"
	ClusterTemplateViewV4ResponseDatalakeRequiredNONE string = "NONE"

	// ClusterTemplateViewV4ResponseDatalakeRequiredOPTIONAL captures enum value "OPTIONAL"
	ClusterTemplateViewV4ResponseDatalakeRequiredOPTIONAL string = "OPTIONAL"

	// ClusterTemplateViewV4ResponseDatalakeRequiredREQUIRED captures enum value "REQUIRED"
	ClusterTemplateViewV4ResponseDatalakeRequiredREQUIRED string = "REQUIRED"
)

// prop value enum
func (m *ClusterTemplateViewV4Response) validateDatalakeRequiredEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateViewV4ResponseTypeDatalakeRequiredPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateViewV4Response) validateDatalakeRequired(formats strfmt.Registry) error {

	if swag.IsZero(m.DatalakeRequired) { // not required
		return nil
	}

	// value enum
	if err := m.validateDatalakeRequiredEnum("datalakeRequired", "body", m.DatalakeRequired); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTemplateViewV4Response) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTemplateViewV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

var clusterTemplateViewV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","DEFAULT_DELETED","USER_MANAGED","OUTDATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateViewV4ResponseTypeStatusPropEnum = append(clusterTemplateViewV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// ClusterTemplateViewV4ResponseStatusDEFAULT captures enum value "DEFAULT"
	ClusterTemplateViewV4ResponseStatusDEFAULT string = "DEFAULT"

	// ClusterTemplateViewV4ResponseStatusDEFAULTDELETED captures enum value "DEFAULT_DELETED"
	ClusterTemplateViewV4ResponseStatusDEFAULTDELETED string = "DEFAULT_DELETED"

	// ClusterTemplateViewV4ResponseStatusUSERMANAGED captures enum value "USER_MANAGED"
	ClusterTemplateViewV4ResponseStatusUSERMANAGED string = "USER_MANAGED"

	// ClusterTemplateViewV4ResponseStatusOUTDATED captures enum value "OUTDATED"
	ClusterTemplateViewV4ResponseStatusOUTDATED string = "OUTDATED"
)

// prop value enum
func (m *ClusterTemplateViewV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateViewV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateViewV4Response) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var clusterTemplateViewV4ResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SPARK","HIVE","DATASCIENCE","EDW","ETL","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateViewV4ResponseTypeTypePropEnum = append(clusterTemplateViewV4ResponseTypeTypePropEnum, v)
	}
}

const (

	// ClusterTemplateViewV4ResponseTypeSPARK captures enum value "SPARK"
	ClusterTemplateViewV4ResponseTypeSPARK string = "SPARK"

	// ClusterTemplateViewV4ResponseTypeHIVE captures enum value "HIVE"
	ClusterTemplateViewV4ResponseTypeHIVE string = "HIVE"

	// ClusterTemplateViewV4ResponseTypeDATASCIENCE captures enum value "DATASCIENCE"
	ClusterTemplateViewV4ResponseTypeDATASCIENCE string = "DATASCIENCE"

	// ClusterTemplateViewV4ResponseTypeEDW captures enum value "EDW"
	ClusterTemplateViewV4ResponseTypeEDW string = "EDW"

	// ClusterTemplateViewV4ResponseTypeETL captures enum value "ETL"
	ClusterTemplateViewV4ResponseTypeETL string = "ETL"

	// ClusterTemplateViewV4ResponseTypeOTHER captures enum value "OTHER"
	ClusterTemplateViewV4ResponseTypeOTHER string = "OTHER"
)

// prop value enum
func (m *ClusterTemplateViewV4Response) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateViewV4ResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateViewV4Response) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterTemplateViewV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTemplateViewV4Response) UnmarshalBinary(b []byte) error {
	var res ClusterTemplateViewV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

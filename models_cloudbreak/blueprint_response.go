package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*BlueprintResponse blueprint response

swagger:model BlueprintResponse
*/
type BlueprintResponse struct {

	// TODO WARNING: do not replace it with string, otherwise it cannot be deserialized
	/* ambari blueprint JSON, set this or the url field

	Required: true
	*/
	AmbariBlueprint `json:"ambariBlueprint"`

	/* gathered from blueprintName field from the blueprint JSON
	 */
	BlueprintName *string `json:"blueprintName,omitempty"`

	/* description of the resource

	Max Length: 1000
	Min Length: 0
	*/
	Description *string `json:"description,omitempty"`

	/* number of host groups
	 */
	HostGroupCount *int32 `json:"hostGroupCount,omitempty"`

	/* id of the resource
	 */
	ID *int64 `json:"id,omitempty"`

	/* input parameters of the blueprint

	Unique: true
	*/
	Inputs []*BlueprintParameter `json:"inputs,omitempty"`

	/* name of the resource

	Required: true
	Max Length: 100
	Min Length: 1
	*/
	Name string `json:"name"`

	/* resource is visible in account
	 */
	Public *bool `json:"public,omitempty"`

	/* status of the blueprint
	 */
	Status *string `json:"status,omitempty"`
}

type Configurations map[string]map[string]interface{}

type Settings map[string][]map[string]interface{}

type AmbariBlueprint struct {
	Blueprint      Blueprint            `json:"Blueprints"`
	Configurations []Configurations     `json:"configurations,omitempty"`
	Settings       []Settings           `json:"settings,omitempty"`
	HostGroups     []BlueprintHostGroup `json:"host_groups"`
}

type Blueprint struct {
	Name         *string `json:"blueprint_name"`
	StackName    string  `json:"stack_name"`
	StackVersion string  `json:"stack_version"`
}

type BlueprintHostGroup struct {
	Name           string           `json:"name"`
	Configurations []Configurations `json:"configurations,omitempty"`
	Components     []Component      `json:"components"`
	Cardinality    string           `json:"cardinality"`
}

type Component struct {
	Name string `json:"name"`
}

// Validate validates this blueprint response
func (m *BlueprintResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlueprintResponse) validateDescription(formats strfmt.Registry) error {

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

func (m *BlueprintResponse) validateInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	if err := validate.UniqueItems("inputs", "body", m.Inputs); err != nil {
		return err
	}

	for i := 0; i < len(m.Inputs); i++ {

		if m.Inputs[i] != nil {

			if err := m.Inputs[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *BlueprintResponse) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 100); err != nil {
		return err
	}

	return nil
}

var blueprintResponseTypeStatusPropEnum []interface{}

func (m *BlueprintResponse) validateStatusEnum(path, location string, value string) error {
	if blueprintResponseTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["DEFAULT","DEFAULT_DELETED","USER_MANAGED"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			blueprintResponseTypeStatusPropEnum = append(blueprintResponseTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, blueprintResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BlueprintResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

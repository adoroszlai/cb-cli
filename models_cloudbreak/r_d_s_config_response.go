package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"

	"encoding/json"
	"github.com/go-swagger/go-swagger/swag"
	"strconv"
)

/*RDSConfigResponse r d s config response

swagger:model RDSConfigResponse
*/
type RDSConfigResponse struct {

	/* list of clusters which use config

	Unique: true
	*/
	ClusterNames []string `json:"clusterNames,omitempty"`

	/* Password to use for the jdbc connection

	Required: true
	*/
	ConnectionPassword string `json:"connectionPassword"`

	/* JDBC connection URL in the form of jdbc:<db-type>://<address>:<port>/<db>

	Required: true
	*/
	ConnectionURL string `json:"connectionURL"`

	/* Username to use for the jdbc connection

	Required: true
	*/
	ConnectionUserName string `json:"connectionUserName"`

	/* creation time of the resource in long
	 */
	CreationDate *int64 `json:"creationDate,omitempty"`

	/* Type of the external database (allowed values: MYSQL, POSTGRES)

	Required: true
	*/
	DatabaseType string `json:"databaseType"`

	/* HDP version for the RDS configuration

	Required: true
	*/
	HdpVersion string `json:"hdpVersion"`

	/* id of the resource
	 */
	ID *int64 `json:"id,omitempty"`

	/* Name of the RDS configuration resource

	Required: true
	*/
	Name string `json:"name"`

	/* custom properties for rds connection

	Unique: true
	*/
	Properties []*RdsConfigProperty `json:"properties,omitempty"`

	/* resource is visible in account
	 */
	PublicInAccount *bool `json:"publicInAccount,omitempty"`

	/* Type of rds (HIVE or RANGER)
	 */
	Type *string `json:"type,omitempty"`

	/* If true, then the RDS configuration will be validated
	 */
	Validated *bool `json:"validated,omitempty"`
}

// Validate validates this r d s config response
func (m *RDSConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterNames(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionPassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionUserName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDatabaseType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHdpVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RDSConfigResponse) validateClusterNames(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusterNames", "body", m.ClusterNames); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterNames); i++ {

		if err := validate.RequiredString("clusterNames"+"."+strconv.Itoa(i), "body", string(m.ClusterNames[i])); err != nil {
			return err
		}

	}

	return nil
}

func (m *RDSConfigResponse) validateConnectionPassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("connectionPassword", "body", string(m.ConnectionPassword)); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("connectionURL", "body", string(m.ConnectionURL)); err != nil {
		return err
	}

	if err := validate.Pattern("connectionURL", "body", string(m.ConnectionURL), `^jdbc:postgresql://[-\w\.]*:?\d*/?\w*`); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateConnectionUserName(formats strfmt.Registry) error {

	if err := validate.RequiredString("connectionUserName", "body", string(m.ConnectionUserName)); err != nil {
		return err
	}

	return nil
}

var rDSConfigResponseTypeDatabaseTypePropEnum []interface{}

func (m *RDSConfigResponse) validateDatabaseTypeEnum(path, location string, value string) error {
	if rDSConfigResponseTypeDatabaseTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["POSTGRES"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			rDSConfigResponseTypeDatabaseTypePropEnum = append(rDSConfigResponseTypeDatabaseTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, rDSConfigResponseTypeDatabaseTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RDSConfigResponse) validateDatabaseType(formats strfmt.Registry) error {

	if err := validate.RequiredString("databaseType", "body", string(m.DatabaseType)); err != nil {
		return err
	}

	if err := m.validateDatabaseTypeEnum("databaseType", "body", m.DatabaseType); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateHdpVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("hdpVersion", "body", string(m.HdpVersion)); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if err := validate.UniqueItems("properties", "body", m.Properties); err != nil {
		return err
	}

	for i := 0; i < len(m.Properties); i++ {

		if m.Properties[i] != nil {

			if err := m.Properties[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var rDSConfigResponseTypeTypePropEnum []interface{}

func (m *RDSConfigResponse) validateTypeEnum(path, location string, value string) error {
	if rDSConfigResponseTypeTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["HIVE","RANGER","DRUID"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			rDSConfigResponseTypeTypePropEnum = append(rDSConfigResponseTypeTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, rDSConfigResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RDSConfigResponse) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

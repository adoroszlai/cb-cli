package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*RdsTestResult rds test result

swagger:model RdsTestResult
*/
type RdsTestResult struct {

	/* result of Ldap connection test

	Required: true
	*/
	ConnectionResult string `json:"connectionResult"`
}

// Validate validates this rds test result
func (m *RdsTestResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionResult(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RdsTestResult) validateConnectionResult(formats strfmt.Registry) error {

	if err := validate.RequiredString("connectionResult", "body", string(m.ConnectionResult)); err != nil {
		return err
	}

	return nil
}

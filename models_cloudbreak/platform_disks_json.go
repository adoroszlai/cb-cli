package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*PlatformDisksJSON platform disks json

swagger:model PlatformDisksJson
*/
type PlatformDisksJSON struct {

	/* default disks
	 */
	DefaultDisks map[string]string `json:"defaultDisks,omitempty"`

	/* disk mappings
	 */
	DiskMappings map[string]map[string]string `json:"diskMappings,omitempty"`

	/* disk types
	 */
	DiskTypes map[string][]string `json:"diskTypes,omitempty"`
}

// Validate validates this platform disks json
func (m *PlatformDisksJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultDisks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDiskMappings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDiskTypes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformDisksJSON) validateDefaultDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultDisks) { // not required
		return nil
	}

	if err := validate.Required("defaultDisks", "body", m.DefaultDisks); err != nil {
		return err
	}

	return nil
}

func (m *PlatformDisksJSON) validateDiskMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.DiskMappings) { // not required
		return nil
	}

	if swag.IsZero(m.DiskMappings) { // not required
		return nil
	}

	return nil
}

func (m *PlatformDisksJSON) validateDiskTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.DiskTypes) { // not required
		return nil
	}

	if err := validate.Required("diskTypes", "body", m.DiskTypes); err != nil {
		return err
	}

	return nil
}
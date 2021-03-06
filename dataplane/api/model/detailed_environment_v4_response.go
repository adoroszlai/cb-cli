// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DetailedEnvironmentV4Response detailed environment v4 response
// swagger:model DetailedEnvironmentV4Response
type DetailedEnvironmentV4Response struct {

	// Cloud platform of the environment.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// Name of the credential of the environment.
	CredentialName string `json:"credentialName,omitempty"`

	// RDS configurations in the environment.
	// Unique: true
	Databases []*DatabaseV4Response `json:"databases"`

	// Names of the datalake clusters created in the environment.
	// Unique: true
	DatalakeClusterNames []string `json:"datalakeClusterNames"`

	// Datalake clusters created in the environment.
	// Unique: true
	DatalakeClusters []*StackViewV4Response `json:"datalakeClusters"`

	// datalake resources
	// Unique: true
	DatalakeResources []*DatalakeResourcesV4Response `json:"datalakeResources"`

	// Datalake cluster resources registered to the environment.
	// Unique: true
	DatalakeResourcesNames []string `json:"datalakeResourcesNames"`

	// description of the resource
	Description string `json:"description,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// Kerberos configs in the environment.
	// Unique: true
	Kerberoses []*KerberosV4Response `json:"kerberoses"`

	// Kubernetes configurations in the environment.
	// Unique: true
	Kubernetes []*KubernetesV4Response `json:"kubernetes"`

	// LDAP configurations in the environment.
	// Unique: true
	Ldaps []*LdapV4Response `json:"ldaps"`

	// Location of the environment.
	Location *LocationV4Response `json:"location,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`

	// Network related specifics of the environment.
	Network *EnvironmentNetworkV4Response `json:"network,omitempty"`

	// Proxy configurations in the environment.
	// Unique: true
	Proxies []*ProxyV4Response `json:"proxies"`

	// Regions of the environment.
	Regions *CompactRegionV4Response `json:"regions,omitempty"`

	// Names of the workload clusters created in the environment.
	// Unique: true
	WorkloadClusterNames []string `json:"workloadClusterNames"`

	// Workload clusters created in the environment.
	// Unique: true
	WorkloadClusters []*StackViewV4Response `json:"workloadClusters"`

	// workspace
	Workspace *WorkspaceResourceV4Response `json:"workspace,omitempty"`
}

// Validate validates this detailed environment v4 response
func (m *DetailedEnvironmentV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalakeClusterNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalakeClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalakeResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalakeResourcesNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKerberoses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLdaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedEnvironmentV4Response) validateDatabases(formats strfmt.Registry) error {

	if swag.IsZero(m.Databases) { // not required
		return nil
	}

	if err := validate.UniqueItems("databases", "body", m.Databases); err != nil {
		return err
	}

	for i := 0; i < len(m.Databases); i++ {
		if swag.IsZero(m.Databases[i]) { // not required
			continue
		}

		if m.Databases[i] != nil {
			if err := m.Databases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("databases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateDatalakeClusterNames(formats strfmt.Registry) error {

	if swag.IsZero(m.DatalakeClusterNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("datalakeClusterNames", "body", m.DatalakeClusterNames); err != nil {
		return err
	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateDatalakeClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.DatalakeClusters) { // not required
		return nil
	}

	if err := validate.UniqueItems("datalakeClusters", "body", m.DatalakeClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.DatalakeClusters); i++ {
		if swag.IsZero(m.DatalakeClusters[i]) { // not required
			continue
		}

		if m.DatalakeClusters[i] != nil {
			if err := m.DatalakeClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datalakeClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateDatalakeResources(formats strfmt.Registry) error {

	if swag.IsZero(m.DatalakeResources) { // not required
		return nil
	}

	if err := validate.UniqueItems("datalakeResources", "body", m.DatalakeResources); err != nil {
		return err
	}

	for i := 0; i < len(m.DatalakeResources); i++ {
		if swag.IsZero(m.DatalakeResources[i]) { // not required
			continue
		}

		if m.DatalakeResources[i] != nil {
			if err := m.DatalakeResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datalakeResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateDatalakeResourcesNames(formats strfmt.Registry) error {

	if swag.IsZero(m.DatalakeResourcesNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("datalakeResourcesNames", "body", m.DatalakeResourcesNames); err != nil {
		return err
	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateKerberoses(formats strfmt.Registry) error {

	if swag.IsZero(m.Kerberoses) { // not required
		return nil
	}

	if err := validate.UniqueItems("kerberoses", "body", m.Kerberoses); err != nil {
		return err
	}

	for i := 0; i < len(m.Kerberoses); i++ {
		if swag.IsZero(m.Kerberoses[i]) { // not required
			continue
		}

		if m.Kerberoses[i] != nil {
			if err := m.Kerberoses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kerberoses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateKubernetes(formats strfmt.Registry) error {

	if swag.IsZero(m.Kubernetes) { // not required
		return nil
	}

	if err := validate.UniqueItems("kubernetes", "body", m.Kubernetes); err != nil {
		return err
	}

	for i := 0; i < len(m.Kubernetes); i++ {
		if swag.IsZero(m.Kubernetes[i]) { // not required
			continue
		}

		if m.Kubernetes[i] != nil {
			if err := m.Kubernetes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kubernetes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateLdaps(formats strfmt.Registry) error {

	if swag.IsZero(m.Ldaps) { // not required
		return nil
	}

	if err := validate.UniqueItems("ldaps", "body", m.Ldaps); err != nil {
		return err
	}

	for i := 0; i < len(m.Ldaps); i++ {
		if swag.IsZero(m.Ldaps[i]) { // not required
			continue
		}

		if m.Ldaps[i] != nil {
			if err := m.Ldaps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ldaps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateProxies(formats strfmt.Registry) error {

	if swag.IsZero(m.Proxies) { // not required
		return nil
	}

	if err := validate.UniqueItems("proxies", "body", m.Proxies); err != nil {
		return err
	}

	for i := 0; i < len(m.Proxies); i++ {
		if swag.IsZero(m.Proxies[i]) { // not required
			continue
		}

		if m.Proxies[i] != nil {
			if err := m.Proxies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if m.Regions != nil {
		if err := m.Regions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regions")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateWorkloadClusterNames(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkloadClusterNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("workloadClusterNames", "body", m.WorkloadClusterNames); err != nil {
		return err
	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateWorkloadClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkloadClusters) { // not required
		return nil
	}

	if err := validate.UniqueItems("workloadClusters", "body", m.WorkloadClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.WorkloadClusters); i++ {
		if swag.IsZero(m.WorkloadClusters[i]) { // not required
			continue
		}

		if m.WorkloadClusters[i] != nil {
			if err := m.WorkloadClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workloadClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedEnvironmentV4Response) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailedEnvironmentV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedEnvironmentV4Response) UnmarshalBinary(b []byte) error {
	var res DetailedEnvironmentV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

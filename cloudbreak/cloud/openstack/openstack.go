package openstack

import "github.com/hortonworks/cb-cli/cloudbreak/v2/cloud"

var name string

func init() {
	name = string(cloud.OPENSTACK)
	cloud.CloudProviders[cloud.OPENSTACK] = new(OpenstackProvider)
}

type OpenstackProvider struct {
	cloud.DefaultCloudProvider
}

func (p *OpenstackProvider) GetName() *string {
	return &name
}

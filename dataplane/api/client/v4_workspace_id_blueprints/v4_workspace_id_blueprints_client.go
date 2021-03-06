// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4 workspace id blueprints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4 workspace id blueprints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateBlueprintInWorkspace creates blueprint in workspace

Cluster definitions are a declarative definition of a Hadoop cluster. With a blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) CreateBlueprintInWorkspace(params *CreateBlueprintInWorkspaceParams) (*CreateBlueprintInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlueprintInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBlueprintInWorkspace",
		Method:             "POST",
		PathPattern:        "/v4/{workspaceId}/blueprints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateBlueprintInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateBlueprintInWorkspaceOK), nil

}

/*
CreateRecommendationForWorkspace creates a recommendation that advises cloud resources for the given blueprint

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) CreateRecommendationForWorkspace(params *CreateRecommendationForWorkspaceParams) (*CreateRecommendationForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRecommendationForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRecommendationForWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/blueprints/recommendation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRecommendationForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRecommendationForWorkspaceOK), nil

}

/*
DeleteBlueprintInWorkspace deletes blueprint by name in workspace

Cluster definitions are a declarative definition of a Hadoop cluster. With a blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) DeleteBlueprintInWorkspace(params *DeleteBlueprintInWorkspaceParams) (*DeleteBlueprintInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBlueprintInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBlueprintInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/blueprints/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBlueprintInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteBlueprintInWorkspaceOK), nil

}

/*
DeleteBlueprintsInWorkspace deletes multiple blueprints by name in workspace

Cluster definitions are a declarative definition of a Hadoop cluster. With a blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) DeleteBlueprintsInWorkspace(params *DeleteBlueprintsInWorkspaceParams) (*DeleteBlueprintsInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBlueprintsInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBlueprintsInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/blueprints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBlueprintsInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteBlueprintsInWorkspaceOK), nil

}

/*
GetBlueprintCustomParameters returns custom parameters
*/
func (a *Client) GetBlueprintCustomParameters(params *GetBlueprintCustomParametersParams) (*GetBlueprintCustomParametersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintCustomParametersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlueprintCustomParameters",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/blueprints/{name}/parameters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBlueprintCustomParametersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBlueprintCustomParametersOK), nil

}

/*
GetBlueprintInWorkspace gets blueprint by name in workspace

Cluster definitions are a declarative definition of a Hadoop cluster. With a blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetBlueprintInWorkspace(params *GetBlueprintInWorkspaceParams) (*GetBlueprintInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlueprintInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/blueprints/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBlueprintInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBlueprintInWorkspaceOK), nil

}

/*
GetBlueprintRequestFromName retrieves validation request by blueprint name

Cluster definitions are a declarative definition of a Hadoop cluster. With a blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) GetBlueprintRequestFromName(params *GetBlueprintRequestFromNameParams) (*GetBlueprintRequestFromNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintRequestFromNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlueprintRequestFromName",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/blueprints/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBlueprintRequestFromNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBlueprintRequestFromNameOK), nil

}

/*
ListBlueprintsByWorkspace lists blueprints for the given workspace

Cluster definitions are a declarative definition of a Hadoop cluster. With a blueprint, you specify a stack, the component layout and the configurations to materialize a Hadoop cluster instance. Hostgroups defined in blueprints can be associated to different templates, thus you can spin up a highly available cluster running on different instance types. This will give you the option to group your Hadoop services based on resource needs (e.g. high I/O, CPU or memory) and create an infrastructure which fits your workload best.
*/
func (a *Client) ListBlueprintsByWorkspace(params *ListBlueprintsByWorkspaceParams) (*ListBlueprintsByWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBlueprintsByWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBlueprintsByWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/blueprints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListBlueprintsByWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListBlueprintsByWorkspaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

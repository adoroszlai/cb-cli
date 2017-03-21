package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new rdsconfigs API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rdsconfigs API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeletePrivateRds deletes private r d s configuration by name

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DeletePrivateRds(params *DeletePrivateRdsParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateRdsParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deletePrivateRds",
		Method:             "DELETE",
		PathPattern:        "/rdsconfigs/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePrivateRdsReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeletePublicRds deletes public owned or private r d s configuration by name

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DeletePublicRds(params *DeletePublicRdsParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicRdsParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deletePublicRds",
		Method:             "DELETE",
		PathPattern:        "/rdsconfigs/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePublicRdsReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteRds deletes r d s configuration by id

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DeleteRds(params *DeleteRdsParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRdsParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteRds",
		Method:             "DELETE",
		PathPattern:        "/rdsconfigs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRdsReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetPrivateRds retrieves a private r d s configuration by name

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetPrivateRds(params *GetPrivateRdsParams) (*GetPrivateRdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateRdsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPrivateRds",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateRdsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateRdsOK), nil
}

/*
GetPrivatesRds retrieves private r d s configurations

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetPrivatesRds(params *GetPrivatesRdsParams) (*GetPrivatesRdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivatesRdsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPrivatesRds",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivatesRdsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivatesRdsOK), nil
}

/*
GetPublicRds retrieves a public or private owned r d s configuration by name

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetPublicRds(params *GetPublicRdsParams) (*GetPublicRdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicRdsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPublicRds",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicRdsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicRdsOK), nil
}

/*
GetPublicsRds retrieves public and private owned r d s configurations

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetPublicsRds(params *GetPublicsRdsParams) (*GetPublicsRdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicsRdsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPublicsRds",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicsRdsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicsRdsOK), nil
}

/*
GetRds retrieves r d s configuration by id

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetRds(params *GetRdsParams) (*GetRdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRdsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getRds",
		Method:             "GET",
		PathPattern:        "/rdsconfigs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRdsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRdsOK), nil
}

/*
PostPrivateRds creates r d s configuration as private resource

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) PostPrivateRds(params *PostPrivateRdsParams) (*PostPrivateRdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrivateRdsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "postPrivateRds",
		Method:             "POST",
		PathPattern:        "/rdsconfigs/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPrivateRdsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrivateRdsOK), nil
}

/*
PotPublicRds creates r d s configuration as public resource

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) PotPublicRds(params *PotPublicRdsParams) (*PotPublicRdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPotPublicRdsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "potPublicRds",
		Method:             "POST",
		PathPattern:        "/rdsconfigs/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PotPublicRdsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PotPublicRdsOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

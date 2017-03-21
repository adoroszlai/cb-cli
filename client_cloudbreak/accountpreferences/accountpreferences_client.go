package accountpreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new accountpreferences API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accountpreferences API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
GetAccountPreferencesEndpoint retrieves account preferences for admin user

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) GetAccountPreferencesEndpoint(params *GetAccountPreferencesEndpointParams) (*GetAccountPreferencesEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountPreferencesEndpointParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getAccountPreferencesEndpoint",
		Method:             "GET",
		PathPattern:        "/accountpreferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccountPreferencesEndpointReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountPreferencesEndpointOK), nil
}

/*
IsPlatformSelectionDisabled is platform selection disabled

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) IsPlatformSelectionDisabled(params *IsPlatformSelectionDisabledParams) (*IsPlatformSelectionDisabledOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsPlatformSelectionDisabledParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "isPlatformSelectionDisabled",
		Method:             "GET",
		PathPattern:        "/accountpreferences/isplatformselectiondisabled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsPlatformSelectionDisabledReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*IsPlatformSelectionDisabledOK), nil
}

/*
PostAccountPreferencesEndpoint posts account preferences of admin user

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) PostAccountPreferencesEndpoint(params *PostAccountPreferencesEndpointParams) (*PostAccountPreferencesEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAccountPreferencesEndpointParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "postAccountPreferencesEndpoint",
		Method:             "POST",
		PathPattern:        "/accountpreferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostAccountPreferencesEndpointReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAccountPreferencesEndpointOK), nil
}

/*
PutAccountPreferencesEndpoint updates account preferences of admin user

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) PutAccountPreferencesEndpoint(params *PutAccountPreferencesEndpointParams) (*PutAccountPreferencesEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAccountPreferencesEndpointParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "putAccountPreferencesEndpoint",
		Method:             "PUT",
		PathPattern:        "/accountpreferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutAccountPreferencesEndpointReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutAccountPreferencesEndpointOK), nil
}

/*
ValidateAccountPreferencesEndpoint validates account preferences of all stacks

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) ValidateAccountPreferencesEndpoint(params *ValidateAccountPreferencesEndpointParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateAccountPreferencesEndpointParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "validateAccountPreferencesEndpoint",
		Method:             "GET",
		PathPattern:        "/accountpreferences/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ValidateAccountPreferencesEndpointReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

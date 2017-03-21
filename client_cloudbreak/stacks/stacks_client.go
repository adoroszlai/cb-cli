package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new stacks API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stacks API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeleteInstanceStack deletes instance resource from stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteInstanceStack(params *DeleteInstanceStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstanceStackParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteInstanceStack",
		Method:             "DELETE",
		PathPattern:        "/stacks/{stackId}/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteInstanceStackReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeletePrivateStack deletes private stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeletePrivateStack(params *DeletePrivateStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateStackParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deletePrivateStack",
		Method:             "DELETE",
		PathPattern:        "/stacks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePrivateStackReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeletePublicStack deletes public owned or private stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeletePublicStack(params *DeletePublicStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicStackParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deletePublicStack",
		Method:             "DELETE",
		PathPattern:        "/stacks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePublicStackReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteStack deletes stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) DeleteStack(params *DeleteStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStackParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteStack",
		Method:             "DELETE",
		PathPattern:        "/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteStackReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetAllStack retrieves all stacks

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetAllStack(params *GetAllStackParams) (*GetAllStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getAllStack",
		Method:             "GET",
		PathPattern:        "/stacks/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllStackOK), nil
}

/*
GetCertificateStack retrieves the TLS certificate used by the gateway

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetCertificateStack(params *GetCertificateStackParams) (*GetCertificateStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCertificateStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getCertificateStack",
		Method:             "GET",
		PathPattern:        "/stacks/{id}/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCertificateStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCertificateStackOK), nil
}

/*
GetPrivateStack retrieves a private stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPrivateStack(params *GetPrivateStackParams) (*GetPrivateStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPrivateStack",
		Method:             "GET",
		PathPattern:        "/stacks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateStackOK), nil
}

/*
GetPrivatesStack retrieves private stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPrivatesStack(params *GetPrivatesStackParams) (*GetPrivatesStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivatesStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPrivatesStack",
		Method:             "GET",
		PathPattern:        "/stacks/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivatesStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivatesStackOK), nil
}

/*
GetPublicStack retrieves a public or private owned stack by name

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPublicStack(params *GetPublicStackParams) (*GetPublicStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPublicStack",
		Method:             "GET",
		PathPattern:        "/stacks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicStackOK), nil
}

/*
GetPublicsStack retrieves public and private owned stacks

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetPublicsStack(params *GetPublicsStackParams) (*GetPublicsStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicsStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPublicsStack",
		Method:             "GET",
		PathPattern:        "/stacks/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicsStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicsStackOK), nil
}

/*
GetStack retrieves stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetStack(params *GetStackParams) (*GetStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getStack",
		Method:             "GET",
		PathPattern:        "/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackOK), nil
}

/*
GetStackForAmbari retrieves stack by ambari address

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetStackForAmbari(params *GetStackForAmbariParams) (*GetStackForAmbariOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackForAmbariParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getStackForAmbari",
		Method:             "POST",
		PathPattern:        "/stacks/ambari",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackForAmbariReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackForAmbariOK), nil
}

/*
PostPrivateStack creates stack as private resource

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PostPrivateStack(params *PostPrivateStackParams) (*PostPrivateStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrivateStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "postPrivateStack",
		Method:             "POST",
		PathPattern:        "/stacks/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPrivateStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrivateStackOK), nil
}

/*
PostPublicStack creates stack as public resource

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PostPublicStack(params *PostPublicStackParams) (*PostPublicStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublicStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "postPublicStack",
		Method:             "POST",
		PathPattern:        "/stacks/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPublicStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPublicStackOK), nil
}

/*
PutStack updates stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutStack(params *PutStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutStackParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "putStack",
		Method:             "PUT",
		PathPattern:        "/stacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutStackReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
StatusStack retrieves stack status by stack id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) StatusStack(params *StatusStackParams) (*StatusStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "statusStack",
		Method:             "GET",
		PathPattern:        "/stacks/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*StatusStackOK), nil
}

/*
ValidateStack validates stack

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) ValidateStack(params *ValidateStackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateStackParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "validateStack",
		Method:             "POST",
		PathPattern:        "/stacks/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ValidateStackReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
VariantsStack retrieves available platform variants

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) VariantsStack(params *VariantsStackParams) (*VariantsStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVariantsStackParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "variantsStack",
		Method:             "GET",
		PathPattern:        "/stacks/platformVariants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VariantsStackReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*VariantsStackOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

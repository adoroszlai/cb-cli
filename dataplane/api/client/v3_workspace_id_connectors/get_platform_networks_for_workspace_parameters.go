// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewGetPlatformNetworksForWorkspaceParams creates a new GetPlatformNetworksForWorkspaceParams object
// with the default values initialized.
func NewGetPlatformNetworksForWorkspaceParams() *GetPlatformNetworksForWorkspaceParams {
	var ()
	return &GetPlatformNetworksForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformNetworksForWorkspaceParamsWithTimeout creates a new GetPlatformNetworksForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformNetworksForWorkspaceParamsWithTimeout(timeout time.Duration) *GetPlatformNetworksForWorkspaceParams {
	var ()
	return &GetPlatformNetworksForWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetPlatformNetworksForWorkspaceParamsWithContext creates a new GetPlatformNetworksForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformNetworksForWorkspaceParamsWithContext(ctx context.Context) *GetPlatformNetworksForWorkspaceParams {
	var ()
	return &GetPlatformNetworksForWorkspaceParams{

		Context: ctx,
	}
}

// NewGetPlatformNetworksForWorkspaceParamsWithHTTPClient creates a new GetPlatformNetworksForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformNetworksForWorkspaceParamsWithHTTPClient(client *http.Client) *GetPlatformNetworksForWorkspaceParams {
	var ()
	return &GetPlatformNetworksForWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetPlatformNetworksForWorkspaceParams contains all the parameters to send to the API endpoint
for the get platform networks for workspace operation typically these are written to a http.Request
*/
type GetPlatformNetworksForWorkspaceParams struct {

	/*Body*/
	Body *model.PlatformResourceRequestJSON
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithTimeout(timeout time.Duration) *GetPlatformNetworksForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithContext(ctx context.Context) *GetPlatformNetworksForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithHTTPClient(client *http.Client) *GetPlatformNetworksForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithBody(body *model.PlatformResourceRequestJSON) *GetPlatformNetworksForWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetBody(body *model.PlatformResourceRequestJSON) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetPlatformNetworksForWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get platform networks for workspace params
func (o *GetPlatformNetworksForWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformNetworksForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(model.PlatformResourceRequestJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
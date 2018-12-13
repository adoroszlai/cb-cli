// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllTenantsParams creates a new GetAllTenantsParams object
// with the default values initialized.
func NewGetAllTenantsParams() *GetAllTenantsParams {

	return &GetAllTenantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllTenantsParamsWithTimeout creates a new GetAllTenantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllTenantsParamsWithTimeout(timeout time.Duration) *GetAllTenantsParams {

	return &GetAllTenantsParams{

		timeout: timeout,
	}
}

// NewGetAllTenantsParamsWithContext creates a new GetAllTenantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllTenantsParamsWithContext(ctx context.Context) *GetAllTenantsParams {

	return &GetAllTenantsParams{

		Context: ctx,
	}
}

// NewGetAllTenantsParamsWithHTTPClient creates a new GetAllTenantsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllTenantsParamsWithHTTPClient(client *http.Client) *GetAllTenantsParams {

	return &GetAllTenantsParams{
		HTTPClient: client,
	}
}

/*GetAllTenantsParams contains all the parameters to send to the API endpoint
for the get all tenants operation typically these are written to a http.Request
*/
type GetAllTenantsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all tenants params
func (o *GetAllTenantsParams) WithTimeout(timeout time.Duration) *GetAllTenantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all tenants params
func (o *GetAllTenantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all tenants params
func (o *GetAllTenantsParams) WithContext(ctx context.Context) *GetAllTenantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all tenants params
func (o *GetAllTenantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all tenants params
func (o *GetAllTenantsParams) WithHTTPClient(client *http.Client) *GetAllTenantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all tenants params
func (o *GetAllTenantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllTenantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

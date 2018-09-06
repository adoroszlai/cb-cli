// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_smartsensesubscriptions

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
)

// NewGetDefaultSmartSenseSubscriptionInOrganizationParams creates a new GetDefaultSmartSenseSubscriptionInOrganizationParams object
// with the default values initialized.
func NewGetDefaultSmartSenseSubscriptionInOrganizationParams() *GetDefaultSmartSenseSubscriptionInOrganizationParams {
	var ()
	return &GetDefaultSmartSenseSubscriptionInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDefaultSmartSenseSubscriptionInOrganizationParamsWithTimeout creates a new GetDefaultSmartSenseSubscriptionInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDefaultSmartSenseSubscriptionInOrganizationParamsWithTimeout(timeout time.Duration) *GetDefaultSmartSenseSubscriptionInOrganizationParams {
	var ()
	return &GetDefaultSmartSenseSubscriptionInOrganizationParams{

		timeout: timeout,
	}
}

// NewGetDefaultSmartSenseSubscriptionInOrganizationParamsWithContext creates a new GetDefaultSmartSenseSubscriptionInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDefaultSmartSenseSubscriptionInOrganizationParamsWithContext(ctx context.Context) *GetDefaultSmartSenseSubscriptionInOrganizationParams {
	var ()
	return &GetDefaultSmartSenseSubscriptionInOrganizationParams{

		Context: ctx,
	}
}

// NewGetDefaultSmartSenseSubscriptionInOrganizationParamsWithHTTPClient creates a new GetDefaultSmartSenseSubscriptionInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDefaultSmartSenseSubscriptionInOrganizationParamsWithHTTPClient(client *http.Client) *GetDefaultSmartSenseSubscriptionInOrganizationParams {
	var ()
	return &GetDefaultSmartSenseSubscriptionInOrganizationParams{
		HTTPClient: client,
	}
}

/*GetDefaultSmartSenseSubscriptionInOrganizationParams contains all the parameters to send to the API endpoint
for the get default smart sense subscription in organization operation typically these are written to a http.Request
*/
type GetDefaultSmartSenseSubscriptionInOrganizationParams struct {

	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get default smart sense subscription in organization params
func (o *GetDefaultSmartSenseSubscriptionInOrganizationParams) WithTimeout(timeout time.Duration) *GetDefaultSmartSenseSubscriptionInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get default smart sense subscription in organization params
func (o *GetDefaultSmartSenseSubscriptionInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get default smart sense subscription in organization params
func (o *GetDefaultSmartSenseSubscriptionInOrganizationParams) WithContext(ctx context.Context) *GetDefaultSmartSenseSubscriptionInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get default smart sense subscription in organization params
func (o *GetDefaultSmartSenseSubscriptionInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get default smart sense subscription in organization params
func (o *GetDefaultSmartSenseSubscriptionInOrganizationParams) WithHTTPClient(client *http.Client) *GetDefaultSmartSenseSubscriptionInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get default smart sense subscription in organization params
func (o *GetDefaultSmartSenseSubscriptionInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get default smart sense subscription in organization params
func (o *GetDefaultSmartSenseSubscriptionInOrganizationParams) WithOrganizationID(organizationID int64) *GetDefaultSmartSenseSubscriptionInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get default smart sense subscription in organization params
func (o *GetDefaultSmartSenseSubscriptionInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDefaultSmartSenseSubscriptionInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
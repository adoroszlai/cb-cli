// Code generated by go-swagger; DO NOT EDIT.

package v1accountpreferences

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

// NewIsPlatformSelectionDisabledParams creates a new IsPlatformSelectionDisabledParams object
// with the default values initialized.
func NewIsPlatformSelectionDisabledParams() *IsPlatformSelectionDisabledParams {

	return &IsPlatformSelectionDisabledParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsPlatformSelectionDisabledParamsWithTimeout creates a new IsPlatformSelectionDisabledParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsPlatformSelectionDisabledParamsWithTimeout(timeout time.Duration) *IsPlatformSelectionDisabledParams {

	return &IsPlatformSelectionDisabledParams{

		timeout: timeout,
	}
}

// NewIsPlatformSelectionDisabledParamsWithContext creates a new IsPlatformSelectionDisabledParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsPlatformSelectionDisabledParamsWithContext(ctx context.Context) *IsPlatformSelectionDisabledParams {

	return &IsPlatformSelectionDisabledParams{

		Context: ctx,
	}
}

// NewIsPlatformSelectionDisabledParamsWithHTTPClient creates a new IsPlatformSelectionDisabledParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsPlatformSelectionDisabledParamsWithHTTPClient(client *http.Client) *IsPlatformSelectionDisabledParams {

	return &IsPlatformSelectionDisabledParams{
		HTTPClient: client,
	}
}

/*IsPlatformSelectionDisabledParams contains all the parameters to send to the API endpoint
for the is platform selection disabled operation typically these are written to a http.Request
*/
type IsPlatformSelectionDisabledParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is platform selection disabled params
func (o *IsPlatformSelectionDisabledParams) WithTimeout(timeout time.Duration) *IsPlatformSelectionDisabledParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is platform selection disabled params
func (o *IsPlatformSelectionDisabledParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is platform selection disabled params
func (o *IsPlatformSelectionDisabledParams) WithContext(ctx context.Context) *IsPlatformSelectionDisabledParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is platform selection disabled params
func (o *IsPlatformSelectionDisabledParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is platform selection disabled params
func (o *IsPlatformSelectionDisabledParams) WithHTTPClient(client *http.Client) *IsPlatformSelectionDisabledParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is platform selection disabled params
func (o *IsPlatformSelectionDisabledParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IsPlatformSelectionDisabledParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetNetworkParams creates a new GetNetworkParams object
// with the default values initialized.
func NewGetNetworkParams() *GetNetworkParams {
	var ()
	return &GetNetworkParams{}
}

/*GetNetworkParams contains all the parameters to send to the API endpoint
for the get network operation typically these are written to a http.Request
*/
type GetNetworkParams struct {

	/*ID*/
	ID int64
}

// WithID adds the id to the get network params
func (o *GetNetworkParams) WithID(id int64) *GetNetworkParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

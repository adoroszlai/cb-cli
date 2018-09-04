// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_mpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// ListManagementPacksByOrganizationReader is a Reader for the ListManagementPacksByOrganization structure.
type ListManagementPacksByOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListManagementPacksByOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListManagementPacksByOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListManagementPacksByOrganizationOK creates a ListManagementPacksByOrganizationOK with default headers values
func NewListManagementPacksByOrganizationOK() *ListManagementPacksByOrganizationOK {
	return &ListManagementPacksByOrganizationOK{}
}

/*ListManagementPacksByOrganizationOK handles this case with default header values.

successful operation
*/
type ListManagementPacksByOrganizationOK struct {
	Payload []*models_cloudbreak.ManagementPackResponse
}

func (o *ListManagementPacksByOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /v3/{organizationId}/mpacks][%d] listManagementPacksByOrganizationOK  %+v", 200, o.Payload)
}

func (o *ListManagementPacksByOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
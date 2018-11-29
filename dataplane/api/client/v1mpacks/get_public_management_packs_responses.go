// Code generated by go-swagger; DO NOT EDIT.

package v1mpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetPublicManagementPacksReader is a Reader for the GetPublicManagementPacks structure.
type GetPublicManagementPacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicManagementPacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicManagementPacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicManagementPacksOK creates a GetPublicManagementPacksOK with default headers values
func NewGetPublicManagementPacksOK() *GetPublicManagementPacksOK {
	return &GetPublicManagementPacksOK{}
}

/*GetPublicManagementPacksOK handles this case with default header values.

successful operation
*/
type GetPublicManagementPacksOK struct {
	Payload []*model.ManagementPackResponse
}

func (o *GetPublicManagementPacksOK) Error() string {
	return fmt.Sprintf("[GET /v1/mpacks/account][%d] getPublicManagementPacksOK  %+v", 200, o.Payload)
}

func (o *GetPublicManagementPacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// SetDefaultImageCatalogByNameInWorkspaceReader is a Reader for the SetDefaultImageCatalogByNameInWorkspace structure.
type SetDefaultImageCatalogByNameInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDefaultImageCatalogByNameInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetDefaultImageCatalogByNameInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetDefaultImageCatalogByNameInWorkspaceOK creates a SetDefaultImageCatalogByNameInWorkspaceOK with default headers values
func NewSetDefaultImageCatalogByNameInWorkspaceOK() *SetDefaultImageCatalogByNameInWorkspaceOK {
	return &SetDefaultImageCatalogByNameInWorkspaceOK{}
}

/*SetDefaultImageCatalogByNameInWorkspaceOK handles this case with default header values.

successful operation
*/
type SetDefaultImageCatalogByNameInWorkspaceOK struct {
	Payload *model.ImageCatalogV4Response
}

func (o *SetDefaultImageCatalogByNameInWorkspaceOK) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/image_catalogs/{name}/set_default][%d] setDefaultImageCatalogByNameInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *SetDefaultImageCatalogByNameInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ImageCatalogV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
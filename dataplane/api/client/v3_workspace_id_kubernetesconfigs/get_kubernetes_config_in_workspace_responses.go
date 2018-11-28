// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_kubernetesconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetKubernetesConfigInWorkspaceReader is a Reader for the GetKubernetesConfigInWorkspace structure.
type GetKubernetesConfigInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesConfigInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKubernetesConfigInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKubernetesConfigInWorkspaceOK creates a GetKubernetesConfigInWorkspaceOK with default headers values
func NewGetKubernetesConfigInWorkspaceOK() *GetKubernetesConfigInWorkspaceOK {
	return &GetKubernetesConfigInWorkspaceOK{}
}

/*GetKubernetesConfigInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetKubernetesConfigInWorkspaceOK struct {
	Payload *model.KubernetesConfigResponse
}

func (o *GetKubernetesConfigInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/kubernetesconfigs/{name}][%d] getKubernetesConfigInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesConfigInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.KubernetesConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

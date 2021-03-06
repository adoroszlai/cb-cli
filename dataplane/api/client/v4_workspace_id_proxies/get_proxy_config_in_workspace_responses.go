// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetProxyConfigInWorkspaceReader is a Reader for the GetProxyConfigInWorkspace structure.
type GetProxyConfigInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProxyConfigInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProxyConfigInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProxyConfigInWorkspaceOK creates a GetProxyConfigInWorkspaceOK with default headers values
func NewGetProxyConfigInWorkspaceOK() *GetProxyConfigInWorkspaceOK {
	return &GetProxyConfigInWorkspaceOK{}
}

/*GetProxyConfigInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetProxyConfigInWorkspaceOK struct {
	Payload *model.ProxyV4Response
}

func (o *GetProxyConfigInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/proxies/{name}][%d] getProxyConfigInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetProxyConfigInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

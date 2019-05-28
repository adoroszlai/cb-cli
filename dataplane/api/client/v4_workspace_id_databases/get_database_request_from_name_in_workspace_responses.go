// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetDatabaseRequestFromNameInWorkspaceReader is a Reader for the GetDatabaseRequestFromNameInWorkspace structure.
type GetDatabaseRequestFromNameInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabaseRequestFromNameInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDatabaseRequestFromNameInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDatabaseRequestFromNameInWorkspaceOK creates a GetDatabaseRequestFromNameInWorkspaceOK with default headers values
func NewGetDatabaseRequestFromNameInWorkspaceOK() *GetDatabaseRequestFromNameInWorkspaceOK {
	return &GetDatabaseRequestFromNameInWorkspaceOK{}
}

/*GetDatabaseRequestFromNameInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetDatabaseRequestFromNameInWorkspaceOK struct {
	Payload *model.DatabaseV4Request
}

func (o *GetDatabaseRequestFromNameInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/databases/{name}/request][%d] getDatabaseRequestFromNameInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetDatabaseRequestFromNameInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseV4Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
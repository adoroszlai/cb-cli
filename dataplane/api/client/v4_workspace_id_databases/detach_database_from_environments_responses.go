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

// DetachDatabaseFromEnvironmentsReader is a Reader for the DetachDatabaseFromEnvironments structure.
type DetachDatabaseFromEnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachDatabaseFromEnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDetachDatabaseFromEnvironmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDetachDatabaseFromEnvironmentsOK creates a DetachDatabaseFromEnvironmentsOK with default headers values
func NewDetachDatabaseFromEnvironmentsOK() *DetachDatabaseFromEnvironmentsOK {
	return &DetachDatabaseFromEnvironmentsOK{}
}

/*DetachDatabaseFromEnvironmentsOK handles this case with default header values.

successful operation
*/
type DetachDatabaseFromEnvironmentsOK struct {
	Payload *model.DatabaseV4Response
}

func (o *DetachDatabaseFromEnvironmentsOK) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/databases/{name}/detach][%d] detachDatabaseFromEnvironmentsOK  %+v", 200, o.Payload)
}

func (o *DetachDatabaseFromEnvironmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

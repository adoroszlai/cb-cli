// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DetachKubernetesResourceFromEnvironmentsReader is a Reader for the DetachKubernetesResourceFromEnvironments structure.
type DetachKubernetesResourceFromEnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachKubernetesResourceFromEnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDetachKubernetesResourceFromEnvironmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDetachKubernetesResourceFromEnvironmentsOK creates a DetachKubernetesResourceFromEnvironmentsOK with default headers values
func NewDetachKubernetesResourceFromEnvironmentsOK() *DetachKubernetesResourceFromEnvironmentsOK {
	return &DetachKubernetesResourceFromEnvironmentsOK{}
}

/*DetachKubernetesResourceFromEnvironmentsOK handles this case with default header values.

successful operation
*/
type DetachKubernetesResourceFromEnvironmentsOK struct {
	Payload *model.KubernetesV4Response
}

func (o *DetachKubernetesResourceFromEnvironmentsOK) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/kubernetes/{name}/detach][%d] detachKubernetesResourceFromEnvironmentsOK  %+v", 200, o.Payload)
}

func (o *DetachKubernetesResourceFromEnvironmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.KubernetesV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

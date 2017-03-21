package constraints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetPrivateConstraintReader is a Reader for the GetPrivateConstraint structure.
type GetPrivateConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPrivateConstraintReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateConstraintOK creates a GetPrivateConstraintOK with default headers values
func NewGetPrivateConstraintOK() *GetPrivateConstraintOK {
	return &GetPrivateConstraintOK{}
}

/*GetPrivateConstraintOK handles this case with default header values.

successful operation
*/
type GetPrivateConstraintOK struct {
	Payload *models_cloudbreak.ConstraintTemplateResponse
}

func (o *GetPrivateConstraintOK) Error() string {
	return fmt.Sprintf("[GET /constraints/user/{name}][%d] getPrivateConstraintOK  %+v", 200, o.Payload)
}

func (o *GetPrivateConstraintOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ConstraintTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

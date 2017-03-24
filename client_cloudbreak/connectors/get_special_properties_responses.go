package connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetSpecialPropertiesReader is a Reader for the GetSpecialProperties structure.
type GetSpecialPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSpecialPropertiesReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSpecialPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSpecialPropertiesOK creates a GetSpecialPropertiesOK with default headers values
func NewGetSpecialPropertiesOK() *GetSpecialPropertiesOK {
	return &GetSpecialPropertiesOK{}
}

/*GetSpecialPropertiesOK handles this case with default header values.

successful operation
*/
type GetSpecialPropertiesOK struct {
	Payload GetSpecialPropertiesOKBodyBody
}

func (o *GetSpecialPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /connectors/custom][%d] getSpecialPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetSpecialPropertiesOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSpecialPropertiesOKBodyBody get special properties o k body body

swagger:model GetSpecialPropertiesOKBodyBody
*/
type GetSpecialPropertiesOKBodyBody map[string]bool

// Validate validates this get special properties o k body body
func (o GetSpecialPropertiesOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("getSpecialPropertiesOK", "body", o); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
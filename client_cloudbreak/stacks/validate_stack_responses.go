package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// ValidateStackReader is a Reader for the ValidateStack structure.
type ValidateStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ValidateStackReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewValidateStackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewValidateStackDefault creates a ValidateStackDefault with default headers values
func NewValidateStackDefault(code int) *ValidateStackDefault {
	return &ValidateStackDefault{
		_statusCode: code,
	}
}

/*ValidateStackDefault handles this case with default header values.

successful operation
*/
type ValidateStackDefault struct {
	_statusCode int
}

// Code gets the status code for the validate stack default response
func (o *ValidateStackDefault) Code() int {
	return o._statusCode
}

func (o *ValidateStackDefault) Error() string {
	return fmt.Sprintf("[POST /stacks/validate][%d] validateStack default ", o._statusCode)
}

func (o *ValidateStackDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

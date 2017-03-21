package constraints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeletePrivateConstraintReader is a Reader for the DeletePrivateConstraint structure.
type DeletePrivateConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeletePrivateConstraintReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeletePrivateConstraintDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeletePrivateConstraintDefault creates a DeletePrivateConstraintDefault with default headers values
func NewDeletePrivateConstraintDefault(code int) *DeletePrivateConstraintDefault {
	return &DeletePrivateConstraintDefault{
		_statusCode: code,
	}
}

/*DeletePrivateConstraintDefault handles this case with default header values.

successful operation
*/
type DeletePrivateConstraintDefault struct {
	_statusCode int
}

// Code gets the status code for the delete private constraint default response
func (o *DeletePrivateConstraintDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateConstraintDefault) Error() string {
	return fmt.Sprintf("[DELETE /constraints/user/{name}][%d] deletePrivateConstraint default ", o._statusCode)
}

func (o *DeletePrivateConstraintDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteRdsReader is a Reader for the DeleteRds structure.
type DeleteRdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteRdsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeleteRdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteRdsDefault creates a DeleteRdsDefault with default headers values
func NewDeleteRdsDefault(code int) *DeleteRdsDefault {
	return &DeleteRdsDefault{
		_statusCode: code,
	}
}

/*DeleteRdsDefault handles this case with default header values.

successful operation
*/
type DeleteRdsDefault struct {
	_statusCode int
}

// Code gets the status code for the delete rds default response
func (o *DeleteRdsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRdsDefault) Error() string {
	return fmt.Sprintf("[DELETE /rdsconfigs/{id}][%d] deleteRds default ", o._statusCode)
}

func (o *DeleteRdsDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

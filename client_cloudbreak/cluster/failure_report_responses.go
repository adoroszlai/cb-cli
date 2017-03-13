package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// FailureReportReader is a Reader for the FailureReport structure.
type FailureReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *FailureReportReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewFailureReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewFailureReportDefault creates a FailureReportDefault with default headers values
func NewFailureReportDefault(code int) *FailureReportDefault {
	return &FailureReportDefault{
		_statusCode: code,
	}
}

/*FailureReportDefault handles this case with default header values.

successful operation
*/
type FailureReportDefault struct {
	_statusCode int
}

// Code gets the status code for the failure report default response
func (o *FailureReportDefault) Code() int {
	return o._statusCode
}

func (o *FailureReportDefault) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/cluster/failurereport][%d] failureReport default ", o._statusCode)
}

func (o *FailureReportDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
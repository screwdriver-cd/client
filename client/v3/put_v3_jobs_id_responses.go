package v3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// PutV3JobsIDReader is a Reader for the PutV3JobsID structure.
type PutV3JobsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PutV3JobsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewPutV3JobsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutV3JobsIDDefault creates a PutV3JobsIDDefault with default headers values
func NewPutV3JobsIDDefault(code int) *PutV3JobsIDDefault {
	return &PutV3JobsIDDefault{
		_statusCode: code,
	}
}

/*PutV3JobsIDDefault handles this case with default header values.

Successful
*/
type PutV3JobsIDDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the put v3 jobs id default response
func (o *PutV3JobsIDDefault) Code() int {
	return o._statusCode
}

func (o *PutV3JobsIDDefault) Error() string {
	return fmt.Sprintf("[PUT /v3/jobs/{id}][%d] putV3JobsId default  %+v", o._statusCode, o.Payload)
}

func (o *PutV3JobsIDDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
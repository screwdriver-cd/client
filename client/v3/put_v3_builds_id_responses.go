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

// PutV3BuildsIDReader is a Reader for the PutV3BuildsID structure.
type PutV3BuildsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PutV3BuildsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewPutV3BuildsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutV3BuildsIDDefault creates a PutV3BuildsIDDefault with default headers values
func NewPutV3BuildsIDDefault(code int) *PutV3BuildsIDDefault {
	return &PutV3BuildsIDDefault{
		_statusCode: code,
	}
}

/*PutV3BuildsIDDefault handles this case with default header values.

Successful
*/
type PutV3BuildsIDDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the put v3 builds id default response
func (o *PutV3BuildsIDDefault) Code() int {
	return o._statusCode
}

func (o *PutV3BuildsIDDefault) Error() string {
	return fmt.Sprintf("[PUT /v3/builds/{id}][%d] putV3BuildsId default  %+v", o._statusCode, o.Payload)
}

func (o *PutV3BuildsIDDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

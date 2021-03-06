// Code generated by go-swagger; DO NOT EDIT.

package relay_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RelayHEADReader is a Reader for the RelayHEAD structure.
type RelayHEADReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RelayHEADReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRelayHEADOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRelayHEADOK creates a RelayHEADOK with default headers values
func NewRelayHEADOK() *RelayHEADOK {
	return &RelayHEADOK{}
}

/* RelayHEADOK describes a response with status code 200, with default header values.

OK
*/
type RelayHEADOK struct {
}

func (o *RelayHEADOK) Error() string {
	return fmt.Sprintf("[HEAD /v1/relay/{key}/{path}][%d] relayHEADOK ", 200)
}

func (o *RelayHEADOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HandleSystemMetricsReader is a Reader for the HandleSystemMetrics structure.
type HandleSystemMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleSystemMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleSystemMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHandleSystemMetricsOK creates a HandleSystemMetricsOK with default headers values
func NewHandleSystemMetricsOK() *HandleSystemMetricsOK {
	return &HandleSystemMetricsOK{}
}

/* HandleSystemMetricsOK describes a response with status code 200, with default header values.

OK
*/
type HandleSystemMetricsOK struct {
}

func (o *HandleSystemMetricsOK) Error() string {
	return fmt.Sprintf("[GET /v1/systems/metrics][%d] handleSystemMetricsOK ", 200)
}

func (o *HandleSystemMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

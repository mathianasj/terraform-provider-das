// Code generated by go-swagger; DO NOT EDIT.

package timeseries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HandleTimeseriesMetricsReader is a Reader for the HandleTimeseriesMetrics structure.
type HandleTimeseriesMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleTimeseriesMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleTimeseriesMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHandleTimeseriesMetricsOK creates a HandleTimeseriesMetricsOK with default headers values
func NewHandleTimeseriesMetricsOK() *HandleTimeseriesMetricsOK {
	return &HandleTimeseriesMetricsOK{}
}

/* HandleTimeseriesMetricsOK describes a response with status code 200, with default header values.

OK
*/
type HandleTimeseriesMetricsOK struct {
}

func (o *HandleTimeseriesMetricsOK) Error() string {
	return fmt.Sprintf("[GET /v1/timeseries/metrics][%d] handleTimeseriesMetricsOK ", 200)
}

func (o *HandleTimeseriesMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

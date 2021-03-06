// Code generated by go-swagger; DO NOT EDIT.

package timeseries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// HandleViolationReader is a Reader for the HandleViolation structure.
type HandleViolationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleViolationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleViolationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHandleViolationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHandleViolationOK creates a HandleViolationOK with default headers values
func NewHandleViolationOK() *HandleViolationOK {
	return &HandleViolationOK{}
}

/* HandleViolationOK describes a response with status code 200, with default header values.

OK
*/
type HandleViolationOK struct {
	Payload *models.TimeseriesV1TimeSeriesPostResponse
}

func (o *HandleViolationOK) Error() string {
	return fmt.Sprintf("[POST /v1/timeseries/violation][%d] handleViolationOK  %+v", 200, o.Payload)
}
func (o *HandleViolationOK) GetPayload() *models.TimeseriesV1TimeSeriesPostResponse {
	return o.Payload
}

func (o *HandleViolationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeseriesV1TimeSeriesPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandleViolationBadRequest creates a HandleViolationBadRequest with default headers values
func NewHandleViolationBadRequest() *HandleViolationBadRequest {
	return &HandleViolationBadRequest{}
}

/* HandleViolationBadRequest describes a response with status code 400, with default header values.

Invalid Parameter
*/
type HandleViolationBadRequest struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *HandleViolationBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/timeseries/violation][%d] handleViolationBadRequest  %+v", 400, o.Payload)
}
func (o *HandleViolationBadRequest) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *HandleViolationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

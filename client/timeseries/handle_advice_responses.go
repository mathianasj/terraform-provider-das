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

// HandleAdviceReader is a Reader for the HandleAdvice structure.
type HandleAdviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleAdviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleAdviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHandleAdviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHandleAdviceOK creates a HandleAdviceOK with default headers values
func NewHandleAdviceOK() *HandleAdviceOK {
	return &HandleAdviceOK{}
}

/* HandleAdviceOK describes a response with status code 200, with default header values.

OK
*/
type HandleAdviceOK struct {
	Payload *models.TimeseriesV1TimeSeriesPostResponse
}

func (o *HandleAdviceOK) Error() string {
	return fmt.Sprintf("[POST /v1/timeseries/advice][%d] handleAdviceOK  %+v", 200, o.Payload)
}
func (o *HandleAdviceOK) GetPayload() *models.TimeseriesV1TimeSeriesPostResponse {
	return o.Payload
}

func (o *HandleAdviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeseriesV1TimeSeriesPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandleAdviceBadRequest creates a HandleAdviceBadRequest with default headers values
func NewHandleAdviceBadRequest() *HandleAdviceBadRequest {
	return &HandleAdviceBadRequest{}
}

/* HandleAdviceBadRequest describes a response with status code 400, with default header values.

Invalid Parameter
*/
type HandleAdviceBadRequest struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *HandleAdviceBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/timeseries/advice][%d] handleAdviceBadRequest  %+v", 400, o.Payload)
}
func (o *HandleAdviceBadRequest) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *HandleAdviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

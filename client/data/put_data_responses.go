// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// PutDataReader is a Reader for the PutData structure.
type PutDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPutDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDataOK creates a PutDataOK with default headers values
func NewPutDataOK() *PutDataOK {
	return &PutDataOK{}
}

/* PutDataOK describes a response with status code 200, with default header values.

OK
*/
type PutDataOK struct {
	Payload *models.DataV1DataPutResponse
}

func (o *PutDataOK) Error() string {
	return fmt.Sprintf("[PUT /v1/data/{name}][%d] putDataOK  %+v", 200, o.Payload)
}
func (o *PutDataOK) GetPayload() *models.DataV1DataPutResponse {
	return o.Payload
}

func (o *PutDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataV1DataPutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDataNotFound creates a PutDataNotFound with default headers values
func NewPutDataNotFound() *PutDataNotFound {
	return &PutDataNotFound{}
}

/* PutDataNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDataNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *PutDataNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/data/{name}][%d] putDataNotFound  %+v", 404, o.Payload)
}
func (o *PutDataNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *PutDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// ValidateSystemTestsReader is a Reader for the ValidateSystemTests structure.
type ValidateSystemTestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateSystemTestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateSystemTestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewValidateSystemTestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateSystemTestsOK creates a ValidateSystemTestsOK with default headers values
func NewValidateSystemTestsOK() *ValidateSystemTestsOK {
	return &ValidateSystemTestsOK{}
}

/* ValidateSystemTestsOK describes a response with status code 200, with default header values.

OK
*/
type ValidateSystemTestsOK struct {
	Payload *models.SystemsV1SystemsTestsResponse
}

func (o *ValidateSystemTestsOK) Error() string {
	return fmt.Sprintf("[POST /v1/systems/{system}/validate/tests][%d] validateSystemTestsOK  %+v", 200, o.Payload)
}
func (o *ValidateSystemTestsOK) GetPayload() *models.SystemsV1SystemsTestsResponse {
	return o.Payload
}

func (o *ValidateSystemTestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsV1SystemsTestsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateSystemTestsNotFound creates a ValidateSystemTestsNotFound with default headers values
func NewValidateSystemTestsNotFound() *ValidateSystemTestsNotFound {
	return &ValidateSystemTestsNotFound{}
}

/* ValidateSystemTestsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ValidateSystemTestsNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *ValidateSystemTestsNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/systems/{system}/validate/tests][%d] validateSystemTestsNotFound  %+v", 404, o.Payload)
}
func (o *ValidateSystemTestsNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *ValidateSystemTestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
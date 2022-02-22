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

// DeleteSystemReader is a Reader for the DeleteSystem structure.
type DeleteSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSystemOK creates a DeleteSystemOK with default headers values
func NewDeleteSystemOK() *DeleteSystemOK {
	return &DeleteSystemOK{}
}

/* DeleteSystemOK describes a response with status code 200, with default header values.

OK
*/
type DeleteSystemOK struct {
	Payload *models.SystemsV1SystemsDeleteResponse
}

func (o *DeleteSystemOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/systems/{system}][%d] deleteSystemOK  %+v", 200, o.Payload)
}
func (o *DeleteSystemOK) GetPayload() *models.SystemsV1SystemsDeleteResponse {
	return o.Payload
}

func (o *DeleteSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsV1SystemsDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSystemNotFound creates a DeleteSystemNotFound with default headers values
func NewDeleteSystemNotFound() *DeleteSystemNotFound {
	return &DeleteSystemNotFound{}
}

/* DeleteSystemNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteSystemNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *DeleteSystemNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/systems/{system}][%d] deleteSystemNotFound  %+v", 404, o.Payload)
}
func (o *DeleteSystemNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *DeleteSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

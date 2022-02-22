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

// UpdateSystemBundleCompileReader is a Reader for the UpdateSystemBundleCompile structure.
type UpdateSystemBundleCompileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSystemBundleCompileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSystemBundleCompileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateSystemBundleCompileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSystemBundleCompileOK creates a UpdateSystemBundleCompileOK with default headers values
func NewUpdateSystemBundleCompileOK() *UpdateSystemBundleCompileOK {
	return &UpdateSystemBundleCompileOK{}
}

/* UpdateSystemBundleCompileOK describes a response with status code 200, with default header values.

OK
*/
type UpdateSystemBundleCompileOK struct {
	Payload *models.SystemsV1SystemsPutBundleCompileResponse
}

func (o *UpdateSystemBundleCompileOK) Error() string {
	return fmt.Sprintf("[PUT /v1/systems/{system}/bundle-compile][%d] updateSystemBundleCompileOK  %+v", 200, o.Payload)
}
func (o *UpdateSystemBundleCompileOK) GetPayload() *models.SystemsV1SystemsPutBundleCompileResponse {
	return o.Payload
}

func (o *UpdateSystemBundleCompileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsV1SystemsPutBundleCompileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSystemBundleCompileNotFound creates a UpdateSystemBundleCompileNotFound with default headers values
func NewUpdateSystemBundleCompileNotFound() *UpdateSystemBundleCompileNotFound {
	return &UpdateSystemBundleCompileNotFound{}
}

/* UpdateSystemBundleCompileNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateSystemBundleCompileNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *UpdateSystemBundleCompileNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/systems/{system}/bundle-compile][%d] updateSystemBundleCompileNotFound  %+v", 404, o.Payload)
}
func (o *UpdateSystemBundleCompileNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *UpdateSystemBundleCompileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

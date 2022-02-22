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

// GetDefaultPolicyReader is a Reader for the GetDefaultPolicy structure.
type GetDefaultPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefaultPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDefaultPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDefaultPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDefaultPolicyOK creates a GetDefaultPolicyOK with default headers values
func NewGetDefaultPolicyOK() *GetDefaultPolicyOK {
	return &GetDefaultPolicyOK{}
}

/* GetDefaultPolicyOK describes a response with status code 200, with default header values.

OK
*/
type GetDefaultPolicyOK struct {
	Payload *models.SystemsV1SystemsGetDefaultPolicyResponse
}

func (o *GetDefaultPolicyOK) Error() string {
	return fmt.Sprintf("[GET /v1/systems/{system}/default-policies/{path}][%d] getDefaultPolicyOK  %+v", 200, o.Payload)
}
func (o *GetDefaultPolicyOK) GetPayload() *models.SystemsV1SystemsGetDefaultPolicyResponse {
	return o.Payload
}

func (o *GetDefaultPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsV1SystemsGetDefaultPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefaultPolicyNotFound creates a GetDefaultPolicyNotFound with default headers values
func NewGetDefaultPolicyNotFound() *GetDefaultPolicyNotFound {
	return &GetDefaultPolicyNotFound{}
}

/* GetDefaultPolicyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDefaultPolicyNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetDefaultPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/systems/{system}/default-policies/{path}][%d] getDefaultPolicyNotFound  %+v", 404, o.Payload)
}
func (o *GetDefaultPolicyNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetDefaultPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

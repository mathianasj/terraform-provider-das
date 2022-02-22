// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// CreatePoliciesAtPathReader is a Reader for the CreatePoliciesAtPath structure.
type CreatePoliciesAtPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePoliciesAtPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePoliciesAtPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePoliciesAtPathOK creates a CreatePoliciesAtPathOK with default headers values
func NewCreatePoliciesAtPathOK() *CreatePoliciesAtPathOK {
	return &CreatePoliciesAtPathOK{}
}

/* CreatePoliciesAtPathOK describes a response with status code 200, with default header values.

OK
*/
type CreatePoliciesAtPathOK struct {
	Payload *models.PoliciesV2PoliciesPostResponse
}

func (o *CreatePoliciesAtPathOK) Error() string {
	return fmt.Sprintf("[POST /v2/policies/{path}][%d] createPoliciesAtPathOK  %+v", 200, o.Payload)
}
func (o *CreatePoliciesAtPathOK) GetPayload() *models.PoliciesV2PoliciesPostResponse {
	return o.Payload
}

func (o *CreatePoliciesAtPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PoliciesV2PoliciesPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

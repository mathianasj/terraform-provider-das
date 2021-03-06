// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// CreateRoleBindingReader is a Reader for the CreateRoleBinding structure.
type CreateRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRoleBindingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRoleBindingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateRoleBindingConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRoleBindingOK creates a CreateRoleBindingOK with default headers values
func NewCreateRoleBindingOK() *CreateRoleBindingOK {
	return &CreateRoleBindingOK{}
}

/* CreateRoleBindingOK describes a response with status code 200, with default header values.

OK
*/
type CreateRoleBindingOK struct {
	Payload *models.AuthzV2RoleBindingsPostResponse
}

func (o *CreateRoleBindingOK) Error() string {
	return fmt.Sprintf("[POST /v2/authz/rolebindings][%d] createRoleBindingOK  %+v", 200, o.Payload)
}
func (o *CreateRoleBindingOK) GetPayload() *models.AuthzV2RoleBindingsPostResponse {
	return o.Payload
}

func (o *CreateRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzV2RoleBindingsPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleBindingBadRequest creates a CreateRoleBindingBadRequest with default headers values
func NewCreateRoleBindingBadRequest() *CreateRoleBindingBadRequest {
	return &CreateRoleBindingBadRequest{}
}

/* CreateRoleBindingBadRequest describes a response with status code 400, with default header values.

Invalid Parameter
*/
type CreateRoleBindingBadRequest struct {
	Payload *models.AuthzV2RoleBindingsPostResponse
}

func (o *CreateRoleBindingBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/authz/rolebindings][%d] createRoleBindingBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRoleBindingBadRequest) GetPayload() *models.AuthzV2RoleBindingsPostResponse {
	return o.Payload
}

func (o *CreateRoleBindingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzV2RoleBindingsPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleBindingNotFound creates a CreateRoleBindingNotFound with default headers values
func NewCreateRoleBindingNotFound() *CreateRoleBindingNotFound {
	return &CreateRoleBindingNotFound{}
}

/* CreateRoleBindingNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateRoleBindingNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *CreateRoleBindingNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/authz/rolebindings][%d] createRoleBindingNotFound  %+v", 404, o.Payload)
}
func (o *CreateRoleBindingNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *CreateRoleBindingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleBindingConflict creates a CreateRoleBindingConflict with default headers values
func NewCreateRoleBindingConflict() *CreateRoleBindingConflict {
	return &CreateRoleBindingConflict{}
}

/* CreateRoleBindingConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateRoleBindingConflict struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *CreateRoleBindingConflict) Error() string {
	return fmt.Sprintf("[POST /v2/authz/rolebindings][%d] createRoleBindingConflict  %+v", 409, o.Payload)
}
func (o *CreateRoleBindingConflict) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *CreateRoleBindingConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

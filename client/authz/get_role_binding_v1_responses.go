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

// GetRoleBindingV1Reader is a Reader for the GetRoleBindingV1 structure.
type GetRoleBindingV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleBindingV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleBindingV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRoleBindingV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoleBindingV1OK creates a GetRoleBindingV1OK with default headers values
func NewGetRoleBindingV1OK() *GetRoleBindingV1OK {
	return &GetRoleBindingV1OK{}
}

/* GetRoleBindingV1OK describes a response with status code 200, with default header values.

OK
*/
type GetRoleBindingV1OK struct {
	Payload *models.AuthzV1RoleBindingsGetResponse
}

func (o *GetRoleBindingV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/authz/rolebindings/{resourcetype}/{resource}/{rolebinding}][%d] getRoleBindingV1OK  %+v", 200, o.Payload)
}
func (o *GetRoleBindingV1OK) GetPayload() *models.AuthzV1RoleBindingsGetResponse {
	return o.Payload
}

func (o *GetRoleBindingV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzV1RoleBindingsGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleBindingV1NotFound creates a GetRoleBindingV1NotFound with default headers values
func NewGetRoleBindingV1NotFound() *GetRoleBindingV1NotFound {
	return &GetRoleBindingV1NotFound{}
}

/* GetRoleBindingV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRoleBindingV1NotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetRoleBindingV1NotFound) Error() string {
	return fmt.Sprintf("[GET /v1/authz/rolebindings/{resourcetype}/{resource}/{rolebinding}][%d] getRoleBindingV1NotFound  %+v", 404, o.Payload)
}
func (o *GetRoleBindingV1NotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetRoleBindingV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

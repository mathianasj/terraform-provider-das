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

// ListRoleBindingsV2Reader is a Reader for the ListRoleBindingsV2 structure.
type ListRoleBindingsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRoleBindingsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRoleBindingsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRoleBindingsV2OK creates a ListRoleBindingsV2OK with default headers values
func NewListRoleBindingsV2OK() *ListRoleBindingsV2OK {
	return &ListRoleBindingsV2OK{}
}

/* ListRoleBindingsV2OK describes a response with status code 200, with default header values.

OK
*/
type ListRoleBindingsV2OK struct {
	Payload *models.AuthzV2RoleBindingsListResponse
}

func (o *ListRoleBindingsV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/authz/rolebindings][%d] listRoleBindingsV2OK  %+v", 200, o.Payload)
}
func (o *ListRoleBindingsV2OK) GetPayload() *models.AuthzV2RoleBindingsListResponse {
	return o.Payload
}

func (o *ListRoleBindingsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzV2RoleBindingsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// MergeRoleBindingSubjectsReader is a Reader for the MergeRoleBindingSubjects structure.
type MergeRoleBindingSubjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MergeRoleBindingSubjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMergeRoleBindingSubjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMergeRoleBindingSubjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMergeRoleBindingSubjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMergeRoleBindingSubjectsOK creates a MergeRoleBindingSubjectsOK with default headers values
func NewMergeRoleBindingSubjectsOK() *MergeRoleBindingSubjectsOK {
	return &MergeRoleBindingSubjectsOK{}
}

/* MergeRoleBindingSubjectsOK describes a response with status code 200, with default header values.

OK
*/
type MergeRoleBindingSubjectsOK struct {
	Payload *models.AuthzV2RoleBindingsPutSubjectsResponse
}

func (o *MergeRoleBindingSubjectsOK) Error() string {
	return fmt.Sprintf("[PUT /v2/authz/rolebindings/{id}/subjects][%d] mergeRoleBindingSubjectsOK  %+v", 200, o.Payload)
}
func (o *MergeRoleBindingSubjectsOK) GetPayload() *models.AuthzV2RoleBindingsPutSubjectsResponse {
	return o.Payload
}

func (o *MergeRoleBindingSubjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzV2RoleBindingsPutSubjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMergeRoleBindingSubjectsBadRequest creates a MergeRoleBindingSubjectsBadRequest with default headers values
func NewMergeRoleBindingSubjectsBadRequest() *MergeRoleBindingSubjectsBadRequest {
	return &MergeRoleBindingSubjectsBadRequest{}
}

/* MergeRoleBindingSubjectsBadRequest describes a response with status code 400, with default header values.

Invalid Parameter
*/
type MergeRoleBindingSubjectsBadRequest struct {
	Payload *models.AuthzV2RoleBindingsPostResponse
}

func (o *MergeRoleBindingSubjectsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v2/authz/rolebindings/{id}/subjects][%d] mergeRoleBindingSubjectsBadRequest  %+v", 400, o.Payload)
}
func (o *MergeRoleBindingSubjectsBadRequest) GetPayload() *models.AuthzV2RoleBindingsPostResponse {
	return o.Payload
}

func (o *MergeRoleBindingSubjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzV2RoleBindingsPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMergeRoleBindingSubjectsNotFound creates a MergeRoleBindingSubjectsNotFound with default headers values
func NewMergeRoleBindingSubjectsNotFound() *MergeRoleBindingSubjectsNotFound {
	return &MergeRoleBindingSubjectsNotFound{}
}

/* MergeRoleBindingSubjectsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type MergeRoleBindingSubjectsNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *MergeRoleBindingSubjectsNotFound) Error() string {
	return fmt.Sprintf("[PUT /v2/authz/rolebindings/{id}/subjects][%d] mergeRoleBindingSubjectsNotFound  %+v", 404, o.Payload)
}
func (o *MergeRoleBindingSubjectsNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *MergeRoleBindingSubjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

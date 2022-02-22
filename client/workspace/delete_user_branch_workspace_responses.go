// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// DeleteUserBranchWorkspaceReader is a Reader for the DeleteUserBranchWorkspace structure.
type DeleteUserBranchWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserBranchWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserBranchWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteUserBranchWorkspaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserBranchWorkspaceOK creates a DeleteUserBranchWorkspaceOK with default headers values
func NewDeleteUserBranchWorkspaceOK() *DeleteUserBranchWorkspaceOK {
	return &DeleteUserBranchWorkspaceOK{}
}

/* DeleteUserBranchWorkspaceOK describes a response with status code 200, with default header values.

OK
*/
type DeleteUserBranchWorkspaceOK struct {
	Payload *models.GitV1DeleteBranchResponse
}

func (o *DeleteUserBranchWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/workspace/{id}/branch][%d] deleteUserBranchWorkspaceOK  %+v", 200, o.Payload)
}
func (o *DeleteUserBranchWorkspaceOK) GetPayload() *models.GitV1DeleteBranchResponse {
	return o.Payload
}

func (o *DeleteUserBranchWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitV1DeleteBranchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserBranchWorkspaceNotFound creates a DeleteUserBranchWorkspaceNotFound with default headers values
func NewDeleteUserBranchWorkspaceNotFound() *DeleteUserBranchWorkspaceNotFound {
	return &DeleteUserBranchWorkspaceNotFound{}
}

/* DeleteUserBranchWorkspaceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserBranchWorkspaceNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *DeleteUserBranchWorkspaceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/workspace/{id}/branch][%d] deleteUserBranchWorkspaceNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserBranchWorkspaceNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *DeleteUserBranchWorkspaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
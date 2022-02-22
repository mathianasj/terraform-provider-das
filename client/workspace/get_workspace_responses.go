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

// GetWorkspaceReader is a Reader for the GetWorkspace structure.
type GetWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetWorkspaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkspaceOK creates a GetWorkspaceOK with default headers values
func NewGetWorkspaceOK() *GetWorkspaceOK {
	return &GetWorkspaceOK{}
}

/* GetWorkspaceOK describes a response with status code 200, with default header values.

OK
*/
type GetWorkspaceOK struct {
	Payload *models.WorkspaceV1WorkspaceGetResponse
}

func (o *GetWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v1/workspace][%d] getWorkspaceOK  %+v", 200, o.Payload)
}
func (o *GetWorkspaceOK) GetPayload() *models.WorkspaceV1WorkspaceGetResponse {
	return o.Payload
}

func (o *GetWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspaceV1WorkspaceGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkspaceNotFound creates a GetWorkspaceNotFound with default headers values
func NewGetWorkspaceNotFound() *GetWorkspaceNotFound {
	return &GetWorkspaceNotFound{}
}

/* GetWorkspaceNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetWorkspaceNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetWorkspaceNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/workspace][%d] getWorkspaceNotFound  %+v", 404, o.Payload)
}
func (o *GetWorkspaceNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetWorkspaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// GetSourceControlFilesBranchStackReader is a Reader for the GetSourceControlFilesBranchStack structure.
type GetSourceControlFilesBranchStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceControlFilesBranchStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceControlFilesBranchStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSourceControlFilesBranchStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSourceControlFilesBranchStackOK creates a GetSourceControlFilesBranchStackOK with default headers values
func NewGetSourceControlFilesBranchStackOK() *GetSourceControlFilesBranchStackOK {
	return &GetSourceControlFilesBranchStackOK{}
}

/* GetSourceControlFilesBranchStackOK describes a response with status code 200, with default header values.

OK
*/
type GetSourceControlFilesBranchStackOK struct {
	Payload *models.GitV1GetFilesResponse
}

func (o *GetSourceControlFilesBranchStackOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/branch][%d] getSourceControlFilesBranchStackOK  %+v", 200, o.Payload)
}
func (o *GetSourceControlFilesBranchStackOK) GetPayload() *models.GitV1GetFilesResponse {
	return o.Payload
}

func (o *GetSourceControlFilesBranchStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitV1GetFilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceControlFilesBranchStackNotFound creates a GetSourceControlFilesBranchStackNotFound with default headers values
func NewGetSourceControlFilesBranchStackNotFound() *GetSourceControlFilesBranchStackNotFound {
	return &GetSourceControlFilesBranchStackNotFound{}
}

/* GetSourceControlFilesBranchStackNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSourceControlFilesBranchStackNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetSourceControlFilesBranchStackNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/branch][%d] getSourceControlFilesBranchStackNotFound  %+v", 404, o.Payload)
}
func (o *GetSourceControlFilesBranchStackNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetSourceControlFilesBranchStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

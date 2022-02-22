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

// GetSourceControlFilesMasterStackReader is a Reader for the GetSourceControlFilesMasterStack structure.
type GetSourceControlFilesMasterStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceControlFilesMasterStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceControlFilesMasterStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSourceControlFilesMasterStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSourceControlFilesMasterStackOK creates a GetSourceControlFilesMasterStackOK with default headers values
func NewGetSourceControlFilesMasterStackOK() *GetSourceControlFilesMasterStackOK {
	return &GetSourceControlFilesMasterStackOK{}
}

/* GetSourceControlFilesMasterStackOK describes a response with status code 200, with default header values.

OK
*/
type GetSourceControlFilesMasterStackOK struct {
	Payload *models.GitV1GetFilesResponse
}

func (o *GetSourceControlFilesMasterStackOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/master][%d] getSourceControlFilesMasterStackOK  %+v", 200, o.Payload)
}
func (o *GetSourceControlFilesMasterStackOK) GetPayload() *models.GitV1GetFilesResponse {
	return o.Payload
}

func (o *GetSourceControlFilesMasterStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitV1GetFilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceControlFilesMasterStackNotFound creates a GetSourceControlFilesMasterStackNotFound with default headers values
func NewGetSourceControlFilesMasterStackNotFound() *GetSourceControlFilesMasterStackNotFound {
	return &GetSourceControlFilesMasterStackNotFound{}
}

/* GetSourceControlFilesMasterStackNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSourceControlFilesMasterStackNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetSourceControlFilesMasterStackNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/master][%d] getSourceControlFilesMasterStackNotFound  %+v", 404, o.Payload)
}
func (o *GetSourceControlFilesMasterStackNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetSourceControlFilesMasterStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
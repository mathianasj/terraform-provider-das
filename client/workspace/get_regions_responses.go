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

// GetRegionsReader is a Reader for the GetRegions structure.
type GetRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRegionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegionsOK creates a GetRegionsOK with default headers values
func NewGetRegionsOK() *GetRegionsOK {
	return &GetRegionsOK{}
}

/* GetRegionsOK describes a response with status code 200, with default header values.

OK
*/
type GetRegionsOK struct {
	Payload *models.WorkspaceV1GetRegionResponse
}

func (o *GetRegionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/workspace/regions/{storagesvc}][%d] getRegionsOK  %+v", 200, o.Payload)
}
func (o *GetRegionsOK) GetPayload() *models.WorkspaceV1GetRegionResponse {
	return o.Payload
}

func (o *GetRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspaceV1GetRegionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegionsNotFound creates a GetRegionsNotFound with default headers values
func NewGetRegionsNotFound() *GetRegionsNotFound {
	return &GetRegionsNotFound{}
}

/* GetRegionsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetRegionsNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetRegionsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/workspace/regions/{storagesvc}][%d] getRegionsNotFound  %+v", 404, o.Payload)
}
func (o *GetRegionsNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetRegionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
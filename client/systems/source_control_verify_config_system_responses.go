// Code generated by go-swagger; DO NOT EDIT.

package systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// SourceControlVerifyConfigSystemReader is a Reader for the SourceControlVerifyConfigSystem structure.
type SourceControlVerifyConfigSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SourceControlVerifyConfigSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSourceControlVerifyConfigSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSourceControlVerifyConfigSystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSourceControlVerifyConfigSystemOK creates a SourceControlVerifyConfigSystemOK with default headers values
func NewSourceControlVerifyConfigSystemOK() *SourceControlVerifyConfigSystemOK {
	return &SourceControlVerifyConfigSystemOK{}
}

/* SourceControlVerifyConfigSystemOK describes a response with status code 200, with default header values.

OK
*/
type SourceControlVerifyConfigSystemOK struct {
	Payload *models.GitV1VerifyConfigResponse
}

func (o *SourceControlVerifyConfigSystemOK) Error() string {
	return fmt.Sprintf("[POST /v1/systems/source-control/verify-config][%d] sourceControlVerifyConfigSystemOK  %+v", 200, o.Payload)
}
func (o *SourceControlVerifyConfigSystemOK) GetPayload() *models.GitV1VerifyConfigResponse {
	return o.Payload
}

func (o *SourceControlVerifyConfigSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitV1VerifyConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSourceControlVerifyConfigSystemBadRequest creates a SourceControlVerifyConfigSystemBadRequest with default headers values
func NewSourceControlVerifyConfigSystemBadRequest() *SourceControlVerifyConfigSystemBadRequest {
	return &SourceControlVerifyConfigSystemBadRequest{}
}

/* SourceControlVerifyConfigSystemBadRequest describes a response with status code 400, with default header values.

Invalid Parameter
*/
type SourceControlVerifyConfigSystemBadRequest struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *SourceControlVerifyConfigSystemBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/systems/source-control/verify-config][%d] sourceControlVerifyConfigSystemBadRequest  %+v", 400, o.Payload)
}
func (o *SourceControlVerifyConfigSystemBadRequest) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *SourceControlVerifyConfigSystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package invitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// GetInvitationReader is a Reader for the GetInvitation structure.
type GetInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvitationOK creates a GetInvitationOK with default headers values
func NewGetInvitationOK() *GetInvitationOK {
	return &GetInvitationOK{}
}

/* GetInvitationOK describes a response with status code 200, with default header values.

OK
*/
type GetInvitationOK struct {
	Payload *models.InvitationsV1InvitationsGetResponse
}

func (o *GetInvitationOK) Error() string {
	return fmt.Sprintf("[GET /v1/invitations/{id}][%d] getInvitationOK  %+v", 200, o.Payload)
}
func (o *GetInvitationOK) GetPayload() *models.InvitationsV1InvitationsGetResponse {
	return o.Payload
}

func (o *GetInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvitationsV1InvitationsGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvitationNotFound creates a GetInvitationNotFound with default headers values
func NewGetInvitationNotFound() *GetInvitationNotFound {
	return &GetInvitationNotFound{}
}

/* GetInvitationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInvitationNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetInvitationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/invitations/{id}][%d] getInvitationNotFound  %+v", 404, o.Payload)
}
func (o *GetInvitationNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

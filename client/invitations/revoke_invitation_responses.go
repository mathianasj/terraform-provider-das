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

// RevokeInvitationReader is a Reader for the RevokeInvitation structure.
type RevokeInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevokeInvitationOK creates a RevokeInvitationOK with default headers values
func NewRevokeInvitationOK() *RevokeInvitationOK {
	return &RevokeInvitationOK{}
}

/* RevokeInvitationOK describes a response with status code 200, with default header values.

OK
*/
type RevokeInvitationOK struct {
	Payload *models.InvitationsV1InvitationsDeleteResponse
}

func (o *RevokeInvitationOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/invitations/{id}][%d] revokeInvitationOK  %+v", 200, o.Payload)
}
func (o *RevokeInvitationOK) GetPayload() *models.InvitationsV1InvitationsDeleteResponse {
	return o.Payload
}

func (o *RevokeInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvitationsV1InvitationsDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// RegisterNotificationToolReader is a Reader for the RegisterNotificationTool structure.
type RegisterNotificationToolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterNotificationToolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 307:
		result := NewRegisterNotificationToolTemporaryRedirect()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterNotificationToolTemporaryRedirect creates a RegisterNotificationToolTemporaryRedirect with default headers values
func NewRegisterNotificationToolTemporaryRedirect() *RegisterNotificationToolTemporaryRedirect {
	return &RegisterNotificationToolTemporaryRedirect{}
}

/* RegisterNotificationToolTemporaryRedirect describes a response with status code 307, with default header values.

OK
*/
type RegisterNotificationToolTemporaryRedirect struct {
	Payload *models.NotificationsV1NotificationIntegrationResponse
}

func (o *RegisterNotificationToolTemporaryRedirect) Error() string {
	return fmt.Sprintf("[GET /v1/notifications-install/callback/{type}][%d] registerNotificationToolTemporaryRedirect  %+v", 307, o.Payload)
}
func (o *RegisterNotificationToolTemporaryRedirect) GetPayload() *models.NotificationsV1NotificationIntegrationResponse {
	return o.Payload
}

func (o *RegisterNotificationToolTemporaryRedirect) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationsV1NotificationIntegrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
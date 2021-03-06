// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotificationsV1NotificationIntegrationResponse notifications v1 notification integration response
//
// swagger:model notifications.v1.NotificationIntegrationResponse
type NotificationsV1NotificationIntegrationResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`

	// response url
	ResponseURL string `json:"response_url,omitempty"`
}

// Validate validates this notifications v1 notification integration response
func (m *NotificationsV1NotificationIntegrationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notifications v1 notification integration response based on context it is used
func (m *NotificationsV1NotificationIntegrationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationsV1NotificationIntegrationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationsV1NotificationIntegrationResponse) UnmarshalBinary(b []byte) error {
	var res NotificationsV1NotificationIntegrationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

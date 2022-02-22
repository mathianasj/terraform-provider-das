// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// PostDecisionsReader is a Reader for the PostDecisions structure.
type PostDecisionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDecisionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDecisionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDecisionsOK creates a PostDecisionsOK with default headers values
func NewPostDecisionsOK() *PostDecisionsOK {
	return &PostDecisionsOK{}
}

/* PostDecisionsOK describes a response with status code 200, with default header values.

OK
*/
type PostDecisionsOK struct {
	Payload *models.LogsV1PostDecisionsResponse
}

func (o *PostDecisionsOK) Error() string {
	return fmt.Sprintf("[POST /v1/logs][%d] postDecisionsOK  %+v", 200, o.Payload)
}
func (o *PostDecisionsOK) GetPayload() *models.LogsV1PostDecisionsResponse {
	return o.Payload
}

func (o *PostDecisionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogsV1PostDecisionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// TranslateExternalIdsReader is a Reader for the TranslateExternalIds structure.
type TranslateExternalIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TranslateExternalIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTranslateExternalIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTranslateExternalIdsOK creates a TranslateExternalIdsOK with default headers values
func NewTranslateExternalIdsOK() *TranslateExternalIdsOK {
	return &TranslateExternalIdsOK{}
}

/* TranslateExternalIdsOK describes a response with status code 200, with default header values.

OK
*/
type TranslateExternalIdsOK struct {
	Payload *models.SystemsV1SystemsTranslateExternalIdsResponse
}

func (o *TranslateExternalIdsOK) Error() string {
	return fmt.Sprintf("[POST /v1/systems/external-ids][%d] translateExternalIdsOK  %+v", 200, o.Payload)
}
func (o *TranslateExternalIdsOK) GetPayload() *models.SystemsV1SystemsTranslateExternalIdsResponse {
	return o.Payload
}

func (o *TranslateExternalIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsV1SystemsTranslateExternalIdsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

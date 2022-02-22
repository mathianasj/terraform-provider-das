// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// ShowAllDataReader is a Reader for the ShowAllData structure.
type ShowAllDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowAllDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowAllDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewShowAllDataNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowAllDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowAllDataOK creates a ShowAllDataOK with default headers values
func NewShowAllDataOK() *ShowAllDataOK {
	return &ShowAllDataOK{}
}

/* ShowAllDataOK describes a response with status code 200, with default header values.

OK
*/
type ShowAllDataOK struct {

	/* The ETag (or entity tag) HTTP response header is an identifier for a specific version of a resource. See <https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag> documentation.
	 */
	Etag string

	Payload *models.DataV1DataResponse
}

func (o *ShowAllDataOK) Error() string {
	return fmt.Sprintf("[POST /v1/data][%d] showAllDataOK  %+v", 200, o.Payload)
}
func (o *ShowAllDataOK) GetPayload() *models.DataV1DataResponse {
	return o.Payload
}

func (o *ShowAllDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Etag
	hdrEtag := response.GetHeader("Etag")

	if hdrEtag != "" {
		o.Etag = hdrEtag
	}

	o.Payload = new(models.DataV1DataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowAllDataNotModified creates a ShowAllDataNotModified with default headers values
func NewShowAllDataNotModified() *ShowAllDataNotModified {
	return &ShowAllDataNotModified{}
}

/* ShowAllDataNotModified describes a response with status code 304, with default header values.

Not Modified
*/
type ShowAllDataNotModified struct {
}

func (o *ShowAllDataNotModified) Error() string {
	return fmt.Sprintf("[POST /v1/data][%d] showAllDataNotModified ", 304)
}

func (o *ShowAllDataNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowAllDataNotFound creates a ShowAllDataNotFound with default headers values
func NewShowAllDataNotFound() *ShowAllDataNotFound {
	return &ShowAllDataNotFound{}
}

/* ShowAllDataNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowAllDataNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *ShowAllDataNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/data][%d] showAllDataNotFound  %+v", 404, o.Payload)
}
func (o *ShowAllDataNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *ShowAllDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

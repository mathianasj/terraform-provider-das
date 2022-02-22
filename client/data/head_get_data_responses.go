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

// HeadGetDataReader is a Reader for the HeadGetData structure.
type HeadGetDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadGetDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadGetDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewHeadGetDataNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHeadGetDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHeadGetDataOK creates a HeadGetDataOK with default headers values
func NewHeadGetDataOK() *HeadGetDataOK {
	return &HeadGetDataOK{}
}

/* HeadGetDataOK describes a response with status code 200, with default header values.

OK
*/
type HeadGetDataOK struct {

	/* The ETag (or entity tag) HTTP response header is an identifier for a specific version of a resource. See <https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag> documentation.
	 */
	Etag string
}

func (o *HeadGetDataOK) Error() string {
	return fmt.Sprintf("[HEAD /v1/data/{name}][%d] headGetDataOK ", 200)
}

func (o *HeadGetDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Etag
	hdrEtag := response.GetHeader("Etag")

	if hdrEtag != "" {
		o.Etag = hdrEtag
	}

	return nil
}

// NewHeadGetDataNotModified creates a HeadGetDataNotModified with default headers values
func NewHeadGetDataNotModified() *HeadGetDataNotModified {
	return &HeadGetDataNotModified{}
}

/* HeadGetDataNotModified describes a response with status code 304, with default header values.

Not Modified
*/
type HeadGetDataNotModified struct {
}

func (o *HeadGetDataNotModified) Error() string {
	return fmt.Sprintf("[HEAD /v1/data/{name}][%d] headGetDataNotModified ", 304)
}

func (o *HeadGetDataNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeadGetDataNotFound creates a HeadGetDataNotFound with default headers values
func NewHeadGetDataNotFound() *HeadGetDataNotFound {
	return &HeadGetDataNotFound{}
}

/* HeadGetDataNotFound describes a response with status code 404, with default header values.

Not Found
*/
type HeadGetDataNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *HeadGetDataNotFound) Error() string {
	return fmt.Sprintf("[HEAD /v1/data/{name}][%d] headGetDataNotFound  %+v", 404, o.Payload)
}
func (o *HeadGetDataNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *HeadGetDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

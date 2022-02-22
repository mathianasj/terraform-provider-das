// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mathianasj/terraform-provider-das/models"
)

// GetDatasourceReader is a Reader for the GetDatasource structure.
type GetDatasourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatasourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatasourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDatasourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDatasourceOK creates a GetDatasourceOK with default headers values
func NewGetDatasourceOK() *GetDatasourceOK {
	return &GetDatasourceOK{}
}

/* GetDatasourceOK describes a response with status code 200, with default header values.

OK
*/
type GetDatasourceOK struct {
	Payload *models.DatasourcesV1DatasourcesGetResponse
}

func (o *GetDatasourceOK) Error() string {
	return fmt.Sprintf("[GET /v1/datasources/{datasource}][%d] getDatasourceOK  %+v", 200, o.Payload)
}
func (o *GetDatasourceOK) GetPayload() *models.DatasourcesV1DatasourcesGetResponse {
	return o.Payload
}

func (o *GetDatasourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatasourcesV1DatasourcesGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasourceNotFound creates a GetDatasourceNotFound with default headers values
func NewGetDatasourceNotFound() *GetDatasourceNotFound {
	return &GetDatasourceNotFound{}
}

/* GetDatasourceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDatasourceNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetDatasourceNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/datasources/{datasource}][%d] getDatasourceNotFound  %+v", 404, o.Payload)
}
func (o *GetDatasourceNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetDatasourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

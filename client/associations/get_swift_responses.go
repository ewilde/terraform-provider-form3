// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetSwiftReader is a Reader for the GetSwift structure.
type GetSwiftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSwiftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSwiftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSwiftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSwiftForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSwiftOK creates a GetSwiftOK with default headers values
func NewGetSwiftOK() *GetSwiftOK {
	return &GetSwiftOK{}
}

/*GetSwiftOK handles this case with default header values.

List of associations
*/
type GetSwiftOK struct {
	Payload *models.SwiftAssociationDetailsListResponse
}

func (o *GetSwiftOK) Error() string {
	return fmt.Sprintf("[GET /swift][%d] getSwiftOK  %+v", 200, o.Payload)
}

func (o *GetSwiftOK) GetPayload() *models.SwiftAssociationDetailsListResponse {
	return o.Payload
}

func (o *GetSwiftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwiftAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSwiftBadRequest creates a GetSwiftBadRequest with default headers values
func NewGetSwiftBadRequest() *GetSwiftBadRequest {
	return &GetSwiftBadRequest{}
}

/*GetSwiftBadRequest handles this case with default header values.

Bad Request
*/
type GetSwiftBadRequest struct {
	Payload *models.APIError
}

func (o *GetSwiftBadRequest) Error() string {
	return fmt.Sprintf("[GET /swift][%d] getSwiftBadRequest  %+v", 400, o.Payload)
}

func (o *GetSwiftBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSwiftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSwiftForbidden creates a GetSwiftForbidden with default headers values
func NewGetSwiftForbidden() *GetSwiftForbidden {
	return &GetSwiftForbidden{}
}

/*GetSwiftForbidden handles this case with default header values.

Forbidden
*/
type GetSwiftForbidden struct {
	Payload *models.APIError
}

func (o *GetSwiftForbidden) Error() string {
	return fmt.Sprintf("[GET /swift][%d] getSwiftForbidden  %+v", 403, o.Payload)
}

func (o *GetSwiftForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSwiftForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

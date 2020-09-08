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

// GetEburyReader is a Reader for the GetEbury structure.
type GetEburyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEburyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEburyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEburyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetEburyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEburyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEburyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEburyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEburyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetEburyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEburyOK creates a GetEburyOK with default headers values
func NewGetEburyOK() *GetEburyOK {
	return &GetEburyOK{}
}

/*GetEburyOK handles this case with default header values.

List of ebury associations
*/
type GetEburyOK struct {
	Payload *models.EburyAssociationListResponse
}

func (o *GetEburyOK) Error() string {
	return fmt.Sprintf("[GET /ebury][%d] getEburyOK  %+v", 200, o.Payload)
}

func (o *GetEburyOK) GetPayload() *models.EburyAssociationListResponse {
	return o.Payload
}

func (o *GetEburyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EburyAssociationListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEburyBadRequest creates a GetEburyBadRequest with default headers values
func NewGetEburyBadRequest() *GetEburyBadRequest {
	return &GetEburyBadRequest{}
}

/*GetEburyBadRequest handles this case with default header values.

Bad Request
*/
type GetEburyBadRequest struct {
	Payload *models.APIError
}

func (o *GetEburyBadRequest) Error() string {
	return fmt.Sprintf("[GET /ebury][%d] getEburyBadRequest  %+v", 400, o.Payload)
}

func (o *GetEburyBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetEburyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEburyUnauthorized creates a GetEburyUnauthorized with default headers values
func NewGetEburyUnauthorized() *GetEburyUnauthorized {
	return &GetEburyUnauthorized{}
}

/*GetEburyUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetEburyUnauthorized struct {
	Payload *models.APIError
}

func (o *GetEburyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ebury][%d] getEburyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetEburyUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetEburyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEburyForbidden creates a GetEburyForbidden with default headers values
func NewGetEburyForbidden() *GetEburyForbidden {
	return &GetEburyForbidden{}
}

/*GetEburyForbidden handles this case with default header values.

Forbidden
*/
type GetEburyForbidden struct {
	Payload *models.APIError
}

func (o *GetEburyForbidden) Error() string {
	return fmt.Sprintf("[GET /ebury][%d] getEburyForbidden  %+v", 403, o.Payload)
}

func (o *GetEburyForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetEburyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEburyNotFound creates a GetEburyNotFound with default headers values
func NewGetEburyNotFound() *GetEburyNotFound {
	return &GetEburyNotFound{}
}

/*GetEburyNotFound handles this case with default header values.

Record not found
*/
type GetEburyNotFound struct {
	Payload *models.APIError
}

func (o *GetEburyNotFound) Error() string {
	return fmt.Sprintf("[GET /ebury][%d] getEburyNotFound  %+v", 404, o.Payload)
}

func (o *GetEburyNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetEburyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEburyTooManyRequests creates a GetEburyTooManyRequests with default headers values
func NewGetEburyTooManyRequests() *GetEburyTooManyRequests {
	return &GetEburyTooManyRequests{}
}

/*GetEburyTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetEburyTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetEburyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ebury][%d] getEburyTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetEburyTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetEburyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEburyInternalServerError creates a GetEburyInternalServerError with default headers values
func NewGetEburyInternalServerError() *GetEburyInternalServerError {
	return &GetEburyInternalServerError{}
}

/*GetEburyInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetEburyInternalServerError struct {
	Payload *models.APIError
}

func (o *GetEburyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ebury][%d] getEburyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEburyInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetEburyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEburyServiceUnavailable creates a GetEburyServiceUnavailable with default headers values
func NewGetEburyServiceUnavailable() *GetEburyServiceUnavailable {
	return &GetEburyServiceUnavailable{}
}

/*GetEburyServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetEburyServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetEburyServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /ebury][%d] getEburyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetEburyServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetEburyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

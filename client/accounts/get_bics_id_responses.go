// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetBicsIDReader is a Reader for the GetBicsID structure.
type GetBicsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBicsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBicsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBicsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBicsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBicsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBicsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetBicsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBicsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBicsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetBicsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBicsIDOK creates a GetBicsIDOK with default headers values
func NewGetBicsIDOK() *GetBicsIDOK {
	return &GetBicsIDOK{}
}

/*GetBicsIDOK handles this case with default header values.

Bic details
*/
type GetBicsIDOK struct {
	Payload *models.BicDetailsResponse
}

func (o *GetBicsIDOK) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdOK  %+v", 200, o.Payload)
}

func (o *GetBicsIDOK) GetPayload() *models.BicDetailsResponse {
	return o.Payload
}

func (o *GetBicsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BicDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBicsIDBadRequest creates a GetBicsIDBadRequest with default headers values
func NewGetBicsIDBadRequest() *GetBicsIDBadRequest {
	return &GetBicsIDBadRequest{}
}

/*GetBicsIDBadRequest handles this case with default header values.

Bad Request
*/
type GetBicsIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetBicsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetBicsIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBicsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBicsIDUnauthorized creates a GetBicsIDUnauthorized with default headers values
func NewGetBicsIDUnauthorized() *GetBicsIDUnauthorized {
	return &GetBicsIDUnauthorized{}
}

/*GetBicsIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetBicsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetBicsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBicsIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBicsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBicsIDForbidden creates a GetBicsIDForbidden with default headers values
func NewGetBicsIDForbidden() *GetBicsIDForbidden {
	return &GetBicsIDForbidden{}
}

/*GetBicsIDForbidden handles this case with default header values.

Forbidden
*/
type GetBicsIDForbidden struct {
	Payload *models.APIError
}

func (o *GetBicsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdForbidden  %+v", 403, o.Payload)
}

func (o *GetBicsIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBicsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBicsIDNotFound creates a GetBicsIDNotFound with default headers values
func NewGetBicsIDNotFound() *GetBicsIDNotFound {
	return &GetBicsIDNotFound{}
}

/*GetBicsIDNotFound handles this case with default header values.

Record not found
*/
type GetBicsIDNotFound struct {
	Payload *models.APIError
}

func (o *GetBicsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetBicsIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBicsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBicsIDConflict creates a GetBicsIDConflict with default headers values
func NewGetBicsIDConflict() *GetBicsIDConflict {
	return &GetBicsIDConflict{}
}

/*GetBicsIDConflict handles this case with default header values.

Conflict
*/
type GetBicsIDConflict struct {
	Payload *models.APIError
}

func (o *GetBicsIDConflict) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdConflict  %+v", 409, o.Payload)
}

func (o *GetBicsIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBicsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBicsIDTooManyRequests creates a GetBicsIDTooManyRequests with default headers values
func NewGetBicsIDTooManyRequests() *GetBicsIDTooManyRequests {
	return &GetBicsIDTooManyRequests{}
}

/*GetBicsIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetBicsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetBicsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBicsIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBicsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBicsIDInternalServerError creates a GetBicsIDInternalServerError with default headers values
func NewGetBicsIDInternalServerError() *GetBicsIDInternalServerError {
	return &GetBicsIDInternalServerError{}
}

/*GetBicsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetBicsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetBicsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBicsIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBicsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBicsIDServiceUnavailable creates a GetBicsIDServiceUnavailable with default headers values
func NewGetBicsIDServiceUnavailable() *GetBicsIDServiceUnavailable {
	return &GetBicsIDServiceUnavailable{}
}

/*GetBicsIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetBicsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetBicsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /bics/{id}][%d] getBicsIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetBicsIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBicsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

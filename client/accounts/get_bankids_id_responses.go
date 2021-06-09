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

// GetBankidsIDReader is a Reader for the GetBankidsID structure.
type GetBankidsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBankidsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBankidsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBankidsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBankidsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBankidsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBankidsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetBankidsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBankidsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBankidsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetBankidsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBankidsIDOK creates a GetBankidsIDOK with default headers values
func NewGetBankidsIDOK() *GetBankidsIDOK {
	return &GetBankidsIDOK{}
}

/* GetBankidsIDOK describes a response with status code 200, with default header values.

BankId details
*/
type GetBankidsIDOK struct {
	Payload *models.BankIDDetailsResponse
}

func (o *GetBankidsIDOK) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdOK  %+v", 200, o.Payload)
}
func (o *GetBankidsIDOK) GetPayload() *models.BankIDDetailsResponse {
	return o.Payload
}

func (o *GetBankidsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BankIDDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsIDBadRequest creates a GetBankidsIDBadRequest with default headers values
func NewGetBankidsIDBadRequest() *GetBankidsIDBadRequest {
	return &GetBankidsIDBadRequest{}
}

/* GetBankidsIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetBankidsIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetBankidsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetBankidsIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsIDUnauthorized creates a GetBankidsIDUnauthorized with default headers values
func NewGetBankidsIDUnauthorized() *GetBankidsIDUnauthorized {
	return &GetBankidsIDUnauthorized{}
}

/* GetBankidsIDUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type GetBankidsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetBankidsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetBankidsIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsIDForbidden creates a GetBankidsIDForbidden with default headers values
func NewGetBankidsIDForbidden() *GetBankidsIDForbidden {
	return &GetBankidsIDForbidden{}
}

/* GetBankidsIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetBankidsIDForbidden struct {
	Payload *models.APIError
}

func (o *GetBankidsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdForbidden  %+v", 403, o.Payload)
}
func (o *GetBankidsIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsIDNotFound creates a GetBankidsIDNotFound with default headers values
func NewGetBankidsIDNotFound() *GetBankidsIDNotFound {
	return &GetBankidsIDNotFound{}
}

/* GetBankidsIDNotFound describes a response with status code 404, with default header values.

Record not found
*/
type GetBankidsIDNotFound struct {
	Payload *models.APIError
}

func (o *GetBankidsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdNotFound  %+v", 404, o.Payload)
}
func (o *GetBankidsIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsIDConflict creates a GetBankidsIDConflict with default headers values
func NewGetBankidsIDConflict() *GetBankidsIDConflict {
	return &GetBankidsIDConflict{}
}

/* GetBankidsIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetBankidsIDConflict struct {
	Payload *models.APIError
}

func (o *GetBankidsIDConflict) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdConflict  %+v", 409, o.Payload)
}
func (o *GetBankidsIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsIDTooManyRequests creates a GetBankidsIDTooManyRequests with default headers values
func NewGetBankidsIDTooManyRequests() *GetBankidsIDTooManyRequests {
	return &GetBankidsIDTooManyRequests{}
}

/* GetBankidsIDTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetBankidsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetBankidsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetBankidsIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsIDInternalServerError creates a GetBankidsIDInternalServerError with default headers values
func NewGetBankidsIDInternalServerError() *GetBankidsIDInternalServerError {
	return &GetBankidsIDInternalServerError{}
}

/* GetBankidsIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetBankidsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetBankidsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBankidsIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsIDServiceUnavailable creates a GetBankidsIDServiceUnavailable with default headers values
func NewGetBankidsIDServiceUnavailable() *GetBankidsIDServiceUnavailable {
	return &GetBankidsIDServiceUnavailable{}
}

/* GetBankidsIDServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetBankidsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetBankidsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /bankids/{id}][%d] getBankidsIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetBankidsIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

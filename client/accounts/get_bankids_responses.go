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

// GetBankidsReader is a Reader for the GetBankids structure.
type GetBankidsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBankidsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBankidsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBankidsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBankidsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBankidsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBankidsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetBankidsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBankidsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBankidsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetBankidsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBankidsOK creates a GetBankidsOK with default headers values
func NewGetBankidsOK() *GetBankidsOK {
	return &GetBankidsOK{}
}

/*GetBankidsOK handles this case with default header values.

List of bankId details
*/
type GetBankidsOK struct {
	Payload *models.BankIDDetailsListResponse
}

func (o *GetBankidsOK) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsOK  %+v", 200, o.Payload)
}

func (o *GetBankidsOK) GetPayload() *models.BankIDDetailsListResponse {
	return o.Payload
}

func (o *GetBankidsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BankIDDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsBadRequest creates a GetBankidsBadRequest with default headers values
func NewGetBankidsBadRequest() *GetBankidsBadRequest {
	return &GetBankidsBadRequest{}
}

/*GetBankidsBadRequest handles this case with default header values.

Bad Request
*/
type GetBankidsBadRequest struct {
	Payload *models.APIError
}

func (o *GetBankidsBadRequest) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsBadRequest  %+v", 400, o.Payload)
}

func (o *GetBankidsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsUnauthorized creates a GetBankidsUnauthorized with default headers values
func NewGetBankidsUnauthorized() *GetBankidsUnauthorized {
	return &GetBankidsUnauthorized{}
}

/*GetBankidsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetBankidsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetBankidsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBankidsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsForbidden creates a GetBankidsForbidden with default headers values
func NewGetBankidsForbidden() *GetBankidsForbidden {
	return &GetBankidsForbidden{}
}

/*GetBankidsForbidden handles this case with default header values.

Forbidden
*/
type GetBankidsForbidden struct {
	Payload *models.APIError
}

func (o *GetBankidsForbidden) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsForbidden  %+v", 403, o.Payload)
}

func (o *GetBankidsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsNotFound creates a GetBankidsNotFound with default headers values
func NewGetBankidsNotFound() *GetBankidsNotFound {
	return &GetBankidsNotFound{}
}

/*GetBankidsNotFound handles this case with default header values.

Record not found
*/
type GetBankidsNotFound struct {
	Payload *models.APIError
}

func (o *GetBankidsNotFound) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsNotFound  %+v", 404, o.Payload)
}

func (o *GetBankidsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsConflict creates a GetBankidsConflict with default headers values
func NewGetBankidsConflict() *GetBankidsConflict {
	return &GetBankidsConflict{}
}

/*GetBankidsConflict handles this case with default header values.

Conflict
*/
type GetBankidsConflict struct {
	Payload *models.APIError
}

func (o *GetBankidsConflict) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsConflict  %+v", 409, o.Payload)
}

func (o *GetBankidsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsTooManyRequests creates a GetBankidsTooManyRequests with default headers values
func NewGetBankidsTooManyRequests() *GetBankidsTooManyRequests {
	return &GetBankidsTooManyRequests{}
}

/*GetBankidsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetBankidsTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetBankidsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBankidsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsInternalServerError creates a GetBankidsInternalServerError with default headers values
func NewGetBankidsInternalServerError() *GetBankidsInternalServerError {
	return &GetBankidsInternalServerError{}
}

/*GetBankidsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetBankidsInternalServerError struct {
	Payload *models.APIError
}

func (o *GetBankidsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBankidsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBankidsServiceUnavailable creates a GetBankidsServiceUnavailable with default headers values
func NewGetBankidsServiceUnavailable() *GetBankidsServiceUnavailable {
	return &GetBankidsServiceUnavailable{}
}

/*GetBankidsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetBankidsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetBankidsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetBankidsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetBankidsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetPaymentsIDReversalsReversalIDReader is a Reader for the GetPaymentsIDReversalsReversalID structure.
type GetPaymentsIDReversalsReversalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDReversalsReversalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPaymentsIDReversalsReversalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPaymentsIDReversalsReversalIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPaymentsIDReversalsReversalIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPaymentsIDReversalsReversalIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPaymentsIDReversalsReversalIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPaymentsIDReversalsReversalIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPaymentsIDReversalsReversalIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPaymentsIDReversalsReversalIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPaymentsIDReversalsReversalIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPaymentsIDReversalsReversalIDOK creates a GetPaymentsIDReversalsReversalIDOK with default headers values
func NewGetPaymentsIDReversalsReversalIDOK() *GetPaymentsIDReversalsReversalIDOK {
	return &GetPaymentsIDReversalsReversalIDOK{}
}

/* GetPaymentsIDReversalsReversalIDOK describes a response with status code 200, with default header values.

Reversal details
*/
type GetPaymentsIDReversalsReversalIDOK struct {
	Payload *models.ReversalDetailsResponse
}

func (o *GetPaymentsIDReversalsReversalIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdOK  %+v", 200, o.Payload)
}
func (o *GetPaymentsIDReversalsReversalIDOK) GetPayload() *models.ReversalDetailsResponse {
	return o.Payload
}

func (o *GetPaymentsIDReversalsReversalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReversalDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDReversalsReversalIDBadRequest creates a GetPaymentsIDReversalsReversalIDBadRequest with default headers values
func NewGetPaymentsIDReversalsReversalIDBadRequest() *GetPaymentsIDReversalsReversalIDBadRequest {
	return &GetPaymentsIDReversalsReversalIDBadRequest{}
}

/* GetPaymentsIDReversalsReversalIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetPaymentsIDReversalsReversalIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDReversalsReversalIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetPaymentsIDReversalsReversalIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentsIDReversalsReversalIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDReversalsReversalIDUnauthorized creates a GetPaymentsIDReversalsReversalIDUnauthorized with default headers values
func NewGetPaymentsIDReversalsReversalIDUnauthorized() *GetPaymentsIDReversalsReversalIDUnauthorized {
	return &GetPaymentsIDReversalsReversalIDUnauthorized{}
}

/* GetPaymentsIDReversalsReversalIDUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type GetPaymentsIDReversalsReversalIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDReversalsReversalIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetPaymentsIDReversalsReversalIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentsIDReversalsReversalIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDReversalsReversalIDForbidden creates a GetPaymentsIDReversalsReversalIDForbidden with default headers values
func NewGetPaymentsIDReversalsReversalIDForbidden() *GetPaymentsIDReversalsReversalIDForbidden {
	return &GetPaymentsIDReversalsReversalIDForbidden{}
}

/* GetPaymentsIDReversalsReversalIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPaymentsIDReversalsReversalIDForbidden struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDReversalsReversalIDForbidden) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdForbidden  %+v", 403, o.Payload)
}
func (o *GetPaymentsIDReversalsReversalIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentsIDReversalsReversalIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDReversalsReversalIDNotFound creates a GetPaymentsIDReversalsReversalIDNotFound with default headers values
func NewGetPaymentsIDReversalsReversalIDNotFound() *GetPaymentsIDReversalsReversalIDNotFound {
	return &GetPaymentsIDReversalsReversalIDNotFound{}
}

/* GetPaymentsIDReversalsReversalIDNotFound describes a response with status code 404, with default header values.

Record not found
*/
type GetPaymentsIDReversalsReversalIDNotFound struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDReversalsReversalIDNotFound) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdNotFound  %+v", 404, o.Payload)
}
func (o *GetPaymentsIDReversalsReversalIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentsIDReversalsReversalIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDReversalsReversalIDConflict creates a GetPaymentsIDReversalsReversalIDConflict with default headers values
func NewGetPaymentsIDReversalsReversalIDConflict() *GetPaymentsIDReversalsReversalIDConflict {
	return &GetPaymentsIDReversalsReversalIDConflict{}
}

/* GetPaymentsIDReversalsReversalIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetPaymentsIDReversalsReversalIDConflict struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDReversalsReversalIDConflict) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdConflict  %+v", 409, o.Payload)
}
func (o *GetPaymentsIDReversalsReversalIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentsIDReversalsReversalIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDReversalsReversalIDTooManyRequests creates a GetPaymentsIDReversalsReversalIDTooManyRequests with default headers values
func NewGetPaymentsIDReversalsReversalIDTooManyRequests() *GetPaymentsIDReversalsReversalIDTooManyRequests {
	return &GetPaymentsIDReversalsReversalIDTooManyRequests{}
}

/* GetPaymentsIDReversalsReversalIDTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetPaymentsIDReversalsReversalIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDReversalsReversalIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetPaymentsIDReversalsReversalIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentsIDReversalsReversalIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDReversalsReversalIDInternalServerError creates a GetPaymentsIDReversalsReversalIDInternalServerError with default headers values
func NewGetPaymentsIDReversalsReversalIDInternalServerError() *GetPaymentsIDReversalsReversalIDInternalServerError {
	return &GetPaymentsIDReversalsReversalIDInternalServerError{}
}

/* GetPaymentsIDReversalsReversalIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetPaymentsIDReversalsReversalIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDReversalsReversalIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetPaymentsIDReversalsReversalIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentsIDReversalsReversalIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsIDReversalsReversalIDServiceUnavailable creates a GetPaymentsIDReversalsReversalIDServiceUnavailable with default headers values
func NewGetPaymentsIDReversalsReversalIDServiceUnavailable() *GetPaymentsIDReversalsReversalIDServiceUnavailable {
	return &GetPaymentsIDReversalsReversalIDServiceUnavailable{}
}

/* GetPaymentsIDReversalsReversalIDServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetPaymentsIDReversalsReversalIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetPaymentsIDReversalsReversalIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/reversals/{reversalId}][%d] getPaymentsIdReversalsReversalIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetPaymentsIDReversalsReversalIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetPaymentsIDReversalsReversalIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package payment_defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// DeletePaymentdefaultsIDReader is a Reader for the DeletePaymentdefaultsID structure.
type DeletePaymentdefaultsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePaymentdefaultsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePaymentdefaultsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePaymentdefaultsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePaymentdefaultsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePaymentdefaultsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePaymentdefaultsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeletePaymentdefaultsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletePaymentdefaultsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePaymentdefaultsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeletePaymentdefaultsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePaymentdefaultsIDNoContent creates a DeletePaymentdefaultsIDNoContent with default headers values
func NewDeletePaymentdefaultsIDNoContent() *DeletePaymentdefaultsIDNoContent {
	return &DeletePaymentdefaultsIDNoContent{}
}

/*DeletePaymentdefaultsIDNoContent handles this case with default header values.

Payment defaults deleted
*/
type DeletePaymentdefaultsIDNoContent struct {
}

func (o *DeletePaymentdefaultsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /paymentdefaults/{id}][%d] deletePaymentdefaultsIdNoContent ", 204)
}

func (o *DeletePaymentdefaultsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePaymentdefaultsIDBadRequest creates a DeletePaymentdefaultsIDBadRequest with default headers values
func NewDeletePaymentdefaultsIDBadRequest() *DeletePaymentdefaultsIDBadRequest {
	return &DeletePaymentdefaultsIDBadRequest{}
}

/*DeletePaymentdefaultsIDBadRequest handles this case with default header values.

Bad Request
*/
type DeletePaymentdefaultsIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeletePaymentdefaultsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /paymentdefaults/{id}][%d] deletePaymentdefaultsIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePaymentdefaultsIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePaymentdefaultsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePaymentdefaultsIDUnauthorized creates a DeletePaymentdefaultsIDUnauthorized with default headers values
func NewDeletePaymentdefaultsIDUnauthorized() *DeletePaymentdefaultsIDUnauthorized {
	return &DeletePaymentdefaultsIDUnauthorized{}
}

/*DeletePaymentdefaultsIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeletePaymentdefaultsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeletePaymentdefaultsIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /paymentdefaults/{id}][%d] deletePaymentdefaultsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePaymentdefaultsIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePaymentdefaultsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePaymentdefaultsIDForbidden creates a DeletePaymentdefaultsIDForbidden with default headers values
func NewDeletePaymentdefaultsIDForbidden() *DeletePaymentdefaultsIDForbidden {
	return &DeletePaymentdefaultsIDForbidden{}
}

/*DeletePaymentdefaultsIDForbidden handles this case with default header values.

Forbidden
*/
type DeletePaymentdefaultsIDForbidden struct {
	Payload *models.APIError
}

func (o *DeletePaymentdefaultsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /paymentdefaults/{id}][%d] deletePaymentdefaultsIdForbidden  %+v", 403, o.Payload)
}

func (o *DeletePaymentdefaultsIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePaymentdefaultsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePaymentdefaultsIDNotFound creates a DeletePaymentdefaultsIDNotFound with default headers values
func NewDeletePaymentdefaultsIDNotFound() *DeletePaymentdefaultsIDNotFound {
	return &DeletePaymentdefaultsIDNotFound{}
}

/*DeletePaymentdefaultsIDNotFound handles this case with default header values.

Record not found
*/
type DeletePaymentdefaultsIDNotFound struct {
	Payload *models.APIError
}

func (o *DeletePaymentdefaultsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /paymentdefaults/{id}][%d] deletePaymentdefaultsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeletePaymentdefaultsIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePaymentdefaultsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePaymentdefaultsIDConflict creates a DeletePaymentdefaultsIDConflict with default headers values
func NewDeletePaymentdefaultsIDConflict() *DeletePaymentdefaultsIDConflict {
	return &DeletePaymentdefaultsIDConflict{}
}

/*DeletePaymentdefaultsIDConflict handles this case with default header values.

Conflict
*/
type DeletePaymentdefaultsIDConflict struct {
	Payload *models.APIError
}

func (o *DeletePaymentdefaultsIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /paymentdefaults/{id}][%d] deletePaymentdefaultsIdConflict  %+v", 409, o.Payload)
}

func (o *DeletePaymentdefaultsIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePaymentdefaultsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePaymentdefaultsIDTooManyRequests creates a DeletePaymentdefaultsIDTooManyRequests with default headers values
func NewDeletePaymentdefaultsIDTooManyRequests() *DeletePaymentdefaultsIDTooManyRequests {
	return &DeletePaymentdefaultsIDTooManyRequests{}
}

/*DeletePaymentdefaultsIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeletePaymentdefaultsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeletePaymentdefaultsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /paymentdefaults/{id}][%d] deletePaymentdefaultsIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeletePaymentdefaultsIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePaymentdefaultsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePaymentdefaultsIDInternalServerError creates a DeletePaymentdefaultsIDInternalServerError with default headers values
func NewDeletePaymentdefaultsIDInternalServerError() *DeletePaymentdefaultsIDInternalServerError {
	return &DeletePaymentdefaultsIDInternalServerError{}
}

/*DeletePaymentdefaultsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeletePaymentdefaultsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeletePaymentdefaultsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /paymentdefaults/{id}][%d] deletePaymentdefaultsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePaymentdefaultsIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePaymentdefaultsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePaymentdefaultsIDServiceUnavailable creates a DeletePaymentdefaultsIDServiceUnavailable with default headers values
func NewDeletePaymentdefaultsIDServiceUnavailable() *DeletePaymentdefaultsIDServiceUnavailable {
	return &DeletePaymentdefaultsIDServiceUnavailable{}
}

/*DeletePaymentdefaultsIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeletePaymentdefaultsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeletePaymentdefaultsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /paymentdefaults/{id}][%d] deletePaymentdefaultsIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeletePaymentdefaultsIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeletePaymentdefaultsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

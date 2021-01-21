// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// DeleteKeysKeyIDReader is a Reader for the DeleteKeysKeyID structure.
type DeleteKeysKeyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKeysKeyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteKeysKeyIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteKeysKeyIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteKeysKeyIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteKeysKeyIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteKeysKeyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteKeysKeyIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewDeleteKeysKeyIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteKeysKeyIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteKeysKeyIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteKeysKeyIDNoContent creates a DeleteKeysKeyIDNoContent with default headers values
func NewDeleteKeysKeyIDNoContent() *DeleteKeysKeyIDNoContent {
	return &DeleteKeysKeyIDNoContent{}
}

/*DeleteKeysKeyIDNoContent handles this case with default header values.

Key deleted
*/
type DeleteKeysKeyIDNoContent struct {
}

func (o *DeleteKeysKeyIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}][%d] deleteKeysKeyIdNoContent ", 204)
}

func (o *DeleteKeysKeyIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKeysKeyIDBadRequest creates a DeleteKeysKeyIDBadRequest with default headers values
func NewDeleteKeysKeyIDBadRequest() *DeleteKeysKeyIDBadRequest {
	return &DeleteKeysKeyIDBadRequest{}
}

/*DeleteKeysKeyIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteKeysKeyIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}][%d] deleteKeysKeyIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKeysKeyIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDUnauthorized creates a DeleteKeysKeyIDUnauthorized with default headers values
func NewDeleteKeysKeyIDUnauthorized() *DeleteKeysKeyIDUnauthorized {
	return &DeleteKeysKeyIDUnauthorized{}
}

/*DeleteKeysKeyIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteKeysKeyIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}][%d] deleteKeysKeyIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKeysKeyIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDForbidden creates a DeleteKeysKeyIDForbidden with default headers values
func NewDeleteKeysKeyIDForbidden() *DeleteKeysKeyIDForbidden {
	return &DeleteKeysKeyIDForbidden{}
}

/*DeleteKeysKeyIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteKeysKeyIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}][%d] deleteKeysKeyIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKeysKeyIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDNotFound creates a DeleteKeysKeyIDNotFound with default headers values
func NewDeleteKeysKeyIDNotFound() *DeleteKeysKeyIDNotFound {
	return &DeleteKeysKeyIDNotFound{}
}

/*DeleteKeysKeyIDNotFound handles this case with default header values.

Record not found
*/
type DeleteKeysKeyIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}][%d] deleteKeysKeyIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKeysKeyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDConflict creates a DeleteKeysKeyIDConflict with default headers values
func NewDeleteKeysKeyIDConflict() *DeleteKeysKeyIDConflict {
	return &DeleteKeysKeyIDConflict{}
}

/*DeleteKeysKeyIDConflict handles this case with default header values.

Conflict
*/
type DeleteKeysKeyIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}][%d] deleteKeysKeyIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteKeysKeyIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDTooManyRequests creates a DeleteKeysKeyIDTooManyRequests with default headers values
func NewDeleteKeysKeyIDTooManyRequests() *DeleteKeysKeyIDTooManyRequests {
	return &DeleteKeysKeyIDTooManyRequests{}
}

/*DeleteKeysKeyIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteKeysKeyIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}][%d] deleteKeysKeyIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKeysKeyIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDInternalServerError creates a DeleteKeysKeyIDInternalServerError with default headers values
func NewDeleteKeysKeyIDInternalServerError() *DeleteKeysKeyIDInternalServerError {
	return &DeleteKeysKeyIDInternalServerError{}
}

/*DeleteKeysKeyIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteKeysKeyIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}][%d] deleteKeysKeyIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKeysKeyIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKeysKeyIDServiceUnavailable creates a DeleteKeysKeyIDServiceUnavailable with default headers values
func NewDeleteKeysKeyIDServiceUnavailable() *DeleteKeysKeyIDServiceUnavailable {
	return &DeleteKeysKeyIDServiceUnavailable{}
}

/*DeleteKeysKeyIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteKeysKeyIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteKeysKeyIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /keys/{key_id}][%d] deleteKeysKeyIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKeysKeyIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

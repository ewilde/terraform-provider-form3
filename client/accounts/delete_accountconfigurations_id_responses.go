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

// DeleteAccountconfigurationsIDReader is a Reader for the DeleteAccountconfigurationsID structure.
type DeleteAccountconfigurationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountconfigurationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAccountconfigurationsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAccountconfigurationsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAccountconfigurationsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAccountconfigurationsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAccountconfigurationsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAccountconfigurationsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAccountconfigurationsIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAccountconfigurationsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteAccountconfigurationsIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAccountconfigurationsIDNoContent creates a DeleteAccountconfigurationsIDNoContent with default headers values
func NewDeleteAccountconfigurationsIDNoContent() *DeleteAccountconfigurationsIDNoContent {
	return &DeleteAccountconfigurationsIDNoContent{}
}

/* DeleteAccountconfigurationsIDNoContent describes a response with status code 204, with default header values.

AccountConfiguration deleted
*/
type DeleteAccountconfigurationsIDNoContent struct {
}

func (o *DeleteAccountconfigurationsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdNoContent ", 204)
}

func (o *DeleteAccountconfigurationsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAccountconfigurationsIDBadRequest creates a DeleteAccountconfigurationsIDBadRequest with default headers values
func NewDeleteAccountconfigurationsIDBadRequest() *DeleteAccountconfigurationsIDBadRequest {
	return &DeleteAccountconfigurationsIDBadRequest{}
}

/* DeleteAccountconfigurationsIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAccountconfigurationsIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteAccountconfigurationsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAccountconfigurationsIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteAccountconfigurationsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountconfigurationsIDUnauthorized creates a DeleteAccountconfigurationsIDUnauthorized with default headers values
func NewDeleteAccountconfigurationsIDUnauthorized() *DeleteAccountconfigurationsIDUnauthorized {
	return &DeleteAccountconfigurationsIDUnauthorized{}
}

/* DeleteAccountconfigurationsIDUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteAccountconfigurationsIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteAccountconfigurationsIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteAccountconfigurationsIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteAccountconfigurationsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountconfigurationsIDForbidden creates a DeleteAccountconfigurationsIDForbidden with default headers values
func NewDeleteAccountconfigurationsIDForbidden() *DeleteAccountconfigurationsIDForbidden {
	return &DeleteAccountconfigurationsIDForbidden{}
}

/* DeleteAccountconfigurationsIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAccountconfigurationsIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteAccountconfigurationsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAccountconfigurationsIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteAccountconfigurationsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountconfigurationsIDNotFound creates a DeleteAccountconfigurationsIDNotFound with default headers values
func NewDeleteAccountconfigurationsIDNotFound() *DeleteAccountconfigurationsIDNotFound {
	return &DeleteAccountconfigurationsIDNotFound{}
}

/* DeleteAccountconfigurationsIDNotFound describes a response with status code 404, with default header values.

Record not found
*/
type DeleteAccountconfigurationsIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteAccountconfigurationsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAccountconfigurationsIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteAccountconfigurationsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountconfigurationsIDConflict creates a DeleteAccountconfigurationsIDConflict with default headers values
func NewDeleteAccountconfigurationsIDConflict() *DeleteAccountconfigurationsIDConflict {
	return &DeleteAccountconfigurationsIDConflict{}
}

/* DeleteAccountconfigurationsIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteAccountconfigurationsIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteAccountconfigurationsIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdConflict  %+v", 409, o.Payload)
}
func (o *DeleteAccountconfigurationsIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteAccountconfigurationsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountconfigurationsIDTooManyRequests creates a DeleteAccountconfigurationsIDTooManyRequests with default headers values
func NewDeleteAccountconfigurationsIDTooManyRequests() *DeleteAccountconfigurationsIDTooManyRequests {
	return &DeleteAccountconfigurationsIDTooManyRequests{}
}

/* DeleteAccountconfigurationsIDTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteAccountconfigurationsIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteAccountconfigurationsIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteAccountconfigurationsIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteAccountconfigurationsIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountconfigurationsIDInternalServerError creates a DeleteAccountconfigurationsIDInternalServerError with default headers values
func NewDeleteAccountconfigurationsIDInternalServerError() *DeleteAccountconfigurationsIDInternalServerError {
	return &DeleteAccountconfigurationsIDInternalServerError{}
}

/* DeleteAccountconfigurationsIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAccountconfigurationsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteAccountconfigurationsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAccountconfigurationsIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteAccountconfigurationsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountconfigurationsIDServiceUnavailable creates a DeleteAccountconfigurationsIDServiceUnavailable with default headers values
func NewDeleteAccountconfigurationsIDServiceUnavailable() *DeleteAccountconfigurationsIDServiceUnavailable {
	return &DeleteAccountconfigurationsIDServiceUnavailable{}
}

/* DeleteAccountconfigurationsIDServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteAccountconfigurationsIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteAccountconfigurationsIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /accountconfigurations/{id}][%d] deleteAccountconfigurationsIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *DeleteAccountconfigurationsIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteAccountconfigurationsIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

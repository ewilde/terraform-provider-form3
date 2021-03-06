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

// DeleteMandatemanagementIDReader is a Reader for the DeleteMandatemanagementID structure.
type DeleteMandatemanagementIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMandatemanagementIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteMandatemanagementIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteMandatemanagementIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteMandatemanagementIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMandatemanagementIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteMandatemanagementIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteMandatemanagementIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteMandatemanagementIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteMandatemanagementIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteMandatemanagementIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMandatemanagementIDNoContent creates a DeleteMandatemanagementIDNoContent with default headers values
func NewDeleteMandatemanagementIDNoContent() *DeleteMandatemanagementIDNoContent {
	return &DeleteMandatemanagementIDNoContent{}
}

/*DeleteMandatemanagementIDNoContent handles this case with default header values.

Mandate management deleted
*/
type DeleteMandatemanagementIDNoContent struct {
}

func (o *DeleteMandatemanagementIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /mandatemanagement/{id}][%d] deleteMandatemanagementIdNoContent ", 204)
}

func (o *DeleteMandatemanagementIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMandatemanagementIDBadRequest creates a DeleteMandatemanagementIDBadRequest with default headers values
func NewDeleteMandatemanagementIDBadRequest() *DeleteMandatemanagementIDBadRequest {
	return &DeleteMandatemanagementIDBadRequest{}
}

/*DeleteMandatemanagementIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteMandatemanagementIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteMandatemanagementIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mandatemanagement/{id}][%d] deleteMandatemanagementIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteMandatemanagementIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMandatemanagementIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMandatemanagementIDUnauthorized creates a DeleteMandatemanagementIDUnauthorized with default headers values
func NewDeleteMandatemanagementIDUnauthorized() *DeleteMandatemanagementIDUnauthorized {
	return &DeleteMandatemanagementIDUnauthorized{}
}

/*DeleteMandatemanagementIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteMandatemanagementIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteMandatemanagementIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /mandatemanagement/{id}][%d] deleteMandatemanagementIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteMandatemanagementIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMandatemanagementIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMandatemanagementIDForbidden creates a DeleteMandatemanagementIDForbidden with default headers values
func NewDeleteMandatemanagementIDForbidden() *DeleteMandatemanagementIDForbidden {
	return &DeleteMandatemanagementIDForbidden{}
}

/*DeleteMandatemanagementIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteMandatemanagementIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteMandatemanagementIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mandatemanagement/{id}][%d] deleteMandatemanagementIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteMandatemanagementIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMandatemanagementIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMandatemanagementIDNotFound creates a DeleteMandatemanagementIDNotFound with default headers values
func NewDeleteMandatemanagementIDNotFound() *DeleteMandatemanagementIDNotFound {
	return &DeleteMandatemanagementIDNotFound{}
}

/*DeleteMandatemanagementIDNotFound handles this case with default header values.

Record not found
*/
type DeleteMandatemanagementIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteMandatemanagementIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /mandatemanagement/{id}][%d] deleteMandatemanagementIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMandatemanagementIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMandatemanagementIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMandatemanagementIDConflict creates a DeleteMandatemanagementIDConflict with default headers values
func NewDeleteMandatemanagementIDConflict() *DeleteMandatemanagementIDConflict {
	return &DeleteMandatemanagementIDConflict{}
}

/*DeleteMandatemanagementIDConflict handles this case with default header values.

Conflict
*/
type DeleteMandatemanagementIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteMandatemanagementIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /mandatemanagement/{id}][%d] deleteMandatemanagementIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteMandatemanagementIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMandatemanagementIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMandatemanagementIDTooManyRequests creates a DeleteMandatemanagementIDTooManyRequests with default headers values
func NewDeleteMandatemanagementIDTooManyRequests() *DeleteMandatemanagementIDTooManyRequests {
	return &DeleteMandatemanagementIDTooManyRequests{}
}

/*DeleteMandatemanagementIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteMandatemanagementIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteMandatemanagementIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /mandatemanagement/{id}][%d] deleteMandatemanagementIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteMandatemanagementIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMandatemanagementIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMandatemanagementIDInternalServerError creates a DeleteMandatemanagementIDInternalServerError with default headers values
func NewDeleteMandatemanagementIDInternalServerError() *DeleteMandatemanagementIDInternalServerError {
	return &DeleteMandatemanagementIDInternalServerError{}
}

/*DeleteMandatemanagementIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteMandatemanagementIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteMandatemanagementIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /mandatemanagement/{id}][%d] deleteMandatemanagementIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteMandatemanagementIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMandatemanagementIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMandatemanagementIDServiceUnavailable creates a DeleteMandatemanagementIDServiceUnavailable with default headers values
func NewDeleteMandatemanagementIDServiceUnavailable() *DeleteMandatemanagementIDServiceUnavailable {
	return &DeleteMandatemanagementIDServiceUnavailable{}
}

/*DeleteMandatemanagementIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteMandatemanagementIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteMandatemanagementIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /mandatemanagement/{id}][%d] deleteMandatemanagementIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteMandatemanagementIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMandatemanagementIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

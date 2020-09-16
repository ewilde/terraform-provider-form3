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

// DeleteSepaddIDReader is a Reader for the DeleteSepaddID structure.
type DeleteSepaddIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSepaddIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSepaddIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSepaddIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSepaddIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSepaddIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSepaddIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteSepaddIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteSepaddIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSepaddIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteSepaddIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSepaddIDNoContent creates a DeleteSepaddIDNoContent with default headers values
func NewDeleteSepaddIDNoContent() *DeleteSepaddIDNoContent {
	return &DeleteSepaddIDNoContent{}
}

/*DeleteSepaddIDNoContent handles this case with default header values.

Association deleted
*/
type DeleteSepaddIDNoContent struct {
}

func (o *DeleteSepaddIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sepadd/{id}][%d] deleteSepaddIdNoContent ", 204)
}

func (o *DeleteSepaddIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSepaddIDBadRequest creates a DeleteSepaddIDBadRequest with default headers values
func NewDeleteSepaddIDBadRequest() *DeleteSepaddIDBadRequest {
	return &DeleteSepaddIDBadRequest{}
}

/*DeleteSepaddIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteSepaddIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteSepaddIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /sepadd/{id}][%d] deleteSepaddIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSepaddIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaddIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaddIDUnauthorized creates a DeleteSepaddIDUnauthorized with default headers values
func NewDeleteSepaddIDUnauthorized() *DeleteSepaddIDUnauthorized {
	return &DeleteSepaddIDUnauthorized{}
}

/*DeleteSepaddIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type DeleteSepaddIDUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteSepaddIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /sepadd/{id}][%d] deleteSepaddIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSepaddIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaddIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaddIDForbidden creates a DeleteSepaddIDForbidden with default headers values
func NewDeleteSepaddIDForbidden() *DeleteSepaddIDForbidden {
	return &DeleteSepaddIDForbidden{}
}

/*DeleteSepaddIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteSepaddIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteSepaddIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /sepadd/{id}][%d] deleteSepaddIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSepaddIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaddIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaddIDNotFound creates a DeleteSepaddIDNotFound with default headers values
func NewDeleteSepaddIDNotFound() *DeleteSepaddIDNotFound {
	return &DeleteSepaddIDNotFound{}
}

/*DeleteSepaddIDNotFound handles this case with default header values.

Record not found
*/
type DeleteSepaddIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteSepaddIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /sepadd/{id}][%d] deleteSepaddIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSepaddIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaddIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaddIDConflict creates a DeleteSepaddIDConflict with default headers values
func NewDeleteSepaddIDConflict() *DeleteSepaddIDConflict {
	return &DeleteSepaddIDConflict{}
}

/*DeleteSepaddIDConflict handles this case with default header values.

Conflict
*/
type DeleteSepaddIDConflict struct {
	Payload *models.APIError
}

func (o *DeleteSepaddIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /sepadd/{id}][%d] deleteSepaddIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteSepaddIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaddIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaddIDTooManyRequests creates a DeleteSepaddIDTooManyRequests with default headers values
func NewDeleteSepaddIDTooManyRequests() *DeleteSepaddIDTooManyRequests {
	return &DeleteSepaddIDTooManyRequests{}
}

/*DeleteSepaddIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type DeleteSepaddIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *DeleteSepaddIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /sepadd/{id}][%d] deleteSepaddIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteSepaddIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaddIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaddIDInternalServerError creates a DeleteSepaddIDInternalServerError with default headers values
func NewDeleteSepaddIDInternalServerError() *DeleteSepaddIDInternalServerError {
	return &DeleteSepaddIDInternalServerError{}
}

/*DeleteSepaddIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteSepaddIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteSepaddIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /sepadd/{id}][%d] deleteSepaddIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSepaddIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaddIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSepaddIDServiceUnavailable creates a DeleteSepaddIDServiceUnavailable with default headers values
func NewDeleteSepaddIDServiceUnavailable() *DeleteSepaddIDServiceUnavailable {
	return &DeleteSepaddIDServiceUnavailable{}
}

/*DeleteSepaddIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type DeleteSepaddIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteSepaddIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /sepadd/{id}][%d] deleteSepaddIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteSepaddIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteSepaddIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
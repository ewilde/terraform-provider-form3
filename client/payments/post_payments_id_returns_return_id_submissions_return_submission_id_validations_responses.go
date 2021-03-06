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

// PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsReader is a Reader for the PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidations structure.
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated handles this case with default header values.

Return submission validation creation response
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated struct {
	Payload *models.ReturnSubmissionValidationCreationResponse
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated) GetPayload() *models.ReturnSubmissionValidationCreationResponse {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnSubmissionValidationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest handles this case with default header values.

Bad Request
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden handles this case with default header values.

Forbidden
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsForbidden  %+v", 403, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound handles this case with default header values.

Record not found
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsNotFound  %+v", 404, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict handles this case with default header values.

Conflict
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsConflict  %+v", 409, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable creates a PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable with default headers values
func NewPostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable() *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable {
	return &PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable{}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/submissions/{returnSubmissionId}/validations][%d] postPaymentsIdReturnsReturnIdSubmissionsReturnSubmissionIdValidationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPaymentsIDReturnsReturnIDSubmissionsReturnSubmissionIDValidationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

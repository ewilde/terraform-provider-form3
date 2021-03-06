// Code generated by go-swagger; DO NOT EDIT.

package limits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostLimitsReader is a Reader for the PostLimits structure.
type PostLimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLimitsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLimitsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLimitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLimitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLimitsCreated creates a PostLimitsCreated with default headers values
func NewPostLimitsCreated() *PostLimitsCreated {
	return &PostLimitsCreated{}
}

/*PostLimitsCreated handles this case with default header values.

Limit creation response
*/
type PostLimitsCreated struct {
	Payload *models.LimitCreationResponse
}

func (o *PostLimitsCreated) Error() string {
	return fmt.Sprintf("[POST /limits][%d] postLimitsCreated  %+v", 201, o.Payload)
}

func (o *PostLimitsCreated) GetPayload() *models.LimitCreationResponse {
	return o.Payload
}

func (o *PostLimitsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLimitsBadRequest creates a PostLimitsBadRequest with default headers values
func NewPostLimitsBadRequest() *PostLimitsBadRequest {
	return &PostLimitsBadRequest{}
}

/*PostLimitsBadRequest handles this case with default header values.

Bad Request
*/
type PostLimitsBadRequest struct {
	Payload *models.APIError
}

func (o *PostLimitsBadRequest) Error() string {
	return fmt.Sprintf("[POST /limits][%d] postLimitsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLimitsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLimitsUnauthorized creates a PostLimitsUnauthorized with default headers values
func NewPostLimitsUnauthorized() *PostLimitsUnauthorized {
	return &PostLimitsUnauthorized{}
}

/*PostLimitsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostLimitsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostLimitsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /limits][%d] postLimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLimitsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLimitsForbidden creates a PostLimitsForbidden with default headers values
func NewPostLimitsForbidden() *PostLimitsForbidden {
	return &PostLimitsForbidden{}
}

/*PostLimitsForbidden handles this case with default header values.

Forbidden
*/
type PostLimitsForbidden struct {
	Payload *models.APIError
}

func (o *PostLimitsForbidden) Error() string {
	return fmt.Sprintf("[POST /limits][%d] postLimitsForbidden  %+v", 403, o.Payload)
}

func (o *PostLimitsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLimitsNotFound creates a PostLimitsNotFound with default headers values
func NewPostLimitsNotFound() *PostLimitsNotFound {
	return &PostLimitsNotFound{}
}

/*PostLimitsNotFound handles this case with default header values.

Record not found
*/
type PostLimitsNotFound struct {
	Payload *models.APIError
}

func (o *PostLimitsNotFound) Error() string {
	return fmt.Sprintf("[POST /limits][%d] postLimitsNotFound  %+v", 404, o.Payload)
}

func (o *PostLimitsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLimitsConflict creates a PostLimitsConflict with default headers values
func NewPostLimitsConflict() *PostLimitsConflict {
	return &PostLimitsConflict{}
}

/*PostLimitsConflict handles this case with default header values.

Conflict
*/
type PostLimitsConflict struct {
	Payload *models.APIError
}

func (o *PostLimitsConflict) Error() string {
	return fmt.Sprintf("[POST /limits][%d] postLimitsConflict  %+v", 409, o.Payload)
}

func (o *PostLimitsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLimitsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLimitsTooManyRequests creates a PostLimitsTooManyRequests with default headers values
func NewPostLimitsTooManyRequests() *PostLimitsTooManyRequests {
	return &PostLimitsTooManyRequests{}
}

/*PostLimitsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostLimitsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostLimitsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /limits][%d] postLimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLimitsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLimitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLimitsInternalServerError creates a PostLimitsInternalServerError with default headers values
func NewPostLimitsInternalServerError() *PostLimitsInternalServerError {
	return &PostLimitsInternalServerError{}
}

/*PostLimitsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostLimitsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostLimitsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /limits][%d] postLimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLimitsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLimitsServiceUnavailable creates a PostLimitsServiceUnavailable with default headers values
func NewPostLimitsServiceUnavailable() *PostLimitsServiceUnavailable {
	return &PostLimitsServiceUnavailable{}
}

/*PostLimitsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostLimitsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostLimitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /limits][%d] postLimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLimitsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLimitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

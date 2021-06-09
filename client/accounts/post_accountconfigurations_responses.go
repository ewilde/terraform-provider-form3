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

// PostAccountconfigurationsReader is a Reader for the PostAccountconfigurations structure.
type PostAccountconfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountconfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAccountconfigurationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAccountconfigurationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAccountconfigurationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAccountconfigurationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAccountconfigurationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAccountconfigurationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAccountconfigurationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAccountconfigurationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAccountconfigurationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAccountconfigurationsCreated creates a PostAccountconfigurationsCreated with default headers values
func NewPostAccountconfigurationsCreated() *PostAccountconfigurationsCreated {
	return &PostAccountconfigurationsCreated{}
}

/* PostAccountconfigurationsCreated describes a response with status code 201, with default header values.

AccountConfiguration creation response
*/
type PostAccountconfigurationsCreated struct {
	Payload *models.AccountConfigurationCreationResponse
}

func (o *PostAccountconfigurationsCreated) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsCreated  %+v", 201, o.Payload)
}
func (o *PostAccountconfigurationsCreated) GetPayload() *models.AccountConfigurationCreationResponse {
	return o.Payload
}

func (o *PostAccountconfigurationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountConfigurationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountconfigurationsBadRequest creates a PostAccountconfigurationsBadRequest with default headers values
func NewPostAccountconfigurationsBadRequest() *PostAccountconfigurationsBadRequest {
	return &PostAccountconfigurationsBadRequest{}
}

/* PostAccountconfigurationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAccountconfigurationsBadRequest struct {
	Payload *models.APIError
}

func (o *PostAccountconfigurationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsBadRequest  %+v", 400, o.Payload)
}
func (o *PostAccountconfigurationsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostAccountconfigurationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountconfigurationsUnauthorized creates a PostAccountconfigurationsUnauthorized with default headers values
func NewPostAccountconfigurationsUnauthorized() *PostAccountconfigurationsUnauthorized {
	return &PostAccountconfigurationsUnauthorized{}
}

/* PostAccountconfigurationsUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type PostAccountconfigurationsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostAccountconfigurationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsUnauthorized  %+v", 401, o.Payload)
}
func (o *PostAccountconfigurationsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostAccountconfigurationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountconfigurationsForbidden creates a PostAccountconfigurationsForbidden with default headers values
func NewPostAccountconfigurationsForbidden() *PostAccountconfigurationsForbidden {
	return &PostAccountconfigurationsForbidden{}
}

/* PostAccountconfigurationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostAccountconfigurationsForbidden struct {
	Payload *models.APIError
}

func (o *PostAccountconfigurationsForbidden) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsForbidden  %+v", 403, o.Payload)
}
func (o *PostAccountconfigurationsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostAccountconfigurationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountconfigurationsNotFound creates a PostAccountconfigurationsNotFound with default headers values
func NewPostAccountconfigurationsNotFound() *PostAccountconfigurationsNotFound {
	return &PostAccountconfigurationsNotFound{}
}

/* PostAccountconfigurationsNotFound describes a response with status code 404, with default header values.

Record not found
*/
type PostAccountconfigurationsNotFound struct {
	Payload *models.APIError
}

func (o *PostAccountconfigurationsNotFound) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsNotFound  %+v", 404, o.Payload)
}
func (o *PostAccountconfigurationsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostAccountconfigurationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountconfigurationsConflict creates a PostAccountconfigurationsConflict with default headers values
func NewPostAccountconfigurationsConflict() *PostAccountconfigurationsConflict {
	return &PostAccountconfigurationsConflict{}
}

/* PostAccountconfigurationsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostAccountconfigurationsConflict struct {
	Payload *models.APIError
}

func (o *PostAccountconfigurationsConflict) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsConflict  %+v", 409, o.Payload)
}
func (o *PostAccountconfigurationsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostAccountconfigurationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountconfigurationsTooManyRequests creates a PostAccountconfigurationsTooManyRequests with default headers values
func NewPostAccountconfigurationsTooManyRequests() *PostAccountconfigurationsTooManyRequests {
	return &PostAccountconfigurationsTooManyRequests{}
}

/* PostAccountconfigurationsTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostAccountconfigurationsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostAccountconfigurationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsTooManyRequests  %+v", 429, o.Payload)
}
func (o *PostAccountconfigurationsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostAccountconfigurationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountconfigurationsInternalServerError creates a PostAccountconfigurationsInternalServerError with default headers values
func NewPostAccountconfigurationsInternalServerError() *PostAccountconfigurationsInternalServerError {
	return &PostAccountconfigurationsInternalServerError{}
}

/* PostAccountconfigurationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAccountconfigurationsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostAccountconfigurationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsInternalServerError  %+v", 500, o.Payload)
}
func (o *PostAccountconfigurationsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostAccountconfigurationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountconfigurationsServiceUnavailable creates a PostAccountconfigurationsServiceUnavailable with default headers values
func NewPostAccountconfigurationsServiceUnavailable() *PostAccountconfigurationsServiceUnavailable {
	return &PostAccountconfigurationsServiceUnavailable{}
}

/* PostAccountconfigurationsServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostAccountconfigurationsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostAccountconfigurationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /accountconfigurations][%d] postAccountconfigurationsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostAccountconfigurationsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostAccountconfigurationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

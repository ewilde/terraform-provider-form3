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

// PostPayportReader is a Reader for the PostPayport structure.
type PostPayportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPayportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostPayportCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPayportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPayportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPayportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPayportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostPayportConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostPayportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPayportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPayportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPayportCreated creates a PostPayportCreated with default headers values
func NewPostPayportCreated() *PostPayportCreated {
	return &PostPayportCreated{}
}

/*PostPayportCreated handles this case with default header values.

creation response
*/
type PostPayportCreated struct {
	Payload *models.PayportAssociationCreationResponse
}

func (o *PostPayportCreated) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportCreated  %+v", 201, o.Payload)
}

func (o *PostPayportCreated) GetPayload() *models.PayportAssociationCreationResponse {
	return o.Payload
}

func (o *PostPayportCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PayportAssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPayportBadRequest creates a PostPayportBadRequest with default headers values
func NewPostPayportBadRequest() *PostPayportBadRequest {
	return &PostPayportBadRequest{}
}

/*PostPayportBadRequest handles this case with default header values.

Bad Request
*/
type PostPayportBadRequest struct {
	Payload *models.APIError
}

func (o *PostPayportBadRequest) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportBadRequest  %+v", 400, o.Payload)
}

func (o *PostPayportBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPayportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPayportUnauthorized creates a PostPayportUnauthorized with default headers values
func NewPostPayportUnauthorized() *PostPayportUnauthorized {
	return &PostPayportUnauthorized{}
}

/*PostPayportUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostPayportUnauthorized struct {
	Payload *models.APIError
}

func (o *PostPayportUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPayportUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPayportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPayportForbidden creates a PostPayportForbidden with default headers values
func NewPostPayportForbidden() *PostPayportForbidden {
	return &PostPayportForbidden{}
}

/*PostPayportForbidden handles this case with default header values.

Forbidden
*/
type PostPayportForbidden struct {
	Payload *models.APIError
}

func (o *PostPayportForbidden) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportForbidden  %+v", 403, o.Payload)
}

func (o *PostPayportForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPayportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPayportNotFound creates a PostPayportNotFound with default headers values
func NewPostPayportNotFound() *PostPayportNotFound {
	return &PostPayportNotFound{}
}

/*PostPayportNotFound handles this case with default header values.

Record not found
*/
type PostPayportNotFound struct {
	Payload *models.APIError
}

func (o *PostPayportNotFound) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportNotFound  %+v", 404, o.Payload)
}

func (o *PostPayportNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPayportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPayportConflict creates a PostPayportConflict with default headers values
func NewPostPayportConflict() *PostPayportConflict {
	return &PostPayportConflict{}
}

/*PostPayportConflict handles this case with default header values.

Conflict
*/
type PostPayportConflict struct {
	Payload *models.APIError
}

func (o *PostPayportConflict) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportConflict  %+v", 409, o.Payload)
}

func (o *PostPayportConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPayportConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPayportTooManyRequests creates a PostPayportTooManyRequests with default headers values
func NewPostPayportTooManyRequests() *PostPayportTooManyRequests {
	return &PostPayportTooManyRequests{}
}

/*PostPayportTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostPayportTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostPayportTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPayportTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPayportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPayportInternalServerError creates a PostPayportInternalServerError with default headers values
func NewPostPayportInternalServerError() *PostPayportInternalServerError {
	return &PostPayportInternalServerError{}
}

/*PostPayportInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostPayportInternalServerError struct {
	Payload *models.APIError
}

func (o *PostPayportInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPayportInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPayportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPayportServiceUnavailable creates a PostPayportServiceUnavailable with default headers values
func NewPostPayportServiceUnavailable() *PostPayportServiceUnavailable {
	return &PostPayportServiceUnavailable{}
}

/*PostPayportServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostPayportServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostPayportServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPayportServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostPayportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

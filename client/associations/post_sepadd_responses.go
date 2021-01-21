// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// PostSepaddReader is a Reader for the PostSepadd structure.
type PostSepaddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSepaddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSepaddCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostSepaddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostSepaddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostSepaddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostSepaddNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostSepaddConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostSepaddTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostSepaddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPostSepaddServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSepaddCreated creates a PostSepaddCreated with default headers values
func NewPostSepaddCreated() *PostSepaddCreated {
	return &PostSepaddCreated{}
}

/*PostSepaddCreated handles this case with default header values.

creation response
*/
type PostSepaddCreated struct {
	Payload *models.SepaDDAssociationCreationResponse
}

func (o *PostSepaddCreated) Error() string {
	return fmt.Sprintf("[POST /sepadd][%d] postSepaddCreated  %+v", 201, o.Payload)
}

func (o *PostSepaddCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaDDAssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSepaddBadRequest creates a PostSepaddBadRequest with default headers values
func NewPostSepaddBadRequest() *PostSepaddBadRequest {
	return &PostSepaddBadRequest{}
}

/*PostSepaddBadRequest handles this case with default header values.

Bad Request
*/
type PostSepaddBadRequest struct {
	Payload *models.APIError
}

func (o *PostSepaddBadRequest) Error() string {
	return fmt.Sprintf("[POST /sepadd][%d] postSepaddBadRequest  %+v", 400, o.Payload)
}

func (o *PostSepaddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSepaddUnauthorized creates a PostSepaddUnauthorized with default headers values
func NewPostSepaddUnauthorized() *PostSepaddUnauthorized {
	return &PostSepaddUnauthorized{}
}

/*PostSepaddUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostSepaddUnauthorized struct {
	Payload *models.APIError
}

func (o *PostSepaddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /sepadd][%d] postSepaddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSepaddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSepaddForbidden creates a PostSepaddForbidden with default headers values
func NewPostSepaddForbidden() *PostSepaddForbidden {
	return &PostSepaddForbidden{}
}

/*PostSepaddForbidden handles this case with default header values.

Forbidden
*/
type PostSepaddForbidden struct {
	Payload *models.APIError
}

func (o *PostSepaddForbidden) Error() string {
	return fmt.Sprintf("[POST /sepadd][%d] postSepaddForbidden  %+v", 403, o.Payload)
}

func (o *PostSepaddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSepaddNotFound creates a PostSepaddNotFound with default headers values
func NewPostSepaddNotFound() *PostSepaddNotFound {
	return &PostSepaddNotFound{}
}

/*PostSepaddNotFound handles this case with default header values.

Record not found
*/
type PostSepaddNotFound struct {
	Payload *models.APIError
}

func (o *PostSepaddNotFound) Error() string {
	return fmt.Sprintf("[POST /sepadd][%d] postSepaddNotFound  %+v", 404, o.Payload)
}

func (o *PostSepaddNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSepaddConflict creates a PostSepaddConflict with default headers values
func NewPostSepaddConflict() *PostSepaddConflict {
	return &PostSepaddConflict{}
}

/*PostSepaddConflict handles this case with default header values.

Conflict
*/
type PostSepaddConflict struct {
	Payload *models.APIError
}

func (o *PostSepaddConflict) Error() string {
	return fmt.Sprintf("[POST /sepadd][%d] postSepaddConflict  %+v", 409, o.Payload)
}

func (o *PostSepaddConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSepaddTooManyRequests creates a PostSepaddTooManyRequests with default headers values
func NewPostSepaddTooManyRequests() *PostSepaddTooManyRequests {
	return &PostSepaddTooManyRequests{}
}

/*PostSepaddTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostSepaddTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostSepaddTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /sepadd][%d] postSepaddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostSepaddTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSepaddInternalServerError creates a PostSepaddInternalServerError with default headers values
func NewPostSepaddInternalServerError() *PostSepaddInternalServerError {
	return &PostSepaddInternalServerError{}
}

/*PostSepaddInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostSepaddInternalServerError struct {
	Payload *models.APIError
}

func (o *PostSepaddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sepadd][%d] postSepaddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSepaddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSepaddServiceUnavailable creates a PostSepaddServiceUnavailable with default headers values
func NewPostSepaddServiceUnavailable() *PostSepaddServiceUnavailable {
	return &PostSepaddServiceUnavailable{}
}

/*PostSepaddServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostSepaddServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostSepaddServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /sepadd][%d] postSepaddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostSepaddServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// PatchGocardlessIDReader is a Reader for the PatchGocardlessID structure.
type PatchGocardlessIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGocardlessIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchGocardlessIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchGocardlessIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchGocardlessIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchGocardlessIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchGocardlessIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchGocardlessIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchGocardlessIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchGocardlessIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchGocardlessIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchGocardlessIDOK creates a PatchGocardlessIDOK with default headers values
func NewPatchGocardlessIDOK() *PatchGocardlessIDOK {
	return &PatchGocardlessIDOK{}
}

/*PatchGocardlessIDOK handles this case with default header values.

Association updated successfully
*/
type PatchGocardlessIDOK struct {
	Payload *models.GocardlessAssociationResponse
}

func (o *PatchGocardlessIDOK) Error() string {
	return fmt.Sprintf("[PATCH /gocardless/{id}][%d] patchGocardlessIdOK  %+v", 200, o.Payload)
}

func (o *PatchGocardlessIDOK) GetPayload() *models.GocardlessAssociationResponse {
	return o.Payload
}

func (o *PatchGocardlessIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GocardlessAssociationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGocardlessIDBadRequest creates a PatchGocardlessIDBadRequest with default headers values
func NewPatchGocardlessIDBadRequest() *PatchGocardlessIDBadRequest {
	return &PatchGocardlessIDBadRequest{}
}

/*PatchGocardlessIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchGocardlessIDBadRequest struct {
	Payload *models.APIError
}

func (o *PatchGocardlessIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /gocardless/{id}][%d] patchGocardlessIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchGocardlessIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchGocardlessIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGocardlessIDUnauthorized creates a PatchGocardlessIDUnauthorized with default headers values
func NewPatchGocardlessIDUnauthorized() *PatchGocardlessIDUnauthorized {
	return &PatchGocardlessIDUnauthorized{}
}

/*PatchGocardlessIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PatchGocardlessIDUnauthorized struct {
	Payload *models.APIError
}

func (o *PatchGocardlessIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /gocardless/{id}][%d] patchGocardlessIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchGocardlessIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchGocardlessIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGocardlessIDForbidden creates a PatchGocardlessIDForbidden with default headers values
func NewPatchGocardlessIDForbidden() *PatchGocardlessIDForbidden {
	return &PatchGocardlessIDForbidden{}
}

/*PatchGocardlessIDForbidden handles this case with default header values.

Forbidden
*/
type PatchGocardlessIDForbidden struct {
	Payload *models.APIError
}

func (o *PatchGocardlessIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /gocardless/{id}][%d] patchGocardlessIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchGocardlessIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchGocardlessIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGocardlessIDNotFound creates a PatchGocardlessIDNotFound with default headers values
func NewPatchGocardlessIDNotFound() *PatchGocardlessIDNotFound {
	return &PatchGocardlessIDNotFound{}
}

/*PatchGocardlessIDNotFound handles this case with default header values.

Record not found
*/
type PatchGocardlessIDNotFound struct {
	Payload *models.APIError
}

func (o *PatchGocardlessIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /gocardless/{id}][%d] patchGocardlessIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchGocardlessIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchGocardlessIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGocardlessIDConflict creates a PatchGocardlessIDConflict with default headers values
func NewPatchGocardlessIDConflict() *PatchGocardlessIDConflict {
	return &PatchGocardlessIDConflict{}
}

/*PatchGocardlessIDConflict handles this case with default header values.

Conflict
*/
type PatchGocardlessIDConflict struct {
	Payload *models.APIError
}

func (o *PatchGocardlessIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /gocardless/{id}][%d] patchGocardlessIdConflict  %+v", 409, o.Payload)
}

func (o *PatchGocardlessIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchGocardlessIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGocardlessIDTooManyRequests creates a PatchGocardlessIDTooManyRequests with default headers values
func NewPatchGocardlessIDTooManyRequests() *PatchGocardlessIDTooManyRequests {
	return &PatchGocardlessIDTooManyRequests{}
}

/*PatchGocardlessIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PatchGocardlessIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *PatchGocardlessIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /gocardless/{id}][%d] patchGocardlessIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchGocardlessIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchGocardlessIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGocardlessIDInternalServerError creates a PatchGocardlessIDInternalServerError with default headers values
func NewPatchGocardlessIDInternalServerError() *PatchGocardlessIDInternalServerError {
	return &PatchGocardlessIDInternalServerError{}
}

/*PatchGocardlessIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type PatchGocardlessIDInternalServerError struct {
	Payload *models.APIError
}

func (o *PatchGocardlessIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /gocardless/{id}][%d] patchGocardlessIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchGocardlessIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchGocardlessIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGocardlessIDServiceUnavailable creates a PatchGocardlessIDServiceUnavailable with default headers values
func NewPatchGocardlessIDServiceUnavailable() *PatchGocardlessIDServiceUnavailable {
	return &PatchGocardlessIDServiceUnavailable{}
}

/*PatchGocardlessIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PatchGocardlessIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PatchGocardlessIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /gocardless/{id}][%d] patchGocardlessIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchGocardlessIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PatchGocardlessIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

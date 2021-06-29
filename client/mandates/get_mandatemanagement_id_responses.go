// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetMandatemanagementIDReader is a Reader for the GetMandatemanagementID structure.
type GetMandatemanagementIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMandatemanagementIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMandatemanagementIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMandatemanagementIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMandatemanagementIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMandatemanagementIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMandatemanagementIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetMandatemanagementIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMandatemanagementIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMandatemanagementIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetMandatemanagementIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMandatemanagementIDOK creates a GetMandatemanagementIDOK with default headers values
func NewGetMandatemanagementIDOK() *GetMandatemanagementIDOK {
	return &GetMandatemanagementIDOK{}
}

/*GetMandatemanagementIDOK handles this case with default header values.

Mandate Management details
*/
type GetMandatemanagementIDOK struct {
	Payload *models.MandateManagementDetailsResponse
}

func (o *GetMandatemanagementIDOK) Error() string {
	return fmt.Sprintf("[GET /mandatemanagement/{id}][%d] getMandatemanagementIdOK  %+v", 200, o.Payload)
}

func (o *GetMandatemanagementIDOK) GetPayload() *models.MandateManagementDetailsResponse {
	return o.Payload
}

func (o *GetMandatemanagementIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MandateManagementDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandatemanagementIDBadRequest creates a GetMandatemanagementIDBadRequest with default headers values
func NewGetMandatemanagementIDBadRequest() *GetMandatemanagementIDBadRequest {
	return &GetMandatemanagementIDBadRequest{}
}

/*GetMandatemanagementIDBadRequest handles this case with default header values.

Bad Request
*/
type GetMandatemanagementIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetMandatemanagementIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /mandatemanagement/{id}][%d] getMandatemanagementIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetMandatemanagementIDBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetMandatemanagementIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandatemanagementIDUnauthorized creates a GetMandatemanagementIDUnauthorized with default headers values
func NewGetMandatemanagementIDUnauthorized() *GetMandatemanagementIDUnauthorized {
	return &GetMandatemanagementIDUnauthorized{}
}

/*GetMandatemanagementIDUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetMandatemanagementIDUnauthorized struct {
	Payload *models.APIError
}

func (o *GetMandatemanagementIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /mandatemanagement/{id}][%d] getMandatemanagementIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMandatemanagementIDUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetMandatemanagementIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandatemanagementIDForbidden creates a GetMandatemanagementIDForbidden with default headers values
func NewGetMandatemanagementIDForbidden() *GetMandatemanagementIDForbidden {
	return &GetMandatemanagementIDForbidden{}
}

/*GetMandatemanagementIDForbidden handles this case with default header values.

Forbidden
*/
type GetMandatemanagementIDForbidden struct {
	Payload *models.APIError
}

func (o *GetMandatemanagementIDForbidden) Error() string {
	return fmt.Sprintf("[GET /mandatemanagement/{id}][%d] getMandatemanagementIdForbidden  %+v", 403, o.Payload)
}

func (o *GetMandatemanagementIDForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetMandatemanagementIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandatemanagementIDNotFound creates a GetMandatemanagementIDNotFound with default headers values
func NewGetMandatemanagementIDNotFound() *GetMandatemanagementIDNotFound {
	return &GetMandatemanagementIDNotFound{}
}

/*GetMandatemanagementIDNotFound handles this case with default header values.

Record not found
*/
type GetMandatemanagementIDNotFound struct {
	Payload *models.APIError
}

func (o *GetMandatemanagementIDNotFound) Error() string {
	return fmt.Sprintf("[GET /mandatemanagement/{id}][%d] getMandatemanagementIdNotFound  %+v", 404, o.Payload)
}

func (o *GetMandatemanagementIDNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetMandatemanagementIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandatemanagementIDConflict creates a GetMandatemanagementIDConflict with default headers values
func NewGetMandatemanagementIDConflict() *GetMandatemanagementIDConflict {
	return &GetMandatemanagementIDConflict{}
}

/*GetMandatemanagementIDConflict handles this case with default header values.

Conflict
*/
type GetMandatemanagementIDConflict struct {
	Payload *models.APIError
}

func (o *GetMandatemanagementIDConflict) Error() string {
	return fmt.Sprintf("[GET /mandatemanagement/{id}][%d] getMandatemanagementIdConflict  %+v", 409, o.Payload)
}

func (o *GetMandatemanagementIDConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetMandatemanagementIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandatemanagementIDTooManyRequests creates a GetMandatemanagementIDTooManyRequests with default headers values
func NewGetMandatemanagementIDTooManyRequests() *GetMandatemanagementIDTooManyRequests {
	return &GetMandatemanagementIDTooManyRequests{}
}

/*GetMandatemanagementIDTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetMandatemanagementIDTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetMandatemanagementIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mandatemanagement/{id}][%d] getMandatemanagementIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMandatemanagementIDTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetMandatemanagementIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandatemanagementIDInternalServerError creates a GetMandatemanagementIDInternalServerError with default headers values
func NewGetMandatemanagementIDInternalServerError() *GetMandatemanagementIDInternalServerError {
	return &GetMandatemanagementIDInternalServerError{}
}

/*GetMandatemanagementIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetMandatemanagementIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetMandatemanagementIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /mandatemanagement/{id}][%d] getMandatemanagementIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMandatemanagementIDInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetMandatemanagementIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandatemanagementIDServiceUnavailable creates a GetMandatemanagementIDServiceUnavailable with default headers values
func NewGetMandatemanagementIDServiceUnavailable() *GetMandatemanagementIDServiceUnavailable {
	return &GetMandatemanagementIDServiceUnavailable{}
}

/*GetMandatemanagementIDServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetMandatemanagementIDServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetMandatemanagementIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /mandatemanagement/{id}][%d] getMandatemanagementIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetMandatemanagementIDServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetMandatemanagementIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

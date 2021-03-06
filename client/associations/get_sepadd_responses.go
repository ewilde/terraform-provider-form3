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

// GetSepaddReader is a Reader for the GetSepadd structure.
type GetSepaddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSepaddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSepaddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSepaddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSepaddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSepaddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSepaddNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetSepaddConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSepaddTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSepaddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSepaddServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSepaddOK creates a GetSepaddOK with default headers values
func NewGetSepaddOK() *GetSepaddOK {
	return &GetSepaddOK{}
}

/*GetSepaddOK handles this case with default header values.

List of associations
*/
type GetSepaddOK struct {
	Payload *models.SepaDDAssociationDetailsListResponse
}

func (o *GetSepaddOK) Error() string {
	return fmt.Sprintf("[GET /sepadd][%d] getSepaddOK  %+v", 200, o.Payload)
}

func (o *GetSepaddOK) GetPayload() *models.SepaDDAssociationDetailsListResponse {
	return o.Payload
}

func (o *GetSepaddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaDDAssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaddBadRequest creates a GetSepaddBadRequest with default headers values
func NewGetSepaddBadRequest() *GetSepaddBadRequest {
	return &GetSepaddBadRequest{}
}

/*GetSepaddBadRequest handles this case with default header values.

Bad Request
*/
type GetSepaddBadRequest struct {
	Payload *models.APIError
}

func (o *GetSepaddBadRequest) Error() string {
	return fmt.Sprintf("[GET /sepadd][%d] getSepaddBadRequest  %+v", 400, o.Payload)
}

func (o *GetSepaddBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaddUnauthorized creates a GetSepaddUnauthorized with default headers values
func NewGetSepaddUnauthorized() *GetSepaddUnauthorized {
	return &GetSepaddUnauthorized{}
}

/*GetSepaddUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetSepaddUnauthorized struct {
	Payload *models.APIError
}

func (o *GetSepaddUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sepadd][%d] getSepaddUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSepaddUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaddForbidden creates a GetSepaddForbidden with default headers values
func NewGetSepaddForbidden() *GetSepaddForbidden {
	return &GetSepaddForbidden{}
}

/*GetSepaddForbidden handles this case with default header values.

Forbidden
*/
type GetSepaddForbidden struct {
	Payload *models.APIError
}

func (o *GetSepaddForbidden) Error() string {
	return fmt.Sprintf("[GET /sepadd][%d] getSepaddForbidden  %+v", 403, o.Payload)
}

func (o *GetSepaddForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaddNotFound creates a GetSepaddNotFound with default headers values
func NewGetSepaddNotFound() *GetSepaddNotFound {
	return &GetSepaddNotFound{}
}

/*GetSepaddNotFound handles this case with default header values.

Record not found
*/
type GetSepaddNotFound struct {
	Payload *models.APIError
}

func (o *GetSepaddNotFound) Error() string {
	return fmt.Sprintf("[GET /sepadd][%d] getSepaddNotFound  %+v", 404, o.Payload)
}

func (o *GetSepaddNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaddNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaddConflict creates a GetSepaddConflict with default headers values
func NewGetSepaddConflict() *GetSepaddConflict {
	return &GetSepaddConflict{}
}

/*GetSepaddConflict handles this case with default header values.

Conflict
*/
type GetSepaddConflict struct {
	Payload *models.APIError
}

func (o *GetSepaddConflict) Error() string {
	return fmt.Sprintf("[GET /sepadd][%d] getSepaddConflict  %+v", 409, o.Payload)
}

func (o *GetSepaddConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaddConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaddTooManyRequests creates a GetSepaddTooManyRequests with default headers values
func NewGetSepaddTooManyRequests() *GetSepaddTooManyRequests {
	return &GetSepaddTooManyRequests{}
}

/*GetSepaddTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetSepaddTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetSepaddTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sepadd][%d] getSepaddTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSepaddTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaddTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaddInternalServerError creates a GetSepaddInternalServerError with default headers values
func NewGetSepaddInternalServerError() *GetSepaddInternalServerError {
	return &GetSepaddInternalServerError{}
}

/*GetSepaddInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSepaddInternalServerError struct {
	Payload *models.APIError
}

func (o *GetSepaddInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sepadd][%d] getSepaddInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSepaddInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSepaddServiceUnavailable creates a GetSepaddServiceUnavailable with default headers values
func NewGetSepaddServiceUnavailable() *GetSepaddServiceUnavailable {
	return &GetSepaddServiceUnavailable{}
}

/*GetSepaddServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetSepaddServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetSepaddServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /sepadd][%d] getSepaddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSepaddServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetSepaddServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

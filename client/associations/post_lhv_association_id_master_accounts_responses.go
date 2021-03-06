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

// PostLhvAssociationIDMasterAccountsReader is a Reader for the PostLhvAssociationIDMasterAccounts structure.
type PostLhvAssociationIDMasterAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLhvAssociationIDMasterAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLhvAssociationIDMasterAccountsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLhvAssociationIDMasterAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLhvAssociationIDMasterAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLhvAssociationIDMasterAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLhvAssociationIDMasterAccountsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLhvAssociationIDMasterAccountsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLhvAssociationIDMasterAccountsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLhvAssociationIDMasterAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLhvAssociationIDMasterAccountsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLhvAssociationIDMasterAccountsCreated creates a PostLhvAssociationIDMasterAccountsCreated with default headers values
func NewPostLhvAssociationIDMasterAccountsCreated() *PostLhvAssociationIDMasterAccountsCreated {
	return &PostLhvAssociationIDMasterAccountsCreated{}
}

/*PostLhvAssociationIDMasterAccountsCreated handles this case with default header values.

creation response
*/
type PostLhvAssociationIDMasterAccountsCreated struct {
	Payload *models.LhvMasterAccountResponse
}

func (o *PostLhvAssociationIDMasterAccountsCreated) Error() string {
	return fmt.Sprintf("[POST /lhv/{associationId}/master_accounts][%d] postLhvAssociationIdMasterAccountsCreated  %+v", 201, o.Payload)
}

func (o *PostLhvAssociationIDMasterAccountsCreated) GetPayload() *models.LhvMasterAccountResponse {
	return o.Payload
}

func (o *PostLhvAssociationIDMasterAccountsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LhvMasterAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLhvAssociationIDMasterAccountsBadRequest creates a PostLhvAssociationIDMasterAccountsBadRequest with default headers values
func NewPostLhvAssociationIDMasterAccountsBadRequest() *PostLhvAssociationIDMasterAccountsBadRequest {
	return &PostLhvAssociationIDMasterAccountsBadRequest{}
}

/*PostLhvAssociationIDMasterAccountsBadRequest handles this case with default header values.

Bad Request
*/
type PostLhvAssociationIDMasterAccountsBadRequest struct {
	Payload *models.APIError
}

func (o *PostLhvAssociationIDMasterAccountsBadRequest) Error() string {
	return fmt.Sprintf("[POST /lhv/{associationId}/master_accounts][%d] postLhvAssociationIdMasterAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLhvAssociationIDMasterAccountsBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLhvAssociationIDMasterAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLhvAssociationIDMasterAccountsUnauthorized creates a PostLhvAssociationIDMasterAccountsUnauthorized with default headers values
func NewPostLhvAssociationIDMasterAccountsUnauthorized() *PostLhvAssociationIDMasterAccountsUnauthorized {
	return &PostLhvAssociationIDMasterAccountsUnauthorized{}
}

/*PostLhvAssociationIDMasterAccountsUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type PostLhvAssociationIDMasterAccountsUnauthorized struct {
	Payload *models.APIError
}

func (o *PostLhvAssociationIDMasterAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lhv/{associationId}/master_accounts][%d] postLhvAssociationIdMasterAccountsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLhvAssociationIDMasterAccountsUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLhvAssociationIDMasterAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLhvAssociationIDMasterAccountsForbidden creates a PostLhvAssociationIDMasterAccountsForbidden with default headers values
func NewPostLhvAssociationIDMasterAccountsForbidden() *PostLhvAssociationIDMasterAccountsForbidden {
	return &PostLhvAssociationIDMasterAccountsForbidden{}
}

/*PostLhvAssociationIDMasterAccountsForbidden handles this case with default header values.

Forbidden
*/
type PostLhvAssociationIDMasterAccountsForbidden struct {
	Payload *models.APIError
}

func (o *PostLhvAssociationIDMasterAccountsForbidden) Error() string {
	return fmt.Sprintf("[POST /lhv/{associationId}/master_accounts][%d] postLhvAssociationIdMasterAccountsForbidden  %+v", 403, o.Payload)
}

func (o *PostLhvAssociationIDMasterAccountsForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLhvAssociationIDMasterAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLhvAssociationIDMasterAccountsNotFound creates a PostLhvAssociationIDMasterAccountsNotFound with default headers values
func NewPostLhvAssociationIDMasterAccountsNotFound() *PostLhvAssociationIDMasterAccountsNotFound {
	return &PostLhvAssociationIDMasterAccountsNotFound{}
}

/*PostLhvAssociationIDMasterAccountsNotFound handles this case with default header values.

Record not found
*/
type PostLhvAssociationIDMasterAccountsNotFound struct {
	Payload *models.APIError
}

func (o *PostLhvAssociationIDMasterAccountsNotFound) Error() string {
	return fmt.Sprintf("[POST /lhv/{associationId}/master_accounts][%d] postLhvAssociationIdMasterAccountsNotFound  %+v", 404, o.Payload)
}

func (o *PostLhvAssociationIDMasterAccountsNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLhvAssociationIDMasterAccountsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLhvAssociationIDMasterAccountsConflict creates a PostLhvAssociationIDMasterAccountsConflict with default headers values
func NewPostLhvAssociationIDMasterAccountsConflict() *PostLhvAssociationIDMasterAccountsConflict {
	return &PostLhvAssociationIDMasterAccountsConflict{}
}

/*PostLhvAssociationIDMasterAccountsConflict handles this case with default header values.

Conflict
*/
type PostLhvAssociationIDMasterAccountsConflict struct {
	Payload *models.APIError
}

func (o *PostLhvAssociationIDMasterAccountsConflict) Error() string {
	return fmt.Sprintf("[POST /lhv/{associationId}/master_accounts][%d] postLhvAssociationIdMasterAccountsConflict  %+v", 409, o.Payload)
}

func (o *PostLhvAssociationIDMasterAccountsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLhvAssociationIDMasterAccountsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLhvAssociationIDMasterAccountsTooManyRequests creates a PostLhvAssociationIDMasterAccountsTooManyRequests with default headers values
func NewPostLhvAssociationIDMasterAccountsTooManyRequests() *PostLhvAssociationIDMasterAccountsTooManyRequests {
	return &PostLhvAssociationIDMasterAccountsTooManyRequests{}
}

/*PostLhvAssociationIDMasterAccountsTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type PostLhvAssociationIDMasterAccountsTooManyRequests struct {
	Payload *models.APIError
}

func (o *PostLhvAssociationIDMasterAccountsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /lhv/{associationId}/master_accounts][%d] postLhvAssociationIdMasterAccountsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLhvAssociationIDMasterAccountsTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLhvAssociationIDMasterAccountsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLhvAssociationIDMasterAccountsInternalServerError creates a PostLhvAssociationIDMasterAccountsInternalServerError with default headers values
func NewPostLhvAssociationIDMasterAccountsInternalServerError() *PostLhvAssociationIDMasterAccountsInternalServerError {
	return &PostLhvAssociationIDMasterAccountsInternalServerError{}
}

/*PostLhvAssociationIDMasterAccountsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostLhvAssociationIDMasterAccountsInternalServerError struct {
	Payload *models.APIError
}

func (o *PostLhvAssociationIDMasterAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /lhv/{associationId}/master_accounts][%d] postLhvAssociationIdMasterAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLhvAssociationIDMasterAccountsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLhvAssociationIDMasterAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLhvAssociationIDMasterAccountsServiceUnavailable creates a PostLhvAssociationIDMasterAccountsServiceUnavailable with default headers values
func NewPostLhvAssociationIDMasterAccountsServiceUnavailable() *PostLhvAssociationIDMasterAccountsServiceUnavailable {
	return &PostLhvAssociationIDMasterAccountsServiceUnavailable{}
}

/*PostLhvAssociationIDMasterAccountsServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type PostLhvAssociationIDMasterAccountsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PostLhvAssociationIDMasterAccountsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /lhv/{associationId}/master_accounts][%d] postLhvAssociationIdMasterAccountsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLhvAssociationIDMasterAccountsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *PostLhvAssociationIDMasterAccountsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

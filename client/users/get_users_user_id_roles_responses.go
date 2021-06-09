// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetUsersUserIDRolesReader is a Reader for the GetUsersUserIDRoles structure.
type GetUsersUserIDRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersUserIDRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersUserIDRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsersUserIDRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersUserIDRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersUserIDRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetUsersUserIDRolesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUsersUserIDRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersUserIDRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUsersUserIDRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersUserIDRolesOK creates a GetUsersUserIDRolesOK with default headers values
func NewGetUsersUserIDRolesOK() *GetUsersUserIDRolesOK {
	return &GetUsersUserIDRolesOK{}
}

/* GetUsersUserIDRolesOK describes a response with status code 200, with default header values.

List of roles for user
*/
type GetUsersUserIDRolesOK struct {
	Payload *models.UserRoleListResponse
}

func (o *GetUsersUserIDRolesOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesOK  %+v", 200, o.Payload)
}
func (o *GetUsersUserIDRolesOK) GetPayload() *models.UserRoleListResponse {
	return o.Payload
}

func (o *GetUsersUserIDRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRoleListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDRolesBadRequest creates a GetUsersUserIDRolesBadRequest with default headers values
func NewGetUsersUserIDRolesBadRequest() *GetUsersUserIDRolesBadRequest {
	return &GetUsersUserIDRolesBadRequest{}
}

/* GetUsersUserIDRolesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetUsersUserIDRolesBadRequest struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesBadRequest  %+v", 400, o.Payload)
}
func (o *GetUsersUserIDRolesBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDRolesUnauthorized creates a GetUsersUserIDRolesUnauthorized with default headers values
func NewGetUsersUserIDRolesUnauthorized() *GetUsersUserIDRolesUnauthorized {
	return &GetUsersUserIDRolesUnauthorized{}
}

/* GetUsersUserIDRolesUnauthorized describes a response with status code 401, with default header values.

Authentication credentials were missing or incorrect
*/
type GetUsersUserIDRolesUnauthorized struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetUsersUserIDRolesUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDRolesForbidden creates a GetUsersUserIDRolesForbidden with default headers values
func NewGetUsersUserIDRolesForbidden() *GetUsersUserIDRolesForbidden {
	return &GetUsersUserIDRolesForbidden{}
}

/* GetUsersUserIDRolesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsersUserIDRolesForbidden struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesForbidden  %+v", 403, o.Payload)
}
func (o *GetUsersUserIDRolesForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDRolesNotFound creates a GetUsersUserIDRolesNotFound with default headers values
func NewGetUsersUserIDRolesNotFound() *GetUsersUserIDRolesNotFound {
	return &GetUsersUserIDRolesNotFound{}
}

/* GetUsersUserIDRolesNotFound describes a response with status code 404, with default header values.

Record not found
*/
type GetUsersUserIDRolesNotFound struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesNotFound  %+v", 404, o.Payload)
}
func (o *GetUsersUserIDRolesNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDRolesConflict creates a GetUsersUserIDRolesConflict with default headers values
func NewGetUsersUserIDRolesConflict() *GetUsersUserIDRolesConflict {
	return &GetUsersUserIDRolesConflict{}
}

/* GetUsersUserIDRolesConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetUsersUserIDRolesConflict struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDRolesConflict) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesConflict  %+v", 409, o.Payload)
}
func (o *GetUsersUserIDRolesConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDRolesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDRolesTooManyRequests creates a GetUsersUserIDRolesTooManyRequests with default headers values
func NewGetUsersUserIDRolesTooManyRequests() *GetUsersUserIDRolesTooManyRequests {
	return &GetUsersUserIDRolesTooManyRequests{}
}

/* GetUsersUserIDRolesTooManyRequests describes a response with status code 429, with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetUsersUserIDRolesTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetUsersUserIDRolesTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDRolesInternalServerError creates a GetUsersUserIDRolesInternalServerError with default headers values
func NewGetUsersUserIDRolesInternalServerError() *GetUsersUserIDRolesInternalServerError {
	return &GetUsersUserIDRolesInternalServerError{}
}

/* GetUsersUserIDRolesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetUsersUserIDRolesInternalServerError struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUsersUserIDRolesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDRolesServiceUnavailable creates a GetUsersUserIDRolesServiceUnavailable with default headers values
func NewGetUsersUserIDRolesServiceUnavailable() *GetUsersUserIDRolesServiceUnavailable {
	return &GetUsersUserIDRolesServiceUnavailable{}
}

/* GetUsersUserIDRolesServiceUnavailable describes a response with status code 503, with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetUsersUserIDRolesServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetUsersUserIDRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] getUsersUserIdRolesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetUsersUserIDRolesServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetUsersUserIDRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

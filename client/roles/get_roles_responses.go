// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// GetRolesReader is a Reader for the GetRoles structure.
type GetRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetRolesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRolesOK creates a GetRolesOK with default headers values
func NewGetRolesOK() *GetRolesOK {
	return &GetRolesOK{}
}

/*GetRolesOK handles this case with default header values.

List of role details
*/
type GetRolesOK struct {
	Payload *models.RoleDetailsListResponse
}

func (o *GetRolesOK) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesOK  %+v", 200, o.Payload)
}

func (o *GetRolesOK) GetPayload() *models.RoleDetailsListResponse {
	return o.Payload
}

func (o *GetRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesBadRequest creates a GetRolesBadRequest with default headers values
func NewGetRolesBadRequest() *GetRolesBadRequest {
	return &GetRolesBadRequest{}
}

/*GetRolesBadRequest handles this case with default header values.

Bad Request
*/
type GetRolesBadRequest struct {
	Payload *models.APIError
}

func (o *GetRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesUnauthorized creates a GetRolesUnauthorized with default headers values
func NewGetRolesUnauthorized() *GetRolesUnauthorized {
	return &GetRolesUnauthorized{}
}

/*GetRolesUnauthorized handles this case with default header values.

Authentication credentials were missing or incorrect
*/
type GetRolesUnauthorized struct {
	Payload *models.APIError
}

func (o *GetRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRolesUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForbidden creates a GetRolesForbidden with default headers values
func NewGetRolesForbidden() *GetRolesForbidden {
	return &GetRolesForbidden{}
}

/*GetRolesForbidden handles this case with default header values.

Forbidden
*/
type GetRolesForbidden struct {
	Payload *models.APIError
}

func (o *GetRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetRolesForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesNotFound creates a GetRolesNotFound with default headers values
func NewGetRolesNotFound() *GetRolesNotFound {
	return &GetRolesNotFound{}
}

/*GetRolesNotFound handles this case with default header values.

Record not found
*/
type GetRolesNotFound struct {
	Payload *models.APIError
}

func (o *GetRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetRolesNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesConflict creates a GetRolesConflict with default headers values
func NewGetRolesConflict() *GetRolesConflict {
	return &GetRolesConflict{}
}

/*GetRolesConflict handles this case with default header values.

Conflict
*/
type GetRolesConflict struct {
	Payload *models.APIError
}

func (o *GetRolesConflict) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesConflict  %+v", 409, o.Payload)
}

func (o *GetRolesConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRolesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesTooManyRequests creates a GetRolesTooManyRequests with default headers values
func NewGetRolesTooManyRequests() *GetRolesTooManyRequests {
	return &GetRolesTooManyRequests{}
}

/*GetRolesTooManyRequests handles this case with default header values.

The request cannot be served due to the application’s rate limit
*/
type GetRolesTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRolesTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesInternalServerError creates a GetRolesInternalServerError with default headers values
func NewGetRolesInternalServerError() *GetRolesInternalServerError {
	return &GetRolesInternalServerError{}
}

/*GetRolesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetRolesInternalServerError struct {
	Payload *models.APIError
}

func (o *GetRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesServiceUnavailable creates a GetRolesServiceUnavailable with default headers values
func NewGetRolesServiceUnavailable() *GetRolesServiceUnavailable {
	return &GetRolesServiceUnavailable{}
}

/*GetRolesServiceUnavailable handles this case with default header values.

The server is up, but overloaded with requests. Try again later.
*/
type GetRolesServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRolesServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

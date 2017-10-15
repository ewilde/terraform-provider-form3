// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ewilde/go-form3/models"
)

// PostUsersUserIDCredentialsReader is a Reader for the PostUsersUserIDCredentials structure.
type PostUsersUserIDCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersUserIDCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostUsersUserIDCredentialsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUsersUserIDCredentialsCreated creates a PostUsersUserIDCredentialsCreated with default headers values
func NewPostUsersUserIDCredentialsCreated() *PostUsersUserIDCredentialsCreated {
	return &PostUsersUserIDCredentialsCreated{}
}

/*PostUsersUserIDCredentialsCreated handles this case with default header values.

Credential creation response
*/
type PostUsersUserIDCredentialsCreated struct {
	Payload *models.CredentialCreationResponse
}

func (o *PostUsersUserIDCredentialsCreated) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials][%d] postUsersUserIdCredentialsCreated  %+v", 201, o.Payload)
}

func (o *PostUsersUserIDCredentialsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

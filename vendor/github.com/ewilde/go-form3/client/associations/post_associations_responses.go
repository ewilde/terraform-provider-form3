// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ewilde/go-form3/models"
)

// PostAssociationsReader is a Reader for the PostAssociations structure.
type PostAssociationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAssociationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAssociationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAssociationsCreated creates a PostAssociationsCreated with default headers values
func NewPostAssociationsCreated() *PostAssociationsCreated {
	return &PostAssociationsCreated{}
}

/*PostAssociationsCreated handles this case with default header values.

creation response
*/
type PostAssociationsCreated struct {
	Payload *models.AssociationCreationResponse
}

func (o *PostAssociationsCreated) Error() string {
	return fmt.Sprintf("[POST /associations][%d] postAssociationsCreated  %+v", 201, o.Payload)
}

func (o *PostAssociationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

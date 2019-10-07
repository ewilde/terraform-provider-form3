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

// GetGocardlessReader is a Reader for the GetGocardless structure.
type GetGocardlessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGocardlessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetGocardlessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGocardlessOK creates a GetGocardlessOK with default headers values
func NewGetGocardlessOK() *GetGocardlessOK {
	return &GetGocardlessOK{}
}

/*GetGocardlessOK handles this case with default header values.

List of gocardless associations
*/
type GetGocardlessOK struct {
	Payload *models.GocardlessAssociationListResponse
}

func (o *GetGocardlessOK) Error() string {
	return fmt.Sprintf("[GET /gocardless][%d] getGocardlessOK  %+v", 200, o.Payload)
}

func (o *GetGocardlessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GocardlessAssociationListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
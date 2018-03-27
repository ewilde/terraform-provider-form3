// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ewilde/go-form3/models"
)

// PatchSubscriptionsIDReader is a Reader for the PatchSubscriptionsID structure.
type PatchSubscriptionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSubscriptionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchSubscriptionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchSubscriptionsIDOK creates a PatchSubscriptionsIDOK with default headers values
func NewPatchSubscriptionsIDOK() *PatchSubscriptionsIDOK {
	return &PatchSubscriptionsIDOK{}
}

/*PatchSubscriptionsIDOK handles this case with default header values.

Subscription details
*/
type PatchSubscriptionsIDOK struct {
	Payload *models.SubscriptionDetailsResponse
}

func (o *PatchSubscriptionsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /subscriptions/{id}][%d] patchSubscriptionsIdOK  %+v", 200, o.Payload)
}

func (o *PatchSubscriptionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

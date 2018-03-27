// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ewilde/go-form3/models"
)

// PatchPaymentsIDSubmissionsSubmissionIDReader is a Reader for the PatchPaymentsIDSubmissionsSubmissionID structure.
type PatchPaymentsIDSubmissionsSubmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPaymentsIDSubmissionsSubmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchPaymentsIDSubmissionsSubmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchPaymentsIDSubmissionsSubmissionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchPaymentsIDSubmissionsSubmissionIDOK creates a PatchPaymentsIDSubmissionsSubmissionIDOK with default headers values
func NewPatchPaymentsIDSubmissionsSubmissionIDOK() *PatchPaymentsIDSubmissionsSubmissionIDOK {
	return &PatchPaymentsIDSubmissionsSubmissionIDOK{}
}

/*PatchPaymentsIDSubmissionsSubmissionIDOK handles this case with default header values.

Submission update response
*/
type PatchPaymentsIDSubmissionsSubmissionIDOK struct {
	Payload *models.PaymentSubmissionDetailsResponse
}

func (o *PatchPaymentsIDSubmissionsSubmissionIDOK) Error() string {
	return fmt.Sprintf("[PATCH /payments/{id}/submissions/{submissionId}][%d] patchPaymentsIdSubmissionsSubmissionIdOK  %+v", 200, o.Payload)
}

func (o *PatchPaymentsIDSubmissionsSubmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSubmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPaymentsIDSubmissionsSubmissionIDBadRequest creates a PatchPaymentsIDSubmissionsSubmissionIDBadRequest with default headers values
func NewPatchPaymentsIDSubmissionsSubmissionIDBadRequest() *PatchPaymentsIDSubmissionsSubmissionIDBadRequest {
	return &PatchPaymentsIDSubmissionsSubmissionIDBadRequest{}
}

/*PatchPaymentsIDSubmissionsSubmissionIDBadRequest handles this case with default header values.

Error
*/
type PatchPaymentsIDSubmissionsSubmissionIDBadRequest struct {
	Payload *models.APIError
}

func (o *PatchPaymentsIDSubmissionsSubmissionIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /payments/{id}/submissions/{submissionId}][%d] patchPaymentsIdSubmissionsSubmissionIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchPaymentsIDSubmissionsSubmissionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

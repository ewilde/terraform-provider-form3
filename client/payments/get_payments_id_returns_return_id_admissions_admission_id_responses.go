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

// GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDReader is a Reader for the GetPaymentsIDReturnsReturnIDAdmissionsAdmissionID structure.
type GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK creates a GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK with default headers values
func NewGetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK() *GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK {
	return &GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK{}
}

/*GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK handles this case with default header values.

Return admission details
*/
type GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK struct {
	Payload *models.ReturnAdmissionDetailsResponse
}

func (o *GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/returns/{returnId}/admissions/{admissionId}][%d] getPaymentsIdReturnsReturnIdAdmissionsAdmissionIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDReturnsReturnIDAdmissionsAdmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnAdmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/models"
)

// GetBankidsReader is a Reader for the GetBankids structure.
type GetBankidsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBankidsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBankidsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBankidsOK creates a GetBankidsOK with default headers values
func NewGetBankidsOK() *GetBankidsOK {
	return &GetBankidsOK{}
}

/*GetBankidsOK handles this case with default header values.

List of bankId details
*/
type GetBankidsOK struct {
	Payload *models.BankIDDetailsListResponse
}

func (o *GetBankidsOK) Error() string {
	return fmt.Sprintf("[GET /bankids][%d] getBankidsOK  %+v", 200, o.Payload)
}

func (o *GetBankidsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BankIDDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

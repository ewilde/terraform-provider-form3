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

// PostSepaLiquidityReader is a Reader for the PostSepaLiquidity structure.
type PostSepaLiquidityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSepaLiquidityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSepaLiquidityCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSepaLiquidityCreated creates a PostSepaLiquidityCreated with default headers values
func NewPostSepaLiquidityCreated() *PostSepaLiquidityCreated {
	return &PostSepaLiquidityCreated{}
}

/*PostSepaLiquidityCreated handles this case with default header values.

creation response
*/
type PostSepaLiquidityCreated struct {
	Payload *models.SepaLiquidityAssociationCreationResponse
}

func (o *PostSepaLiquidityCreated) Error() string {
	return fmt.Sprintf("[POST /sepa-liquidity][%d] postSepaLiquidityCreated  %+v", 201, o.Payload)
}

func (o *PostSepaLiquidityCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SepaLiquidityAssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
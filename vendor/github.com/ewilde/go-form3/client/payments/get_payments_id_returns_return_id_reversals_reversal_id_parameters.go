// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPaymentsIDReturnsReturnIDReversalsReversalIDParams creates a new GetPaymentsIDReturnsReturnIDReversalsReversalIDParams object
// with the default values initialized.
func NewGetPaymentsIDReturnsReturnIDReversalsReversalIDParams() *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	var ()
	return &GetPaymentsIDReturnsReturnIDReversalsReversalIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsIDReturnsReturnIDReversalsReversalIDParamsWithTimeout creates a new GetPaymentsIDReturnsReturnIDReversalsReversalIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsIDReturnsReturnIDReversalsReversalIDParamsWithTimeout(timeout time.Duration) *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	var ()
	return &GetPaymentsIDReturnsReturnIDReversalsReversalIDParams{

		timeout: timeout,
	}
}

// NewGetPaymentsIDReturnsReturnIDReversalsReversalIDParamsWithContext creates a new GetPaymentsIDReturnsReturnIDReversalsReversalIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentsIDReturnsReturnIDReversalsReversalIDParamsWithContext(ctx context.Context) *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	var ()
	return &GetPaymentsIDReturnsReturnIDReversalsReversalIDParams{

		Context: ctx,
	}
}

// NewGetPaymentsIDReturnsReturnIDReversalsReversalIDParamsWithHTTPClient creates a new GetPaymentsIDReturnsReturnIDReversalsReversalIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentsIDReturnsReturnIDReversalsReversalIDParamsWithHTTPClient(client *http.Client) *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	var ()
	return &GetPaymentsIDReturnsReturnIDReversalsReversalIDParams{
		HTTPClient: client,
	}
}

/*GetPaymentsIDReturnsReturnIDReversalsReversalIDParams contains all the parameters to send to the API endpoint
for the get payments ID returns return ID reversals reversal ID operation typically these are written to a http.Request
*/
type GetPaymentsIDReturnsReturnIDReversalsReversalIDParams struct {

	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*ReturnID
	  Return Id

	*/
	ReturnID strfmt.UUID
	/*ReversalID
	  Reversal Id

	*/
	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) WithTimeout(timeout time.Duration) *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) WithContext(ctx context.Context) *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) WithHTTPClient(client *http.Client) *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) WithID(id strfmt.UUID) *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithReturnID adds the returnID to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) WithReturnID(returnID strfmt.UUID) *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	o.SetReturnID(returnID)
	return o
}

// SetReturnID adds the returnId to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) SetReturnID(returnID strfmt.UUID) {
	o.ReturnID = returnID
}

// WithReversalID adds the reversalID to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) WithReversalID(reversalID strfmt.UUID) *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams {
	o.SetReversalID(reversalID)
	return o
}

// SetReversalID adds the reversalId to the get payments ID returns return ID reversals reversal ID params
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) SetReversalID(reversalID strfmt.UUID) {
	o.ReversalID = reversalID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param returnId
	if err := r.SetPathParam("returnId", o.ReturnID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
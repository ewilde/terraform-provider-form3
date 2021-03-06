// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetBankidsIDParams creates a new GetBankidsIDParams object
// with the default values initialized.
func NewGetBankidsIDParams() *GetBankidsIDParams {
	var ()
	return &GetBankidsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBankidsIDParamsWithTimeout creates a new GetBankidsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBankidsIDParamsWithTimeout(timeout time.Duration) *GetBankidsIDParams {
	var ()
	return &GetBankidsIDParams{

		timeout: timeout,
	}
}

// NewGetBankidsIDParamsWithContext creates a new GetBankidsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBankidsIDParamsWithContext(ctx context.Context) *GetBankidsIDParams {
	var ()
	return &GetBankidsIDParams{

		Context: ctx,
	}
}

// NewGetBankidsIDParamsWithHTTPClient creates a new GetBankidsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBankidsIDParamsWithHTTPClient(client *http.Client) *GetBankidsIDParams {
	var ()
	return &GetBankidsIDParams{
		HTTPClient: client,
	}
}

/*GetBankidsIDParams contains all the parameters to send to the API endpoint
for the get bankids ID operation typically these are written to a http.Request
*/
type GetBankidsIDParams struct {

	/*ID
	  Bank Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bankids ID params
func (o *GetBankidsIDParams) WithTimeout(timeout time.Duration) *GetBankidsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bankids ID params
func (o *GetBankidsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bankids ID params
func (o *GetBankidsIDParams) WithContext(ctx context.Context) *GetBankidsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bankids ID params
func (o *GetBankidsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bankids ID params
func (o *GetBankidsIDParams) WithHTTPClient(client *http.Client) *GetBankidsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bankids ID params
func (o *GetBankidsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get bankids ID params
func (o *GetBankidsIDParams) WithID(id strfmt.UUID) *GetBankidsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get bankids ID params
func (o *GetBankidsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBankidsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

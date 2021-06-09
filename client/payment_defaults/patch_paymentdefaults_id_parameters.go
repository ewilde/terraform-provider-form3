// Code generated by go-swagger; DO NOT EDIT.

package payment_defaults

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

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPatchPaymentdefaultsIDParams creates a new PatchPaymentdefaultsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchPaymentdefaultsIDParams() *PatchPaymentdefaultsIDParams {
	return &PatchPaymentdefaultsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPaymentdefaultsIDParamsWithTimeout creates a new PatchPaymentdefaultsIDParams object
// with the ability to set a timeout on a request.
func NewPatchPaymentdefaultsIDParamsWithTimeout(timeout time.Duration) *PatchPaymentdefaultsIDParams {
	return &PatchPaymentdefaultsIDParams{
		timeout: timeout,
	}
}

// NewPatchPaymentdefaultsIDParamsWithContext creates a new PatchPaymentdefaultsIDParams object
// with the ability to set a context for a request.
func NewPatchPaymentdefaultsIDParamsWithContext(ctx context.Context) *PatchPaymentdefaultsIDParams {
	return &PatchPaymentdefaultsIDParams{
		Context: ctx,
	}
}

// NewPatchPaymentdefaultsIDParamsWithHTTPClient creates a new PatchPaymentdefaultsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchPaymentdefaultsIDParamsWithHTTPClient(client *http.Client) *PatchPaymentdefaultsIDParams {
	return &PatchPaymentdefaultsIDParams{
		HTTPClient: client,
	}
}

/* PatchPaymentdefaultsIDParams contains all the parameters to send to the API endpoint
   for the patch paymentdefaults ID operation.

   Typically these are written to a http.Request.
*/
type PatchPaymentdefaultsIDParams struct {

	// DefaultsUpdateRequest.
	DefaultsUpdateRequest *models.PaymentDefaultsAmendment

	/* ID.

	   Payment defaults Id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch paymentdefaults ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPaymentdefaultsIDParams) WithDefaults() *PatchPaymentdefaultsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch paymentdefaults ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPaymentdefaultsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) WithTimeout(timeout time.Duration) *PatchPaymentdefaultsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) WithContext(ctx context.Context) *PatchPaymentdefaultsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) WithHTTPClient(client *http.Client) *PatchPaymentdefaultsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDefaultsUpdateRequest adds the defaultsUpdateRequest to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) WithDefaultsUpdateRequest(defaultsUpdateRequest *models.PaymentDefaultsAmendment) *PatchPaymentdefaultsIDParams {
	o.SetDefaultsUpdateRequest(defaultsUpdateRequest)
	return o
}

// SetDefaultsUpdateRequest adds the defaultsUpdateRequest to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) SetDefaultsUpdateRequest(defaultsUpdateRequest *models.PaymentDefaultsAmendment) {
	o.DefaultsUpdateRequest = defaultsUpdateRequest
}

// WithID adds the id to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) WithID(id strfmt.UUID) *PatchPaymentdefaultsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch paymentdefaults ID params
func (o *PatchPaymentdefaultsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPaymentdefaultsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.DefaultsUpdateRequest != nil {
		if err := r.SetBodyParam(o.DefaultsUpdateRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package associations

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

// NewGetEburyAssociationIDAccountsParams creates a new GetEburyAssociationIDAccountsParams object
// with the default values initialized.
func NewGetEburyAssociationIDAccountsParams() *GetEburyAssociationIDAccountsParams {
	var ()
	return &GetEburyAssociationIDAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEburyAssociationIDAccountsParamsWithTimeout creates a new GetEburyAssociationIDAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEburyAssociationIDAccountsParamsWithTimeout(timeout time.Duration) *GetEburyAssociationIDAccountsParams {
	var ()
	return &GetEburyAssociationIDAccountsParams{

		timeout: timeout,
	}
}

// NewGetEburyAssociationIDAccountsParamsWithContext creates a new GetEburyAssociationIDAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEburyAssociationIDAccountsParamsWithContext(ctx context.Context) *GetEburyAssociationIDAccountsParams {
	var ()
	return &GetEburyAssociationIDAccountsParams{

		Context: ctx,
	}
}

// NewGetEburyAssociationIDAccountsParamsWithHTTPClient creates a new GetEburyAssociationIDAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEburyAssociationIDAccountsParamsWithHTTPClient(client *http.Client) *GetEburyAssociationIDAccountsParams {
	var ()
	return &GetEburyAssociationIDAccountsParams{
		HTTPClient: client,
	}
}

/*GetEburyAssociationIDAccountsParams contains all the parameters to send to the API endpoint
for the get ebury association ID accounts operation typically these are written to a http.Request
*/
type GetEburyAssociationIDAccountsParams struct {

	/*AssociationID
	  Association Id

	*/
	AssociationID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ebury association ID accounts params
func (o *GetEburyAssociationIDAccountsParams) WithTimeout(timeout time.Duration) *GetEburyAssociationIDAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ebury association ID accounts params
func (o *GetEburyAssociationIDAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ebury association ID accounts params
func (o *GetEburyAssociationIDAccountsParams) WithContext(ctx context.Context) *GetEburyAssociationIDAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ebury association ID accounts params
func (o *GetEburyAssociationIDAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ebury association ID accounts params
func (o *GetEburyAssociationIDAccountsParams) WithHTTPClient(client *http.Client) *GetEburyAssociationIDAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ebury association ID accounts params
func (o *GetEburyAssociationIDAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the get ebury association ID accounts params
func (o *GetEburyAssociationIDAccountsParams) WithAssociationID(associationID strfmt.UUID) *GetEburyAssociationIDAccountsParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the get ebury association ID accounts params
func (o *GetEburyAssociationIDAccountsParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEburyAssociationIDAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param association_id
	if err := r.SetPathParam("association_id", o.AssociationID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
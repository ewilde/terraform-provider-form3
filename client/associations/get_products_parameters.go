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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProductsParams creates a new GetProductsParams object
// with the default values initialized.
func NewGetProductsParams() *GetProductsParams {
	var ()
	return &GetProductsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsParamsWithTimeout creates a new GetProductsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductsParamsWithTimeout(timeout time.Duration) *GetProductsParams {
	var ()
	return &GetProductsParams{

		timeout: timeout,
	}
}

// NewGetProductsParamsWithContext creates a new GetProductsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductsParamsWithContext(ctx context.Context) *GetProductsParams {
	var ()
	return &GetProductsParams{

		Context: ctx,
	}
}

// NewGetProductsParamsWithHTTPClient creates a new GetProductsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductsParamsWithHTTPClient(client *http.Client) *GetProductsParams {
	var ()
	return &GetProductsParams{
		HTTPClient: client,
	}
}

/*GetProductsParams contains all the parameters to send to the API endpoint
for the get products operation typically these are written to a http.Request
*/
type GetProductsParams struct {

	/*FilterOrganisationID
	  Organisation id

	*/
	FilterOrganisationID []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get products params
func (o *GetProductsParams) WithTimeout(timeout time.Duration) *GetProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get products params
func (o *GetProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get products params
func (o *GetProductsParams) WithContext(ctx context.Context) *GetProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get products params
func (o *GetProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get products params
func (o *GetProductsParams) WithHTTPClient(client *http.Client) *GetProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get products params
func (o *GetProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get products params
func (o *GetProductsParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetProductsParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get products params
func (o *GetProductsParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

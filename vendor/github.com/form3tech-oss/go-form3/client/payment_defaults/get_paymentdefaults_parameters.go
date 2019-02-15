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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPaymentdefaultsParams creates a new GetPaymentdefaultsParams object
// with the default values initialized.
func NewGetPaymentdefaultsParams() *GetPaymentdefaultsParams {
	var ()
	return &GetPaymentdefaultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentdefaultsParamsWithTimeout creates a new GetPaymentdefaultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentdefaultsParamsWithTimeout(timeout time.Duration) *GetPaymentdefaultsParams {
	var ()
	return &GetPaymentdefaultsParams{

		timeout: timeout,
	}
}

// NewGetPaymentdefaultsParamsWithContext creates a new GetPaymentdefaultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentdefaultsParamsWithContext(ctx context.Context) *GetPaymentdefaultsParams {
	var ()
	return &GetPaymentdefaultsParams{

		Context: ctx,
	}
}

// NewGetPaymentdefaultsParamsWithHTTPClient creates a new GetPaymentdefaultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentdefaultsParamsWithHTTPClient(client *http.Client) *GetPaymentdefaultsParams {
	var ()
	return &GetPaymentdefaultsParams{
		HTTPClient: client,
	}
}

/*GetPaymentdefaultsParams contains all the parameters to send to the API endpoint
for the get paymentdefaults operation typically these are written to a http.Request
*/
type GetPaymentdefaultsParams struct {

	/*FilterOrganisationID
	  Filter by organisation id

	*/
	FilterOrganisationID []strfmt.UUID
	/*PageNumber
	  Which page to select

	*/
	PageNumber *string
	/*PageSize
	  Number of items to select

	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) WithTimeout(timeout time.Duration) *GetPaymentdefaultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) WithContext(ctx context.Context) *GetPaymentdefaultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) WithHTTPClient(client *http.Client) *GetPaymentdefaultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetPaymentdefaultsParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WithPageNumber adds the pageNumber to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) WithPageNumber(pageNumber *string) *GetPaymentdefaultsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) SetPageNumber(pageNumber *string) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) WithPageSize(pageSize *int64) *GetPaymentdefaultsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get paymentdefaults params
func (o *GetPaymentdefaultsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentdefaultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "csv")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

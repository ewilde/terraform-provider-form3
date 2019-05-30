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

// NewGetSepainstantParams creates a new GetSepainstantParams object
// with the default values initialized.
func NewGetSepainstantParams() *GetSepainstantParams {
	var ()
	return &GetSepainstantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSepainstantParamsWithTimeout creates a new GetSepainstantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSepainstantParamsWithTimeout(timeout time.Duration) *GetSepainstantParams {
	var ()
	return &GetSepainstantParams{

		timeout: timeout,
	}
}

// NewGetSepainstantParamsWithContext creates a new GetSepainstantParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSepainstantParamsWithContext(ctx context.Context) *GetSepainstantParams {
	var ()
	return &GetSepainstantParams{

		Context: ctx,
	}
}

// NewGetSepainstantParamsWithHTTPClient creates a new GetSepainstantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSepainstantParamsWithHTTPClient(client *http.Client) *GetSepainstantParams {
	var ()
	return &GetSepainstantParams{
		HTTPClient: client,
	}
}

/*GetSepainstantParams contains all the parameters to send to the API endpoint
for the get sepainstant operation typically these are written to a http.Request
*/
type GetSepainstantParams struct {

	/*FilterOrganisationID
	  Organisation id

	*/
	FilterOrganisationID []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sepainstant params
func (o *GetSepainstantParams) WithTimeout(timeout time.Duration) *GetSepainstantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sepainstant params
func (o *GetSepainstantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sepainstant params
func (o *GetSepainstantParams) WithContext(ctx context.Context) *GetSepainstantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sepainstant params
func (o *GetSepainstantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sepainstant params
func (o *GetSepainstantParams) WithHTTPClient(client *http.Client) *GetSepainstantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sepainstant params
func (o *GetSepainstantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterOrganisationID adds the filterOrganisationID to the get sepainstant params
func (o *GetSepainstantParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetSepainstantParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get sepainstant params
func (o *GetSepainstantParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSepainstantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
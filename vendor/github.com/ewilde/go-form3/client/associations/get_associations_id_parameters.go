// Code generated by go-swagger; DO NOT EDIT.

package associations

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

// NewGetAssociationsIDParams creates a new GetAssociationsIDParams object
// with the default values initialized.
func NewGetAssociationsIDParams() *GetAssociationsIDParams {
	var ()
	return &GetAssociationsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssociationsIDParamsWithTimeout creates a new GetAssociationsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAssociationsIDParamsWithTimeout(timeout time.Duration) *GetAssociationsIDParams {
	var ()
	return &GetAssociationsIDParams{

		timeout: timeout,
	}
}

// NewGetAssociationsIDParamsWithContext creates a new GetAssociationsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAssociationsIDParamsWithContext(ctx context.Context) *GetAssociationsIDParams {
	var ()
	return &GetAssociationsIDParams{

		Context: ctx,
	}
}

// NewGetAssociationsIDParamsWithHTTPClient creates a new GetAssociationsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAssociationsIDParamsWithHTTPClient(client *http.Client) *GetAssociationsIDParams {
	var ()
	return &GetAssociationsIDParams{
		HTTPClient: client,
	}
}

/*GetAssociationsIDParams contains all the parameters to send to the API endpoint
for the get associations ID operation typically these are written to a http.Request
*/
type GetAssociationsIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get associations ID params
func (o *GetAssociationsIDParams) WithTimeout(timeout time.Duration) *GetAssociationsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get associations ID params
func (o *GetAssociationsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get associations ID params
func (o *GetAssociationsIDParams) WithContext(ctx context.Context) *GetAssociationsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get associations ID params
func (o *GetAssociationsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get associations ID params
func (o *GetAssociationsIDParams) WithHTTPClient(client *http.Client) *GetAssociationsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get associations ID params
func (o *GetAssociationsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get associations ID params
func (o *GetAssociationsIDParams) WithID(id strfmt.UUID) *GetAssociationsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get associations ID params
func (o *GetAssociationsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssociationsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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
	"github.com/go-openapi/swag"
)

// NewDeleteSwiftIDParams creates a new DeleteSwiftIDParams object
// with the default values initialized.
func NewDeleteSwiftIDParams() *DeleteSwiftIDParams {
	var ()
	return &DeleteSwiftIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSwiftIDParamsWithTimeout creates a new DeleteSwiftIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSwiftIDParamsWithTimeout(timeout time.Duration) *DeleteSwiftIDParams {
	var ()
	return &DeleteSwiftIDParams{

		timeout: timeout,
	}
}

// NewDeleteSwiftIDParamsWithContext creates a new DeleteSwiftIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSwiftIDParamsWithContext(ctx context.Context) *DeleteSwiftIDParams {
	var ()
	return &DeleteSwiftIDParams{

		Context: ctx,
	}
}

// NewDeleteSwiftIDParamsWithHTTPClient creates a new DeleteSwiftIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSwiftIDParamsWithHTTPClient(client *http.Client) *DeleteSwiftIDParams {
	var ()
	return &DeleteSwiftIDParams{
		HTTPClient: client,
	}
}

/*DeleteSwiftIDParams contains all the parameters to send to the API endpoint
for the delete swift ID operation typically these are written to a http.Request
*/
type DeleteSwiftIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete swift ID params
func (o *DeleteSwiftIDParams) WithTimeout(timeout time.Duration) *DeleteSwiftIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete swift ID params
func (o *DeleteSwiftIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete swift ID params
func (o *DeleteSwiftIDParams) WithContext(ctx context.Context) *DeleteSwiftIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete swift ID params
func (o *DeleteSwiftIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete swift ID params
func (o *DeleteSwiftIDParams) WithHTTPClient(client *http.Client) *DeleteSwiftIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete swift ID params
func (o *DeleteSwiftIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete swift ID params
func (o *DeleteSwiftIDParams) WithID(id strfmt.UUID) *DeleteSwiftIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete swift ID params
func (o *DeleteSwiftIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersion adds the version to the delete swift ID params
func (o *DeleteSwiftIDParams) WithVersion(version int64) *DeleteSwiftIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete swift ID params
func (o *DeleteSwiftIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSwiftIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
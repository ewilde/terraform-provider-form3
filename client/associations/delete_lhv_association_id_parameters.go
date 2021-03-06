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

// NewDeleteLhvAssociationIDParams creates a new DeleteLhvAssociationIDParams object
// with the default values initialized.
func NewDeleteLhvAssociationIDParams() *DeleteLhvAssociationIDParams {
	var ()
	return &DeleteLhvAssociationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLhvAssociationIDParamsWithTimeout creates a new DeleteLhvAssociationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLhvAssociationIDParamsWithTimeout(timeout time.Duration) *DeleteLhvAssociationIDParams {
	var ()
	return &DeleteLhvAssociationIDParams{

		timeout: timeout,
	}
}

// NewDeleteLhvAssociationIDParamsWithContext creates a new DeleteLhvAssociationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLhvAssociationIDParamsWithContext(ctx context.Context) *DeleteLhvAssociationIDParams {
	var ()
	return &DeleteLhvAssociationIDParams{

		Context: ctx,
	}
}

// NewDeleteLhvAssociationIDParamsWithHTTPClient creates a new DeleteLhvAssociationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLhvAssociationIDParamsWithHTTPClient(client *http.Client) *DeleteLhvAssociationIDParams {
	var ()
	return &DeleteLhvAssociationIDParams{
		HTTPClient: client,
	}
}

/*DeleteLhvAssociationIDParams contains all the parameters to send to the API endpoint
for the delete lhv association ID operation typically these are written to a http.Request
*/
type DeleteLhvAssociationIDParams struct {

	/*AssociationID
	  Association Id

	*/
	AssociationID strfmt.UUID
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) WithTimeout(timeout time.Duration) *DeleteLhvAssociationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) WithContext(ctx context.Context) *DeleteLhvAssociationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) WithHTTPClient(client *http.Client) *DeleteLhvAssociationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) WithAssociationID(associationID strfmt.UUID) *DeleteLhvAssociationIDParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WithVersion adds the version to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) WithVersion(version int64) *DeleteLhvAssociationIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete lhv association ID params
func (o *DeleteLhvAssociationIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLhvAssociationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param associationId
	if err := r.SetPathParam("associationId", o.AssociationID.String()); err != nil {
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

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

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams creates a new DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams object
// with the default values initialized.
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams() *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	var ()
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParamsWithTimeout creates a new DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParamsWithTimeout(timeout time.Duration) *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	var ()
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams{

		timeout: timeout,
	}
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParamsWithContext creates a new DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParamsWithContext(ctx context.Context) *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	var ()
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams{

		Context: ctx,
	}
}

// NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParamsWithHTTPClient creates a new DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParamsWithHTTPClient(client *http.Client) *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	var ()
	return &DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams{
		HTTPClient: client,
	}
}

/*DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams contains all the parameters to send to the API endpoint
for the delete lhv association ID agency synchronisations agency synchronisation ID operation typically these are written to a http.Request
*/
type DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams struct {

	/*AgencySynchronisationID
	  Agency synchronisation details Id

	*/
	AgencySynchronisationID strfmt.UUID
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

// WithTimeout adds the timeout to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) WithTimeout(timeout time.Duration) *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) WithContext(ctx context.Context) *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) WithHTTPClient(client *http.Client) *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgencySynchronisationID adds the agencySynchronisationID to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) WithAgencySynchronisationID(agencySynchronisationID strfmt.UUID) *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	o.SetAgencySynchronisationID(agencySynchronisationID)
	return o
}

// SetAgencySynchronisationID adds the agencySynchronisationId to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) SetAgencySynchronisationID(agencySynchronisationID strfmt.UUID) {
	o.AgencySynchronisationID = agencySynchronisationID
}

// WithAssociationID adds the associationID to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) WithAssociationID(associationID strfmt.UUID) *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WithVersion adds the version to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) WithVersion(version int64) *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete lhv association ID agency synchronisations agency synchronisation ID params
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agencySynchronisationId
	if err := r.SetPathParam("agencySynchronisationId", o.AgencySynchronisationID.String()); err != nil {
		return err
	}

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

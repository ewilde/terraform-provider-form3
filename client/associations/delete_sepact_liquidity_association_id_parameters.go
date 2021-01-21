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

// NewDeleteSepactLiquidityAssociationIDParams creates a new DeleteSepactLiquidityAssociationIDParams object
// with the default values initialized.
func NewDeleteSepactLiquidityAssociationIDParams() *DeleteSepactLiquidityAssociationIDParams {
	var ()
	return &DeleteSepactLiquidityAssociationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSepactLiquidityAssociationIDParamsWithTimeout creates a new DeleteSepactLiquidityAssociationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSepactLiquidityAssociationIDParamsWithTimeout(timeout time.Duration) *DeleteSepactLiquidityAssociationIDParams {
	var ()
	return &DeleteSepactLiquidityAssociationIDParams{

		timeout: timeout,
	}
}

// NewDeleteSepactLiquidityAssociationIDParamsWithContext creates a new DeleteSepactLiquidityAssociationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSepactLiquidityAssociationIDParamsWithContext(ctx context.Context) *DeleteSepactLiquidityAssociationIDParams {
	var ()
	return &DeleteSepactLiquidityAssociationIDParams{

		Context: ctx,
	}
}

// NewDeleteSepactLiquidityAssociationIDParamsWithHTTPClient creates a new DeleteSepactLiquidityAssociationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSepactLiquidityAssociationIDParamsWithHTTPClient(client *http.Client) *DeleteSepactLiquidityAssociationIDParams {
	var ()
	return &DeleteSepactLiquidityAssociationIDParams{
		HTTPClient: client,
	}
}

/*DeleteSepactLiquidityAssociationIDParams contains all the parameters to send to the API endpoint
for the delete sepact liquidity association ID operation typically these are written to a http.Request
*/
type DeleteSepactLiquidityAssociationIDParams struct {

	/*AssociationID
	  Association id

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

// WithTimeout adds the timeout to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) WithTimeout(timeout time.Duration) *DeleteSepactLiquidityAssociationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) WithContext(ctx context.Context) *DeleteSepactLiquidityAssociationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) WithHTTPClient(client *http.Client) *DeleteSepactLiquidityAssociationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) WithAssociationID(associationID strfmt.UUID) *DeleteSepactLiquidityAssociationIDParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WithVersion adds the version to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) WithVersion(version int64) *DeleteSepactLiquidityAssociationIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete sepact liquidity association ID params
func (o *DeleteSepactLiquidityAssociationIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSepactLiquidityAssociationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
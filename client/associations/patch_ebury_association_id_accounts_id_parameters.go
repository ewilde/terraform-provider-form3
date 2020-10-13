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

	"github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPatchEburyAssociationIDAccountsIDParams creates a new PatchEburyAssociationIDAccountsIDParams object
// with the default values initialized.
func NewPatchEburyAssociationIDAccountsIDParams() *PatchEburyAssociationIDAccountsIDParams {
	var ()
	return &PatchEburyAssociationIDAccountsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEburyAssociationIDAccountsIDParamsWithTimeout creates a new PatchEburyAssociationIDAccountsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchEburyAssociationIDAccountsIDParamsWithTimeout(timeout time.Duration) *PatchEburyAssociationIDAccountsIDParams {
	var ()
	return &PatchEburyAssociationIDAccountsIDParams{

		timeout: timeout,
	}
}

// NewPatchEburyAssociationIDAccountsIDParamsWithContext creates a new PatchEburyAssociationIDAccountsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchEburyAssociationIDAccountsIDParamsWithContext(ctx context.Context) *PatchEburyAssociationIDAccountsIDParams {
	var ()
	return &PatchEburyAssociationIDAccountsIDParams{

		Context: ctx,
	}
}

// NewPatchEburyAssociationIDAccountsIDParamsWithHTTPClient creates a new PatchEburyAssociationIDAccountsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchEburyAssociationIDAccountsIDParamsWithHTTPClient(client *http.Client) *PatchEburyAssociationIDAccountsIDParams {
	var ()
	return &PatchEburyAssociationIDAccountsIDParams{
		HTTPClient: client,
	}
}

/*PatchEburyAssociationIDAccountsIDParams contains all the parameters to send to the API endpoint
for the patch ebury association ID accounts ID operation typically these are written to a http.Request
*/
type PatchEburyAssociationIDAccountsIDParams struct {

	/*AssociationID
	  Association Id

	*/
	AssociationID strfmt.UUID
	/*ID
	  Association Account Id

	*/
	ID strfmt.UUID
	/*PatchBody*/
	PatchBody *models.EburyAssociationAccountAmendment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) WithTimeout(timeout time.Duration) *PatchEburyAssociationIDAccountsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) WithContext(ctx context.Context) *PatchEburyAssociationIDAccountsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) WithHTTPClient(client *http.Client) *PatchEburyAssociationIDAccountsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) WithAssociationID(associationID strfmt.UUID) *PatchEburyAssociationIDAccountsIDParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WithID adds the id to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) WithID(id strfmt.UUID) *PatchEburyAssociationIDAccountsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPatchBody adds the patchBody to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) WithPatchBody(patchBody *models.EburyAssociationAccountAmendment) *PatchEburyAssociationIDAccountsIDParams {
	o.SetPatchBody(patchBody)
	return o
}

// SetPatchBody adds the patchBody to the patch ebury association ID accounts ID params
func (o *PatchEburyAssociationIDAccountsIDParams) SetPatchBody(patchBody *models.EburyAssociationAccountAmendment) {
	o.PatchBody = patchBody
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEburyAssociationIDAccountsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param association_id
	if err := r.SetPathParam("association_id", o.AssociationID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.PatchBody != nil {
		if err := r.SetBodyParam(o.PatchBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
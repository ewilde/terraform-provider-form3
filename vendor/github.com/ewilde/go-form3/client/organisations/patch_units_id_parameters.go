// Code generated by go-swagger; DO NOT EDIT.

package organisations

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

	"github.com/ewilde/go-form3/models"
)

// NewPatchUnitsIDParams creates a new PatchUnitsIDParams object
// with the default values initialized.
func NewPatchUnitsIDParams() *PatchUnitsIDParams {
	var ()
	return &PatchUnitsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUnitsIDParamsWithTimeout creates a new PatchUnitsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchUnitsIDParamsWithTimeout(timeout time.Duration) *PatchUnitsIDParams {
	var ()
	return &PatchUnitsIDParams{

		timeout: timeout,
	}
}

// NewPatchUnitsIDParamsWithContext creates a new PatchUnitsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchUnitsIDParamsWithContext(ctx context.Context) *PatchUnitsIDParams {
	var ()
	return &PatchUnitsIDParams{

		Context: ctx,
	}
}

// NewPatchUnitsIDParamsWithHTTPClient creates a new PatchUnitsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchUnitsIDParamsWithHTTPClient(client *http.Client) *PatchUnitsIDParams {
	var ()
	return &PatchUnitsIDParams{
		HTTPClient: client,
	}
}

/*PatchUnitsIDParams contains all the parameters to send to the API endpoint
for the patch units ID operation typically these are written to a http.Request
*/
type PatchUnitsIDParams struct {

	/*OrganisationUpdateRequest*/
	OrganisationUpdateRequest *models.OrganisationUpdate
	/*ID
	  Organisation Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch units ID params
func (o *PatchUnitsIDParams) WithTimeout(timeout time.Duration) *PatchUnitsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch units ID params
func (o *PatchUnitsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch units ID params
func (o *PatchUnitsIDParams) WithContext(ctx context.Context) *PatchUnitsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch units ID params
func (o *PatchUnitsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch units ID params
func (o *PatchUnitsIDParams) WithHTTPClient(client *http.Client) *PatchUnitsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch units ID params
func (o *PatchUnitsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganisationUpdateRequest adds the organisationUpdateRequest to the patch units ID params
func (o *PatchUnitsIDParams) WithOrganisationUpdateRequest(organisationUpdateRequest *models.OrganisationUpdate) *PatchUnitsIDParams {
	o.SetOrganisationUpdateRequest(organisationUpdateRequest)
	return o
}

// SetOrganisationUpdateRequest adds the organisationUpdateRequest to the patch units ID params
func (o *PatchUnitsIDParams) SetOrganisationUpdateRequest(organisationUpdateRequest *models.OrganisationUpdate) {
	o.OrganisationUpdateRequest = organisationUpdateRequest
}

// WithID adds the id to the patch units ID params
func (o *PatchUnitsIDParams) WithID(id strfmt.UUID) *PatchUnitsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch units ID params
func (o *PatchUnitsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUnitsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrganisationUpdateRequest != nil {
		if err := r.SetBodyParam(o.OrganisationUpdateRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

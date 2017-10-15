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

// NewPostOrganisationsParams creates a new PostOrganisationsParams object
// with the default values initialized.
func NewPostOrganisationsParams() *PostOrganisationsParams {
	var ()
	return &PostOrganisationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrganisationsParamsWithTimeout creates a new PostOrganisationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOrganisationsParamsWithTimeout(timeout time.Duration) *PostOrganisationsParams {
	var ()
	return &PostOrganisationsParams{

		timeout: timeout,
	}
}

// NewPostOrganisationsParamsWithContext creates a new PostOrganisationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOrganisationsParamsWithContext(ctx context.Context) *PostOrganisationsParams {
	var ()
	return &PostOrganisationsParams{

		Context: ctx,
	}
}

// NewPostOrganisationsParamsWithHTTPClient creates a new PostOrganisationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOrganisationsParamsWithHTTPClient(client *http.Client) *PostOrganisationsParams {
	var ()
	return &PostOrganisationsParams{
		HTTPClient: client,
	}
}

/*PostOrganisationsParams contains all the parameters to send to the API endpoint
for the post organisations operation typically these are written to a http.Request
*/
type PostOrganisationsParams struct {

	/*OrganisationCreationRequest*/
	OrganisationCreationRequest *models.OrganisationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post organisations params
func (o *PostOrganisationsParams) WithTimeout(timeout time.Duration) *PostOrganisationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post organisations params
func (o *PostOrganisationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post organisations params
func (o *PostOrganisationsParams) WithContext(ctx context.Context) *PostOrganisationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post organisations params
func (o *PostOrganisationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post organisations params
func (o *PostOrganisationsParams) WithHTTPClient(client *http.Client) *PostOrganisationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post organisations params
func (o *PostOrganisationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganisationCreationRequest adds the organisationCreationRequest to the post organisations params
func (o *PostOrganisationsParams) WithOrganisationCreationRequest(organisationCreationRequest *models.OrganisationCreation) *PostOrganisationsParams {
	o.SetOrganisationCreationRequest(organisationCreationRequest)
	return o
}

// SetOrganisationCreationRequest adds the organisationCreationRequest to the post organisations params
func (o *PostOrganisationsParams) SetOrganisationCreationRequest(organisationCreationRequest *models.OrganisationCreation) {
	o.OrganisationCreationRequest = organisationCreationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrganisationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrganisationCreationRequest == nil {
		o.OrganisationCreationRequest = new(models.OrganisationCreation)
	}

	if err := r.SetBodyParam(o.OrganisationCreationRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

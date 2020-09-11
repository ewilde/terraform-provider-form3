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

// NewPostEburyAssociationIDAccountsParams creates a new PostEburyAssociationIDAccountsParams object
// with the default values initialized.
func NewPostEburyAssociationIDAccountsParams() *PostEburyAssociationIDAccountsParams {
	var ()
	return &PostEburyAssociationIDAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEburyAssociationIDAccountsParamsWithTimeout creates a new PostEburyAssociationIDAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEburyAssociationIDAccountsParamsWithTimeout(timeout time.Duration) *PostEburyAssociationIDAccountsParams {
	var ()
	return &PostEburyAssociationIDAccountsParams{

		timeout: timeout,
	}
}

// NewPostEburyAssociationIDAccountsParamsWithContext creates a new PostEburyAssociationIDAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEburyAssociationIDAccountsParamsWithContext(ctx context.Context) *PostEburyAssociationIDAccountsParams {
	var ()
	return &PostEburyAssociationIDAccountsParams{

		Context: ctx,
	}
}

// NewPostEburyAssociationIDAccountsParamsWithHTTPClient creates a new PostEburyAssociationIDAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEburyAssociationIDAccountsParamsWithHTTPClient(client *http.Client) *PostEburyAssociationIDAccountsParams {
	var ()
	return &PostEburyAssociationIDAccountsParams{
		HTTPClient: client,
	}
}

/*PostEburyAssociationIDAccountsParams contains all the parameters to send to the API endpoint
for the post ebury association ID accounts operation typically these are written to a http.Request
*/
type PostEburyAssociationIDAccountsParams struct {

	/*AssociationID
	  Association Id

	*/
	AssociationID strfmt.UUID
	/*CreationRequest*/
	CreationRequest *models.EburyAssociationAccountCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) WithTimeout(timeout time.Duration) *PostEburyAssociationIDAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) WithContext(ctx context.Context) *PostEburyAssociationIDAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) WithHTTPClient(client *http.Client) *PostEburyAssociationIDAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociationID adds the associationID to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) WithAssociationID(associationID strfmt.UUID) *PostEburyAssociationIDAccountsParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) SetAssociationID(associationID strfmt.UUID) {
	o.AssociationID = associationID
}

// WithCreationRequest adds the creationRequest to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) WithCreationRequest(creationRequest *models.EburyAssociationAccountCreation) *PostEburyAssociationIDAccountsParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post ebury association ID accounts params
func (o *PostEburyAssociationIDAccountsParams) SetCreationRequest(creationRequest *models.EburyAssociationAccountCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostEburyAssociationIDAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param association_id
	if err := r.SetPathParam("association_id", o.AssociationID.String()); err != nil {
		return err
	}

	if o.CreationRequest != nil {
		if err := r.SetBodyParam(o.CreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
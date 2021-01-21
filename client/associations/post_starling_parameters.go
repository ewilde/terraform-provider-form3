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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/terraform-provider-form3/models"
)

// NewPostStarlingParams creates a new PostStarlingParams object
// with the default values initialized.
func NewPostStarlingParams() *PostStarlingParams {
	var ()
	return &PostStarlingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStarlingParamsWithTimeout creates a new PostStarlingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStarlingParamsWithTimeout(timeout time.Duration) *PostStarlingParams {
	var ()
	return &PostStarlingParams{

		timeout: timeout,
	}
}

// NewPostStarlingParamsWithContext creates a new PostStarlingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostStarlingParamsWithContext(ctx context.Context) *PostStarlingParams {
	var ()
	return &PostStarlingParams{

		Context: ctx,
	}
}

// NewPostStarlingParamsWithHTTPClient creates a new PostStarlingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostStarlingParamsWithHTTPClient(client *http.Client) *PostStarlingParams {
	var ()
	return &PostStarlingParams{
		HTTPClient: client,
	}
}

/*PostStarlingParams contains all the parameters to send to the API endpoint
for the post starling operation typically these are written to a http.Request
*/
type PostStarlingParams struct {

	/*CreationRequest*/
	CreationRequest *models.AssociationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post starling params
func (o *PostStarlingParams) WithTimeout(timeout time.Duration) *PostStarlingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post starling params
func (o *PostStarlingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post starling params
func (o *PostStarlingParams) WithContext(ctx context.Context) *PostStarlingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post starling params
func (o *PostStarlingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post starling params
func (o *PostStarlingParams) WithHTTPClient(client *http.Client) *PostStarlingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post starling params
func (o *PostStarlingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreationRequest adds the creationRequest to the post starling params
func (o *PostStarlingParams) WithCreationRequest(creationRequest *models.AssociationCreation) *PostStarlingParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post starling params
func (o *PostStarlingParams) SetCreationRequest(creationRequest *models.AssociationCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostStarlingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

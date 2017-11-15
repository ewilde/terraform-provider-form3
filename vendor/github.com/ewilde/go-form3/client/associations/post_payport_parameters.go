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

	"github.com/ewilde/go-form3/models"
)

// NewPostPayportParams creates a new PostPayportParams object
// with the default values initialized.
func NewPostPayportParams() *PostPayportParams {
	var ()
	return &PostPayportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPayportParamsWithTimeout creates a new PostPayportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPayportParamsWithTimeout(timeout time.Duration) *PostPayportParams {
	var ()
	return &PostPayportParams{

		timeout: timeout,
	}
}

// NewPostPayportParamsWithContext creates a new PostPayportParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPayportParamsWithContext(ctx context.Context) *PostPayportParams {
	var ()
	return &PostPayportParams{

		Context: ctx,
	}
}

// NewPostPayportParamsWithHTTPClient creates a new PostPayportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPayportParamsWithHTTPClient(client *http.Client) *PostPayportParams {
	var ()
	return &PostPayportParams{
		HTTPClient: client,
	}
}

/*PostPayportParams contains all the parameters to send to the API endpoint
for the post payport operation typically these are written to a http.Request
*/
type PostPayportParams struct {

	/*CreationRequest*/
	CreationRequest *models.PayportAssociationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payport params
func (o *PostPayportParams) WithTimeout(timeout time.Duration) *PostPayportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payport params
func (o *PostPayportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payport params
func (o *PostPayportParams) WithContext(ctx context.Context) *PostPayportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payport params
func (o *PostPayportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payport params
func (o *PostPayportParams) WithHTTPClient(client *http.Client) *PostPayportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payport params
func (o *PostPayportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreationRequest adds the creationRequest to the post payport params
func (o *PostPayportParams) WithCreationRequest(creationRequest *models.PayportAssociationCreation) *PostPayportParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post payport params
func (o *PostPayportParams) SetCreationRequest(creationRequest *models.PayportAssociationCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostPayportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreationRequest == nil {
		o.CreationRequest = new(models.PayportAssociationCreation)
	}

	if err := r.SetBodyParam(o.CreationRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

	models "github.com/form3tech-oss/go-form3/models"
)

// NewPostBicsParams creates a new PostBicsParams object
// with the default values initialized.
func NewPostBicsParams() *PostBicsParams {
	var ()
	return &PostBicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostBicsParamsWithTimeout creates a new PostBicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostBicsParamsWithTimeout(timeout time.Duration) *PostBicsParams {
	var ()
	return &PostBicsParams{

		timeout: timeout,
	}
}

// NewPostBicsParamsWithContext creates a new PostBicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostBicsParamsWithContext(ctx context.Context) *PostBicsParams {
	var ()
	return &PostBicsParams{

		Context: ctx,
	}
}

// NewPostBicsParamsWithHTTPClient creates a new PostBicsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostBicsParamsWithHTTPClient(client *http.Client) *PostBicsParams {
	var ()
	return &PostBicsParams{
		HTTPClient: client,
	}
}

/*PostBicsParams contains all the parameters to send to the API endpoint
for the post bics operation typically these are written to a http.Request
*/
type PostBicsParams struct {

	/*BicCreationRequest*/
	BicCreationRequest *models.BicCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post bics params
func (o *PostBicsParams) WithTimeout(timeout time.Duration) *PostBicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post bics params
func (o *PostBicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post bics params
func (o *PostBicsParams) WithContext(ctx context.Context) *PostBicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post bics params
func (o *PostBicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post bics params
func (o *PostBicsParams) WithHTTPClient(client *http.Client) *PostBicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post bics params
func (o *PostBicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBicCreationRequest adds the bicCreationRequest to the post bics params
func (o *PostBicsParams) WithBicCreationRequest(bicCreationRequest *models.BicCreation) *PostBicsParams {
	o.SetBicCreationRequest(bicCreationRequest)
	return o
}

// SetBicCreationRequest adds the bicCreationRequest to the post bics params
func (o *PostBicsParams) SetBicCreationRequest(bicCreationRequest *models.BicCreation) {
	o.BicCreationRequest = bicCreationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostBicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BicCreationRequest != nil {
		if err := r.SetBodyParam(o.BicCreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewPostSepactLiquidityParams creates a new PostSepactLiquidityParams object
// with the default values initialized.
func NewPostSepactLiquidityParams() *PostSepactLiquidityParams {
	var ()
	return &PostSepactLiquidityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSepactLiquidityParamsWithTimeout creates a new PostSepactLiquidityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSepactLiquidityParamsWithTimeout(timeout time.Duration) *PostSepactLiquidityParams {
	var ()
	return &PostSepactLiquidityParams{

		timeout: timeout,
	}
}

// NewPostSepactLiquidityParamsWithContext creates a new PostSepactLiquidityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSepactLiquidityParamsWithContext(ctx context.Context) *PostSepactLiquidityParams {
	var ()
	return &PostSepactLiquidityParams{

		Context: ctx,
	}
}

// NewPostSepactLiquidityParamsWithHTTPClient creates a new PostSepactLiquidityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSepactLiquidityParamsWithHTTPClient(client *http.Client) *PostSepactLiquidityParams {
	var ()
	return &PostSepactLiquidityParams{
		HTTPClient: client,
	}
}

/*PostSepactLiquidityParams contains all the parameters to send to the API endpoint
for the post sepact liquidity operation typically these are written to a http.Request
*/
type PostSepactLiquidityParams struct {

	/*CreationRequest*/
	CreationRequest *models.SepactLiquidityAssociationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sepact liquidity params
func (o *PostSepactLiquidityParams) WithTimeout(timeout time.Duration) *PostSepactLiquidityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sepact liquidity params
func (o *PostSepactLiquidityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sepact liquidity params
func (o *PostSepactLiquidityParams) WithContext(ctx context.Context) *PostSepactLiquidityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sepact liquidity params
func (o *PostSepactLiquidityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sepact liquidity params
func (o *PostSepactLiquidityParams) WithHTTPClient(client *http.Client) *PostSepactLiquidityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sepact liquidity params
func (o *PostSepactLiquidityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreationRequest adds the creationRequest to the post sepact liquidity params
func (o *PostSepactLiquidityParams) WithCreationRequest(creationRequest *models.SepactLiquidityAssociationCreation) *PostSepactLiquidityParams {
	o.SetCreationRequest(creationRequest)
	return o
}

// SetCreationRequest adds the creationRequest to the post sepact liquidity params
func (o *PostSepactLiquidityParams) SetCreationRequest(creationRequest *models.SepactLiquidityAssociationCreation) {
	o.CreationRequest = creationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostSepactLiquidityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

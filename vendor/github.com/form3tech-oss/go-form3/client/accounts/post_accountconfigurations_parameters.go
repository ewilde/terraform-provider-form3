// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

	models "github.com/form3tech-oss/go-form3/models"
)

// NewPostAccountconfigurationsParams creates a new PostAccountconfigurationsParams object
// with the default values initialized.
func NewPostAccountconfigurationsParams() *PostAccountconfigurationsParams {
	var ()
	return &PostAccountconfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountconfigurationsParamsWithTimeout creates a new PostAccountconfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAccountconfigurationsParamsWithTimeout(timeout time.Duration) *PostAccountconfigurationsParams {
	var ()
	return &PostAccountconfigurationsParams{

		timeout: timeout,
	}
}

// NewPostAccountconfigurationsParamsWithContext creates a new PostAccountconfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAccountconfigurationsParamsWithContext(ctx context.Context) *PostAccountconfigurationsParams {
	var ()
	return &PostAccountconfigurationsParams{

		Context: ctx,
	}
}

// NewPostAccountconfigurationsParamsWithHTTPClient creates a new PostAccountconfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAccountconfigurationsParamsWithHTTPClient(client *http.Client) *PostAccountconfigurationsParams {
	var ()
	return &PostAccountconfigurationsParams{
		HTTPClient: client,
	}
}

/*PostAccountconfigurationsParams contains all the parameters to send to the API endpoint
for the post accountconfigurations operation typically these are written to a http.Request
*/
type PostAccountconfigurationsParams struct {

	/*AccountConfigurationCreationRequest*/
	AccountConfigurationCreationRequest *models.AccountConfigurationCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post accountconfigurations params
func (o *PostAccountconfigurationsParams) WithTimeout(timeout time.Duration) *PostAccountconfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post accountconfigurations params
func (o *PostAccountconfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post accountconfigurations params
func (o *PostAccountconfigurationsParams) WithContext(ctx context.Context) *PostAccountconfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post accountconfigurations params
func (o *PostAccountconfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post accountconfigurations params
func (o *PostAccountconfigurationsParams) WithHTTPClient(client *http.Client) *PostAccountconfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post accountconfigurations params
func (o *PostAccountconfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountConfigurationCreationRequest adds the accountConfigurationCreationRequest to the post accountconfigurations params
func (o *PostAccountconfigurationsParams) WithAccountConfigurationCreationRequest(accountConfigurationCreationRequest *models.AccountConfigurationCreation) *PostAccountconfigurationsParams {
	o.SetAccountConfigurationCreationRequest(accountConfigurationCreationRequest)
	return o
}

// SetAccountConfigurationCreationRequest adds the accountConfigurationCreationRequest to the post accountconfigurations params
func (o *PostAccountconfigurationsParams) SetAccountConfigurationCreationRequest(accountConfigurationCreationRequest *models.AccountConfigurationCreation) {
	o.AccountConfigurationCreationRequest = accountConfigurationCreationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountconfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountConfigurationCreationRequest != nil {
		if err := r.SetBodyParam(o.AccountConfigurationCreationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package users

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
)

// NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams creates a new DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams object
// with the default values initialized.
func NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams() *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	var ()
	return &DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithTimeout creates a new DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithTimeout(timeout time.Duration) *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	var ()
	return &DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams{

		timeout: timeout,
	}
}

// NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithContext creates a new DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithContext(ctx context.Context) *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	var ()
	return &DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams{

		Context: ctx,
	}
}

// NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithHTTPClient creates a new DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParamsWithHTTPClient(client *http.Client) *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	var ()
	return &DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams{
		HTTPClient: client,
	}
}

/*DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams contains all the parameters to send to the API endpoint
for the delete users user ID credentials public key public key ID operation typically these are written to a http.Request
*/
type DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams struct {

	/*PublicKeyID
	  public_key_id

	*/
	PublicKeyID strfmt.UUID
	/*UserID
	  User Id

	*/
	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithTimeout(timeout time.Duration) *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithContext(ctx context.Context) *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithHTTPClient(client *http.Client) *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPublicKeyID adds the publicKeyID to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithPublicKeyID(publicKeyID strfmt.UUID) *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetPublicKeyID(publicKeyID)
	return o
}

// SetPublicKeyID adds the publicKeyId to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetPublicKeyID(publicKeyID strfmt.UUID) {
	o.PublicKeyID = publicKeyID
}

// WithUserID adds the userID to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WithUserID(userID strfmt.UUID) *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete users user ID credentials public key public key ID params
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) SetUserID(userID strfmt.UUID) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param public_key_id
	if err := r.SetPathParam("public_key_id", o.PublicKeyID.String()); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

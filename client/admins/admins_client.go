// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new admins API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for admins API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteAdminsUserIDCredentialsClientID deletes admin credential for user
*/
func (a *Client) DeleteAdminsUserIDCredentialsClientID(params *DeleteAdminsUserIDCredentialsClientIDParams) (*DeleteAdminsUserIDCredentialsClientIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAdminsUserIDCredentialsClientIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAdminsUserIDCredentialsClientID",
		Method:             "DELETE",
		PathPattern:        "/admins/{user_id}/credentials/{client_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAdminsUserIDCredentialsClientIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAdminsUserIDCredentialsClientIDNoContent), nil

}

/*
GetAdminsUserIDCredentials fetches admin credentials for user
*/
func (a *Client) GetAdminsUserIDCredentials(params *GetAdminsUserIDCredentialsParams) (*GetAdminsUserIDCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdminsUserIDCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAdminsUserIDCredentials",
		Method:             "GET",
		PathPattern:        "/admins/{user_id}/credentials",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdminsUserIDCredentialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAdminsUserIDCredentialsOK), nil

}

/*
PostAdminsUserIDCredentials generates new admin credentials for a user
*/
func (a *Client) PostAdminsUserIDCredentials(params *PostAdminsUserIDCredentialsParams) (*PostAdminsUserIDCredentialsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAdminsUserIDCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAdminsUserIDCredentials",
		Method:             "POST",
		PathPattern:        "/admins/{user_id}/credentials",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAdminsUserIDCredentialsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAdminsUserIDCredentialsCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
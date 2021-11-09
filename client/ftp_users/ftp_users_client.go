// Code generated by go-swagger; DO NOT EDIT.

package ftp_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ftp users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ftp users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteFtpusersName(params *DeleteFtpusersNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFtpusersNameOK, error)

	GetFtpusers(params *GetFtpusersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFtpusersOK, error)

	PostFtpusers(params *PostFtpusersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFtpusersOK, error)

	PutFtpusersName(params *PutFtpusersNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFtpusersNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteFtpusersName deletes f t p user
*/
func (a *Client) DeleteFtpusersName(params *DeleteFtpusersNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFtpusersNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFtpusersNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteFtpusersName",
		Method:             "DELETE",
		PathPattern:        "/ftpusers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFtpusersNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteFtpusersNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteFtpusersName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFtpusers gets f t p users
*/
func (a *Client) GetFtpusers(params *GetFtpusersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFtpusersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFtpusersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFtpusers",
		Method:             "GET",
		PathPattern:        "/ftpusers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFtpusersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFtpusersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetFtpusers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostFtpusers creates f t p user
*/
func (a *Client) PostFtpusers(params *PostFtpusersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFtpusersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFtpusersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostFtpusers",
		Method:             "POST",
		PathPattern:        "/ftpusers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFtpusersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostFtpusersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostFtpusers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutFtpusersName updates f t p user
*/
func (a *Client) PutFtpusersName(params *PutFtpusersNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFtpusersNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFtpusersNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutFtpusersName",
		Method:             "PUT",
		PathPattern:        "/ftpusers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutFtpusersNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutFtpusersNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutFtpusersName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

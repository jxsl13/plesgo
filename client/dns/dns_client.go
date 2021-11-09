// Code generated by go-swagger; DO NOT EDIT.

package dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dns API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dns API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDNSRecords(params *GetDNSRecordsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDNSRecordsOK, error)

	GetDNSRecordsID(params *GetDNSRecordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDNSRecordsIDOK, error)

	PutDNSRecordsID(params *PutDNSRecordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutDNSRecordsIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDNSRecords gets domain DNS records
*/
func (a *Client) GetDNSRecords(params *GetDNSRecordsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDNSRecordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDNSRecordsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDNSRecords",
		Method:             "GET",
		PathPattern:        "/dns/records",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDNSRecordsReader{formats: a.formats},
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
	success, ok := result.(*GetDNSRecordsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDNSRecords: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDNSRecordsID gets DNS record
*/
func (a *Client) GetDNSRecordsID(params *GetDNSRecordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDNSRecordsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDNSRecordsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDNSRecordsID",
		Method:             "GET",
		PathPattern:        "/dns/records/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDNSRecordsIDReader{formats: a.formats},
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
	success, ok := result.(*GetDNSRecordsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDNSRecordsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutDNSRecordsID updates DNS record
*/
func (a *Client) PutDNSRecordsID(params *PutDNSRecordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutDNSRecordsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDNSRecordsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutDNSRecordsID",
		Method:             "PUT",
		PathPattern:        "/dns/records/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutDNSRecordsIDReader{formats: a.formats},
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
	success, ok := result.(*PutDNSRecordsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutDNSRecordsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
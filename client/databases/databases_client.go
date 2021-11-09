// Code generated by go-swagger; DO NOT EDIT.

package databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new databases API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for databases API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteDatabasesID(params *DeleteDatabasesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDatabasesIDOK, error)

	DeleteDbusersID(params *DeleteDbusersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDbusersIDOK, error)

	GetDatabases(params *GetDatabasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDatabasesOK, error)

	GetDbservers(params *GetDbserversParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDbserversOK, error)

	GetDbusers(params *GetDbusersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDbusersOK, error)

	PostDatabases(params *PostDatabasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDatabasesOK, error)

	PostDbusers(params *PostDbusersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDbusersOK, error)

	PutDbusersID(params *PutDbusersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutDbusersIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteDatabasesID deletes database
*/
func (a *Client) DeleteDatabasesID(params *DeleteDatabasesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDatabasesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDatabasesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDatabasesID",
		Method:             "DELETE",
		PathPattern:        "/databases/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDatabasesIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteDatabasesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteDatabasesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDbusersID deletes database user
*/
func (a *Client) DeleteDbusersID(params *DeleteDbusersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDbusersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDbusersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDbusersID",
		Method:             "DELETE",
		PathPattern:        "/dbusers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDbusersIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteDbusersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteDbusersID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDatabases gets databases
*/
func (a *Client) GetDatabases(params *GetDatabasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDatabasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatabasesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDatabases",
		Method:             "GET",
		PathPattern:        "/databases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatabasesReader{formats: a.formats},
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
	success, ok := result.(*GetDatabasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDatabases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDbservers gets database servers
*/
func (a *Client) GetDbservers(params *GetDbserversParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDbserversOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDbserversParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDbservers",
		Method:             "GET",
		PathPattern:        "/dbservers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDbserversReader{formats: a.formats},
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
	success, ok := result.(*GetDbserversOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDbservers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDbusers gets database users
*/
func (a *Client) GetDbusers(params *GetDbusersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDbusersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDbusersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDbusers",
		Method:             "GET",
		PathPattern:        "/dbusers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDbusersReader{formats: a.formats},
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
	success, ok := result.(*GetDbusersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDbusers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostDatabases creates database
*/
func (a *Client) PostDatabases(params *PostDatabasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDatabasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDatabasesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostDatabases",
		Method:             "POST",
		PathPattern:        "/databases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostDatabasesReader{formats: a.formats},
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
	success, ok := result.(*PostDatabasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostDatabases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostDbusers creates database user
*/
func (a *Client) PostDbusers(params *PostDbusersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDbusersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDbusersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostDbusers",
		Method:             "POST",
		PathPattern:        "/dbusers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostDbusersReader{formats: a.formats},
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
	success, ok := result.(*PostDbusersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostDbusers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutDbusersID updates database user
*/
func (a *Client) PutDbusersID(params *PutDbusersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutDbusersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDbusersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutDbusersID",
		Method:             "PUT",
		PathPattern:        "/dbusers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutDbusersIDReader{formats: a.formats},
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
	success, ok := result.(*PutDbusersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutDbusersID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
// Code generated by go-swagger; DO NOT EDIT.

package clients

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
	"github.com/go-openapi/swag"

	"github.com/jxsl13/plesgo/models"
)

// NewPutClientsIDParams creates a new PutClientsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutClientsIDParams() *PutClientsIDParams {
	return &PutClientsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutClientsIDParamsWithTimeout creates a new PutClientsIDParams object
// with the ability to set a timeout on a request.
func NewPutClientsIDParamsWithTimeout(timeout time.Duration) *PutClientsIDParams {
	return &PutClientsIDParams{
		timeout: timeout,
	}
}

// NewPutClientsIDParamsWithContext creates a new PutClientsIDParams object
// with the ability to set a context for a request.
func NewPutClientsIDParamsWithContext(ctx context.Context) *PutClientsIDParams {
	return &PutClientsIDParams{
		Context: ctx,
	}
}

// NewPutClientsIDParamsWithHTTPClient creates a new PutClientsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutClientsIDParamsWithHTTPClient(client *http.Client) *PutClientsIDParams {
	return &PutClientsIDParams{
		HTTPClient: client,
	}
}

/* PutClientsIDParams contains all the parameters to send to the API endpoint
   for the put clients ID operation.

   Typically these are written to a http.Request.
*/
type PutClientsIDParams struct {

	/* Body.

	   Client data
	*/
	Body *models.ClientUpdateRequest

	/* ID.

	   Client ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put clients ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutClientsIDParams) WithDefaults() *PutClientsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put clients ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutClientsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put clients ID params
func (o *PutClientsIDParams) WithTimeout(timeout time.Duration) *PutClientsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put clients ID params
func (o *PutClientsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put clients ID params
func (o *PutClientsIDParams) WithContext(ctx context.Context) *PutClientsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put clients ID params
func (o *PutClientsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put clients ID params
func (o *PutClientsIDParams) WithHTTPClient(client *http.Client) *PutClientsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put clients ID params
func (o *PutClientsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put clients ID params
func (o *PutClientsIDParams) WithBody(body *models.ClientUpdateRequest) *PutClientsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put clients ID params
func (o *PutClientsIDParams) SetBody(body *models.ClientUpdateRequest) {
	o.Body = body
}

// WithID adds the id to the put clients ID params
func (o *PutClientsIDParams) WithID(id int64) *PutClientsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put clients ID params
func (o *PutClientsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutClientsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

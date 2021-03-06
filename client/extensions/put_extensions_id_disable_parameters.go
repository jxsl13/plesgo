// Code generated by go-swagger; DO NOT EDIT.

package extensions

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
)

// NewPutExtensionsIDDisableParams creates a new PutExtensionsIDDisableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutExtensionsIDDisableParams() *PutExtensionsIDDisableParams {
	return &PutExtensionsIDDisableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutExtensionsIDDisableParamsWithTimeout creates a new PutExtensionsIDDisableParams object
// with the ability to set a timeout on a request.
func NewPutExtensionsIDDisableParamsWithTimeout(timeout time.Duration) *PutExtensionsIDDisableParams {
	return &PutExtensionsIDDisableParams{
		timeout: timeout,
	}
}

// NewPutExtensionsIDDisableParamsWithContext creates a new PutExtensionsIDDisableParams object
// with the ability to set a context for a request.
func NewPutExtensionsIDDisableParamsWithContext(ctx context.Context) *PutExtensionsIDDisableParams {
	return &PutExtensionsIDDisableParams{
		Context: ctx,
	}
}

// NewPutExtensionsIDDisableParamsWithHTTPClient creates a new PutExtensionsIDDisableParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutExtensionsIDDisableParamsWithHTTPClient(client *http.Client) *PutExtensionsIDDisableParams {
	return &PutExtensionsIDDisableParams{
		HTTPClient: client,
	}
}

/* PutExtensionsIDDisableParams contains all the parameters to send to the API endpoint
   for the put extensions ID disable operation.

   Typically these are written to a http.Request.
*/
type PutExtensionsIDDisableParams struct {

	/* ID.

	   Extension identifier
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put extensions ID disable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutExtensionsIDDisableParams) WithDefaults() *PutExtensionsIDDisableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put extensions ID disable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutExtensionsIDDisableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put extensions ID disable params
func (o *PutExtensionsIDDisableParams) WithTimeout(timeout time.Duration) *PutExtensionsIDDisableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put extensions ID disable params
func (o *PutExtensionsIDDisableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put extensions ID disable params
func (o *PutExtensionsIDDisableParams) WithContext(ctx context.Context) *PutExtensionsIDDisableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put extensions ID disable params
func (o *PutExtensionsIDDisableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put extensions ID disable params
func (o *PutExtensionsIDDisableParams) WithHTTPClient(client *http.Client) *PutExtensionsIDDisableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put extensions ID disable params
func (o *PutExtensionsIDDisableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put extensions ID disable params
func (o *PutExtensionsIDDisableParams) WithID(id string) *PutExtensionsIDDisableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put extensions ID disable params
func (o *PutExtensionsIDDisableParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutExtensionsIDDisableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

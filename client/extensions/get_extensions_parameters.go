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

// NewGetExtensionsParams creates a new GetExtensionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExtensionsParams() *GetExtensionsParams {
	return &GetExtensionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExtensionsParamsWithTimeout creates a new GetExtensionsParams object
// with the ability to set a timeout on a request.
func NewGetExtensionsParamsWithTimeout(timeout time.Duration) *GetExtensionsParams {
	return &GetExtensionsParams{
		timeout: timeout,
	}
}

// NewGetExtensionsParamsWithContext creates a new GetExtensionsParams object
// with the ability to set a context for a request.
func NewGetExtensionsParamsWithContext(ctx context.Context) *GetExtensionsParams {
	return &GetExtensionsParams{
		Context: ctx,
	}
}

// NewGetExtensionsParamsWithHTTPClient creates a new GetExtensionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExtensionsParamsWithHTTPClient(client *http.Client) *GetExtensionsParams {
	return &GetExtensionsParams{
		HTTPClient: client,
	}
}

/* GetExtensionsParams contains all the parameters to send to the API endpoint
   for the get extensions operation.

   Typically these are written to a http.Request.
*/
type GetExtensionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get extensions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExtensionsParams) WithDefaults() *GetExtensionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get extensions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExtensionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get extensions params
func (o *GetExtensionsParams) WithTimeout(timeout time.Duration) *GetExtensionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get extensions params
func (o *GetExtensionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get extensions params
func (o *GetExtensionsParams) WithContext(ctx context.Context) *GetExtensionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get extensions params
func (o *GetExtensionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get extensions params
func (o *GetExtensionsParams) WithHTTPClient(client *http.Client) *GetExtensionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get extensions params
func (o *GetExtensionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetExtensionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

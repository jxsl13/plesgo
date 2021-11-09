// Code generated by go-swagger; DO NOT EDIT.

package domains

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
)

// NewGetDomainsIDClientParams creates a new GetDomainsIDClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomainsIDClientParams() *GetDomainsIDClientParams {
	return &GetDomainsIDClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainsIDClientParamsWithTimeout creates a new GetDomainsIDClientParams object
// with the ability to set a timeout on a request.
func NewGetDomainsIDClientParamsWithTimeout(timeout time.Duration) *GetDomainsIDClientParams {
	return &GetDomainsIDClientParams{
		timeout: timeout,
	}
}

// NewGetDomainsIDClientParamsWithContext creates a new GetDomainsIDClientParams object
// with the ability to set a context for a request.
func NewGetDomainsIDClientParamsWithContext(ctx context.Context) *GetDomainsIDClientParams {
	return &GetDomainsIDClientParams{
		Context: ctx,
	}
}

// NewGetDomainsIDClientParamsWithHTTPClient creates a new GetDomainsIDClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomainsIDClientParamsWithHTTPClient(client *http.Client) *GetDomainsIDClientParams {
	return &GetDomainsIDClientParams{
		HTTPClient: client,
	}
}

/* GetDomainsIDClientParams contains all the parameters to send to the API endpoint
   for the get domains ID client operation.

   Typically these are written to a http.Request.
*/
type GetDomainsIDClientParams struct {

	/* ID.

	   Domain ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domains ID client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainsIDClientParams) WithDefaults() *GetDomainsIDClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domains ID client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainsIDClientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domains ID client params
func (o *GetDomainsIDClientParams) WithTimeout(timeout time.Duration) *GetDomainsIDClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domains ID client params
func (o *GetDomainsIDClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domains ID client params
func (o *GetDomainsIDClientParams) WithContext(ctx context.Context) *GetDomainsIDClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domains ID client params
func (o *GetDomainsIDClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domains ID client params
func (o *GetDomainsIDClientParams) WithHTTPClient(client *http.Client) *GetDomainsIDClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domains ID client params
func (o *GetDomainsIDClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get domains ID client params
func (o *GetDomainsIDClientParams) WithID(id int64) *GetDomainsIDClientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get domains ID client params
func (o *GetDomainsIDClientParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainsIDClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

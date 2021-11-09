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

// NewDeleteDomainsIDParams creates a new DeleteDomainsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDomainsIDParams() *DeleteDomainsIDParams {
	return &DeleteDomainsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDomainsIDParamsWithTimeout creates a new DeleteDomainsIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDomainsIDParamsWithTimeout(timeout time.Duration) *DeleteDomainsIDParams {
	return &DeleteDomainsIDParams{
		timeout: timeout,
	}
}

// NewDeleteDomainsIDParamsWithContext creates a new DeleteDomainsIDParams object
// with the ability to set a context for a request.
func NewDeleteDomainsIDParamsWithContext(ctx context.Context) *DeleteDomainsIDParams {
	return &DeleteDomainsIDParams{
		Context: ctx,
	}
}

// NewDeleteDomainsIDParamsWithHTTPClient creates a new DeleteDomainsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDomainsIDParamsWithHTTPClient(client *http.Client) *DeleteDomainsIDParams {
	return &DeleteDomainsIDParams{
		HTTPClient: client,
	}
}

/* DeleteDomainsIDParams contains all the parameters to send to the API endpoint
   for the delete domains ID operation.

   Typically these are written to a http.Request.
*/
type DeleteDomainsIDParams struct {

	/* ID.

	   Domain ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete domains ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDomainsIDParams) WithDefaults() *DeleteDomainsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete domains ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDomainsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete domains ID params
func (o *DeleteDomainsIDParams) WithTimeout(timeout time.Duration) *DeleteDomainsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete domains ID params
func (o *DeleteDomainsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete domains ID params
func (o *DeleteDomainsIDParams) WithContext(ctx context.Context) *DeleteDomainsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete domains ID params
func (o *DeleteDomainsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete domains ID params
func (o *DeleteDomainsIDParams) WithHTTPClient(client *http.Client) *DeleteDomainsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete domains ID params
func (o *DeleteDomainsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete domains ID params
func (o *DeleteDomainsIDParams) WithID(id int64) *DeleteDomainsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete domains ID params
func (o *DeleteDomainsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDomainsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

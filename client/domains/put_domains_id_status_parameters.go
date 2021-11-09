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

	"github.com/jxsl13/plesgo/models"
)

// NewPutDomainsIDStatusParams creates a new PutDomainsIDStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDomainsIDStatusParams() *PutDomainsIDStatusParams {
	return &PutDomainsIDStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDomainsIDStatusParamsWithTimeout creates a new PutDomainsIDStatusParams object
// with the ability to set a timeout on a request.
func NewPutDomainsIDStatusParamsWithTimeout(timeout time.Duration) *PutDomainsIDStatusParams {
	return &PutDomainsIDStatusParams{
		timeout: timeout,
	}
}

// NewPutDomainsIDStatusParamsWithContext creates a new PutDomainsIDStatusParams object
// with the ability to set a context for a request.
func NewPutDomainsIDStatusParamsWithContext(ctx context.Context) *PutDomainsIDStatusParams {
	return &PutDomainsIDStatusParams{
		Context: ctx,
	}
}

// NewPutDomainsIDStatusParamsWithHTTPClient creates a new PutDomainsIDStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDomainsIDStatusParamsWithHTTPClient(client *http.Client) *PutDomainsIDStatusParams {
	return &PutDomainsIDStatusParams{
		HTTPClient: client,
	}
}

/* PutDomainsIDStatusParams contains all the parameters to send to the API endpoint
   for the put domains ID status operation.

   Typically these are written to a http.Request.
*/
type PutDomainsIDStatusParams struct {

	/* Body.

	   Domain status
	*/
	Body *models.DomainStatus

	/* ID.

	   Domain ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put domains ID status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDomainsIDStatusParams) WithDefaults() *PutDomainsIDStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put domains ID status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDomainsIDStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put domains ID status params
func (o *PutDomainsIDStatusParams) WithTimeout(timeout time.Duration) *PutDomainsIDStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put domains ID status params
func (o *PutDomainsIDStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put domains ID status params
func (o *PutDomainsIDStatusParams) WithContext(ctx context.Context) *PutDomainsIDStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put domains ID status params
func (o *PutDomainsIDStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put domains ID status params
func (o *PutDomainsIDStatusParams) WithHTTPClient(client *http.Client) *PutDomainsIDStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put domains ID status params
func (o *PutDomainsIDStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put domains ID status params
func (o *PutDomainsIDStatusParams) WithBody(body *models.DomainStatus) *PutDomainsIDStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put domains ID status params
func (o *PutDomainsIDStatusParams) SetBody(body *models.DomainStatus) {
	o.Body = body
}

// WithID adds the id to the put domains ID status params
func (o *PutDomainsIDStatusParams) WithID(id int64) *PutDomainsIDStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put domains ID status params
func (o *PutDomainsIDStatusParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDomainsIDStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package dns

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

// NewPutDNSRecordsIDParams creates a new PutDNSRecordsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDNSRecordsIDParams() *PutDNSRecordsIDParams {
	return &PutDNSRecordsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDNSRecordsIDParamsWithTimeout creates a new PutDNSRecordsIDParams object
// with the ability to set a timeout on a request.
func NewPutDNSRecordsIDParamsWithTimeout(timeout time.Duration) *PutDNSRecordsIDParams {
	return &PutDNSRecordsIDParams{
		timeout: timeout,
	}
}

// NewPutDNSRecordsIDParamsWithContext creates a new PutDNSRecordsIDParams object
// with the ability to set a context for a request.
func NewPutDNSRecordsIDParamsWithContext(ctx context.Context) *PutDNSRecordsIDParams {
	return &PutDNSRecordsIDParams{
		Context: ctx,
	}
}

// NewPutDNSRecordsIDParamsWithHTTPClient creates a new PutDNSRecordsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDNSRecordsIDParamsWithHTTPClient(client *http.Client) *PutDNSRecordsIDParams {
	return &PutDNSRecordsIDParams{
		HTTPClient: client,
	}
}

/* PutDNSRecordsIDParams contains all the parameters to send to the API endpoint
   for the put DNS records ID operation.

   Typically these are written to a http.Request.
*/
type PutDNSRecordsIDParams struct {

	/* Body.

	   DNS record data
	*/
	Body *models.DNSRecord

	/* ID.

	   DNS record ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put DNS records ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDNSRecordsIDParams) WithDefaults() *PutDNSRecordsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put DNS records ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDNSRecordsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put DNS records ID params
func (o *PutDNSRecordsIDParams) WithTimeout(timeout time.Duration) *PutDNSRecordsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put DNS records ID params
func (o *PutDNSRecordsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put DNS records ID params
func (o *PutDNSRecordsIDParams) WithContext(ctx context.Context) *PutDNSRecordsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put DNS records ID params
func (o *PutDNSRecordsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put DNS records ID params
func (o *PutDNSRecordsIDParams) WithHTTPClient(client *http.Client) *PutDNSRecordsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put DNS records ID params
func (o *PutDNSRecordsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put DNS records ID params
func (o *PutDNSRecordsIDParams) WithBody(body *models.DNSRecord) *PutDNSRecordsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put DNS records ID params
func (o *PutDNSRecordsIDParams) SetBody(body *models.DNSRecord) {
	o.Body = body
}

// WithID adds the id to the put DNS records ID params
func (o *PutDNSRecordsIDParams) WithID(id int64) *PutDNSRecordsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put DNS records ID params
func (o *PutDNSRecordsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutDNSRecordsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package databases

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

	"github.com/jxsl13/plesgo/models"
)

// NewPostDatabasesParams creates a new PostDatabasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDatabasesParams() *PostDatabasesParams {
	return &PostDatabasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDatabasesParamsWithTimeout creates a new PostDatabasesParams object
// with the ability to set a timeout on a request.
func NewPostDatabasesParamsWithTimeout(timeout time.Duration) *PostDatabasesParams {
	return &PostDatabasesParams{
		timeout: timeout,
	}
}

// NewPostDatabasesParamsWithContext creates a new PostDatabasesParams object
// with the ability to set a context for a request.
func NewPostDatabasesParamsWithContext(ctx context.Context) *PostDatabasesParams {
	return &PostDatabasesParams{
		Context: ctx,
	}
}

// NewPostDatabasesParamsWithHTTPClient creates a new PostDatabasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDatabasesParamsWithHTTPClient(client *http.Client) *PostDatabasesParams {
	return &PostDatabasesParams{
		HTTPClient: client,
	}
}

/* PostDatabasesParams contains all the parameters to send to the API endpoint
   for the post databases operation.

   Typically these are written to a http.Request.
*/
type PostDatabasesParams struct {

	/* Body.

	   Database data
	*/
	Body *models.DatabaseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post databases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDatabasesParams) WithDefaults() *PostDatabasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post databases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDatabasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post databases params
func (o *PostDatabasesParams) WithTimeout(timeout time.Duration) *PostDatabasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post databases params
func (o *PostDatabasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post databases params
func (o *PostDatabasesParams) WithContext(ctx context.Context) *PostDatabasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post databases params
func (o *PostDatabasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post databases params
func (o *PostDatabasesParams) WithHTTPClient(client *http.Client) *PostDatabasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post databases params
func (o *PostDatabasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post databases params
func (o *PostDatabasesParams) WithBody(body *models.DatabaseRequest) *PostDatabasesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post databases params
func (o *PostDatabasesParams) SetBody(body *models.DatabaseRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDatabasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

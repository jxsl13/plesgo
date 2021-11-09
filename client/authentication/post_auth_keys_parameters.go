// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewPostAuthKeysParams creates a new PostAuthKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAuthKeysParams() *PostAuthKeysParams {
	return &PostAuthKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthKeysParamsWithTimeout creates a new PostAuthKeysParams object
// with the ability to set a timeout on a request.
func NewPostAuthKeysParamsWithTimeout(timeout time.Duration) *PostAuthKeysParams {
	return &PostAuthKeysParams{
		timeout: timeout,
	}
}

// NewPostAuthKeysParamsWithContext creates a new PostAuthKeysParams object
// with the ability to set a context for a request.
func NewPostAuthKeysParamsWithContext(ctx context.Context) *PostAuthKeysParams {
	return &PostAuthKeysParams{
		Context: ctx,
	}
}

// NewPostAuthKeysParamsWithHTTPClient creates a new PostAuthKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAuthKeysParamsWithHTTPClient(client *http.Client) *PostAuthKeysParams {
	return &PostAuthKeysParams{
		HTTPClient: client,
	}
}

/* PostAuthKeysParams contains all the parameters to send to the API endpoint
   for the post auth keys operation.

   Typically these are written to a http.Request.
*/
type PostAuthKeysParams struct {

	/* Body.

	   Key parameters
	*/
	Body *models.SecretKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post auth keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthKeysParams) WithDefaults() *PostAuthKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post auth keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post auth keys params
func (o *PostAuthKeysParams) WithTimeout(timeout time.Duration) *PostAuthKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auth keys params
func (o *PostAuthKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auth keys params
func (o *PostAuthKeysParams) WithContext(ctx context.Context) *PostAuthKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auth keys params
func (o *PostAuthKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auth keys params
func (o *PostAuthKeysParams) WithHTTPClient(client *http.Client) *PostAuthKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auth keys params
func (o *PostAuthKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post auth keys params
func (o *PostAuthKeysParams) WithBody(body *models.SecretKeyRequest) *PostAuthKeysParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post auth keys params
func (o *PostAuthKeysParams) SetBody(body *models.SecretKeyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRepositoriesMyParams creates a new GetRepositoriesMyParams object
// with the default values initialized.
func NewGetRepositoriesMyParams() *GetRepositoriesMyParams {

	return &GetRepositoriesMyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesMyParamsWithTimeout creates a new GetRepositoriesMyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesMyParamsWithTimeout(timeout time.Duration) *GetRepositoriesMyParams {

	return &GetRepositoriesMyParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesMyParamsWithContext creates a new GetRepositoriesMyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesMyParamsWithContext(ctx context.Context) *GetRepositoriesMyParams {

	return &GetRepositoriesMyParams{

		Context: ctx,
	}
}

// NewGetRepositoriesMyParamsWithHTTPClient creates a new GetRepositoriesMyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesMyParamsWithHTTPClient(client *http.Client) *GetRepositoriesMyParams {

	return &GetRepositoriesMyParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesMyParams contains all the parameters to send to the API endpoint
for the get repositories my operation typically these are written to a http.Request
*/
type GetRepositoriesMyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repositories my params
func (o *GetRepositoriesMyParams) WithTimeout(timeout time.Duration) *GetRepositoriesMyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories my params
func (o *GetRepositoriesMyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories my params
func (o *GetRepositoriesMyParams) WithContext(ctx context.Context) *GetRepositoriesMyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories my params
func (o *GetRepositoriesMyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories my params
func (o *GetRepositoriesMyParams) WithHTTPClient(client *http.Client) *GetRepositoriesMyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories my params
func (o *GetRepositoriesMyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesMyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

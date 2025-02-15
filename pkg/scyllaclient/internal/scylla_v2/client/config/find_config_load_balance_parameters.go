// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigLoadBalanceParams creates a new FindConfigLoadBalanceParams object
// with the default values initialized.
func NewFindConfigLoadBalanceParams() *FindConfigLoadBalanceParams {

	return &FindConfigLoadBalanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigLoadBalanceParamsWithTimeout creates a new FindConfigLoadBalanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigLoadBalanceParamsWithTimeout(timeout time.Duration) *FindConfigLoadBalanceParams {

	return &FindConfigLoadBalanceParams{

		timeout: timeout,
	}
}

// NewFindConfigLoadBalanceParamsWithContext creates a new FindConfigLoadBalanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigLoadBalanceParamsWithContext(ctx context.Context) *FindConfigLoadBalanceParams {

	return &FindConfigLoadBalanceParams{

		Context: ctx,
	}
}

// NewFindConfigLoadBalanceParamsWithHTTPClient creates a new FindConfigLoadBalanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigLoadBalanceParamsWithHTTPClient(client *http.Client) *FindConfigLoadBalanceParams {

	return &FindConfigLoadBalanceParams{
		HTTPClient: client,
	}
}

/*
FindConfigLoadBalanceParams contains all the parameters to send to the API endpoint
for the find config load balance operation typically these are written to a http.Request
*/
type FindConfigLoadBalanceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config load balance params
func (o *FindConfigLoadBalanceParams) WithTimeout(timeout time.Duration) *FindConfigLoadBalanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config load balance params
func (o *FindConfigLoadBalanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config load balance params
func (o *FindConfigLoadBalanceParams) WithContext(ctx context.Context) *FindConfigLoadBalanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config load balance params
func (o *FindConfigLoadBalanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config load balance params
func (o *FindConfigLoadBalanceParams) WithHTTPClient(client *http.Client) *FindConfigLoadBalanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config load balance params
func (o *FindConfigLoadBalanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigLoadBalanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

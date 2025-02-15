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

// NewFindConfigCommitlogUseoDsyncParams creates a new FindConfigCommitlogUseoDsyncParams object
// with the default values initialized.
func NewFindConfigCommitlogUseoDsyncParams() *FindConfigCommitlogUseoDsyncParams {

	return &FindConfigCommitlogUseoDsyncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCommitlogUseoDsyncParamsWithTimeout creates a new FindConfigCommitlogUseoDsyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigCommitlogUseoDsyncParamsWithTimeout(timeout time.Duration) *FindConfigCommitlogUseoDsyncParams {

	return &FindConfigCommitlogUseoDsyncParams{

		timeout: timeout,
	}
}

// NewFindConfigCommitlogUseoDsyncParamsWithContext creates a new FindConfigCommitlogUseoDsyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigCommitlogUseoDsyncParamsWithContext(ctx context.Context) *FindConfigCommitlogUseoDsyncParams {

	return &FindConfigCommitlogUseoDsyncParams{

		Context: ctx,
	}
}

// NewFindConfigCommitlogUseoDsyncParamsWithHTTPClient creates a new FindConfigCommitlogUseoDsyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigCommitlogUseoDsyncParamsWithHTTPClient(client *http.Client) *FindConfigCommitlogUseoDsyncParams {

	return &FindConfigCommitlogUseoDsyncParams{
		HTTPClient: client,
	}
}

/*
FindConfigCommitlogUseoDsyncParams contains all the parameters to send to the API endpoint
for the find config commitlog use o dsync operation typically these are written to a http.Request
*/
type FindConfigCommitlogUseoDsyncParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseoDsyncParams) WithTimeout(timeout time.Duration) *FindConfigCommitlogUseoDsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseoDsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseoDsyncParams) WithContext(ctx context.Context) *FindConfigCommitlogUseoDsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseoDsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseoDsyncParams) WithHTTPClient(client *http.Client) *FindConfigCommitlogUseoDsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseoDsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCommitlogUseoDsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

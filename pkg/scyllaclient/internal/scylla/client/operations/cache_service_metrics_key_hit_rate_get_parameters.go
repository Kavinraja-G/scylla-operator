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

// NewCacheServiceMetricsKeyHitRateGetParams creates a new CacheServiceMetricsKeyHitRateGetParams object
// with the default values initialized.
func NewCacheServiceMetricsKeyHitRateGetParams() *CacheServiceMetricsKeyHitRateGetParams {

	return &CacheServiceMetricsKeyHitRateGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsKeyHitRateGetParamsWithTimeout creates a new CacheServiceMetricsKeyHitRateGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsKeyHitRateGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsKeyHitRateGetParams {

	return &CacheServiceMetricsKeyHitRateGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsKeyHitRateGetParamsWithContext creates a new CacheServiceMetricsKeyHitRateGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsKeyHitRateGetParamsWithContext(ctx context.Context) *CacheServiceMetricsKeyHitRateGetParams {

	return &CacheServiceMetricsKeyHitRateGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsKeyHitRateGetParamsWithHTTPClient creates a new CacheServiceMetricsKeyHitRateGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsKeyHitRateGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsKeyHitRateGetParams {

	return &CacheServiceMetricsKeyHitRateGetParams{
		HTTPClient: client,
	}
}

/*
CacheServiceMetricsKeyHitRateGetParams contains all the parameters to send to the API endpoint
for the cache service metrics key hit rate get operation typically these are written to a http.Request
*/
type CacheServiceMetricsKeyHitRateGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics key hit rate get params
func (o *CacheServiceMetricsKeyHitRateGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsKeyHitRateGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics key hit rate get params
func (o *CacheServiceMetricsKeyHitRateGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics key hit rate get params
func (o *CacheServiceMetricsKeyHitRateGetParams) WithContext(ctx context.Context) *CacheServiceMetricsKeyHitRateGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics key hit rate get params
func (o *CacheServiceMetricsKeyHitRateGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics key hit rate get params
func (o *CacheServiceMetricsKeyHitRateGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsKeyHitRateGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics key hit rate get params
func (o *CacheServiceMetricsKeyHitRateGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsKeyHitRateGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

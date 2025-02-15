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

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams creates a new ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams() *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams {
	var ()
	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams {
	var ()
	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParamsWithContext creates a new ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams {
	var ()
	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams {
	var ()
	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics all memtables off heap size by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics all memtables off heap size by name get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics all memtables off heap size by name get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics all memtables off heap size by name get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics all memtables off heap size by name get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics all memtables off heap size by name get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics all memtables off heap size by name get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics all memtables off heap size by name get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams) WithName(name string) *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics all memtables off heap size by name get params
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

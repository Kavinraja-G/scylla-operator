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

// NewHintedHandoffMetricsCreateHintByAddrGetParams creates a new HintedHandoffMetricsCreateHintByAddrGetParams object
// with the default values initialized.
func NewHintedHandoffMetricsCreateHintByAddrGetParams() *HintedHandoffMetricsCreateHintByAddrGetParams {
	var ()
	return &HintedHandoffMetricsCreateHintByAddrGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHintedHandoffMetricsCreateHintByAddrGetParamsWithTimeout creates a new HintedHandoffMetricsCreateHintByAddrGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHintedHandoffMetricsCreateHintByAddrGetParamsWithTimeout(timeout time.Duration) *HintedHandoffMetricsCreateHintByAddrGetParams {
	var ()
	return &HintedHandoffMetricsCreateHintByAddrGetParams{

		timeout: timeout,
	}
}

// NewHintedHandoffMetricsCreateHintByAddrGetParamsWithContext creates a new HintedHandoffMetricsCreateHintByAddrGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewHintedHandoffMetricsCreateHintByAddrGetParamsWithContext(ctx context.Context) *HintedHandoffMetricsCreateHintByAddrGetParams {
	var ()
	return &HintedHandoffMetricsCreateHintByAddrGetParams{

		Context: ctx,
	}
}

// NewHintedHandoffMetricsCreateHintByAddrGetParamsWithHTTPClient creates a new HintedHandoffMetricsCreateHintByAddrGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHintedHandoffMetricsCreateHintByAddrGetParamsWithHTTPClient(client *http.Client) *HintedHandoffMetricsCreateHintByAddrGetParams {
	var ()
	return &HintedHandoffMetricsCreateHintByAddrGetParams{
		HTTPClient: client,
	}
}

/*
HintedHandoffMetricsCreateHintByAddrGetParams contains all the parameters to send to the API endpoint
for the hinted handoff metrics create hint by addr get operation typically these are written to a http.Request
*/
type HintedHandoffMetricsCreateHintByAddrGetParams struct {

	/*Addr
	  The peer address

	*/
	Addr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WithTimeout(timeout time.Duration) *HintedHandoffMetricsCreateHintByAddrGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WithContext(ctx context.Context) *HintedHandoffMetricsCreateHintByAddrGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WithHTTPClient(client *http.Client) *HintedHandoffMetricsCreateHintByAddrGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddr adds the addr to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WithAddr(addr string) *HintedHandoffMetricsCreateHintByAddrGetParams {
	o.SetAddr(addr)
	return o
}

// SetAddr adds the addr to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) SetAddr(addr string) {
	o.Addr = addr
}

// WriteToRequest writes these params to a swagger request
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addr
	if err := r.SetPathParam("addr", o.Addr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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
	"github.com/go-openapi/strfmt"
)

// NewCacheServiceMetricsCounterCapacityGetParams creates a new CacheServiceMetricsCounterCapacityGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCacheServiceMetricsCounterCapacityGetParams() *CacheServiceMetricsCounterCapacityGetParams {
	return &CacheServiceMetricsCounterCapacityGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsCounterCapacityGetParamsWithTimeout creates a new CacheServiceMetricsCounterCapacityGetParams object
// with the ability to set a timeout on a request.
func NewCacheServiceMetricsCounterCapacityGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsCounterCapacityGetParams {
	return &CacheServiceMetricsCounterCapacityGetParams{
		timeout: timeout,
	}
}

// NewCacheServiceMetricsCounterCapacityGetParamsWithContext creates a new CacheServiceMetricsCounterCapacityGetParams object
// with the ability to set a context for a request.
func NewCacheServiceMetricsCounterCapacityGetParamsWithContext(ctx context.Context) *CacheServiceMetricsCounterCapacityGetParams {
	return &CacheServiceMetricsCounterCapacityGetParams{
		Context: ctx,
	}
}

// NewCacheServiceMetricsCounterCapacityGetParamsWithHTTPClient creates a new CacheServiceMetricsCounterCapacityGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCacheServiceMetricsCounterCapacityGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsCounterCapacityGetParams {
	return &CacheServiceMetricsCounterCapacityGetParams{
		HTTPClient: client,
	}
}

/*
CacheServiceMetricsCounterCapacityGetParams contains all the parameters to send to the API endpoint

	for the cache service metrics counter capacity get operation.

	Typically these are written to a http.Request.
*/
type CacheServiceMetricsCounterCapacityGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cache service metrics counter capacity get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CacheServiceMetricsCounterCapacityGetParams) WithDefaults() *CacheServiceMetricsCounterCapacityGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cache service metrics counter capacity get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CacheServiceMetricsCounterCapacityGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cache service metrics counter capacity get params
func (o *CacheServiceMetricsCounterCapacityGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsCounterCapacityGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics counter capacity get params
func (o *CacheServiceMetricsCounterCapacityGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics counter capacity get params
func (o *CacheServiceMetricsCounterCapacityGetParams) WithContext(ctx context.Context) *CacheServiceMetricsCounterCapacityGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics counter capacity get params
func (o *CacheServiceMetricsCounterCapacityGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics counter capacity get params
func (o *CacheServiceMetricsCounterCapacityGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsCounterCapacityGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics counter capacity get params
func (o *CacheServiceMetricsCounterCapacityGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsCounterCapacityGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
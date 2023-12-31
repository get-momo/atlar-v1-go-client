// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// NewGetV1TransactionsIDParams creates a new GetV1TransactionsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TransactionsIDParams() *GetV1TransactionsIDParams {
	return &GetV1TransactionsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TransactionsIDParamsWithTimeout creates a new GetV1TransactionsIDParams object
// with the ability to set a timeout on a request.
func NewGetV1TransactionsIDParamsWithTimeout(timeout time.Duration) *GetV1TransactionsIDParams {
	return &GetV1TransactionsIDParams{
		timeout: timeout,
	}
}

// NewGetV1TransactionsIDParamsWithContext creates a new GetV1TransactionsIDParams object
// with the ability to set a context for a request.
func NewGetV1TransactionsIDParamsWithContext(ctx context.Context) *GetV1TransactionsIDParams {
	return &GetV1TransactionsIDParams{
		Context: ctx,
	}
}

// NewGetV1TransactionsIDParamsWithHTTPClient creates a new GetV1TransactionsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TransactionsIDParamsWithHTTPClient(client *http.Client) *GetV1TransactionsIDParams {
	return &GetV1TransactionsIDParams{
		HTTPClient: client,
	}
}

/*
GetV1TransactionsIDParams contains all the parameters to send to the API endpoint

	for the get v1 transactions ID operation.

	Typically these are written to a http.Request.
*/
type GetV1TransactionsIDParams struct {

	/* ID.

	   Transaction ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 transactions ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TransactionsIDParams) WithDefaults() *GetV1TransactionsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 transactions ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TransactionsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 transactions ID params
func (o *GetV1TransactionsIDParams) WithTimeout(timeout time.Duration) *GetV1TransactionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 transactions ID params
func (o *GetV1TransactionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 transactions ID params
func (o *GetV1TransactionsIDParams) WithContext(ctx context.Context) *GetV1TransactionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 transactions ID params
func (o *GetV1TransactionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 transactions ID params
func (o *GetV1TransactionsIDParams) WithHTTPClient(client *http.Client) *GetV1TransactionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 transactions ID params
func (o *GetV1TransactionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 transactions ID params
func (o *GetV1TransactionsIDParams) WithID(id string) *GetV1TransactionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 transactions ID params
func (o *GetV1TransactionsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TransactionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

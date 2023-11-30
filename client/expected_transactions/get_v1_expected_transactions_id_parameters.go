// Code generated by go-swagger; DO NOT EDIT.

package expected_transactions

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

// NewGetV1ExpectedTransactionsIDParams creates a new GetV1ExpectedTransactionsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ExpectedTransactionsIDParams() *GetV1ExpectedTransactionsIDParams {
	return &GetV1ExpectedTransactionsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ExpectedTransactionsIDParamsWithTimeout creates a new GetV1ExpectedTransactionsIDParams object
// with the ability to set a timeout on a request.
func NewGetV1ExpectedTransactionsIDParamsWithTimeout(timeout time.Duration) *GetV1ExpectedTransactionsIDParams {
	return &GetV1ExpectedTransactionsIDParams{
		timeout: timeout,
	}
}

// NewGetV1ExpectedTransactionsIDParamsWithContext creates a new GetV1ExpectedTransactionsIDParams object
// with the ability to set a context for a request.
func NewGetV1ExpectedTransactionsIDParamsWithContext(ctx context.Context) *GetV1ExpectedTransactionsIDParams {
	return &GetV1ExpectedTransactionsIDParams{
		Context: ctx,
	}
}

// NewGetV1ExpectedTransactionsIDParamsWithHTTPClient creates a new GetV1ExpectedTransactionsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ExpectedTransactionsIDParamsWithHTTPClient(client *http.Client) *GetV1ExpectedTransactionsIDParams {
	return &GetV1ExpectedTransactionsIDParams{
		HTTPClient: client,
	}
}

/*
GetV1ExpectedTransactionsIDParams contains all the parameters to send to the API endpoint

	for the get v1 expected transactions ID operation.

	Typically these are written to a http.Request.
*/
type GetV1ExpectedTransactionsIDParams struct {

	/* ID.

	   Expected Transaction ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 expected transactions ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ExpectedTransactionsIDParams) WithDefaults() *GetV1ExpectedTransactionsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 expected transactions ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ExpectedTransactionsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 expected transactions ID params
func (o *GetV1ExpectedTransactionsIDParams) WithTimeout(timeout time.Duration) *GetV1ExpectedTransactionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 expected transactions ID params
func (o *GetV1ExpectedTransactionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 expected transactions ID params
func (o *GetV1ExpectedTransactionsIDParams) WithContext(ctx context.Context) *GetV1ExpectedTransactionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 expected transactions ID params
func (o *GetV1ExpectedTransactionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 expected transactions ID params
func (o *GetV1ExpectedTransactionsIDParams) WithHTTPClient(client *http.Client) *GetV1ExpectedTransactionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 expected transactions ID params
func (o *GetV1ExpectedTransactionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 expected transactions ID params
func (o *GetV1ExpectedTransactionsIDParams) WithID(id string) *GetV1ExpectedTransactionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 expected transactions ID params
func (o *GetV1ExpectedTransactionsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ExpectedTransactionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
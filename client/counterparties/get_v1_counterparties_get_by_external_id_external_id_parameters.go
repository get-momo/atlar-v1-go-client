// Code generated by go-swagger; DO NOT EDIT.

package counterparties

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

// NewGetV1CounterpartiesGetByExternalIDExternalIDParams creates a new GetV1CounterpartiesGetByExternalIDExternalIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1CounterpartiesGetByExternalIDExternalIDParams() *GetV1CounterpartiesGetByExternalIDExternalIDParams {
	return &GetV1CounterpartiesGetByExternalIDExternalIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1CounterpartiesGetByExternalIDExternalIDParamsWithTimeout creates a new GetV1CounterpartiesGetByExternalIDExternalIDParams object
// with the ability to set a timeout on a request.
func NewGetV1CounterpartiesGetByExternalIDExternalIDParamsWithTimeout(timeout time.Duration) *GetV1CounterpartiesGetByExternalIDExternalIDParams {
	return &GetV1CounterpartiesGetByExternalIDExternalIDParams{
		timeout: timeout,
	}
}

// NewGetV1CounterpartiesGetByExternalIDExternalIDParamsWithContext creates a new GetV1CounterpartiesGetByExternalIDExternalIDParams object
// with the ability to set a context for a request.
func NewGetV1CounterpartiesGetByExternalIDExternalIDParamsWithContext(ctx context.Context) *GetV1CounterpartiesGetByExternalIDExternalIDParams {
	return &GetV1CounterpartiesGetByExternalIDExternalIDParams{
		Context: ctx,
	}
}

// NewGetV1CounterpartiesGetByExternalIDExternalIDParamsWithHTTPClient creates a new GetV1CounterpartiesGetByExternalIDExternalIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1CounterpartiesGetByExternalIDExternalIDParamsWithHTTPClient(client *http.Client) *GetV1CounterpartiesGetByExternalIDExternalIDParams {
	return &GetV1CounterpartiesGetByExternalIDExternalIDParams{
		HTTPClient: client,
	}
}

/*
GetV1CounterpartiesGetByExternalIDExternalIDParams contains all the parameters to send to the API endpoint

	for the get v1 counterparties get by external ID external ID operation.

	Typically these are written to a http.Request.
*/
type GetV1CounterpartiesGetByExternalIDExternalIDParams struct {

	/* ExternalID.

	   The externalId of the counterparty. Supplied when creating counterparties.
	*/
	ExternalID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 counterparties get by external ID external ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) WithDefaults() *GetV1CounterpartiesGetByExternalIDExternalIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 counterparties get by external ID external ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 counterparties get by external ID external ID params
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) WithTimeout(timeout time.Duration) *GetV1CounterpartiesGetByExternalIDExternalIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 counterparties get by external ID external ID params
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 counterparties get by external ID external ID params
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) WithContext(ctx context.Context) *GetV1CounterpartiesGetByExternalIDExternalIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 counterparties get by external ID external ID params
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 counterparties get by external ID external ID params
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) WithHTTPClient(client *http.Client) *GetV1CounterpartiesGetByExternalIDExternalIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 counterparties get by external ID external ID params
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalID adds the externalID to the get v1 counterparties get by external ID external ID params
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) WithExternalID(externalID string) *GetV1CounterpartiesGetByExternalIDExternalIDParams {
	o.SetExternalID(externalID)
	return o
}

// SetExternalID adds the externalId to the get v1 counterparties get by external ID external ID params
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) SetExternalID(externalID string) {
	o.ExternalID = externalID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1CounterpartiesGetByExternalIDExternalIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param externalId
	if err := r.SetPathParam("externalId", o.ExternalID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

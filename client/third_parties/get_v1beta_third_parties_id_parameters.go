// Code generated by go-swagger; DO NOT EDIT.

package third_parties

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

// NewGetV1betaThirdPartiesIDParams creates a new GetV1betaThirdPartiesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1betaThirdPartiesIDParams() *GetV1betaThirdPartiesIDParams {
	return &GetV1betaThirdPartiesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1betaThirdPartiesIDParamsWithTimeout creates a new GetV1betaThirdPartiesIDParams object
// with the ability to set a timeout on a request.
func NewGetV1betaThirdPartiesIDParamsWithTimeout(timeout time.Duration) *GetV1betaThirdPartiesIDParams {
	return &GetV1betaThirdPartiesIDParams{
		timeout: timeout,
	}
}

// NewGetV1betaThirdPartiesIDParamsWithContext creates a new GetV1betaThirdPartiesIDParams object
// with the ability to set a context for a request.
func NewGetV1betaThirdPartiesIDParamsWithContext(ctx context.Context) *GetV1betaThirdPartiesIDParams {
	return &GetV1betaThirdPartiesIDParams{
		Context: ctx,
	}
}

// NewGetV1betaThirdPartiesIDParamsWithHTTPClient creates a new GetV1betaThirdPartiesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1betaThirdPartiesIDParamsWithHTTPClient(client *http.Client) *GetV1betaThirdPartiesIDParams {
	return &GetV1betaThirdPartiesIDParams{
		HTTPClient: client,
	}
}

/*
GetV1betaThirdPartiesIDParams contains all the parameters to send to the API endpoint

	for the get v1beta third parties ID operation.

	Typically these are written to a http.Request.
*/
type GetV1betaThirdPartiesIDParams struct {

	/* ID.

	   ThirdParty ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1beta third parties ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1betaThirdPartiesIDParams) WithDefaults() *GetV1betaThirdPartiesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1beta third parties ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1betaThirdPartiesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1beta third parties ID params
func (o *GetV1betaThirdPartiesIDParams) WithTimeout(timeout time.Duration) *GetV1betaThirdPartiesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1beta third parties ID params
func (o *GetV1betaThirdPartiesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1beta third parties ID params
func (o *GetV1betaThirdPartiesIDParams) WithContext(ctx context.Context) *GetV1betaThirdPartiesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1beta third parties ID params
func (o *GetV1betaThirdPartiesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1beta third parties ID params
func (o *GetV1betaThirdPartiesIDParams) WithHTTPClient(client *http.Client) *GetV1betaThirdPartiesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1beta third parties ID params
func (o *GetV1betaThirdPartiesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1beta third parties ID params
func (o *GetV1betaThirdPartiesIDParams) WithID(id string) *GetV1betaThirdPartiesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1beta third parties ID params
func (o *GetV1betaThirdPartiesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1betaThirdPartiesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
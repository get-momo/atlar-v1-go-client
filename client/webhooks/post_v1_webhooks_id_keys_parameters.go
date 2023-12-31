// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

// NewPostV1WebhooksIDKeysParams creates a new PostV1WebhooksIDKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1WebhooksIDKeysParams() *PostV1WebhooksIDKeysParams {
	return &PostV1WebhooksIDKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1WebhooksIDKeysParamsWithTimeout creates a new PostV1WebhooksIDKeysParams object
// with the ability to set a timeout on a request.
func NewPostV1WebhooksIDKeysParamsWithTimeout(timeout time.Duration) *PostV1WebhooksIDKeysParams {
	return &PostV1WebhooksIDKeysParams{
		timeout: timeout,
	}
}

// NewPostV1WebhooksIDKeysParamsWithContext creates a new PostV1WebhooksIDKeysParams object
// with the ability to set a context for a request.
func NewPostV1WebhooksIDKeysParamsWithContext(ctx context.Context) *PostV1WebhooksIDKeysParams {
	return &PostV1WebhooksIDKeysParams{
		Context: ctx,
	}
}

// NewPostV1WebhooksIDKeysParamsWithHTTPClient creates a new PostV1WebhooksIDKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1WebhooksIDKeysParamsWithHTTPClient(client *http.Client) *PostV1WebhooksIDKeysParams {
	return &PostV1WebhooksIDKeysParams{
		HTTPClient: client,
	}
}

/*
PostV1WebhooksIDKeysParams contains all the parameters to send to the API endpoint

	for the post v1 webhooks ID keys operation.

	Typically these are written to a http.Request.
*/
type PostV1WebhooksIDKeysParams struct {

	/* ID.

	   Webhook ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 webhooks ID keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1WebhooksIDKeysParams) WithDefaults() *PostV1WebhooksIDKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 webhooks ID keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1WebhooksIDKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 webhooks ID keys params
func (o *PostV1WebhooksIDKeysParams) WithTimeout(timeout time.Duration) *PostV1WebhooksIDKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 webhooks ID keys params
func (o *PostV1WebhooksIDKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 webhooks ID keys params
func (o *PostV1WebhooksIDKeysParams) WithContext(ctx context.Context) *PostV1WebhooksIDKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 webhooks ID keys params
func (o *PostV1WebhooksIDKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 webhooks ID keys params
func (o *PostV1WebhooksIDKeysParams) WithHTTPClient(client *http.Client) *PostV1WebhooksIDKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 webhooks ID keys params
func (o *PostV1WebhooksIDKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post v1 webhooks ID keys params
func (o *PostV1WebhooksIDKeysParams) WithID(id string) *PostV1WebhooksIDKeysParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post v1 webhooks ID keys params
func (o *PostV1WebhooksIDKeysParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1WebhooksIDKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

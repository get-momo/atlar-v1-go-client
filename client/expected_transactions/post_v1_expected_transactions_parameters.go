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

	"github.com/get-momo/atlar-v1-go-client/models"
)

// NewPostV1ExpectedTransactionsParams creates a new PostV1ExpectedTransactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ExpectedTransactionsParams() *PostV1ExpectedTransactionsParams {
	return &PostV1ExpectedTransactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ExpectedTransactionsParamsWithTimeout creates a new PostV1ExpectedTransactionsParams object
// with the ability to set a timeout on a request.
func NewPostV1ExpectedTransactionsParamsWithTimeout(timeout time.Duration) *PostV1ExpectedTransactionsParams {
	return &PostV1ExpectedTransactionsParams{
		timeout: timeout,
	}
}

// NewPostV1ExpectedTransactionsParamsWithContext creates a new PostV1ExpectedTransactionsParams object
// with the ability to set a context for a request.
func NewPostV1ExpectedTransactionsParamsWithContext(ctx context.Context) *PostV1ExpectedTransactionsParams {
	return &PostV1ExpectedTransactionsParams{
		Context: ctx,
	}
}

// NewPostV1ExpectedTransactionsParamsWithHTTPClient creates a new PostV1ExpectedTransactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ExpectedTransactionsParamsWithHTTPClient(client *http.Client) *PostV1ExpectedTransactionsParams {
	return &PostV1ExpectedTransactionsParams{
		HTTPClient: client,
	}
}

/*
PostV1ExpectedTransactionsParams contains all the parameters to send to the API endpoint

	for the post v1 expected transactions operation.

	Typically these are written to a http.Request.
*/
type PostV1ExpectedTransactionsParams struct {

	/* Transaction.

	   Expected Transaction
	*/
	Transaction *models.CreateExpectedTransactionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 expected transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ExpectedTransactionsParams) WithDefaults() *PostV1ExpectedTransactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 expected transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ExpectedTransactionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 expected transactions params
func (o *PostV1ExpectedTransactionsParams) WithTimeout(timeout time.Duration) *PostV1ExpectedTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 expected transactions params
func (o *PostV1ExpectedTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 expected transactions params
func (o *PostV1ExpectedTransactionsParams) WithContext(ctx context.Context) *PostV1ExpectedTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 expected transactions params
func (o *PostV1ExpectedTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 expected transactions params
func (o *PostV1ExpectedTransactionsParams) WithHTTPClient(client *http.Client) *PostV1ExpectedTransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 expected transactions params
func (o *PostV1ExpectedTransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransaction adds the transaction to the post v1 expected transactions params
func (o *PostV1ExpectedTransactionsParams) WithTransaction(transaction *models.CreateExpectedTransactionRequest) *PostV1ExpectedTransactionsParams {
	o.SetTransaction(transaction)
	return o
}

// SetTransaction adds the transaction to the post v1 expected transactions params
func (o *PostV1ExpectedTransactionsParams) SetTransaction(transaction *models.CreateExpectedTransactionRequest) {
	o.Transaction = transaction
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ExpectedTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Transaction != nil {
		if err := r.SetBodyParam(o.Transaction); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

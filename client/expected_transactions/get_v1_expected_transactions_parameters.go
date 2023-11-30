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
	"github.com/go-openapi/swag"
)

// NewGetV1ExpectedTransactionsParams creates a new GetV1ExpectedTransactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ExpectedTransactionsParams() *GetV1ExpectedTransactionsParams {
	return &GetV1ExpectedTransactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ExpectedTransactionsParamsWithTimeout creates a new GetV1ExpectedTransactionsParams object
// with the ability to set a timeout on a request.
func NewGetV1ExpectedTransactionsParamsWithTimeout(timeout time.Duration) *GetV1ExpectedTransactionsParams {
	return &GetV1ExpectedTransactionsParams{
		timeout: timeout,
	}
}

// NewGetV1ExpectedTransactionsParamsWithContext creates a new GetV1ExpectedTransactionsParams object
// with the ability to set a context for a request.
func NewGetV1ExpectedTransactionsParamsWithContext(ctx context.Context) *GetV1ExpectedTransactionsParams {
	return &GetV1ExpectedTransactionsParams{
		Context: ctx,
	}
}

// NewGetV1ExpectedTransactionsParamsWithHTTPClient creates a new GetV1ExpectedTransactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ExpectedTransactionsParamsWithHTTPClient(client *http.Client) *GetV1ExpectedTransactionsParams {
	return &GetV1ExpectedTransactionsParams{
		HTTPClient: client,
	}
}

/*
GetV1ExpectedTransactionsParams contains all the parameters to send to the API endpoint

	for the get v1 expected transactions operation.

	Typically these are written to a http.Request.
*/
type GetV1ExpectedTransactionsParams struct {

	/* AccountID.

	   Filters out transactions of the given account.
	*/
	AccountID *string

	/* FromDate.

	   Filters out transactions with date after or equal to the given date. Format as `yyyy-mm-dd`.
	*/
	FromDate *string

	/* Limit.

	   Limit resulting response list. Defaults to 100.
	*/
	Limit *int64

	/* Status.

	   Filters out transactions of the given status (`CREATED`, `RECONCILED` or `ARCHIVED`).
	*/
	Status *string

	/* ToDate.

	   Filters out transactions with date before or equal to the given date. Format as `yyyy-mm-dd`.
	*/
	ToDate *string

	/* Token.

	   If response comes back and `nextToken` is populated, use that value in this `token` query to continue pagination.
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 expected transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ExpectedTransactionsParams) WithDefaults() *GetV1ExpectedTransactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 expected transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ExpectedTransactionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) WithTimeout(timeout time.Duration) *GetV1ExpectedTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) WithContext(ctx context.Context) *GetV1ExpectedTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) WithHTTPClient(client *http.Client) *GetV1ExpectedTransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) WithAccountID(accountID *string) *GetV1ExpectedTransactionsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithFromDate adds the fromDate to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) WithFromDate(fromDate *string) *GetV1ExpectedTransactionsParams {
	o.SetFromDate(fromDate)
	return o
}

// SetFromDate adds the fromDate to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) SetFromDate(fromDate *string) {
	o.FromDate = fromDate
}

// WithLimit adds the limit to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) WithLimit(limit *int64) *GetV1ExpectedTransactionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStatus adds the status to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) WithStatus(status *string) *GetV1ExpectedTransactionsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) SetStatus(status *string) {
	o.Status = status
}

// WithToDate adds the toDate to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) WithToDate(toDate *string) *GetV1ExpectedTransactionsParams {
	o.SetToDate(toDate)
	return o
}

// SetToDate adds the toDate to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) SetToDate(toDate *string) {
	o.ToDate = toDate
}

// WithToken adds the token to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) WithToken(token *string) *GetV1ExpectedTransactionsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get v1 expected transactions params
func (o *GetV1ExpectedTransactionsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ExpectedTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID string

		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {

			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}
	}

	if o.FromDate != nil {

		// query param fromDate
		var qrFromDate string

		if o.FromDate != nil {
			qrFromDate = *o.FromDate
		}
		qFromDate := qrFromDate
		if qFromDate != "" {

			if err := r.SetQueryParam("fromDate", qFromDate); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.ToDate != nil {

		// query param toDate
		var qrToDate string

		if o.ToDate != nil {
			qrToDate = *o.ToDate
		}
		qToDate := qrToDate
		if qToDate != "" {

			if err := r.SetQueryParam("toDate", qToDate); err != nil {
				return err
			}
		}
	}

	if o.Token != nil {

		// query param token
		var qrToken string

		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {

			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

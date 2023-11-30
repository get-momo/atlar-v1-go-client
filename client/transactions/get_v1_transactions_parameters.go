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
	"github.com/go-openapi/swag"
)

// NewGetV1TransactionsParams creates a new GetV1TransactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TransactionsParams() *GetV1TransactionsParams {
	return &GetV1TransactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TransactionsParamsWithTimeout creates a new GetV1TransactionsParams object
// with the ability to set a timeout on a request.
func NewGetV1TransactionsParamsWithTimeout(timeout time.Duration) *GetV1TransactionsParams {
	return &GetV1TransactionsParams{
		timeout: timeout,
	}
}

// NewGetV1TransactionsParamsWithContext creates a new GetV1TransactionsParams object
// with the ability to set a context for a request.
func NewGetV1TransactionsParamsWithContext(ctx context.Context) *GetV1TransactionsParams {
	return &GetV1TransactionsParams{
		Context: ctx,
	}
}

// NewGetV1TransactionsParamsWithHTTPClient creates a new GetV1TransactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TransactionsParamsWithHTTPClient(client *http.Client) *GetV1TransactionsParams {
	return &GetV1TransactionsParams{
		HTTPClient: client,
	}
}

/*
GetV1TransactionsParams contains all the parameters to send to the API endpoint

	for the get v1 transactions operation.

	Typically these are written to a http.Request.
*/
type GetV1TransactionsParams struct {

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

	   Filters out transactions of the given reconciliation status.
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

// WithDefaults hydrates default values in the get v1 transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TransactionsParams) WithDefaults() *GetV1TransactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TransactionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 transactions params
func (o *GetV1TransactionsParams) WithTimeout(timeout time.Duration) *GetV1TransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 transactions params
func (o *GetV1TransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 transactions params
func (o *GetV1TransactionsParams) WithContext(ctx context.Context) *GetV1TransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 transactions params
func (o *GetV1TransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 transactions params
func (o *GetV1TransactionsParams) WithHTTPClient(client *http.Client) *GetV1TransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 transactions params
func (o *GetV1TransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get v1 transactions params
func (o *GetV1TransactionsParams) WithAccountID(accountID *string) *GetV1TransactionsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get v1 transactions params
func (o *GetV1TransactionsParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithFromDate adds the fromDate to the get v1 transactions params
func (o *GetV1TransactionsParams) WithFromDate(fromDate *string) *GetV1TransactionsParams {
	o.SetFromDate(fromDate)
	return o
}

// SetFromDate adds the fromDate to the get v1 transactions params
func (o *GetV1TransactionsParams) SetFromDate(fromDate *string) {
	o.FromDate = fromDate
}

// WithLimit adds the limit to the get v1 transactions params
func (o *GetV1TransactionsParams) WithLimit(limit *int64) *GetV1TransactionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get v1 transactions params
func (o *GetV1TransactionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStatus adds the status to the get v1 transactions params
func (o *GetV1TransactionsParams) WithStatus(status *string) *GetV1TransactionsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get v1 transactions params
func (o *GetV1TransactionsParams) SetStatus(status *string) {
	o.Status = status
}

// WithToDate adds the toDate to the get v1 transactions params
func (o *GetV1TransactionsParams) WithToDate(toDate *string) *GetV1TransactionsParams {
	o.SetToDate(toDate)
	return o
}

// SetToDate adds the toDate to the get v1 transactions params
func (o *GetV1TransactionsParams) SetToDate(toDate *string) {
	o.ToDate = toDate
}

// WithToken adds the token to the get v1 transactions params
func (o *GetV1TransactionsParams) WithToken(token *string) *GetV1TransactionsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get v1 transactions params
func (o *GetV1TransactionsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
// Code generated by go-swagger; DO NOT EDIT.

package external_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/get-momo/atlar-v1-go-client/models"
)

// PostV1ExternalAccountsReader is a Reader for the PostV1ExternalAccounts structure.
type PostV1ExternalAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ExternalAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ExternalAccountsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1ExternalAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/external-accounts] PostV1ExternalAccounts", response, response.Code())
	}
}

// NewPostV1ExternalAccountsCreated creates a PostV1ExternalAccountsCreated with default headers values
func NewPostV1ExternalAccountsCreated() *PostV1ExternalAccountsCreated {
	return &PostV1ExternalAccountsCreated{}
}

/*
PostV1ExternalAccountsCreated describes a response with status code 201, with default header values.

The created external account
*/
type PostV1ExternalAccountsCreated struct {
	Payload *models.ExternalAccount
}

// IsSuccess returns true when this post v1 external accounts created response has a 2xx status code
func (o *PostV1ExternalAccountsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 external accounts created response has a 3xx status code
func (o *PostV1ExternalAccountsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 external accounts created response has a 4xx status code
func (o *PostV1ExternalAccountsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 external accounts created response has a 5xx status code
func (o *PostV1ExternalAccountsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 external accounts created response a status code equal to that given
func (o *PostV1ExternalAccountsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 external accounts created response
func (o *PostV1ExternalAccountsCreated) Code() int {
	return 201
}

func (o *PostV1ExternalAccountsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/external-accounts][%d] postV1ExternalAccountsCreated  %+v", 201, o.Payload)
}

func (o *PostV1ExternalAccountsCreated) String() string {
	return fmt.Sprintf("[POST /v1/external-accounts][%d] postV1ExternalAccountsCreated  %+v", 201, o.Payload)
}

func (o *PostV1ExternalAccountsCreated) GetPayload() *models.ExternalAccount {
	return o.Payload
}

func (o *PostV1ExternalAccountsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ExternalAccountsBadRequest creates a PostV1ExternalAccountsBadRequest with default headers values
func NewPostV1ExternalAccountsBadRequest() *PostV1ExternalAccountsBadRequest {
	return &PostV1ExternalAccountsBadRequest{}
}

/*
PostV1ExternalAccountsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostV1ExternalAccountsBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this post v1 external accounts bad request response has a 2xx status code
func (o *PostV1ExternalAccountsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 external accounts bad request response has a 3xx status code
func (o *PostV1ExternalAccountsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 external accounts bad request response has a 4xx status code
func (o *PostV1ExternalAccountsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 external accounts bad request response has a 5xx status code
func (o *PostV1ExternalAccountsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 external accounts bad request response a status code equal to that given
func (o *PostV1ExternalAccountsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 external accounts bad request response
func (o *PostV1ExternalAccountsBadRequest) Code() int {
	return 400
}

func (o *PostV1ExternalAccountsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/external-accounts][%d] postV1ExternalAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ExternalAccountsBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/external-accounts][%d] postV1ExternalAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ExternalAccountsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostV1ExternalAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

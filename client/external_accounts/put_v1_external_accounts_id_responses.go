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

// PutV1ExternalAccountsIDReader is a Reader for the PutV1ExternalAccountsID structure.
type PutV1ExternalAccountsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1ExternalAccountsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1ExternalAccountsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutV1ExternalAccountsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutV1ExternalAccountsIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/external-accounts/{id}] PutV1ExternalAccountsID", response, response.Code())
	}
}

// NewPutV1ExternalAccountsIDOK creates a PutV1ExternalAccountsIDOK with default headers values
func NewPutV1ExternalAccountsIDOK() *PutV1ExternalAccountsIDOK {
	return &PutV1ExternalAccountsIDOK{}
}

/*
PutV1ExternalAccountsIDOK describes a response with status code 200, with default header values.

The updated external account
*/
type PutV1ExternalAccountsIDOK struct {
	Payload *models.ExternalAccount
}

// IsSuccess returns true when this put v1 external accounts Id o k response has a 2xx status code
func (o *PutV1ExternalAccountsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 external accounts Id o k response has a 3xx status code
func (o *PutV1ExternalAccountsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 external accounts Id o k response has a 4xx status code
func (o *PutV1ExternalAccountsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 external accounts Id o k response has a 5xx status code
func (o *PutV1ExternalAccountsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 external accounts Id o k response a status code equal to that given
func (o *PutV1ExternalAccountsIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put v1 external accounts Id o k response
func (o *PutV1ExternalAccountsIDOK) Code() int {
	return 200
}

func (o *PutV1ExternalAccountsIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/external-accounts/{id}][%d] putV1ExternalAccountsIdOK  %+v", 200, o.Payload)
}

func (o *PutV1ExternalAccountsIDOK) String() string {
	return fmt.Sprintf("[PUT /v1/external-accounts/{id}][%d] putV1ExternalAccountsIdOK  %+v", 200, o.Payload)
}

func (o *PutV1ExternalAccountsIDOK) GetPayload() *models.ExternalAccount {
	return o.Payload
}

func (o *PutV1ExternalAccountsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1ExternalAccountsIDBadRequest creates a PutV1ExternalAccountsIDBadRequest with default headers values
func NewPutV1ExternalAccountsIDBadRequest() *PutV1ExternalAccountsIDBadRequest {
	return &PutV1ExternalAccountsIDBadRequest{}
}

/*
PutV1ExternalAccountsIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutV1ExternalAccountsIDBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this put v1 external accounts Id bad request response has a 2xx status code
func (o *PutV1ExternalAccountsIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 external accounts Id bad request response has a 3xx status code
func (o *PutV1ExternalAccountsIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 external accounts Id bad request response has a 4xx status code
func (o *PutV1ExternalAccountsIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 external accounts Id bad request response has a 5xx status code
func (o *PutV1ExternalAccountsIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 external accounts Id bad request response a status code equal to that given
func (o *PutV1ExternalAccountsIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put v1 external accounts Id bad request response
func (o *PutV1ExternalAccountsIDBadRequest) Code() int {
	return 400
}

func (o *PutV1ExternalAccountsIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/external-accounts/{id}][%d] putV1ExternalAccountsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1ExternalAccountsIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/external-accounts/{id}][%d] putV1ExternalAccountsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1ExternalAccountsIDBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutV1ExternalAccountsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1ExternalAccountsIDConflict creates a PutV1ExternalAccountsIDConflict with default headers values
func NewPutV1ExternalAccountsIDConflict() *PutV1ExternalAccountsIDConflict {
	return &PutV1ExternalAccountsIDConflict{}
}

/*
PutV1ExternalAccountsIDConflict describes a response with status code 409, with default header values.

Conflict with the current state of the target resource
*/
type PutV1ExternalAccountsIDConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this put v1 external accounts Id conflict response has a 2xx status code
func (o *PutV1ExternalAccountsIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 external accounts Id conflict response has a 3xx status code
func (o *PutV1ExternalAccountsIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 external accounts Id conflict response has a 4xx status code
func (o *PutV1ExternalAccountsIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 external accounts Id conflict response has a 5xx status code
func (o *PutV1ExternalAccountsIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 external accounts Id conflict response a status code equal to that given
func (o *PutV1ExternalAccountsIDConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the put v1 external accounts Id conflict response
func (o *PutV1ExternalAccountsIDConflict) Code() int {
	return 409
}

func (o *PutV1ExternalAccountsIDConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/external-accounts/{id}][%d] putV1ExternalAccountsIdConflict  %+v", 409, o.Payload)
}

func (o *PutV1ExternalAccountsIDConflict) String() string {
	return fmt.Sprintf("[PUT /v1/external-accounts/{id}][%d] putV1ExternalAccountsIdConflict  %+v", 409, o.Payload)
}

func (o *PutV1ExternalAccountsIDConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutV1ExternalAccountsIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

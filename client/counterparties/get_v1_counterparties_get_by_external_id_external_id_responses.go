// Code generated by go-swagger; DO NOT EDIT.

package counterparties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/get-momo/atlar-v1-go-client/models"
)

// GetV1CounterpartiesGetByExternalIDExternalIDReader is a Reader for the GetV1CounterpartiesGetByExternalIDExternalID structure.
type GetV1CounterpartiesGetByExternalIDExternalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1CounterpartiesGetByExternalIDExternalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1CounterpartiesGetByExternalIDExternalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetV1CounterpartiesGetByExternalIDExternalIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/counterparties:getByExternalId/{externalId}] GetV1CounterpartiesGetByExternalIDExternalID", response, response.Code())
	}
}

// NewGetV1CounterpartiesGetByExternalIDExternalIDOK creates a GetV1CounterpartiesGetByExternalIDExternalIDOK with default headers values
func NewGetV1CounterpartiesGetByExternalIDExternalIDOK() *GetV1CounterpartiesGetByExternalIDExternalIDOK {
	return &GetV1CounterpartiesGetByExternalIDExternalIDOK{}
}

/*
GetV1CounterpartiesGetByExternalIDExternalIDOK describes a response with status code 200, with default header values.

The identified counterparty.
*/
type GetV1CounterpartiesGetByExternalIDExternalIDOK struct {
	Payload *models.Counterparty
}

// IsSuccess returns true when this get v1 counterparties get by external Id external Id o k response has a 2xx status code
func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 counterparties get by external Id external Id o k response has a 3xx status code
func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 counterparties get by external Id external Id o k response has a 4xx status code
func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 counterparties get by external Id external Id o k response has a 5xx status code
func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 counterparties get by external Id external Id o k response a status code equal to that given
func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 counterparties get by external Id external Id o k response
func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) Code() int {
	return 200
}

func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/counterparties:getByExternalId/{externalId}][%d] getV1CounterpartiesGetByExternalIdExternalIdOK  %+v", 200, o.Payload)
}

func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) String() string {
	return fmt.Sprintf("[GET /v1/counterparties:getByExternalId/{externalId}][%d] getV1CounterpartiesGetByExternalIdExternalIdOK  %+v", 200, o.Payload)
}

func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) GetPayload() *models.Counterparty {
	return o.Payload
}

func (o *GetV1CounterpartiesGetByExternalIDExternalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Counterparty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1CounterpartiesGetByExternalIDExternalIDNotFound creates a GetV1CounterpartiesGetByExternalIDExternalIDNotFound with default headers values
func NewGetV1CounterpartiesGetByExternalIDExternalIDNotFound() *GetV1CounterpartiesGetByExternalIDExternalIDNotFound {
	return &GetV1CounterpartiesGetByExternalIDExternalIDNotFound{}
}

/*
GetV1CounterpartiesGetByExternalIDExternalIDNotFound describes a response with status code 404, with default header values.

The identified counterparty doesn't exist.
*/
type GetV1CounterpartiesGetByExternalIDExternalIDNotFound struct {
	Payload string
}

// IsSuccess returns true when this get v1 counterparties get by external Id external Id not found response has a 2xx status code
func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 counterparties get by external Id external Id not found response has a 3xx status code
func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 counterparties get by external Id external Id not found response has a 4xx status code
func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 counterparties get by external Id external Id not found response has a 5xx status code
func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 counterparties get by external Id external Id not found response a status code equal to that given
func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 counterparties get by external Id external Id not found response
func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) Code() int {
	return 404
}

func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/counterparties:getByExternalId/{externalId}][%d] getV1CounterpartiesGetByExternalIdExternalIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/counterparties:getByExternalId/{externalId}][%d] getV1CounterpartiesGetByExternalIdExternalIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetV1CounterpartiesGetByExternalIDExternalIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

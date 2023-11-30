// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/get-momo/atlar-v1-go-client/models"
)

// PutV1WebhooksIDReader is a Reader for the PutV1WebhooksID structure.
type PutV1WebhooksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1WebhooksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1WebhooksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPutV1WebhooksIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/webhooks/{id}] PutV1WebhooksID", response, response.Code())
	}
}

// NewPutV1WebhooksIDOK creates a PutV1WebhooksIDOK with default headers values
func NewPutV1WebhooksIDOK() *PutV1WebhooksIDOK {
	return &PutV1WebhooksIDOK{}
}

/*
PutV1WebhooksIDOK describes a response with status code 200, with default header values.

The updated webhook
*/
type PutV1WebhooksIDOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this put v1 webhooks Id o k response has a 2xx status code
func (o *PutV1WebhooksIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 webhooks Id o k response has a 3xx status code
func (o *PutV1WebhooksIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 webhooks Id o k response has a 4xx status code
func (o *PutV1WebhooksIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 webhooks Id o k response has a 5xx status code
func (o *PutV1WebhooksIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 webhooks Id o k response a status code equal to that given
func (o *PutV1WebhooksIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put v1 webhooks Id o k response
func (o *PutV1WebhooksIDOK) Code() int {
	return 200
}

func (o *PutV1WebhooksIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/webhooks/{id}][%d] putV1WebhooksIdOK  %+v", 200, o.Payload)
}

func (o *PutV1WebhooksIDOK) String() string {
	return fmt.Sprintf("[PUT /v1/webhooks/{id}][%d] putV1WebhooksIdOK  %+v", 200, o.Payload)
}

func (o *PutV1WebhooksIDOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *PutV1WebhooksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1WebhooksIDNotFound creates a PutV1WebhooksIDNotFound with default headers values
func NewPutV1WebhooksIDNotFound() *PutV1WebhooksIDNotFound {
	return &PutV1WebhooksIDNotFound{}
}

/*
PutV1WebhooksIDNotFound describes a response with status code 404, with default header values.

The identified webhook doesn't exist
*/
type PutV1WebhooksIDNotFound struct {
	Payload string
}

// IsSuccess returns true when this put v1 webhooks Id not found response has a 2xx status code
func (o *PutV1WebhooksIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 webhooks Id not found response has a 3xx status code
func (o *PutV1WebhooksIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 webhooks Id not found response has a 4xx status code
func (o *PutV1WebhooksIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 webhooks Id not found response has a 5xx status code
func (o *PutV1WebhooksIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 webhooks Id not found response a status code equal to that given
func (o *PutV1WebhooksIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put v1 webhooks Id not found response
func (o *PutV1WebhooksIDNotFound) Code() int {
	return 404
}

func (o *PutV1WebhooksIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/webhooks/{id}][%d] putV1WebhooksIdNotFound  %+v", 404, o.Payload)
}

func (o *PutV1WebhooksIDNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/webhooks/{id}][%d] putV1WebhooksIdNotFound  %+v", 404, o.Payload)
}

func (o *PutV1WebhooksIDNotFound) GetPayload() string {
	return o.Payload
}

func (o *PutV1WebhooksIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

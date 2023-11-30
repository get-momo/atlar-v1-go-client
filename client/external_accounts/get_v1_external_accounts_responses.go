// Code generated by go-swagger; DO NOT EDIT.

package external_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/get-momo/atlar-v1-go-client/models"
)

// GetV1ExternalAccountsReader is a Reader for the GetV1ExternalAccounts structure.
type GetV1ExternalAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ExternalAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ExternalAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/external-accounts] GetV1ExternalAccounts", response, response.Code())
	}
}

// NewGetV1ExternalAccountsOK creates a GetV1ExternalAccountsOK with default headers values
func NewGetV1ExternalAccountsOK() *GetV1ExternalAccountsOK {
	return &GetV1ExternalAccountsOK{}
}

/*
GetV1ExternalAccountsOK describes a response with status code 200, with default header values.

QueryResponse with list of External Accounts
*/
type GetV1ExternalAccountsOK struct {
	Payload *GetV1ExternalAccountsOKBody
}

// IsSuccess returns true when this get v1 external accounts o k response has a 2xx status code
func (o *GetV1ExternalAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 external accounts o k response has a 3xx status code
func (o *GetV1ExternalAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 external accounts o k response has a 4xx status code
func (o *GetV1ExternalAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 external accounts o k response has a 5xx status code
func (o *GetV1ExternalAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 external accounts o k response a status code equal to that given
func (o *GetV1ExternalAccountsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 external accounts o k response
func (o *GetV1ExternalAccountsOK) Code() int {
	return 200
}

func (o *GetV1ExternalAccountsOK) Error() string {
	return fmt.Sprintf("[GET /v1/external-accounts][%d] getV1ExternalAccountsOK  %+v", 200, o.Payload)
}

func (o *GetV1ExternalAccountsOK) String() string {
	return fmt.Sprintf("[GET /v1/external-accounts][%d] getV1ExternalAccountsOK  %+v", 200, o.Payload)
}

func (o *GetV1ExternalAccountsOK) GetPayload() *GetV1ExternalAccountsOKBody {
	return o.Payload
}

func (o *GetV1ExternalAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1ExternalAccountsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1ExternalAccountsOKBody get v1 external accounts o k body
swagger:model GetV1ExternalAccountsOKBody
*/
type GetV1ExternalAccountsOKBody struct {
	models.QueryResponse

	// items
	Items []*models.ExternalAccount `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1ExternalAccountsOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1ExternalAccountsOKBodyAO0
	var getV1ExternalAccountsOKBodyAO0 models.QueryResponse
	if err := swag.ReadJSON(raw, &getV1ExternalAccountsOKBodyAO0); err != nil {
		return err
	}
	o.QueryResponse = getV1ExternalAccountsOKBodyAO0

	// GetV1ExternalAccountsOKBodyAO1
	var dataGetV1ExternalAccountsOKBodyAO1 struct {
		Items []*models.ExternalAccount `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1ExternalAccountsOKBodyAO1); err != nil {
		return err
	}

	o.Items = dataGetV1ExternalAccountsOKBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1ExternalAccountsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1ExternalAccountsOKBodyAO0, err := swag.WriteJSON(o.QueryResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1ExternalAccountsOKBodyAO0)
	var dataGetV1ExternalAccountsOKBodyAO1 struct {
		Items []*models.ExternalAccount `json:"items"`
	}

	dataGetV1ExternalAccountsOKBodyAO1.Items = o.Items

	jsonDataGetV1ExternalAccountsOKBodyAO1, errGetV1ExternalAccountsOKBodyAO1 := swag.WriteJSON(dataGetV1ExternalAccountsOKBodyAO1)
	if errGetV1ExternalAccountsOKBodyAO1 != nil {
		return nil, errGetV1ExternalAccountsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1ExternalAccountsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 external accounts o k body
func (o *GetV1ExternalAccountsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.QueryResponse
	if err := o.QueryResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV1ExternalAccountsOKBody) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1ExternalAccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1ExternalAccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1 external accounts o k body based on the context it is used
func (o *GetV1ExternalAccountsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.QueryResponse
	if err := o.QueryResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV1ExternalAccountsOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {

			if swag.IsZero(o.Items[i]) { // not required
				return nil
			}

			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1ExternalAccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1ExternalAccountsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1ExternalAccountsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1ExternalAccountsOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1ExternalAccountsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

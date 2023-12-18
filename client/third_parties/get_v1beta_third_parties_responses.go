// Code generated by go-swagger; DO NOT EDIT.

package third_parties

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

// GetV1betaThirdPartiesReader is a Reader for the GetV1betaThirdParties structure.
type GetV1betaThirdPartiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1betaThirdPartiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1betaThirdPartiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1beta/third-parties] GetV1betaThirdParties", response, response.Code())
	}
}

// NewGetV1betaThirdPartiesOK creates a GetV1betaThirdPartiesOK with default headers values
func NewGetV1betaThirdPartiesOK() *GetV1betaThirdPartiesOK {
	return &GetV1betaThirdPartiesOK{}
}

/*
GetV1betaThirdPartiesOK describes a response with status code 200, with default header values.

QueryResponse with list of ThirdParties
*/
type GetV1betaThirdPartiesOK struct {
	Payload *GetV1betaThirdPartiesOKBody
}

// IsSuccess returns true when this get v1beta third parties o k response has a 2xx status code
func (o *GetV1betaThirdPartiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1beta third parties o k response has a 3xx status code
func (o *GetV1betaThirdPartiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1beta third parties o k response has a 4xx status code
func (o *GetV1betaThirdPartiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1beta third parties o k response has a 5xx status code
func (o *GetV1betaThirdPartiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1beta third parties o k response a status code equal to that given
func (o *GetV1betaThirdPartiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1beta third parties o k response
func (o *GetV1betaThirdPartiesOK) Code() int {
	return 200
}

func (o *GetV1betaThirdPartiesOK) Error() string {
	return fmt.Sprintf("[GET /v1beta/third-parties][%d] getV1betaThirdPartiesOK  %+v", 200, o.Payload)
}

func (o *GetV1betaThirdPartiesOK) String() string {
	return fmt.Sprintf("[GET /v1beta/third-parties][%d] getV1betaThirdPartiesOK  %+v", 200, o.Payload)
}

func (o *GetV1betaThirdPartiesOK) GetPayload() *GetV1betaThirdPartiesOKBody {
	return o.Payload
}

func (o *GetV1betaThirdPartiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1betaThirdPartiesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1betaThirdPartiesOKBody get v1beta third parties o k body
swagger:model GetV1betaThirdPartiesOKBody
*/
type GetV1betaThirdPartiesOKBody struct {
	models.QueryResponse

	// items
	Items []*models.ThirdParty `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1betaThirdPartiesOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1betaThirdPartiesOKBodyAO0
	var getV1betaThirdPartiesOKBodyAO0 models.QueryResponse
	if err := swag.ReadJSON(raw, &getV1betaThirdPartiesOKBodyAO0); err != nil {
		return err
	}
	o.QueryResponse = getV1betaThirdPartiesOKBodyAO0

	// GetV1betaThirdPartiesOKBodyAO1
	var dataGetV1betaThirdPartiesOKBodyAO1 struct {
		Items []*models.ThirdParty `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1betaThirdPartiesOKBodyAO1); err != nil {
		return err
	}

	o.Items = dataGetV1betaThirdPartiesOKBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1betaThirdPartiesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1betaThirdPartiesOKBodyAO0, err := swag.WriteJSON(o.QueryResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1betaThirdPartiesOKBodyAO0)
	var dataGetV1betaThirdPartiesOKBodyAO1 struct {
		Items []*models.ThirdParty `json:"items"`
	}

	dataGetV1betaThirdPartiesOKBodyAO1.Items = o.Items

	jsonDataGetV1betaThirdPartiesOKBodyAO1, errGetV1betaThirdPartiesOKBodyAO1 := swag.WriteJSON(dataGetV1betaThirdPartiesOKBodyAO1)
	if errGetV1betaThirdPartiesOKBodyAO1 != nil {
		return nil, errGetV1betaThirdPartiesOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1betaThirdPartiesOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1beta third parties o k body
func (o *GetV1betaThirdPartiesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetV1betaThirdPartiesOKBody) validateItems(formats strfmt.Registry) error {

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
					return ve.ValidateName("getV1betaThirdPartiesOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1betaThirdPartiesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1beta third parties o k body based on the context it is used
func (o *GetV1betaThirdPartiesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetV1betaThirdPartiesOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {

			if swag.IsZero(o.Items[i]) { // not required
				return nil
			}

			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1betaThirdPartiesOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1betaThirdPartiesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1betaThirdPartiesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1betaThirdPartiesOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1betaThirdPartiesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
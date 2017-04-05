package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JAccountAcceptInvitationReader is a Reader for the JAccountAcceptInvitation structure.
type JAccountAcceptInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JAccountAcceptInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJAccountAcceptInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJAccountAcceptInvitationOK creates a JAccountAcceptInvitationOK with default headers values
func NewJAccountAcceptInvitationOK() *JAccountAcceptInvitationOK {
	return &JAccountAcceptInvitationOK{}
}

/*JAccountAcceptInvitationOK handles this case with default header values.

OK
*/
type JAccountAcceptInvitationOK struct {
	Payload JAccountAcceptInvitationOKBody
}

func (o *JAccountAcceptInvitationOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.acceptInvitation/{id}][%d] jAccountAcceptInvitationOK  %+v", 200, o.Payload)
}

func (o *JAccountAcceptInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JAccountAcceptInvitationOKBody j account accept invitation o k body
swagger:model JAccountAcceptInvitationOKBody
*/
type JAccountAcceptInvitationOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JAccountAcceptInvitationOKBody) UnmarshalJSON(raw []byte) error {

	var jAccountAcceptInvitationOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &jAccountAcceptInvitationOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = jAccountAcceptInvitationOKBodyAO0

	var jAccountAcceptInvitationOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jAccountAcceptInvitationOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jAccountAcceptInvitationOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JAccountAcceptInvitationOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jAccountAcceptInvitationOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountAcceptInvitationOKBodyAO0)

	jAccountAcceptInvitationOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountAcceptInvitationOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j account accept invitation o k body
func (o *JAccountAcceptInvitationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

package j_stack_template

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

// JStackTemplateDeleteReader is a Reader for the JStackTemplateDelete structure.
type JStackTemplateDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JStackTemplateDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJStackTemplateDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJStackTemplateDeleteOK creates a JStackTemplateDeleteOK with default headers values
func NewJStackTemplateDeleteOK() *JStackTemplateDeleteOK {
	return &JStackTemplateDeleteOK{}
}

/*JStackTemplateDeleteOK handles this case with default header values.

OK
*/
type JStackTemplateDeleteOK struct {
	Payload JStackTemplateDeleteOKBody
}

func (o *JStackTemplateDeleteOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.delete/{id}][%d] jStackTemplateDeleteOK  %+v", 200, o.Payload)
}

func (o *JStackTemplateDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JStackTemplateDeleteOKBody j stack template delete o k body
swagger:model JStackTemplateDeleteOKBody
*/
type JStackTemplateDeleteOKBody struct {
	models.JStackTemplate

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JStackTemplateDeleteOKBody) UnmarshalJSON(raw []byte) error {

	var jStackTemplateDeleteOKBodyAO0 models.JStackTemplate
	if err := swag.ReadJSON(raw, &jStackTemplateDeleteOKBodyAO0); err != nil {
		return err
	}
	o.JStackTemplate = jStackTemplateDeleteOKBodyAO0

	var jStackTemplateDeleteOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jStackTemplateDeleteOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jStackTemplateDeleteOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JStackTemplateDeleteOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jStackTemplateDeleteOKBodyAO0, err := swag.WriteJSON(o.JStackTemplate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jStackTemplateDeleteOKBodyAO0)

	jStackTemplateDeleteOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jStackTemplateDeleteOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j stack template delete o k body
func (o *JStackTemplateDeleteOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JStackTemplate.Validate(formats); err != nil {
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

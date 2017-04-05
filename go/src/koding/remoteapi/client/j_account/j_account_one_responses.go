package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JAccountOneReader is a Reader for the JAccountOne structure.
type JAccountOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JAccountOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJAccountOneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJAccountOneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJAccountOneOK creates a JAccountOneOK with default headers values
func NewJAccountOneOK() *JAccountOneOK {
	return &JAccountOneOK{}
}

/*JAccountOneOK handles this case with default header values.

Request processed successfully
*/
type JAccountOneOK struct {
	Payload *models.DefaultResponse
}

func (o *JAccountOneOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.one][%d] jAccountOneOK  %+v", 200, o.Payload)
}

func (o *JAccountOneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJAccountOneUnauthorized creates a JAccountOneUnauthorized with default headers values
func NewJAccountOneUnauthorized() *JAccountOneUnauthorized {
	return &JAccountOneUnauthorized{}
}

/*JAccountOneUnauthorized handles this case with default header values.

Unauthorized request
*/
type JAccountOneUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JAccountOneUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.one][%d] jAccountOneUnauthorized  %+v", 401, o.Payload)
}

func (o *JAccountOneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

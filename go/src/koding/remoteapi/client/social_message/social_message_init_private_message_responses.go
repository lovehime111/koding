package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// SocialMessageInitPrivateMessageReader is a Reader for the SocialMessageInitPrivateMessage structure.
type SocialMessageInitPrivateMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SocialMessageInitPrivateMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSocialMessageInitPrivateMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSocialMessageInitPrivateMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSocialMessageInitPrivateMessageOK creates a SocialMessageInitPrivateMessageOK with default headers values
func NewSocialMessageInitPrivateMessageOK() *SocialMessageInitPrivateMessageOK {
	return &SocialMessageInitPrivateMessageOK{}
}

/*SocialMessageInitPrivateMessageOK handles this case with default header values.

Request processed successfully
*/
type SocialMessageInitPrivateMessageOK struct {
	Payload *models.DefaultResponse
}

func (o *SocialMessageInitPrivateMessageOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.initPrivateMessage][%d] socialMessageInitPrivateMessageOK  %+v", 200, o.Payload)
}

func (o *SocialMessageInitPrivateMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSocialMessageInitPrivateMessageUnauthorized creates a SocialMessageInitPrivateMessageUnauthorized with default headers values
func NewSocialMessageInitPrivateMessageUnauthorized() *SocialMessageInitPrivateMessageUnauthorized {
	return &SocialMessageInitPrivateMessageUnauthorized{}
}

/*SocialMessageInitPrivateMessageUnauthorized handles this case with default header values.

Unauthorized request
*/
type SocialMessageInitPrivateMessageUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SocialMessageInitPrivateMessageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.initPrivateMessage][%d] socialMessageInitPrivateMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *SocialMessageInitPrivateMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

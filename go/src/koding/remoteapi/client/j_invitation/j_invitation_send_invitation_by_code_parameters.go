package j_invitation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewJInvitationSendInvitationByCodeParams creates a new JInvitationSendInvitationByCodeParams object
// with the default values initialized.
func NewJInvitationSendInvitationByCodeParams() *JInvitationSendInvitationByCodeParams {
	var ()
	return &JInvitationSendInvitationByCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJInvitationSendInvitationByCodeParamsWithTimeout creates a new JInvitationSendInvitationByCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJInvitationSendInvitationByCodeParamsWithTimeout(timeout time.Duration) *JInvitationSendInvitationByCodeParams {
	var ()
	return &JInvitationSendInvitationByCodeParams{

		timeout: timeout,
	}
}

// NewJInvitationSendInvitationByCodeParamsWithContext creates a new JInvitationSendInvitationByCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewJInvitationSendInvitationByCodeParamsWithContext(ctx context.Context) *JInvitationSendInvitationByCodeParams {
	var ()
	return &JInvitationSendInvitationByCodeParams{

		Context: ctx,
	}
}

/*JInvitationSendInvitationByCodeParams contains all the parameters to send to the API endpoint
for the j invitation send invitation by code operation typically these are written to a http.Request
*/
type JInvitationSendInvitationByCodeParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j invitation send invitation by code params
func (o *JInvitationSendInvitationByCodeParams) WithTimeout(timeout time.Duration) *JInvitationSendInvitationByCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j invitation send invitation by code params
func (o *JInvitationSendInvitationByCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j invitation send invitation by code params
func (o *JInvitationSendInvitationByCodeParams) WithContext(ctx context.Context) *JInvitationSendInvitationByCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j invitation send invitation by code params
func (o *JInvitationSendInvitationByCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j invitation send invitation by code params
func (o *JInvitationSendInvitationByCodeParams) WithBody(body models.DefaultSelector) *JInvitationSendInvitationByCodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j invitation send invitation by code params
func (o *JInvitationSendInvitationByCodeParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JInvitationSendInvitationByCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

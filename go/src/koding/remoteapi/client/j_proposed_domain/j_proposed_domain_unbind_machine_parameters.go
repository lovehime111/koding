package j_proposed_domain

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

// NewJProposedDomainUnbindMachineParams creates a new JProposedDomainUnbindMachineParams object
// with the default values initialized.
func NewJProposedDomainUnbindMachineParams() *JProposedDomainUnbindMachineParams {
	var ()
	return &JProposedDomainUnbindMachineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJProposedDomainUnbindMachineParamsWithTimeout creates a new JProposedDomainUnbindMachineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJProposedDomainUnbindMachineParamsWithTimeout(timeout time.Duration) *JProposedDomainUnbindMachineParams {
	var ()
	return &JProposedDomainUnbindMachineParams{

		timeout: timeout,
	}
}

// NewJProposedDomainUnbindMachineParamsWithContext creates a new JProposedDomainUnbindMachineParams object
// with the default values initialized, and the ability to set a context for a request
func NewJProposedDomainUnbindMachineParamsWithContext(ctx context.Context) *JProposedDomainUnbindMachineParams {
	var ()
	return &JProposedDomainUnbindMachineParams{

		Context: ctx,
	}
}

/*JProposedDomainUnbindMachineParams contains all the parameters to send to the API endpoint
for the j proposed domain unbind machine operation typically these are written to a http.Request
*/
type JProposedDomainUnbindMachineParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j proposed domain unbind machine params
func (o *JProposedDomainUnbindMachineParams) WithTimeout(timeout time.Duration) *JProposedDomainUnbindMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j proposed domain unbind machine params
func (o *JProposedDomainUnbindMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j proposed domain unbind machine params
func (o *JProposedDomainUnbindMachineParams) WithContext(ctx context.Context) *JProposedDomainUnbindMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j proposed domain unbind machine params
func (o *JProposedDomainUnbindMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j proposed domain unbind machine params
func (o *JProposedDomainUnbindMachineParams) WithBody(body models.DefaultSelector) *JProposedDomainUnbindMachineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j proposed domain unbind machine params
func (o *JProposedDomainUnbindMachineParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j proposed domain unbind machine params
func (o *JProposedDomainUnbindMachineParams) WithID(id string) *JProposedDomainUnbindMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j proposed domain unbind machine params
func (o *JProposedDomainUnbindMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JProposedDomainUnbindMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

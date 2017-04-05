package j_group

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

// NewJGroupFetchDataAtParams creates a new JGroupFetchDataAtParams object
// with the default values initialized.
func NewJGroupFetchDataAtParams() *JGroupFetchDataAtParams {
	var ()
	return &JGroupFetchDataAtParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupFetchDataAtParamsWithTimeout creates a new JGroupFetchDataAtParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupFetchDataAtParamsWithTimeout(timeout time.Duration) *JGroupFetchDataAtParams {
	var ()
	return &JGroupFetchDataAtParams{

		timeout: timeout,
	}
}

// NewJGroupFetchDataAtParamsWithContext creates a new JGroupFetchDataAtParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupFetchDataAtParamsWithContext(ctx context.Context) *JGroupFetchDataAtParams {
	var ()
	return &JGroupFetchDataAtParams{

		Context: ctx,
	}
}

/*JGroupFetchDataAtParams contains all the parameters to send to the API endpoint
for the j group fetch data at operation typically these are written to a http.Request
*/
type JGroupFetchDataAtParams struct {

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

// WithTimeout adds the timeout to the j group fetch data at params
func (o *JGroupFetchDataAtParams) WithTimeout(timeout time.Duration) *JGroupFetchDataAtParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group fetch data at params
func (o *JGroupFetchDataAtParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group fetch data at params
func (o *JGroupFetchDataAtParams) WithContext(ctx context.Context) *JGroupFetchDataAtParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group fetch data at params
func (o *JGroupFetchDataAtParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group fetch data at params
func (o *JGroupFetchDataAtParams) WithBody(body models.DefaultSelector) *JGroupFetchDataAtParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group fetch data at params
func (o *JGroupFetchDataAtParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group fetch data at params
func (o *JGroupFetchDataAtParams) WithID(id string) *JGroupFetchDataAtParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group fetch data at params
func (o *JGroupFetchDataAtParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupFetchDataAtParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

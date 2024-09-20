// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/arcfraofficial/aoc-go-sdk/models"
)

// NewUpdateVMIoPolicyParams creates a new UpdateVMIoPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMIoPolicyParams() *UpdateVMIoPolicyParams {
	return &UpdateVMIoPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMIoPolicyParamsWithTimeout creates a new UpdateVMIoPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateVMIoPolicyParamsWithTimeout(timeout time.Duration) *UpdateVMIoPolicyParams {
	return &UpdateVMIoPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateVMIoPolicyParamsWithContext creates a new UpdateVMIoPolicyParams object
// with the ability to set a context for a request.
func NewUpdateVMIoPolicyParamsWithContext(ctx context.Context) *UpdateVMIoPolicyParams {
	return &UpdateVMIoPolicyParams{
		Context: ctx,
	}
}

// NewUpdateVMIoPolicyParamsWithHTTPClient creates a new UpdateVMIoPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMIoPolicyParamsWithHTTPClient(client *http.Client) *UpdateVMIoPolicyParams {
	return &UpdateVMIoPolicyParams{
		HTTPClient: client,
	}
}

/* UpdateVMIoPolicyParams contains all the parameters to send to the API endpoint
   for the update Vm io policy operation.

   Typically these are written to a http.Request.
*/
type UpdateVMIoPolicyParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMUpdateIoPolicyParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm io policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMIoPolicyParams) WithDefaults() *UpdateVMIoPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm io policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMIoPolicyParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVMIoPolicyParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) WithTimeout(timeout time.Duration) *UpdateVMIoPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) WithContext(ctx context.Context) *UpdateVMIoPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) WithHTTPClient(client *http.Client) *UpdateVMIoPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) WithContentLanguage(contentLanguage *string) *UpdateVMIoPolicyParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) WithRequestBody(requestBody *models.VMUpdateIoPolicyParams) *UpdateVMIoPolicyParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm io policy params
func (o *UpdateVMIoPolicyParams) SetRequestBody(requestBody *models.VMUpdateIoPolicyParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMIoPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentLanguage != nil {

		// header param content-language
		if err := r.SetHeaderParam("content-language", *o.ContentLanguage); err != nil {
			return err
		}
	}
	if o.RequestBody != nil {
		if err := r.SetBodyParam(o.RequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

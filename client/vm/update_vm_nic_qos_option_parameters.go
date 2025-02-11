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

// NewUpdateVMNicQosOptionParams creates a new UpdateVMNicQosOptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMNicQosOptionParams() *UpdateVMNicQosOptionParams {
	return &UpdateVMNicQosOptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMNicQosOptionParamsWithTimeout creates a new UpdateVMNicQosOptionParams object
// with the ability to set a timeout on a request.
func NewUpdateVMNicQosOptionParamsWithTimeout(timeout time.Duration) *UpdateVMNicQosOptionParams {
	return &UpdateVMNicQosOptionParams{
		timeout: timeout,
	}
}

// NewUpdateVMNicQosOptionParamsWithContext creates a new UpdateVMNicQosOptionParams object
// with the ability to set a context for a request.
func NewUpdateVMNicQosOptionParamsWithContext(ctx context.Context) *UpdateVMNicQosOptionParams {
	return &UpdateVMNicQosOptionParams{
		Context: ctx,
	}
}

// NewUpdateVMNicQosOptionParamsWithHTTPClient creates a new UpdateVMNicQosOptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMNicQosOptionParamsWithHTTPClient(client *http.Client) *UpdateVMNicQosOptionParams {
	return &UpdateVMNicQosOptionParams{
		HTTPClient: client,
	}
}

/* UpdateVMNicQosOptionParams contains all the parameters to send to the API endpoint
   for the update Vm nic qos option operation.

   Typically these are written to a http.Request.
*/
type UpdateVMNicQosOptionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMUpdateNicQosOptionsParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm nic qos option params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMNicQosOptionParams) WithDefaults() *UpdateVMNicQosOptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm nic qos option params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMNicQosOptionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVMNicQosOptionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) WithTimeout(timeout time.Duration) *UpdateVMNicQosOptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) WithContext(ctx context.Context) *UpdateVMNicQosOptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) WithHTTPClient(client *http.Client) *UpdateVMNicQosOptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) WithContentLanguage(contentLanguage *string) *UpdateVMNicQosOptionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) WithRequestBody(requestBody *models.VMUpdateNicQosOptionsParams) *UpdateVMNicQosOptionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm nic qos option params
func (o *UpdateVMNicQosOptionParams) SetRequestBody(requestBody *models.VMUpdateNicQosOptionsParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMNicQosOptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

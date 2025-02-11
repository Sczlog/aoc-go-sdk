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

// NewAddVMGpuDeviceParams creates a new AddVMGpuDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddVMGpuDeviceParams() *AddVMGpuDeviceParams {
	return &AddVMGpuDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddVMGpuDeviceParamsWithTimeout creates a new AddVMGpuDeviceParams object
// with the ability to set a timeout on a request.
func NewAddVMGpuDeviceParamsWithTimeout(timeout time.Duration) *AddVMGpuDeviceParams {
	return &AddVMGpuDeviceParams{
		timeout: timeout,
	}
}

// NewAddVMGpuDeviceParamsWithContext creates a new AddVMGpuDeviceParams object
// with the ability to set a context for a request.
func NewAddVMGpuDeviceParamsWithContext(ctx context.Context) *AddVMGpuDeviceParams {
	return &AddVMGpuDeviceParams{
		Context: ctx,
	}
}

// NewAddVMGpuDeviceParamsWithHTTPClient creates a new AddVMGpuDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddVMGpuDeviceParamsWithHTTPClient(client *http.Client) *AddVMGpuDeviceParams {
	return &AddVMGpuDeviceParams{
		HTTPClient: client,
	}
}

/* AddVMGpuDeviceParams contains all the parameters to send to the API endpoint
   for the add Vm gpu device operation.

   Typically these are written to a http.Request.
*/
type AddVMGpuDeviceParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMAddGpuDeviceParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add Vm gpu device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVMGpuDeviceParams) WithDefaults() *AddVMGpuDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add Vm gpu device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVMGpuDeviceParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := AddVMGpuDeviceParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) WithTimeout(timeout time.Duration) *AddVMGpuDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) WithContext(ctx context.Context) *AddVMGpuDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) WithHTTPClient(client *http.Client) *AddVMGpuDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) WithContentLanguage(contentLanguage *string) *AddVMGpuDeviceParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) WithRequestBody(requestBody *models.VMAddGpuDeviceParams) *AddVMGpuDeviceParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the add Vm gpu device params
func (o *AddVMGpuDeviceParams) SetRequestBody(requestBody *models.VMAddGpuDeviceParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *AddVMGpuDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

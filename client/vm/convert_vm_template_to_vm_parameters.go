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

// NewConvertVMTemplateToVMParams creates a new ConvertVMTemplateToVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConvertVMTemplateToVMParams() *ConvertVMTemplateToVMParams {
	return &ConvertVMTemplateToVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConvertVMTemplateToVMParamsWithTimeout creates a new ConvertVMTemplateToVMParams object
// with the ability to set a timeout on a request.
func NewConvertVMTemplateToVMParamsWithTimeout(timeout time.Duration) *ConvertVMTemplateToVMParams {
	return &ConvertVMTemplateToVMParams{
		timeout: timeout,
	}
}

// NewConvertVMTemplateToVMParamsWithContext creates a new ConvertVMTemplateToVMParams object
// with the ability to set a context for a request.
func NewConvertVMTemplateToVMParamsWithContext(ctx context.Context) *ConvertVMTemplateToVMParams {
	return &ConvertVMTemplateToVMParams{
		Context: ctx,
	}
}

// NewConvertVMTemplateToVMParamsWithHTTPClient creates a new ConvertVMTemplateToVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewConvertVMTemplateToVMParamsWithHTTPClient(client *http.Client) *ConvertVMTemplateToVMParams {
	return &ConvertVMTemplateToVMParams{
		HTTPClient: client,
	}
}

/* ConvertVMTemplateToVMParams contains all the parameters to send to the API endpoint
   for the convert Vm template to Vm operation.

   Typically these are written to a http.Request.
*/
type ConvertVMTemplateToVMParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.ConvertVMTemplateToVMParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the convert Vm template to Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConvertVMTemplateToVMParams) WithDefaults() *ConvertVMTemplateToVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the convert Vm template to Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConvertVMTemplateToVMParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := ConvertVMTemplateToVMParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) WithTimeout(timeout time.Duration) *ConvertVMTemplateToVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) WithContext(ctx context.Context) *ConvertVMTemplateToVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) WithHTTPClient(client *http.Client) *ConvertVMTemplateToVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) WithContentLanguage(contentLanguage *string) *ConvertVMTemplateToVMParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) WithRequestBody(requestBody []*models.ConvertVMTemplateToVMParams) *ConvertVMTemplateToVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the convert Vm template to Vm params
func (o *ConvertVMTemplateToVMParams) SetRequestBody(requestBody []*models.ConvertVMTemplateToVMParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *ConvertVMTemplateToVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

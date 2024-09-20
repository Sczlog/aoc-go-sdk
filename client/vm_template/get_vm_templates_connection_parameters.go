// Code generated by go-swagger; DO NOT EDIT.

package vm_template

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

// NewGetVMTemplatesConnectionParams creates a new GetVMTemplatesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMTemplatesConnectionParams() *GetVMTemplatesConnectionParams {
	return &GetVMTemplatesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMTemplatesConnectionParamsWithTimeout creates a new GetVMTemplatesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVMTemplatesConnectionParamsWithTimeout(timeout time.Duration) *GetVMTemplatesConnectionParams {
	return &GetVMTemplatesConnectionParams{
		timeout: timeout,
	}
}

// NewGetVMTemplatesConnectionParamsWithContext creates a new GetVMTemplatesConnectionParams object
// with the ability to set a context for a request.
func NewGetVMTemplatesConnectionParamsWithContext(ctx context.Context) *GetVMTemplatesConnectionParams {
	return &GetVMTemplatesConnectionParams{
		Context: ctx,
	}
}

// NewGetVMTemplatesConnectionParamsWithHTTPClient creates a new GetVMTemplatesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMTemplatesConnectionParamsWithHTTPClient(client *http.Client) *GetVMTemplatesConnectionParams {
	return &GetVMTemplatesConnectionParams{
		HTTPClient: client,
	}
}

/* GetVMTemplatesConnectionParams contains all the parameters to send to the API endpoint
   for the get Vm templates connection operation.

   Typically these are written to a http.Request.
*/
type GetVMTemplatesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMTemplatesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm templates connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMTemplatesConnectionParams) WithDefaults() *GetVMTemplatesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm templates connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMTemplatesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMTemplatesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) WithTimeout(timeout time.Duration) *GetVMTemplatesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) WithContext(ctx context.Context) *GetVMTemplatesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) WithHTTPClient(client *http.Client) *GetVMTemplatesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) WithContentLanguage(contentLanguage *string) *GetVMTemplatesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) WithRequestBody(requestBody *models.GetVMTemplatesConnectionRequestBody) *GetVMTemplatesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm templates connection params
func (o *GetVMTemplatesConnectionParams) SetRequestBody(requestBody *models.GetVMTemplatesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMTemplatesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

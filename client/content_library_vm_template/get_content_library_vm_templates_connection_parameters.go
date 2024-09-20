// Code generated by go-swagger; DO NOT EDIT.

package content_library_vm_template

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

// NewGetContentLibraryVMTemplatesConnectionParams creates a new GetContentLibraryVMTemplatesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContentLibraryVMTemplatesConnectionParams() *GetContentLibraryVMTemplatesConnectionParams {
	return &GetContentLibraryVMTemplatesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentLibraryVMTemplatesConnectionParamsWithTimeout creates a new GetContentLibraryVMTemplatesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetContentLibraryVMTemplatesConnectionParamsWithTimeout(timeout time.Duration) *GetContentLibraryVMTemplatesConnectionParams {
	return &GetContentLibraryVMTemplatesConnectionParams{
		timeout: timeout,
	}
}

// NewGetContentLibraryVMTemplatesConnectionParamsWithContext creates a new GetContentLibraryVMTemplatesConnectionParams object
// with the ability to set a context for a request.
func NewGetContentLibraryVMTemplatesConnectionParamsWithContext(ctx context.Context) *GetContentLibraryVMTemplatesConnectionParams {
	return &GetContentLibraryVMTemplatesConnectionParams{
		Context: ctx,
	}
}

// NewGetContentLibraryVMTemplatesConnectionParamsWithHTTPClient creates a new GetContentLibraryVMTemplatesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContentLibraryVMTemplatesConnectionParamsWithHTTPClient(client *http.Client) *GetContentLibraryVMTemplatesConnectionParams {
	return &GetContentLibraryVMTemplatesConnectionParams{
		HTTPClient: client,
	}
}

/* GetContentLibraryVMTemplatesConnectionParams contains all the parameters to send to the API endpoint
   for the get content library Vm templates connection operation.

   Typically these are written to a http.Request.
*/
type GetContentLibraryVMTemplatesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetContentLibraryVMTemplatesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get content library Vm templates connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentLibraryVMTemplatesConnectionParams) WithDefaults() *GetContentLibraryVMTemplatesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get content library Vm templates connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentLibraryVMTemplatesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetContentLibraryVMTemplatesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) WithTimeout(timeout time.Duration) *GetContentLibraryVMTemplatesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) WithContext(ctx context.Context) *GetContentLibraryVMTemplatesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) WithHTTPClient(client *http.Client) *GetContentLibraryVMTemplatesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) WithContentLanguage(contentLanguage *string) *GetContentLibraryVMTemplatesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) WithRequestBody(requestBody *models.GetContentLibraryVMTemplatesConnectionRequestBody) *GetContentLibraryVMTemplatesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get content library Vm templates connection params
func (o *GetContentLibraryVMTemplatesConnectionParams) SetRequestBody(requestBody *models.GetContentLibraryVMTemplatesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentLibraryVMTemplatesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

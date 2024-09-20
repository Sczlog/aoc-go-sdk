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

// NewUpdateContentLibraryVMTemplateParams creates a new UpdateContentLibraryVMTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateContentLibraryVMTemplateParams() *UpdateContentLibraryVMTemplateParams {
	return &UpdateContentLibraryVMTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateContentLibraryVMTemplateParamsWithTimeout creates a new UpdateContentLibraryVMTemplateParams object
// with the ability to set a timeout on a request.
func NewUpdateContentLibraryVMTemplateParamsWithTimeout(timeout time.Duration) *UpdateContentLibraryVMTemplateParams {
	return &UpdateContentLibraryVMTemplateParams{
		timeout: timeout,
	}
}

// NewUpdateContentLibraryVMTemplateParamsWithContext creates a new UpdateContentLibraryVMTemplateParams object
// with the ability to set a context for a request.
func NewUpdateContentLibraryVMTemplateParamsWithContext(ctx context.Context) *UpdateContentLibraryVMTemplateParams {
	return &UpdateContentLibraryVMTemplateParams{
		Context: ctx,
	}
}

// NewUpdateContentLibraryVMTemplateParamsWithHTTPClient creates a new UpdateContentLibraryVMTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateContentLibraryVMTemplateParamsWithHTTPClient(client *http.Client) *UpdateContentLibraryVMTemplateParams {
	return &UpdateContentLibraryVMTemplateParams{
		HTTPClient: client,
	}
}

/* UpdateContentLibraryVMTemplateParams contains all the parameters to send to the API endpoint
   for the update content library Vm template operation.

   Typically these are written to a http.Request.
*/
type UpdateContentLibraryVMTemplateParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ContentLibraryVMTemplateUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update content library Vm template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateContentLibraryVMTemplateParams) WithDefaults() *UpdateContentLibraryVMTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update content library Vm template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateContentLibraryVMTemplateParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateContentLibraryVMTemplateParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) WithTimeout(timeout time.Duration) *UpdateContentLibraryVMTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) WithContext(ctx context.Context) *UpdateContentLibraryVMTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) WithHTTPClient(client *http.Client) *UpdateContentLibraryVMTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) WithContentLanguage(contentLanguage *string) *UpdateContentLibraryVMTemplateParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) WithRequestBody(requestBody *models.ContentLibraryVMTemplateUpdationParams) *UpdateContentLibraryVMTemplateParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update content library Vm template params
func (o *UpdateContentLibraryVMTemplateParams) SetRequestBody(requestBody *models.ContentLibraryVMTemplateUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateContentLibraryVMTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

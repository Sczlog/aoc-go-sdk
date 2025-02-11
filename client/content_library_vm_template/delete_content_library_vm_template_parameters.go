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

// NewDeleteContentLibraryVMTemplateParams creates a new DeleteContentLibraryVMTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteContentLibraryVMTemplateParams() *DeleteContentLibraryVMTemplateParams {
	return &DeleteContentLibraryVMTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContentLibraryVMTemplateParamsWithTimeout creates a new DeleteContentLibraryVMTemplateParams object
// with the ability to set a timeout on a request.
func NewDeleteContentLibraryVMTemplateParamsWithTimeout(timeout time.Duration) *DeleteContentLibraryVMTemplateParams {
	return &DeleteContentLibraryVMTemplateParams{
		timeout: timeout,
	}
}

// NewDeleteContentLibraryVMTemplateParamsWithContext creates a new DeleteContentLibraryVMTemplateParams object
// with the ability to set a context for a request.
func NewDeleteContentLibraryVMTemplateParamsWithContext(ctx context.Context) *DeleteContentLibraryVMTemplateParams {
	return &DeleteContentLibraryVMTemplateParams{
		Context: ctx,
	}
}

// NewDeleteContentLibraryVMTemplateParamsWithHTTPClient creates a new DeleteContentLibraryVMTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteContentLibraryVMTemplateParamsWithHTTPClient(client *http.Client) *DeleteContentLibraryVMTemplateParams {
	return &DeleteContentLibraryVMTemplateParams{
		HTTPClient: client,
	}
}

/* DeleteContentLibraryVMTemplateParams contains all the parameters to send to the API endpoint
   for the delete content library Vm template operation.

   Typically these are written to a http.Request.
*/
type DeleteContentLibraryVMTemplateParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ContentLibraryVMTemplateDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete content library Vm template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContentLibraryVMTemplateParams) WithDefaults() *DeleteContentLibraryVMTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete content library Vm template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContentLibraryVMTemplateParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteContentLibraryVMTemplateParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) WithTimeout(timeout time.Duration) *DeleteContentLibraryVMTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) WithContext(ctx context.Context) *DeleteContentLibraryVMTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) WithHTTPClient(client *http.Client) *DeleteContentLibraryVMTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) WithContentLanguage(contentLanguage *string) *DeleteContentLibraryVMTemplateParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) WithRequestBody(requestBody *models.ContentLibraryVMTemplateDeletionParams) *DeleteContentLibraryVMTemplateParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete content library Vm template params
func (o *DeleteContentLibraryVMTemplateParams) SetRequestBody(requestBody *models.ContentLibraryVMTemplateDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContentLibraryVMTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

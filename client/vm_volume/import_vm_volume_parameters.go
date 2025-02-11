// Code generated by go-swagger; DO NOT EDIT.

package vm_volume

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

// NewImportVMVolumeParams creates a new ImportVMVolumeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportVMVolumeParams() *ImportVMVolumeParams {
	return &ImportVMVolumeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportVMVolumeParamsWithTimeout creates a new ImportVMVolumeParams object
// with the ability to set a timeout on a request.
func NewImportVMVolumeParamsWithTimeout(timeout time.Duration) *ImportVMVolumeParams {
	return &ImportVMVolumeParams{
		timeout: timeout,
	}
}

// NewImportVMVolumeParamsWithContext creates a new ImportVMVolumeParams object
// with the ability to set a context for a request.
func NewImportVMVolumeParamsWithContext(ctx context.Context) *ImportVMVolumeParams {
	return &ImportVMVolumeParams{
		Context: ctx,
	}
}

// NewImportVMVolumeParamsWithHTTPClient creates a new ImportVMVolumeParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportVMVolumeParamsWithHTTPClient(client *http.Client) *ImportVMVolumeParams {
	return &ImportVMVolumeParams{
		HTTPClient: client,
	}
}

/* ImportVMVolumeParams contains all the parameters to send to the API endpoint
   for the import Vm volume operation.

   Typically these are written to a http.Request.
*/
type ImportVMVolumeParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.ImportVMVolumeParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import Vm volume params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportVMVolumeParams) WithDefaults() *ImportVMVolumeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import Vm volume params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportVMVolumeParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := ImportVMVolumeParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the import Vm volume params
func (o *ImportVMVolumeParams) WithTimeout(timeout time.Duration) *ImportVMVolumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import Vm volume params
func (o *ImportVMVolumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import Vm volume params
func (o *ImportVMVolumeParams) WithContext(ctx context.Context) *ImportVMVolumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import Vm volume params
func (o *ImportVMVolumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import Vm volume params
func (o *ImportVMVolumeParams) WithHTTPClient(client *http.Client) *ImportVMVolumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import Vm volume params
func (o *ImportVMVolumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the import Vm volume params
func (o *ImportVMVolumeParams) WithContentLanguage(contentLanguage *string) *ImportVMVolumeParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the import Vm volume params
func (o *ImportVMVolumeParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the import Vm volume params
func (o *ImportVMVolumeParams) WithRequestBody(requestBody []*models.ImportVMVolumeParams) *ImportVMVolumeParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the import Vm volume params
func (o *ImportVMVolumeParams) SetRequestBody(requestBody []*models.ImportVMVolumeParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *ImportVMVolumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewDeleteVMParams creates a new DeleteVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVMParams() *DeleteVMParams {
	return &DeleteVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVMParamsWithTimeout creates a new DeleteVMParams object
// with the ability to set a timeout on a request.
func NewDeleteVMParamsWithTimeout(timeout time.Duration) *DeleteVMParams {
	return &DeleteVMParams{
		timeout: timeout,
	}
}

// NewDeleteVMParamsWithContext creates a new DeleteVMParams object
// with the ability to set a context for a request.
func NewDeleteVMParamsWithContext(ctx context.Context) *DeleteVMParams {
	return &DeleteVMParams{
		Context: ctx,
	}
}

// NewDeleteVMParamsWithHTTPClient creates a new DeleteVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVMParamsWithHTTPClient(client *http.Client) *DeleteVMParams {
	return &DeleteVMParams{
		HTTPClient: client,
	}
}

/* DeleteVMParams contains all the parameters to send to the API endpoint
   for the delete Vm operation.

   Typically these are written to a http.Request.
*/
type DeleteVMParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMDeleteParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMParams) WithDefaults() *DeleteVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteVMParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete Vm params
func (o *DeleteVMParams) WithTimeout(timeout time.Duration) *DeleteVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Vm params
func (o *DeleteVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Vm params
func (o *DeleteVMParams) WithContext(ctx context.Context) *DeleteVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Vm params
func (o *DeleteVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Vm params
func (o *DeleteVMParams) WithHTTPClient(client *http.Client) *DeleteVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Vm params
func (o *DeleteVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete Vm params
func (o *DeleteVMParams) WithContentLanguage(contentLanguage *string) *DeleteVMParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete Vm params
func (o *DeleteVMParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete Vm params
func (o *DeleteVMParams) WithRequestBody(requestBody *models.VMDeleteParams) *DeleteVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete Vm params
func (o *DeleteVMParams) SetRequestBody(requestBody *models.VMDeleteParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

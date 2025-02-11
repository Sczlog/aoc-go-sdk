// Code generated by go-swagger; DO NOT EDIT.

package gpu_device

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

// NewGetGpuDevicesParams creates a new GetGpuDevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGpuDevicesParams() *GetGpuDevicesParams {
	return &GetGpuDevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGpuDevicesParamsWithTimeout creates a new GetGpuDevicesParams object
// with the ability to set a timeout on a request.
func NewGetGpuDevicesParamsWithTimeout(timeout time.Duration) *GetGpuDevicesParams {
	return &GetGpuDevicesParams{
		timeout: timeout,
	}
}

// NewGetGpuDevicesParamsWithContext creates a new GetGpuDevicesParams object
// with the ability to set a context for a request.
func NewGetGpuDevicesParamsWithContext(ctx context.Context) *GetGpuDevicesParams {
	return &GetGpuDevicesParams{
		Context: ctx,
	}
}

// NewGetGpuDevicesParamsWithHTTPClient creates a new GetGpuDevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGpuDevicesParamsWithHTTPClient(client *http.Client) *GetGpuDevicesParams {
	return &GetGpuDevicesParams{
		HTTPClient: client,
	}
}

/* GetGpuDevicesParams contains all the parameters to send to the API endpoint
   for the get gpu devices operation.

   Typically these are written to a http.Request.
*/
type GetGpuDevicesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetGpuDevicesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gpu devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGpuDevicesParams) WithDefaults() *GetGpuDevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gpu devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGpuDevicesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetGpuDevicesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get gpu devices params
func (o *GetGpuDevicesParams) WithTimeout(timeout time.Duration) *GetGpuDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gpu devices params
func (o *GetGpuDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gpu devices params
func (o *GetGpuDevicesParams) WithContext(ctx context.Context) *GetGpuDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gpu devices params
func (o *GetGpuDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gpu devices params
func (o *GetGpuDevicesParams) WithHTTPClient(client *http.Client) *GetGpuDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gpu devices params
func (o *GetGpuDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get gpu devices params
func (o *GetGpuDevicesParams) WithContentLanguage(contentLanguage *string) *GetGpuDevicesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get gpu devices params
func (o *GetGpuDevicesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get gpu devices params
func (o *GetGpuDevicesParams) WithRequestBody(requestBody *models.GetGpuDevicesRequestBody) *GetGpuDevicesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get gpu devices params
func (o *GetGpuDevicesParams) SetRequestBody(requestBody *models.GetGpuDevicesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetGpuDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

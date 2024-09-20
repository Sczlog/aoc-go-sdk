// Code generated by go-swagger; DO NOT EDIT.

package usb_device

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

// NewGetUsbDevicesParams creates a new GetUsbDevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsbDevicesParams() *GetUsbDevicesParams {
	return &GetUsbDevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsbDevicesParamsWithTimeout creates a new GetUsbDevicesParams object
// with the ability to set a timeout on a request.
func NewGetUsbDevicesParamsWithTimeout(timeout time.Duration) *GetUsbDevicesParams {
	return &GetUsbDevicesParams{
		timeout: timeout,
	}
}

// NewGetUsbDevicesParamsWithContext creates a new GetUsbDevicesParams object
// with the ability to set a context for a request.
func NewGetUsbDevicesParamsWithContext(ctx context.Context) *GetUsbDevicesParams {
	return &GetUsbDevicesParams{
		Context: ctx,
	}
}

// NewGetUsbDevicesParamsWithHTTPClient creates a new GetUsbDevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsbDevicesParamsWithHTTPClient(client *http.Client) *GetUsbDevicesParams {
	return &GetUsbDevicesParams{
		HTTPClient: client,
	}
}

/* GetUsbDevicesParams contains all the parameters to send to the API endpoint
   for the get usb devices operation.

   Typically these are written to a http.Request.
*/
type GetUsbDevicesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetUsbDevicesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get usb devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsbDevicesParams) WithDefaults() *GetUsbDevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get usb devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsbDevicesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetUsbDevicesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get usb devices params
func (o *GetUsbDevicesParams) WithTimeout(timeout time.Duration) *GetUsbDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get usb devices params
func (o *GetUsbDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get usb devices params
func (o *GetUsbDevicesParams) WithContext(ctx context.Context) *GetUsbDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get usb devices params
func (o *GetUsbDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get usb devices params
func (o *GetUsbDevicesParams) WithHTTPClient(client *http.Client) *GetUsbDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get usb devices params
func (o *GetUsbDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get usb devices params
func (o *GetUsbDevicesParams) WithContentLanguage(contentLanguage *string) *GetUsbDevicesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get usb devices params
func (o *GetUsbDevicesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get usb devices params
func (o *GetUsbDevicesParams) WithRequestBody(requestBody *models.GetUsbDevicesRequestBody) *GetUsbDevicesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get usb devices params
func (o *GetUsbDevicesParams) SetRequestBody(requestBody *models.GetUsbDevicesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsbDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

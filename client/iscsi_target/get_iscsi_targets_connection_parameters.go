// Code generated by go-swagger; DO NOT EDIT.

package iscsi_target

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

// NewGetIscsiTargetsConnectionParams creates a new GetIscsiTargetsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIscsiTargetsConnectionParams() *GetIscsiTargetsConnectionParams {
	return &GetIscsiTargetsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIscsiTargetsConnectionParamsWithTimeout creates a new GetIscsiTargetsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetIscsiTargetsConnectionParamsWithTimeout(timeout time.Duration) *GetIscsiTargetsConnectionParams {
	return &GetIscsiTargetsConnectionParams{
		timeout: timeout,
	}
}

// NewGetIscsiTargetsConnectionParamsWithContext creates a new GetIscsiTargetsConnectionParams object
// with the ability to set a context for a request.
func NewGetIscsiTargetsConnectionParamsWithContext(ctx context.Context) *GetIscsiTargetsConnectionParams {
	return &GetIscsiTargetsConnectionParams{
		Context: ctx,
	}
}

// NewGetIscsiTargetsConnectionParamsWithHTTPClient creates a new GetIscsiTargetsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIscsiTargetsConnectionParamsWithHTTPClient(client *http.Client) *GetIscsiTargetsConnectionParams {
	return &GetIscsiTargetsConnectionParams{
		HTTPClient: client,
	}
}

/* GetIscsiTargetsConnectionParams contains all the parameters to send to the API endpoint
   for the get iscsi targets connection operation.

   Typically these are written to a http.Request.
*/
type GetIscsiTargetsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetIscsiTargetsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get iscsi targets connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIscsiTargetsConnectionParams) WithDefaults() *GetIscsiTargetsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get iscsi targets connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIscsiTargetsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetIscsiTargetsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) WithTimeout(timeout time.Duration) *GetIscsiTargetsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) WithContext(ctx context.Context) *GetIscsiTargetsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) WithHTTPClient(client *http.Client) *GetIscsiTargetsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) WithContentLanguage(contentLanguage *string) *GetIscsiTargetsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) WithRequestBody(requestBody *models.GetIscsiTargetsConnectionRequestBody) *GetIscsiTargetsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get iscsi targets connection params
func (o *GetIscsiTargetsConnectionParams) SetRequestBody(requestBody *models.GetIscsiTargetsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetIscsiTargetsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

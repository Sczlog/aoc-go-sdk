// Code generated by go-swagger; DO NOT EDIT.

package user_audit_log

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

// NewGetUserAuditLogsParams creates a new GetUserAuditLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserAuditLogsParams() *GetUserAuditLogsParams {
	return &GetUserAuditLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserAuditLogsParamsWithTimeout creates a new GetUserAuditLogsParams object
// with the ability to set a timeout on a request.
func NewGetUserAuditLogsParamsWithTimeout(timeout time.Duration) *GetUserAuditLogsParams {
	return &GetUserAuditLogsParams{
		timeout: timeout,
	}
}

// NewGetUserAuditLogsParamsWithContext creates a new GetUserAuditLogsParams object
// with the ability to set a context for a request.
func NewGetUserAuditLogsParamsWithContext(ctx context.Context) *GetUserAuditLogsParams {
	return &GetUserAuditLogsParams{
		Context: ctx,
	}
}

// NewGetUserAuditLogsParamsWithHTTPClient creates a new GetUserAuditLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserAuditLogsParamsWithHTTPClient(client *http.Client) *GetUserAuditLogsParams {
	return &GetUserAuditLogsParams{
		HTTPClient: client,
	}
}

/* GetUserAuditLogsParams contains all the parameters to send to the API endpoint
   for the get user audit logs operation.

   Typically these are written to a http.Request.
*/
type GetUserAuditLogsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetUserAuditLogsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user audit logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAuditLogsParams) WithDefaults() *GetUserAuditLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user audit logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAuditLogsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetUserAuditLogsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get user audit logs params
func (o *GetUserAuditLogsParams) WithTimeout(timeout time.Duration) *GetUserAuditLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user audit logs params
func (o *GetUserAuditLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user audit logs params
func (o *GetUserAuditLogsParams) WithContext(ctx context.Context) *GetUserAuditLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user audit logs params
func (o *GetUserAuditLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user audit logs params
func (o *GetUserAuditLogsParams) WithHTTPClient(client *http.Client) *GetUserAuditLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user audit logs params
func (o *GetUserAuditLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get user audit logs params
func (o *GetUserAuditLogsParams) WithContentLanguage(contentLanguage *string) *GetUserAuditLogsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get user audit logs params
func (o *GetUserAuditLogsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get user audit logs params
func (o *GetUserAuditLogsParams) WithRequestBody(requestBody *models.GetUserAuditLogsRequestBody) *GetUserAuditLogsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get user audit logs params
func (o *GetUserAuditLogsParams) SetRequestBody(requestBody *models.GetUserAuditLogsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserAuditLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

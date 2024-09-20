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

// NewGetUserAuditLogsConnectionParams creates a new GetUserAuditLogsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserAuditLogsConnectionParams() *GetUserAuditLogsConnectionParams {
	return &GetUserAuditLogsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserAuditLogsConnectionParamsWithTimeout creates a new GetUserAuditLogsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetUserAuditLogsConnectionParamsWithTimeout(timeout time.Duration) *GetUserAuditLogsConnectionParams {
	return &GetUserAuditLogsConnectionParams{
		timeout: timeout,
	}
}

// NewGetUserAuditLogsConnectionParamsWithContext creates a new GetUserAuditLogsConnectionParams object
// with the ability to set a context for a request.
func NewGetUserAuditLogsConnectionParamsWithContext(ctx context.Context) *GetUserAuditLogsConnectionParams {
	return &GetUserAuditLogsConnectionParams{
		Context: ctx,
	}
}

// NewGetUserAuditLogsConnectionParamsWithHTTPClient creates a new GetUserAuditLogsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserAuditLogsConnectionParamsWithHTTPClient(client *http.Client) *GetUserAuditLogsConnectionParams {
	return &GetUserAuditLogsConnectionParams{
		HTTPClient: client,
	}
}

/* GetUserAuditLogsConnectionParams contains all the parameters to send to the API endpoint
   for the get user audit logs connection operation.

   Typically these are written to a http.Request.
*/
type GetUserAuditLogsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetUserAuditLogsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user audit logs connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAuditLogsConnectionParams) WithDefaults() *GetUserAuditLogsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user audit logs connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAuditLogsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetUserAuditLogsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) WithTimeout(timeout time.Duration) *GetUserAuditLogsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) WithContext(ctx context.Context) *GetUserAuditLogsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) WithHTTPClient(client *http.Client) *GetUserAuditLogsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) WithContentLanguage(contentLanguage *string) *GetUserAuditLogsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) WithRequestBody(requestBody *models.GetUserAuditLogsConnectionRequestBody) *GetUserAuditLogsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get user audit logs connection params
func (o *GetUserAuditLogsConnectionParams) SetRequestBody(requestBody *models.GetUserAuditLogsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserAuditLogsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

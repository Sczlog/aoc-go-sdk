// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGetUsersParams creates a new GetUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersParams() *GetUsersParams {
	return &GetUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersParamsWithTimeout creates a new GetUsersParams object
// with the ability to set a timeout on a request.
func NewGetUsersParamsWithTimeout(timeout time.Duration) *GetUsersParams {
	return &GetUsersParams{
		timeout: timeout,
	}
}

// NewGetUsersParamsWithContext creates a new GetUsersParams object
// with the ability to set a context for a request.
func NewGetUsersParamsWithContext(ctx context.Context) *GetUsersParams {
	return &GetUsersParams{
		Context: ctx,
	}
}

// NewGetUsersParamsWithHTTPClient creates a new GetUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersParamsWithHTTPClient(client *http.Client) *GetUsersParams {
	return &GetUsersParams{
		HTTPClient: client,
	}
}

/* GetUsersParams contains all the parameters to send to the API endpoint
   for the get users operation.

   Typically these are written to a http.Request.
*/
type GetUsersParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetUsersRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersParams) WithDefaults() *GetUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetUsersParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get users params
func (o *GetUsersParams) WithTimeout(timeout time.Duration) *GetUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users params
func (o *GetUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users params
func (o *GetUsersParams) WithContext(ctx context.Context) *GetUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users params
func (o *GetUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) WithHTTPClient(client *http.Client) *GetUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get users params
func (o *GetUsersParams) WithContentLanguage(contentLanguage *string) *GetUsersParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get users params
func (o *GetUsersParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get users params
func (o *GetUsersParams) WithRequestBody(requestBody *models.GetUsersRequestBody) *GetUsersParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get users params
func (o *GetUsersParams) SetRequestBody(requestBody *models.GetUsersRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewGetOrganizationsParams creates a new GetOrganizationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationsParams() *GetOrganizationsParams {
	return &GetOrganizationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsParamsWithTimeout creates a new GetOrganizationsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationsParamsWithTimeout(timeout time.Duration) *GetOrganizationsParams {
	return &GetOrganizationsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationsParamsWithContext creates a new GetOrganizationsParams object
// with the ability to set a context for a request.
func NewGetOrganizationsParamsWithContext(ctx context.Context) *GetOrganizationsParams {
	return &GetOrganizationsParams{
		Context: ctx,
	}
}

// NewGetOrganizationsParamsWithHTTPClient creates a new GetOrganizationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationsParamsWithHTTPClient(client *http.Client) *GetOrganizationsParams {
	return &GetOrganizationsParams{
		HTTPClient: client,
	}
}

/* GetOrganizationsParams contains all the parameters to send to the API endpoint
   for the get organizations operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetOrganizationsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsParams) WithDefaults() *GetOrganizationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetOrganizationsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) WithTimeout(timeout time.Duration) *GetOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations params
func (o *GetOrganizationsParams) WithContext(ctx context.Context) *GetOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations params
func (o *GetOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations params
func (o *GetOrganizationsParams) WithHTTPClient(client *http.Client) *GetOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations params
func (o *GetOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get organizations params
func (o *GetOrganizationsParams) WithContentLanguage(contentLanguage *string) *GetOrganizationsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get organizations params
func (o *GetOrganizationsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get organizations params
func (o *GetOrganizationsParams) WithRequestBody(requestBody *models.GetOrganizationsRequestBody) *GetOrganizationsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get organizations params
func (o *GetOrganizationsParams) SetRequestBody(requestBody *models.GetOrganizationsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

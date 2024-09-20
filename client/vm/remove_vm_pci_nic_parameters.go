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

// NewRemoveVMPciNicParams creates a new RemoveVMPciNicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveVMPciNicParams() *RemoveVMPciNicParams {
	return &RemoveVMPciNicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveVMPciNicParamsWithTimeout creates a new RemoveVMPciNicParams object
// with the ability to set a timeout on a request.
func NewRemoveVMPciNicParamsWithTimeout(timeout time.Duration) *RemoveVMPciNicParams {
	return &RemoveVMPciNicParams{
		timeout: timeout,
	}
}

// NewRemoveVMPciNicParamsWithContext creates a new RemoveVMPciNicParams object
// with the ability to set a context for a request.
func NewRemoveVMPciNicParamsWithContext(ctx context.Context) *RemoveVMPciNicParams {
	return &RemoveVMPciNicParams{
		Context: ctx,
	}
}

// NewRemoveVMPciNicParamsWithHTTPClient creates a new RemoveVMPciNicParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveVMPciNicParamsWithHTTPClient(client *http.Client) *RemoveVMPciNicParams {
	return &RemoveVMPciNicParams{
		HTTPClient: client,
	}
}

/* RemoveVMPciNicParams contains all the parameters to send to the API endpoint
   for the remove Vm pci nic operation.

   Typically these are written to a http.Request.
*/
type RemoveVMPciNicParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMOperatePciNicParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove Vm pci nic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVMPciNicParams) WithDefaults() *RemoveVMPciNicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove Vm pci nic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVMPciNicParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := RemoveVMPciNicParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) WithTimeout(timeout time.Duration) *RemoveVMPciNicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) WithContext(ctx context.Context) *RemoveVMPciNicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) WithHTTPClient(client *http.Client) *RemoveVMPciNicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) WithContentLanguage(contentLanguage *string) *RemoveVMPciNicParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) WithRequestBody(requestBody *models.VMOperatePciNicParams) *RemoveVMPciNicParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the remove Vm pci nic params
func (o *RemoveVMPciNicParams) SetRequestBody(requestBody *models.VMOperatePciNicParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveVMPciNicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package vm_volume_snapshot

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

// NewGetVMVolumeSnapshotsParams creates a new GetVMVolumeSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMVolumeSnapshotsParams() *GetVMVolumeSnapshotsParams {
	return &GetVMVolumeSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMVolumeSnapshotsParamsWithTimeout creates a new GetVMVolumeSnapshotsParams object
// with the ability to set a timeout on a request.
func NewGetVMVolumeSnapshotsParamsWithTimeout(timeout time.Duration) *GetVMVolumeSnapshotsParams {
	return &GetVMVolumeSnapshotsParams{
		timeout: timeout,
	}
}

// NewGetVMVolumeSnapshotsParamsWithContext creates a new GetVMVolumeSnapshotsParams object
// with the ability to set a context for a request.
func NewGetVMVolumeSnapshotsParamsWithContext(ctx context.Context) *GetVMVolumeSnapshotsParams {
	return &GetVMVolumeSnapshotsParams{
		Context: ctx,
	}
}

// NewGetVMVolumeSnapshotsParamsWithHTTPClient creates a new GetVMVolumeSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMVolumeSnapshotsParamsWithHTTPClient(client *http.Client) *GetVMVolumeSnapshotsParams {
	return &GetVMVolumeSnapshotsParams{
		HTTPClient: client,
	}
}

/* GetVMVolumeSnapshotsParams contains all the parameters to send to the API endpoint
   for the get Vm volume snapshots operation.

   Typically these are written to a http.Request.
*/
type GetVMVolumeSnapshotsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMVolumeSnapshotsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm volume snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMVolumeSnapshotsParams) WithDefaults() *GetVMVolumeSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm volume snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMVolumeSnapshotsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMVolumeSnapshotsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) WithTimeout(timeout time.Duration) *GetVMVolumeSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) WithContext(ctx context.Context) *GetVMVolumeSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) WithHTTPClient(client *http.Client) *GetVMVolumeSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) WithContentLanguage(contentLanguage *string) *GetVMVolumeSnapshotsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) WithRequestBody(requestBody *models.GetVMVolumeSnapshotsRequestBody) *GetVMVolumeSnapshotsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm volume snapshots params
func (o *GetVMVolumeSnapshotsParams) SetRequestBody(requestBody *models.GetVMVolumeSnapshotsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMVolumeSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

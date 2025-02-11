// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new datacenter API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for datacenter API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDatacenter(params *CreateDatacenterParams, opts ...ClientOption) (*CreateDatacenterOK, error)

	GetDatacenters(params *GetDatacentersParams, opts ...ClientOption) (*GetDatacentersOK, error)

	GetDatacentersConnection(params *GetDatacentersConnectionParams, opts ...ClientOption) (*GetDatacentersConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDatacenter create datacenter API
*/
func (a *Client) CreateDatacenter(params *CreateDatacenterParams, opts ...ClientOption) (*CreateDatacenterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDatacenterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateDatacenter",
		Method:             "POST",
		PathPattern:        "/create-datacenter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDatacenterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDatacenterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDatacenter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDatacenters get datacenters API
*/
func (a *Client) GetDatacenters(params *GetDatacentersParams, opts ...ClientOption) (*GetDatacentersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatacentersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDatacenters",
		Method:             "POST",
		PathPattern:        "/get-datacenters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDatacentersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDatacentersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDatacenters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDatacentersConnection get datacenters connection API
*/
func (a *Client) GetDatacentersConnection(params *GetDatacentersConnectionParams, opts ...ClientOption) (*GetDatacentersConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatacentersConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDatacentersConnection",
		Method:             "POST",
		PathPattern:        "/get-datacenters-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDatacentersConnectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDatacentersConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDatacentersConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

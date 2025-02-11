// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ClusterType cluster type
//
// swagger:model ClusterType
type ClusterType string

func NewClusterType(value ClusterType) *ClusterType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterType.
func (m ClusterType) Pointer() *ClusterType {
	return &m
}

const (

	// ClusterTypeACOS captures enum value "ACOS"
	ClusterTypeACOS ClusterType = "ACOS"
)

// for schema
var clusterTypeEnum []interface{}

func init() {
	var res []ClusterType
	if err := json.Unmarshal([]byte(`["ACOS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTypeEnum = append(clusterTypeEnum, v)
	}
}

func (m ClusterType) validateClusterTypeEnum(path, location string, value ClusterType) error {
	if err := validate.EnumCase(path, location, value, clusterTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster type
func (m ClusterType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster type based on context it is used
func (m ClusterType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

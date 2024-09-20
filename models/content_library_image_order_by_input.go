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

// ContentLibraryImageOrderByInput content library image order by input
//
// swagger:model ContentLibraryImageOrderByInput
type ContentLibraryImageOrderByInput string

func NewContentLibraryImageOrderByInput(value ContentLibraryImageOrderByInput) *ContentLibraryImageOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ContentLibraryImageOrderByInput.
func (m ContentLibraryImageOrderByInput) Pointer() *ContentLibraryImageOrderByInput {
	return &m
}

const (

	// ContentLibraryImageOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	ContentLibraryImageOrderByInputCreatedAtASC ContentLibraryImageOrderByInput = "createdAt_ASC"

	// ContentLibraryImageOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	ContentLibraryImageOrderByInputCreatedAtDESC ContentLibraryImageOrderByInput = "createdAt_DESC"

	// ContentLibraryImageOrderByInputDescriptionASC captures enum value "description_ASC"
	ContentLibraryImageOrderByInputDescriptionASC ContentLibraryImageOrderByInput = "description_ASC"

	// ContentLibraryImageOrderByInputDescriptionDESC captures enum value "description_DESC"
	ContentLibraryImageOrderByInputDescriptionDESC ContentLibraryImageOrderByInput = "description_DESC"

	// ContentLibraryImageOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	ContentLibraryImageOrderByInputEntityAsyncStatusASC ContentLibraryImageOrderByInput = "entityAsyncStatus_ASC"

	// ContentLibraryImageOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	ContentLibraryImageOrderByInputEntityAsyncStatusDESC ContentLibraryImageOrderByInput = "entityAsyncStatus_DESC"

	// ContentLibraryImageOrderByInputIDASC captures enum value "id_ASC"
	ContentLibraryImageOrderByInputIDASC ContentLibraryImageOrderByInput = "id_ASC"

	// ContentLibraryImageOrderByInputIDDESC captures enum value "id_DESC"
	ContentLibraryImageOrderByInputIDDESC ContentLibraryImageOrderByInput = "id_DESC"

	// ContentLibraryImageOrderByInputNameASC captures enum value "name_ASC"
	ContentLibraryImageOrderByInputNameASC ContentLibraryImageOrderByInput = "name_ASC"

	// ContentLibraryImageOrderByInputNameDESC captures enum value "name_DESC"
	ContentLibraryImageOrderByInputNameDESC ContentLibraryImageOrderByInput = "name_DESC"

	// ContentLibraryImageOrderByInputPathASC captures enum value "path_ASC"
	ContentLibraryImageOrderByInputPathASC ContentLibraryImageOrderByInput = "path_ASC"

	// ContentLibraryImageOrderByInputPathDESC captures enum value "path_DESC"
	ContentLibraryImageOrderByInputPathDESC ContentLibraryImageOrderByInput = "path_DESC"

	// ContentLibraryImageOrderByInputSizeASC captures enum value "size_ASC"
	ContentLibraryImageOrderByInputSizeASC ContentLibraryImageOrderByInput = "size_ASC"

	// ContentLibraryImageOrderByInputSizeDESC captures enum value "size_DESC"
	ContentLibraryImageOrderByInputSizeDESC ContentLibraryImageOrderByInput = "size_DESC"
)

// for schema
var contentLibraryImageOrderByInputEnum []interface{}

func init() {
	var res []ContentLibraryImageOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","description_ASC","description_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","name_ASC","name_DESC","path_ASC","path_DESC","size_ASC","size_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentLibraryImageOrderByInputEnum = append(contentLibraryImageOrderByInputEnum, v)
	}
}

func (m ContentLibraryImageOrderByInput) validateContentLibraryImageOrderByInputEnum(path, location string, value ContentLibraryImageOrderByInput) error {
	if err := validate.EnumCase(path, location, value, contentLibraryImageOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this content library image order by input
func (m ContentLibraryImageOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContentLibraryImageOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this content library image order by input based on context it is used
func (m ContentLibraryImageOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

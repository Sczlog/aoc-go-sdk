// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserRoleNext user role next
//
// swagger:model UserRoleNext
type UserRoleNext struct {

	// actions
	// Required: true
	Actions []string `json:"actions"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// platform
	// Required: true
	Platform *UserRolePlatform `json:"platform"`

	// preset
	Preset *UserRolePreset `json:"preset,omitempty"`

	// users
	Users []*NestedUser `json:"users,omitempty"`
}

// Validate validates this user role next
func (m *UserRoleNext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRoleNext) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("actions", "body", m.Actions); err != nil {
		return err
	}

	return nil
}

func (m *UserRoleNext) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UserRoleNext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UserRoleNext) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNext) validatePreset(formats strfmt.Registry) error {
	if swag.IsZero(m.Preset) { // not required
		return nil
	}

	if m.Preset != nil {
		if err := m.Preset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNext) validateUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this user role next based on the context it is used
func (m *UserRoleNext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRoleNext) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {

	if m.Platform != nil {
		if err := m.Platform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNext) contextValidatePreset(ctx context.Context, formats strfmt.Registry) error {

	if m.Preset != nil {
		if err := m.Preset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNext) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Users); i++ {

		if m.Users[i] != nil {
			if err := m.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserRoleNext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRoleNext) UnmarshalBinary(b []byte) error {
	var res UserRoleNext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

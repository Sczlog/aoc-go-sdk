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
)

// AlertWhereInput alert where input
//
// swagger:model AlertWhereInput
type AlertWhereInput struct {

	// a n d
	AND []*AlertWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*AlertWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*AlertWhereInput `json:"OR,omitempty"`

	// cause
	Cause *string `json:"cause,omitempty"`

	// cause contains
	CauseContains *string `json:"cause_contains,omitempty"`

	// cause ends with
	CauseEndsWith *string `json:"cause_ends_with,omitempty"`

	// cause gt
	CauseGt *string `json:"cause_gt,omitempty"`

	// cause gte
	CauseGte *string `json:"cause_gte,omitempty"`

	// cause in
	CauseIn []string `json:"cause_in,omitempty"`

	// cause lt
	CauseLt *string `json:"cause_lt,omitempty"`

	// cause lte
	CauseLte *string `json:"cause_lte,omitempty"`

	// cause not
	CauseNot *string `json:"cause_not,omitempty"`

	// cause not contains
	CauseNotContains *string `json:"cause_not_contains,omitempty"`

	// cause not ends with
	CauseNotEndsWith *string `json:"cause_not_ends_with,omitempty"`

	// cause not in
	CauseNotIn []string `json:"cause_not_in,omitempty"`

	// cause not starts with
	CauseNotStartsWith *string `json:"cause_not_starts_with,omitempty"`

	// cause starts with
	CauseStartsWith *string `json:"cause_starts_with,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// create time
	CreateTime *string `json:"create_time,omitempty"`

	// create time gt
	CreateTimeGt *string `json:"create_time_gt,omitempty"`

	// create time gte
	CreateTimeGte *string `json:"create_time_gte,omitempty"`

	// create time in
	CreateTimeIn []string `json:"create_time_in,omitempty"`

	// create time lt
	CreateTimeLt *string `json:"create_time_lt,omitempty"`

	// create time lte
	CreateTimeLte *string `json:"create_time_lte,omitempty"`

	// create time not
	CreateTimeNot *string `json:"create_time_not,omitempty"`

	// create time not in
	CreateTimeNotIn []string `json:"create_time_not_in,omitempty"`

	// ended
	Ended *bool `json:"ended,omitempty"`

	// ended not
	EndedNot *bool `json:"ended_not,omitempty"`

	// host
	Host *HostWhereInput `json:"host,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// impact
	Impact *string `json:"impact,omitempty"`

	// impact contains
	ImpactContains *string `json:"impact_contains,omitempty"`

	// impact ends with
	ImpactEndsWith *string `json:"impact_ends_with,omitempty"`

	// impact gt
	ImpactGt *string `json:"impact_gt,omitempty"`

	// impact gte
	ImpactGte *string `json:"impact_gte,omitempty"`

	// impact in
	ImpactIn []string `json:"impact_in,omitempty"`

	// impact lt
	ImpactLt *string `json:"impact_lt,omitempty"`

	// impact lte
	ImpactLte *string `json:"impact_lte,omitempty"`

	// impact not
	ImpactNot *string `json:"impact_not,omitempty"`

	// impact not contains
	ImpactNotContains *string `json:"impact_not_contains,omitempty"`

	// impact not ends with
	ImpactNotEndsWith *string `json:"impact_not_ends_with,omitempty"`

	// impact not in
	ImpactNotIn []string `json:"impact_not_in,omitempty"`

	// impact not starts with
	ImpactNotStartsWith *string `json:"impact_not_starts_with,omitempty"`

	// impact starts with
	ImpactStartsWith *string `json:"impact_starts_with,omitempty"`

	// local create time
	LocalCreateTime *string `json:"local_create_time,omitempty"`

	// local create time gt
	LocalCreateTimeGt *string `json:"local_create_time_gt,omitempty"`

	// local create time gte
	LocalCreateTimeGte *string `json:"local_create_time_gte,omitempty"`

	// local create time in
	LocalCreateTimeIn []string `json:"local_create_time_in,omitempty"`

	// local create time lt
	LocalCreateTimeLt *string `json:"local_create_time_lt,omitempty"`

	// local create time lte
	LocalCreateTimeLte *string `json:"local_create_time_lte,omitempty"`

	// local create time not
	LocalCreateTimeNot *string `json:"local_create_time_not,omitempty"`

	// local create time not in
	LocalCreateTimeNotIn []string `json:"local_create_time_not_in,omitempty"`

	// local end time
	LocalEndTime *string `json:"local_end_time,omitempty"`

	// local end time contains
	LocalEndTimeContains *string `json:"local_end_time_contains,omitempty"`

	// local end time ends with
	LocalEndTimeEndsWith *string `json:"local_end_time_ends_with,omitempty"`

	// local end time gt
	LocalEndTimeGt *string `json:"local_end_time_gt,omitempty"`

	// local end time gte
	LocalEndTimeGte *string `json:"local_end_time_gte,omitempty"`

	// local end time in
	LocalEndTimeIn []string `json:"local_end_time_in,omitempty"`

	// local end time lt
	LocalEndTimeLt *string `json:"local_end_time_lt,omitempty"`

	// local end time lte
	LocalEndTimeLte *string `json:"local_end_time_lte,omitempty"`

	// local end time not
	LocalEndTimeNot *string `json:"local_end_time_not,omitempty"`

	// local end time not contains
	LocalEndTimeNotContains *string `json:"local_end_time_not_contains,omitempty"`

	// local end time not ends with
	LocalEndTimeNotEndsWith *string `json:"local_end_time_not_ends_with,omitempty"`

	// local end time not in
	LocalEndTimeNotIn []string `json:"local_end_time_not_in,omitempty"`

	// local end time not starts with
	LocalEndTimeNotStartsWith *string `json:"local_end_time_not_starts_with,omitempty"`

	// local end time starts with
	LocalEndTimeStartsWith *string `json:"local_end_time_starts_with,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// local start time
	LocalStartTime *string `json:"local_start_time,omitempty"`

	// local start time contains
	LocalStartTimeContains *string `json:"local_start_time_contains,omitempty"`

	// local start time ends with
	LocalStartTimeEndsWith *string `json:"local_start_time_ends_with,omitempty"`

	// local start time gt
	LocalStartTimeGt *string `json:"local_start_time_gt,omitempty"`

	// local start time gte
	LocalStartTimeGte *string `json:"local_start_time_gte,omitempty"`

	// local start time in
	LocalStartTimeIn []string `json:"local_start_time_in,omitempty"`

	// local start time lt
	LocalStartTimeLt *string `json:"local_start_time_lt,omitempty"`

	// local start time lte
	LocalStartTimeLte *string `json:"local_start_time_lte,omitempty"`

	// local start time not
	LocalStartTimeNot *string `json:"local_start_time_not,omitempty"`

	// local start time not contains
	LocalStartTimeNotContains *string `json:"local_start_time_not_contains,omitempty"`

	// local start time not ends with
	LocalStartTimeNotEndsWith *string `json:"local_start_time_not_ends_with,omitempty"`

	// local start time not in
	LocalStartTimeNotIn []string `json:"local_start_time_not_in,omitempty"`

	// local start time not starts with
	LocalStartTimeNotStartsWith *string `json:"local_start_time_not_starts_with,omitempty"`

	// local start time starts with
	LocalStartTimeStartsWith *string `json:"local_start_time_starts_with,omitempty"`

	// local update time
	LocalUpdateTime *string `json:"local_update_time,omitempty"`

	// local update time contains
	LocalUpdateTimeContains *string `json:"local_update_time_contains,omitempty"`

	// local update time ends with
	LocalUpdateTimeEndsWith *string `json:"local_update_time_ends_with,omitempty"`

	// local update time gt
	LocalUpdateTimeGt *string `json:"local_update_time_gt,omitempty"`

	// local update time gte
	LocalUpdateTimeGte *string `json:"local_update_time_gte,omitempty"`

	// local update time in
	LocalUpdateTimeIn []string `json:"local_update_time_in,omitempty"`

	// local update time lt
	LocalUpdateTimeLt *string `json:"local_update_time_lt,omitempty"`

	// local update time lte
	LocalUpdateTimeLte *string `json:"local_update_time_lte,omitempty"`

	// local update time not
	LocalUpdateTimeNot *string `json:"local_update_time_not,omitempty"`

	// local update time not contains
	LocalUpdateTimeNotContains *string `json:"local_update_time_not_contains,omitempty"`

	// local update time not ends with
	LocalUpdateTimeNotEndsWith *string `json:"local_update_time_not_ends_with,omitempty"`

	// local update time not in
	LocalUpdateTimeNotIn []string `json:"local_update_time_not_in,omitempty"`

	// local update time not starts with
	LocalUpdateTimeNotStartsWith *string `json:"local_update_time_not_starts_with,omitempty"`

	// local update time starts with
	LocalUpdateTimeStartsWith *string `json:"local_update_time_starts_with,omitempty"`

	// message
	Message *string `json:"message,omitempty"`

	// message contains
	MessageContains *string `json:"message_contains,omitempty"`

	// message ends with
	MessageEndsWith *string `json:"message_ends_with,omitempty"`

	// message gt
	MessageGt *string `json:"message_gt,omitempty"`

	// message gte
	MessageGte *string `json:"message_gte,omitempty"`

	// message in
	MessageIn []string `json:"message_in,omitempty"`

	// message lt
	MessageLt *string `json:"message_lt,omitempty"`

	// message lte
	MessageLte *string `json:"message_lte,omitempty"`

	// message not
	MessageNot *string `json:"message_not,omitempty"`

	// message not contains
	MessageNotContains *string `json:"message_not_contains,omitempty"`

	// message not ends with
	MessageNotEndsWith *string `json:"message_not_ends_with,omitempty"`

	// message not in
	MessageNotIn []string `json:"message_not_in,omitempty"`

	// message not starts with
	MessageNotStartsWith *string `json:"message_not_starts_with,omitempty"`

	// message starts with
	MessageStartsWith *string `json:"message_starts_with,omitempty"`

	// severity
	Severity *string `json:"severity,omitempty"`

	// severity contains
	SeverityContains *string `json:"severity_contains,omitempty"`

	// severity ends with
	SeverityEndsWith *string `json:"severity_ends_with,omitempty"`

	// severity gt
	SeverityGt *string `json:"severity_gt,omitempty"`

	// severity gte
	SeverityGte *string `json:"severity_gte,omitempty"`

	// severity in
	SeverityIn []string `json:"severity_in,omitempty"`

	// severity lt
	SeverityLt *string `json:"severity_lt,omitempty"`

	// severity lte
	SeverityLte *string `json:"severity_lte,omitempty"`

	// severity not
	SeverityNot *string `json:"severity_not,omitempty"`

	// severity not contains
	SeverityNotContains *string `json:"severity_not_contains,omitempty"`

	// severity not ends with
	SeverityNotEndsWith *string `json:"severity_not_ends_with,omitempty"`

	// severity not in
	SeverityNotIn []string `json:"severity_not_in,omitempty"`

	// severity not starts with
	SeverityNotStartsWith *string `json:"severity_not_starts_with,omitempty"`

	// severity starts with
	SeverityStartsWith *string `json:"severity_starts_with,omitempty"`

	// solution
	Solution *string `json:"solution,omitempty"`

	// solution contains
	SolutionContains *string `json:"solution_contains,omitempty"`

	// solution ends with
	SolutionEndsWith *string `json:"solution_ends_with,omitempty"`

	// solution gt
	SolutionGt *string `json:"solution_gt,omitempty"`

	// solution gte
	SolutionGte *string `json:"solution_gte,omitempty"`

	// solution in
	SolutionIn []string `json:"solution_in,omitempty"`

	// solution lt
	SolutionLt *string `json:"solution_lt,omitempty"`

	// solution lte
	SolutionLte *string `json:"solution_lte,omitempty"`

	// solution not
	SolutionNot *string `json:"solution_not,omitempty"`

	// solution not contains
	SolutionNotContains *string `json:"solution_not_contains,omitempty"`

	// solution not ends with
	SolutionNotEndsWith *string `json:"solution_not_ends_with,omitempty"`

	// solution not in
	SolutionNotIn []string `json:"solution_not_in,omitempty"`

	// solution not starts with
	SolutionNotStartsWith *string `json:"solution_not_starts_with,omitempty"`

	// solution starts with
	SolutionStartsWith *string `json:"solution_starts_with,omitempty"`

	// threshold
	Threshold *float64 `json:"threshold,omitempty"`

	// threshold gt
	ThresholdGt *float64 `json:"threshold_gt,omitempty"`

	// threshold gte
	ThresholdGte *float64 `json:"threshold_gte,omitempty"`

	// threshold in
	ThresholdIn []float64 `json:"threshold_in,omitempty"`

	// threshold lt
	ThresholdLt *float64 `json:"threshold_lt,omitempty"`

	// threshold lte
	ThresholdLte *float64 `json:"threshold_lte,omitempty"`

	// threshold not
	ThresholdNot *float64 `json:"threshold_not,omitempty"`

	// threshold not in
	ThresholdNotIn []float64 `json:"threshold_not_in,omitempty"`

	// value
	Value *float64 `json:"value,omitempty"`

	// value gt
	ValueGt *float64 `json:"value_gt,omitempty"`

	// value gte
	ValueGte *float64 `json:"value_gte,omitempty"`

	// value in
	ValueIn []float64 `json:"value_in,omitempty"`

	// value lt
	ValueLt *float64 `json:"value_lt,omitempty"`

	// value lte
	ValueLte *float64 `json:"value_lte,omitempty"`

	// value not
	ValueNot *float64 `json:"value_not,omitempty"`

	// value not in
	ValueNotIn []float64 `json:"value_not_in,omitempty"`

	// vms every
	VmsEvery *VMWhereInput `json:"vms_every,omitempty"`

	// vms none
	VmsNone *VMWhereInput `json:"vms_none,omitempty"`

	// vms some
	VmsSome *VMWhereInput `json:"vms_some,omitempty"`
}

// Validate validates this alert where input
func (m *AlertWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmsSome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertWhereInput) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *AlertWhereInput) validateHost(formats strfmt.Registry) error {
	if swag.IsZero(m.Host) { // not required
		return nil
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *AlertWhereInput) validateVmsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.VmsEvery) { // not required
		return nil
	}

	if m.VmsEvery != nil {
		if err := m.VmsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vms_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vms_every")
			}
			return err
		}
	}

	return nil
}

func (m *AlertWhereInput) validateVmsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.VmsNone) { // not required
		return nil
	}

	if m.VmsNone != nil {
		if err := m.VmsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vms_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vms_none")
			}
			return err
		}
	}

	return nil
}

func (m *AlertWhereInput) validateVmsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.VmsSome) { // not required
		return nil
	}

	if m.VmsSome != nil {
		if err := m.VmsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vms_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vms_some")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this alert where input based on the context it is used
func (m *AlertWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *AlertWhereInput) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *AlertWhereInput) contextValidateVmsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.VmsEvery != nil {
		if err := m.VmsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vms_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vms_every")
			}
			return err
		}
	}

	return nil
}

func (m *AlertWhereInput) contextValidateVmsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.VmsNone != nil {
		if err := m.VmsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vms_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vms_none")
			}
			return err
		}
	}

	return nil
}

func (m *AlertWhereInput) contextValidateVmsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.VmsSome != nil {
		if err := m.VmsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vms_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vms_some")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertWhereInput) UnmarshalBinary(b []byte) error {
	var res AlertWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

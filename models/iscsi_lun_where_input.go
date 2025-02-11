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

// IscsiLunWhereInput iscsi lun where input
//
// swagger:model IscsiLunWhereInput
type IscsiLunWhereInput struct {

	// a n d
	AND []*IscsiLunWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*IscsiLunWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*IscsiLunWhereInput `json:"OR,omitempty"`

	// allowed initiators
	AllowedInitiators *string `json:"allowed_initiators,omitempty"`

	// allowed initiators contains
	AllowedInitiatorsContains *string `json:"allowed_initiators_contains,omitempty"`

	// allowed initiators ends with
	AllowedInitiatorsEndsWith *string `json:"allowed_initiators_ends_with,omitempty"`

	// allowed initiators gt
	AllowedInitiatorsGt *string `json:"allowed_initiators_gt,omitempty"`

	// allowed initiators gte
	AllowedInitiatorsGte *string `json:"allowed_initiators_gte,omitempty"`

	// allowed initiators in
	AllowedInitiatorsIn []string `json:"allowed_initiators_in,omitempty"`

	// allowed initiators lt
	AllowedInitiatorsLt *string `json:"allowed_initiators_lt,omitempty"`

	// allowed initiators lte
	AllowedInitiatorsLte *string `json:"allowed_initiators_lte,omitempty"`

	// allowed initiators not
	AllowedInitiatorsNot *string `json:"allowed_initiators_not,omitempty"`

	// allowed initiators not contains
	AllowedInitiatorsNotContains *string `json:"allowed_initiators_not_contains,omitempty"`

	// allowed initiators not ends with
	AllowedInitiatorsNotEndsWith *string `json:"allowed_initiators_not_ends_with,omitempty"`

	// allowed initiators not in
	AllowedInitiatorsNotIn []string `json:"allowed_initiators_not_in,omitempty"`

	// allowed initiators not starts with
	AllowedInitiatorsNotStartsWith *string `json:"allowed_initiators_not_starts_with,omitempty"`

	// allowed initiators starts with
	AllowedInitiatorsStartsWith *string `json:"allowed_initiators_starts_with,omitempty"`

	// assigned size
	AssignedSize *int64 `json:"assigned_size,omitempty"`

	// assigned size gt
	AssignedSizeGt *int64 `json:"assigned_size_gt,omitempty"`

	// assigned size gte
	AssignedSizeGte *int64 `json:"assigned_size_gte,omitempty"`

	// assigned size in
	AssignedSizeIn []int64 `json:"assigned_size_in,omitempty"`

	// assigned size lt
	AssignedSizeLt *int64 `json:"assigned_size_lt,omitempty"`

	// assigned size lte
	AssignedSizeLte *int64 `json:"assigned_size_lte,omitempty"`

	// assigned size not
	AssignedSizeNot *int64 `json:"assigned_size_not,omitempty"`

	// assigned size not in
	AssignedSizeNotIn []int64 `json:"assigned_size_not_in,omitempty"`

	// bps
	Bps *int64 `json:"bps,omitempty"`

	// bps gt
	BpsGt *int64 `json:"bps_gt,omitempty"`

	// bps gte
	BpsGte *int64 `json:"bps_gte,omitempty"`

	// bps in
	BpsIn []int64 `json:"bps_in,omitempty"`

	// bps lt
	BpsLt *int64 `json:"bps_lt,omitempty"`

	// bps lte
	BpsLte *int64 `json:"bps_lte,omitempty"`

	// bps max
	BpsMax *int64 `json:"bps_max,omitempty"`

	// bps max gt
	BpsMaxGt *int64 `json:"bps_max_gt,omitempty"`

	// bps max gte
	BpsMaxGte *int64 `json:"bps_max_gte,omitempty"`

	// bps max in
	BpsMaxIn []int64 `json:"bps_max_in,omitempty"`

	// bps max length
	BpsMaxLength *int64 `json:"bps_max_length,omitempty"`

	// bps max length gt
	BpsMaxLengthGt *int64 `json:"bps_max_length_gt,omitempty"`

	// bps max length gte
	BpsMaxLengthGte *int64 `json:"bps_max_length_gte,omitempty"`

	// bps max length in
	BpsMaxLengthIn []int64 `json:"bps_max_length_in,omitempty"`

	// bps max length lt
	BpsMaxLengthLt *int64 `json:"bps_max_length_lt,omitempty"`

	// bps max length lte
	BpsMaxLengthLte *int64 `json:"bps_max_length_lte,omitempty"`

	// bps max length not
	BpsMaxLengthNot *int64 `json:"bps_max_length_not,omitempty"`

	// bps max length not in
	BpsMaxLengthNotIn []int64 `json:"bps_max_length_not_in,omitempty"`

	// bps max lt
	BpsMaxLt *int64 `json:"bps_max_lt,omitempty"`

	// bps max lte
	BpsMaxLte *int64 `json:"bps_max_lte,omitempty"`

	// bps max not
	BpsMaxNot *int64 `json:"bps_max_not,omitempty"`

	// bps max not in
	BpsMaxNotIn []int64 `json:"bps_max_not_in,omitempty"`

	// bps not
	BpsNot *int64 `json:"bps_not,omitempty"`

	// bps not in
	BpsNotIn []int64 `json:"bps_not_in,omitempty"`

	// bps rd
	BpsRd *int64 `json:"bps_rd,omitempty"`

	// bps rd gt
	BpsRdGt *int64 `json:"bps_rd_gt,omitempty"`

	// bps rd gte
	BpsRdGte *int64 `json:"bps_rd_gte,omitempty"`

	// bps rd in
	BpsRdIn []int64 `json:"bps_rd_in,omitempty"`

	// bps rd lt
	BpsRdLt *int64 `json:"bps_rd_lt,omitempty"`

	// bps rd lte
	BpsRdLte *int64 `json:"bps_rd_lte,omitempty"`

	// bps rd max
	BpsRdMax *int64 `json:"bps_rd_max,omitempty"`

	// bps rd max gt
	BpsRdMaxGt *int64 `json:"bps_rd_max_gt,omitempty"`

	// bps rd max gte
	BpsRdMaxGte *int64 `json:"bps_rd_max_gte,omitempty"`

	// bps rd max in
	BpsRdMaxIn []int64 `json:"bps_rd_max_in,omitempty"`

	// bps rd max length
	BpsRdMaxLength *int64 `json:"bps_rd_max_length,omitempty"`

	// bps rd max length gt
	BpsRdMaxLengthGt *int64 `json:"bps_rd_max_length_gt,omitempty"`

	// bps rd max length gte
	BpsRdMaxLengthGte *int64 `json:"bps_rd_max_length_gte,omitempty"`

	// bps rd max length in
	BpsRdMaxLengthIn []int64 `json:"bps_rd_max_length_in,omitempty"`

	// bps rd max length lt
	BpsRdMaxLengthLt *int64 `json:"bps_rd_max_length_lt,omitempty"`

	// bps rd max length lte
	BpsRdMaxLengthLte *int64 `json:"bps_rd_max_length_lte,omitempty"`

	// bps rd max length not
	BpsRdMaxLengthNot *int64 `json:"bps_rd_max_length_not,omitempty"`

	// bps rd max length not in
	BpsRdMaxLengthNotIn []int64 `json:"bps_rd_max_length_not_in,omitempty"`

	// bps rd max lt
	BpsRdMaxLt *int64 `json:"bps_rd_max_lt,omitempty"`

	// bps rd max lte
	BpsRdMaxLte *int64 `json:"bps_rd_max_lte,omitempty"`

	// bps rd max not
	BpsRdMaxNot *int64 `json:"bps_rd_max_not,omitempty"`

	// bps rd max not in
	BpsRdMaxNotIn []int64 `json:"bps_rd_max_not_in,omitempty"`

	// bps rd not
	BpsRdNot *int64 `json:"bps_rd_not,omitempty"`

	// bps rd not in
	BpsRdNotIn []int64 `json:"bps_rd_not_in,omitempty"`

	// bps wr
	BpsWr *int64 `json:"bps_wr,omitempty"`

	// bps wr gt
	BpsWrGt *int64 `json:"bps_wr_gt,omitempty"`

	// bps wr gte
	BpsWrGte *int64 `json:"bps_wr_gte,omitempty"`

	// bps wr in
	BpsWrIn []int64 `json:"bps_wr_in,omitempty"`

	// bps wr lt
	BpsWrLt *int64 `json:"bps_wr_lt,omitempty"`

	// bps wr lte
	BpsWrLte *int64 `json:"bps_wr_lte,omitempty"`

	// bps wr max
	BpsWrMax *int64 `json:"bps_wr_max,omitempty"`

	// bps wr max gt
	BpsWrMaxGt *int64 `json:"bps_wr_max_gt,omitempty"`

	// bps wr max gte
	BpsWrMaxGte *int64 `json:"bps_wr_max_gte,omitempty"`

	// bps wr max in
	BpsWrMaxIn []int64 `json:"bps_wr_max_in,omitempty"`

	// bps wr max length
	BpsWrMaxLength *int64 `json:"bps_wr_max_length,omitempty"`

	// bps wr max length gt
	BpsWrMaxLengthGt *int64 `json:"bps_wr_max_length_gt,omitempty"`

	// bps wr max length gte
	BpsWrMaxLengthGte *int64 `json:"bps_wr_max_length_gte,omitempty"`

	// bps wr max length in
	BpsWrMaxLengthIn []int64 `json:"bps_wr_max_length_in,omitempty"`

	// bps wr max length lt
	BpsWrMaxLengthLt *int64 `json:"bps_wr_max_length_lt,omitempty"`

	// bps wr max length lte
	BpsWrMaxLengthLte *int64 `json:"bps_wr_max_length_lte,omitempty"`

	// bps wr max length not
	BpsWrMaxLengthNot *int64 `json:"bps_wr_max_length_not,omitempty"`

	// bps wr max length not in
	BpsWrMaxLengthNotIn []int64 `json:"bps_wr_max_length_not_in,omitempty"`

	// bps wr max lt
	BpsWrMaxLt *int64 `json:"bps_wr_max_lt,omitempty"`

	// bps wr max lte
	BpsWrMaxLte *int64 `json:"bps_wr_max_lte,omitempty"`

	// bps wr max not
	BpsWrMaxNot *int64 `json:"bps_wr_max_not,omitempty"`

	// bps wr max not in
	BpsWrMaxNotIn []int64 `json:"bps_wr_max_not_in,omitempty"`

	// bps wr not
	BpsWrNot *int64 `json:"bps_wr_not,omitempty"`

	// bps wr not in
	BpsWrNotIn []int64 `json:"bps_wr_not_in,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

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

	// io size
	IoSize *int64 `json:"io_size,omitempty"`

	// io size gt
	IoSizeGt *int64 `json:"io_size_gt,omitempty"`

	// io size gte
	IoSizeGte *int64 `json:"io_size_gte,omitempty"`

	// io size in
	IoSizeIn []int64 `json:"io_size_in,omitempty"`

	// io size lt
	IoSizeLt *int64 `json:"io_size_lt,omitempty"`

	// io size lte
	IoSizeLte *int64 `json:"io_size_lte,omitempty"`

	// io size not
	IoSizeNot *int64 `json:"io_size_not,omitempty"`

	// io size not in
	IoSizeNotIn []int64 `json:"io_size_not_in,omitempty"`

	// iops
	Iops *int64 `json:"iops,omitempty"`

	// iops gt
	IopsGt *int64 `json:"iops_gt,omitempty"`

	// iops gte
	IopsGte *int64 `json:"iops_gte,omitempty"`

	// iops in
	IopsIn []int64 `json:"iops_in,omitempty"`

	// iops lt
	IopsLt *int64 `json:"iops_lt,omitempty"`

	// iops lte
	IopsLte *int64 `json:"iops_lte,omitempty"`

	// iops max
	IopsMax *int64 `json:"iops_max,omitempty"`

	// iops max gt
	IopsMaxGt *int64 `json:"iops_max_gt,omitempty"`

	// iops max gte
	IopsMaxGte *int64 `json:"iops_max_gte,omitempty"`

	// iops max in
	IopsMaxIn []int64 `json:"iops_max_in,omitempty"`

	// iops max length
	IopsMaxLength *int64 `json:"iops_max_length,omitempty"`

	// iops max length gt
	IopsMaxLengthGt *int64 `json:"iops_max_length_gt,omitempty"`

	// iops max length gte
	IopsMaxLengthGte *int64 `json:"iops_max_length_gte,omitempty"`

	// iops max length in
	IopsMaxLengthIn []int64 `json:"iops_max_length_in,omitempty"`

	// iops max length lt
	IopsMaxLengthLt *int64 `json:"iops_max_length_lt,omitempty"`

	// iops max length lte
	IopsMaxLengthLte *int64 `json:"iops_max_length_lte,omitempty"`

	// iops max length not
	IopsMaxLengthNot *int64 `json:"iops_max_length_not,omitempty"`

	// iops max length not in
	IopsMaxLengthNotIn []int64 `json:"iops_max_length_not_in,omitempty"`

	// iops max lt
	IopsMaxLt *int64 `json:"iops_max_lt,omitempty"`

	// iops max lte
	IopsMaxLte *int64 `json:"iops_max_lte,omitempty"`

	// iops max not
	IopsMaxNot *int64 `json:"iops_max_not,omitempty"`

	// iops max not in
	IopsMaxNotIn []int64 `json:"iops_max_not_in,omitempty"`

	// iops not
	IopsNot *int64 `json:"iops_not,omitempty"`

	// iops not in
	IopsNotIn []int64 `json:"iops_not_in,omitempty"`

	// iops rd
	IopsRd *int64 `json:"iops_rd,omitempty"`

	// iops rd gt
	IopsRdGt *int64 `json:"iops_rd_gt,omitempty"`

	// iops rd gte
	IopsRdGte *int64 `json:"iops_rd_gte,omitempty"`

	// iops rd in
	IopsRdIn []int64 `json:"iops_rd_in,omitempty"`

	// iops rd lt
	IopsRdLt *int64 `json:"iops_rd_lt,omitempty"`

	// iops rd lte
	IopsRdLte *int64 `json:"iops_rd_lte,omitempty"`

	// iops rd max
	IopsRdMax *int64 `json:"iops_rd_max,omitempty"`

	// iops rd max gt
	IopsRdMaxGt *int64 `json:"iops_rd_max_gt,omitempty"`

	// iops rd max gte
	IopsRdMaxGte *int64 `json:"iops_rd_max_gte,omitempty"`

	// iops rd max in
	IopsRdMaxIn []int64 `json:"iops_rd_max_in,omitempty"`

	// iops rd max length
	IopsRdMaxLength *int64 `json:"iops_rd_max_length,omitempty"`

	// iops rd max length gt
	IopsRdMaxLengthGt *int64 `json:"iops_rd_max_length_gt,omitempty"`

	// iops rd max length gte
	IopsRdMaxLengthGte *int64 `json:"iops_rd_max_length_gte,omitempty"`

	// iops rd max length in
	IopsRdMaxLengthIn []int64 `json:"iops_rd_max_length_in,omitempty"`

	// iops rd max length lt
	IopsRdMaxLengthLt *int64 `json:"iops_rd_max_length_lt,omitempty"`

	// iops rd max length lte
	IopsRdMaxLengthLte *int64 `json:"iops_rd_max_length_lte,omitempty"`

	// iops rd max length not
	IopsRdMaxLengthNot *int64 `json:"iops_rd_max_length_not,omitempty"`

	// iops rd max length not in
	IopsRdMaxLengthNotIn []int64 `json:"iops_rd_max_length_not_in,omitempty"`

	// iops rd max lt
	IopsRdMaxLt *int64 `json:"iops_rd_max_lt,omitempty"`

	// iops rd max lte
	IopsRdMaxLte *int64 `json:"iops_rd_max_lte,omitempty"`

	// iops rd max not
	IopsRdMaxNot *int64 `json:"iops_rd_max_not,omitempty"`

	// iops rd max not in
	IopsRdMaxNotIn []int64 `json:"iops_rd_max_not_in,omitempty"`

	// iops rd not
	IopsRdNot *int64 `json:"iops_rd_not,omitempty"`

	// iops rd not in
	IopsRdNotIn []int64 `json:"iops_rd_not_in,omitempty"`

	// iops wr
	IopsWr *int64 `json:"iops_wr,omitempty"`

	// iops wr gt
	IopsWrGt *int64 `json:"iops_wr_gt,omitempty"`

	// iops wr gte
	IopsWrGte *int64 `json:"iops_wr_gte,omitempty"`

	// iops wr in
	IopsWrIn []int64 `json:"iops_wr_in,omitempty"`

	// iops wr lt
	IopsWrLt *int64 `json:"iops_wr_lt,omitempty"`

	// iops wr lte
	IopsWrLte *int64 `json:"iops_wr_lte,omitempty"`

	// iops wr max
	IopsWrMax *int64 `json:"iops_wr_max,omitempty"`

	// iops wr max gt
	IopsWrMaxGt *int64 `json:"iops_wr_max_gt,omitempty"`

	// iops wr max gte
	IopsWrMaxGte *int64 `json:"iops_wr_max_gte,omitempty"`

	// iops wr max in
	IopsWrMaxIn []int64 `json:"iops_wr_max_in,omitempty"`

	// iops wr max length
	IopsWrMaxLength *int64 `json:"iops_wr_max_length,omitempty"`

	// iops wr max length gt
	IopsWrMaxLengthGt *int64 `json:"iops_wr_max_length_gt,omitempty"`

	// iops wr max length gte
	IopsWrMaxLengthGte *int64 `json:"iops_wr_max_length_gte,omitempty"`

	// iops wr max length in
	IopsWrMaxLengthIn []int64 `json:"iops_wr_max_length_in,omitempty"`

	// iops wr max length lt
	IopsWrMaxLengthLt *int64 `json:"iops_wr_max_length_lt,omitempty"`

	// iops wr max length lte
	IopsWrMaxLengthLte *int64 `json:"iops_wr_max_length_lte,omitempty"`

	// iops wr max length not
	IopsWrMaxLengthNot *int64 `json:"iops_wr_max_length_not,omitempty"`

	// iops wr max length not in
	IopsWrMaxLengthNotIn []int64 `json:"iops_wr_max_length_not_in,omitempty"`

	// iops wr max lt
	IopsWrMaxLt *int64 `json:"iops_wr_max_lt,omitempty"`

	// iops wr max lte
	IopsWrMaxLte *int64 `json:"iops_wr_max_lte,omitempty"`

	// iops wr max not
	IopsWrMaxNot *int64 `json:"iops_wr_max_not,omitempty"`

	// iops wr max not in
	IopsWrMaxNotIn []int64 `json:"iops_wr_max_not_in,omitempty"`

	// iops wr not
	IopsWrNot *int64 `json:"iops_wr_not,omitempty"`

	// iops wr not in
	IopsWrNotIn []int64 `json:"iops_wr_not_in,omitempty"`

	// iscsi target
	IscsiTarget *IscsiTargetWhereInput `json:"iscsi_target,omitempty"`

	// labels every
	LabelsEvery *LabelWhereInput `json:"labels_every,omitempty"`

	// labels none
	LabelsNone *LabelWhereInput `json:"labels_none,omitempty"`

	// labels some
	LabelsSome *LabelWhereInput `json:"labels_some,omitempty"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local created at gt
	LocalCreatedAtGt *string `json:"local_created_at_gt,omitempty"`

	// local created at gte
	LocalCreatedAtGte *string `json:"local_created_at_gte,omitempty"`

	// local created at in
	LocalCreatedAtIn []string `json:"local_created_at_in,omitempty"`

	// local created at lt
	LocalCreatedAtLt *string `json:"local_created_at_lt,omitempty"`

	// local created at lte
	LocalCreatedAtLte *string `json:"local_created_at_lte,omitempty"`

	// local created at not
	LocalCreatedAtNot *string `json:"local_created_at_not,omitempty"`

	// local created at not in
	LocalCreatedAtNotIn []string `json:"local_created_at_not_in,omitempty"`

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

	// lun id
	LunID *int32 `json:"lun_id,omitempty"`

	// lun id gt
	LunIDGt *int32 `json:"lun_id_gt,omitempty"`

	// lun id gte
	LunIDGte *int32 `json:"lun_id_gte,omitempty"`

	// lun id in
	LunIDIn []int32 `json:"lun_id_in,omitempty"`

	// lun id lt
	LunIDLt *int32 `json:"lun_id_lt,omitempty"`

	// lun id lte
	LunIDLte *int32 `json:"lun_id_lte,omitempty"`

	// lun id not
	LunIDNot *int32 `json:"lun_id_not,omitempty"`

	// lun id not in
	LunIDNotIn []int32 `json:"lun_id_not_in,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// replica num
	ReplicaNum *int32 `json:"replica_num,omitempty"`

	// replica num gt
	ReplicaNumGt *int32 `json:"replica_num_gt,omitempty"`

	// replica num gte
	ReplicaNumGte *int32 `json:"replica_num_gte,omitempty"`

	// replica num in
	ReplicaNumIn []int32 `json:"replica_num_in,omitempty"`

	// replica num lt
	ReplicaNumLt *int32 `json:"replica_num_lt,omitempty"`

	// replica num lte
	ReplicaNumLte *int32 `json:"replica_num_lte,omitempty"`

	// replica num not
	ReplicaNumNot *int32 `json:"replica_num_not,omitempty"`

	// replica num not in
	ReplicaNumNotIn []int32 `json:"replica_num_not_in,omitempty"`

	// shared size
	SharedSize *int64 `json:"shared_size,omitempty"`

	// shared size gt
	SharedSizeGt *int64 `json:"shared_size_gt,omitempty"`

	// shared size gte
	SharedSizeGte *int64 `json:"shared_size_gte,omitempty"`

	// shared size in
	SharedSizeIn []int64 `json:"shared_size_in,omitempty"`

	// shared size lt
	SharedSizeLt *int64 `json:"shared_size_lt,omitempty"`

	// shared size lte
	SharedSizeLte *int64 `json:"shared_size_lte,omitempty"`

	// shared size not
	SharedSizeNot *int64 `json:"shared_size_not,omitempty"`

	// shared size not in
	SharedSizeNotIn []int64 `json:"shared_size_not_in,omitempty"`

	// snapshot num
	SnapshotNum *int32 `json:"snapshot_num,omitempty"`

	// snapshot num gt
	SnapshotNumGt *int32 `json:"snapshot_num_gt,omitempty"`

	// snapshot num gte
	SnapshotNumGte *int32 `json:"snapshot_num_gte,omitempty"`

	// snapshot num in
	SnapshotNumIn []int32 `json:"snapshot_num_in,omitempty"`

	// snapshot num lt
	SnapshotNumLt *int32 `json:"snapshot_num_lt,omitempty"`

	// snapshot num lte
	SnapshotNumLte *int32 `json:"snapshot_num_lte,omitempty"`

	// snapshot num not
	SnapshotNumNot *int32 `json:"snapshot_num_not,omitempty"`

	// snapshot num not in
	SnapshotNumNotIn []int32 `json:"snapshot_num_not_in,omitempty"`

	// stripe num
	StripeNum *int32 `json:"stripe_num,omitempty"`

	// stripe num gt
	StripeNumGt *int32 `json:"stripe_num_gt,omitempty"`

	// stripe num gte
	StripeNumGte *int32 `json:"stripe_num_gte,omitempty"`

	// stripe num in
	StripeNumIn []int32 `json:"stripe_num_in,omitempty"`

	// stripe num lt
	StripeNumLt *int32 `json:"stripe_num_lt,omitempty"`

	// stripe num lte
	StripeNumLte *int32 `json:"stripe_num_lte,omitempty"`

	// stripe num not
	StripeNumNot *int32 `json:"stripe_num_not,omitempty"`

	// stripe num not in
	StripeNumNotIn []int32 `json:"stripe_num_not_in,omitempty"`

	// stripe size
	StripeSize *int64 `json:"stripe_size,omitempty"`

	// stripe size gt
	StripeSizeGt *int64 `json:"stripe_size_gt,omitempty"`

	// stripe size gte
	StripeSizeGte *int64 `json:"stripe_size_gte,omitempty"`

	// stripe size in
	StripeSizeIn []int64 `json:"stripe_size_in,omitempty"`

	// stripe size lt
	StripeSizeLt *int64 `json:"stripe_size_lt,omitempty"`

	// stripe size lte
	StripeSizeLte *int64 `json:"stripe_size_lte,omitempty"`

	// stripe size not
	StripeSizeNot *int64 `json:"stripe_size_not,omitempty"`

	// stripe size not in
	StripeSizeNotIn []int64 `json:"stripe_size_not_in,omitempty"`

	// thin provision
	ThinProvision *bool `json:"thin_provision,omitempty"`

	// thin provision not
	ThinProvisionNot *bool `json:"thin_provision_not,omitempty"`

	// unique logical size
	UniqueLogicalSize *float64 `json:"unique_logical_size,omitempty"`

	// unique logical size gt
	UniqueLogicalSizeGt *float64 `json:"unique_logical_size_gt,omitempty"`

	// unique logical size gte
	UniqueLogicalSizeGte *float64 `json:"unique_logical_size_gte,omitempty"`

	// unique logical size in
	UniqueLogicalSizeIn []float64 `json:"unique_logical_size_in,omitempty"`

	// unique logical size lt
	UniqueLogicalSizeLt *float64 `json:"unique_logical_size_lt,omitempty"`

	// unique logical size lte
	UniqueLogicalSizeLte *float64 `json:"unique_logical_size_lte,omitempty"`

	// unique logical size not
	UniqueLogicalSizeNot *float64 `json:"unique_logical_size_not,omitempty"`

	// unique logical size not in
	UniqueLogicalSizeNotIn []float64 `json:"unique_logical_size_not_in,omitempty"`

	// unique size
	UniqueSize *int64 `json:"unique_size,omitempty"`

	// unique size gt
	UniqueSizeGt *int64 `json:"unique_size_gt,omitempty"`

	// unique size gte
	UniqueSizeGte *int64 `json:"unique_size_gte,omitempty"`

	// unique size in
	UniqueSizeIn []int64 `json:"unique_size_in,omitempty"`

	// unique size lt
	UniqueSizeLt *int64 `json:"unique_size_lt,omitempty"`

	// unique size lte
	UniqueSizeLte *int64 `json:"unique_size_lte,omitempty"`

	// unique size not
	UniqueSizeNot *int64 `json:"unique_size_not,omitempty"`

	// unique size not in
	UniqueSizeNotIn []int64 `json:"unique_size_not_in,omitempty"`

	// zbs volume id
	ZbsVolumeID *string `json:"zbs_volume_id,omitempty"`

	// zbs volume id contains
	ZbsVolumeIDContains *string `json:"zbs_volume_id_contains,omitempty"`

	// zbs volume id ends with
	ZbsVolumeIDEndsWith *string `json:"zbs_volume_id_ends_with,omitempty"`

	// zbs volume id gt
	ZbsVolumeIDGt *string `json:"zbs_volume_id_gt,omitempty"`

	// zbs volume id gte
	ZbsVolumeIDGte *string `json:"zbs_volume_id_gte,omitempty"`

	// zbs volume id in
	ZbsVolumeIDIn []string `json:"zbs_volume_id_in,omitempty"`

	// zbs volume id lt
	ZbsVolumeIDLt *string `json:"zbs_volume_id_lt,omitempty"`

	// zbs volume id lte
	ZbsVolumeIDLte *string `json:"zbs_volume_id_lte,omitempty"`

	// zbs volume id not
	ZbsVolumeIDNot *string `json:"zbs_volume_id_not,omitempty"`

	// zbs volume id not contains
	ZbsVolumeIDNotContains *string `json:"zbs_volume_id_not_contains,omitempty"`

	// zbs volume id not ends with
	ZbsVolumeIDNotEndsWith *string `json:"zbs_volume_id_not_ends_with,omitempty"`

	// zbs volume id not in
	ZbsVolumeIDNotIn []string `json:"zbs_volume_id_not_in,omitempty"`

	// zbs volume id not starts with
	ZbsVolumeIDNotStartsWith *string `json:"zbs_volume_id_not_starts_with,omitempty"`

	// zbs volume id starts with
	ZbsVolumeIDStartsWith *string `json:"zbs_volume_id_starts_with,omitempty"`
}

// Validate validates this iscsi lun where input
func (m *IscsiLunWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsSome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiLunWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *IscsiLunWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *IscsiLunWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *IscsiLunWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *IscsiLunWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *IscsiLunWhereInput) validateIscsiTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.IscsiTarget) { // not required
		return nil
	}

	if m.IscsiTarget != nil {
		if err := m.IscsiTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iscsi_target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iscsi_target")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) validateLabelsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsEvery) { // not required
		return nil
	}

	if m.LabelsEvery != nil {
		if err := m.LabelsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_every")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) validateLabelsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsNone) { // not required
		return nil
	}

	if m.LabelsNone != nil {
		if err := m.LabelsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_none")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) validateLabelsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsSome) { // not required
		return nil
	}

	if m.LabelsSome != nil {
		if err := m.LabelsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_some")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi lun where input based on the context it is used
func (m *IscsiLunWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiLunWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiLunWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiLunWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiLunWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *IscsiLunWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *IscsiLunWhereInput) contextValidateIscsiTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.IscsiTarget != nil {
		if err := m.IscsiTarget.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iscsi_target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iscsi_target")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) contextValidateLabelsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsEvery != nil {
		if err := m.LabelsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_every")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) contextValidateLabelsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsNone != nil {
		if err := m.LabelsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_none")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLunWhereInput) contextValidateLabelsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsSome != nil {
		if err := m.LabelsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_some")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiLunWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiLunWhereInput) UnmarshalBinary(b []byte) error {
	var res IscsiLunWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

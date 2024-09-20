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

// LabelOrderByInput label order by input
//
// swagger:model LabelOrderByInput
type LabelOrderByInput string

func NewLabelOrderByInput(value LabelOrderByInput) *LabelOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated LabelOrderByInput.
func (m LabelOrderByInput) Pointer() *LabelOrderByInput {
	return &m
}

const (

	// LabelOrderByInputClusterNumASC captures enum value "cluster_num_ASC"
	LabelOrderByInputClusterNumASC LabelOrderByInput = "cluster_num_ASC"

	// LabelOrderByInputClusterNumDESC captures enum value "cluster_num_DESC"
	LabelOrderByInputClusterNumDESC LabelOrderByInput = "cluster_num_DESC"

	// LabelOrderByInputContentLibraryImageNumASC captures enum value "content_library_image_num_ASC"
	LabelOrderByInputContentLibraryImageNumASC LabelOrderByInput = "content_library_image_num_ASC"

	// LabelOrderByInputContentLibraryImageNumDESC captures enum value "content_library_image_num_DESC"
	LabelOrderByInputContentLibraryImageNumDESC LabelOrderByInput = "content_library_image_num_DESC"

	// LabelOrderByInputContentLibraryVMTemplateNumASC captures enum value "content_library_vm_template_num_ASC"
	LabelOrderByInputContentLibraryVMTemplateNumASC LabelOrderByInput = "content_library_vm_template_num_ASC"

	// LabelOrderByInputContentLibraryVMTemplateNumDESC captures enum value "content_library_vm_template_num_DESC"
	LabelOrderByInputContentLibraryVMTemplateNumDESC LabelOrderByInput = "content_library_vm_template_num_DESC"

	// LabelOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	LabelOrderByInputCreatedAtASC LabelOrderByInput = "createdAt_ASC"

	// LabelOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	LabelOrderByInputCreatedAtDESC LabelOrderByInput = "createdAt_DESC"

	// LabelOrderByInputDatacenterNumASC captures enum value "datacenter_num_ASC"
	LabelOrderByInputDatacenterNumASC LabelOrderByInput = "datacenter_num_ASC"

	// LabelOrderByInputDatacenterNumDESC captures enum value "datacenter_num_DESC"
	LabelOrderByInputDatacenterNumDESC LabelOrderByInput = "datacenter_num_DESC"

	// LabelOrderByInputDiskNumASC captures enum value "disk_num_ASC"
	LabelOrderByInputDiskNumASC LabelOrderByInput = "disk_num_ASC"

	// LabelOrderByInputDiskNumDESC captures enum value "disk_num_DESC"
	LabelOrderByInputDiskNumDESC LabelOrderByInput = "disk_num_DESC"

	// LabelOrderByInputElfImageNumASC captures enum value "elf_image_num_ASC"
	LabelOrderByInputElfImageNumASC LabelOrderByInput = "elf_image_num_ASC"

	// LabelOrderByInputElfImageNumDESC captures enum value "elf_image_num_DESC"
	LabelOrderByInputElfImageNumDESC LabelOrderByInput = "elf_image_num_DESC"

	// LabelOrderByInputGpuDeviceNumASC captures enum value "gpu_device_num_ASC"
	LabelOrderByInputGpuDeviceNumASC LabelOrderByInput = "gpu_device_num_ASC"

	// LabelOrderByInputGpuDeviceNumDESC captures enum value "gpu_device_num_DESC"
	LabelOrderByInputGpuDeviceNumDESC LabelOrderByInput = "gpu_device_num_DESC"

	// LabelOrderByInputHostNumASC captures enum value "host_num_ASC"
	LabelOrderByInputHostNumASC LabelOrderByInput = "host_num_ASC"

	// LabelOrderByInputHostNumDESC captures enum value "host_num_DESC"
	LabelOrderByInputHostNumDESC LabelOrderByInput = "host_num_DESC"

	// LabelOrderByInputIDASC captures enum value "id_ASC"
	LabelOrderByInputIDASC LabelOrderByInput = "id_ASC"

	// LabelOrderByInputIDDESC captures enum value "id_DESC"
	LabelOrderByInputIDDESC LabelOrderByInput = "id_DESC"

	// LabelOrderByInputIscsiLunNumASC captures enum value "iscsi_lun_num_ASC"
	LabelOrderByInputIscsiLunNumASC LabelOrderByInput = "iscsi_lun_num_ASC"

	// LabelOrderByInputIscsiLunNumDESC captures enum value "iscsi_lun_num_DESC"
	LabelOrderByInputIscsiLunNumDESC LabelOrderByInput = "iscsi_lun_num_DESC"

	// LabelOrderByInputIscsiLunSnapshotNumASC captures enum value "iscsi_lun_snapshot_num_ASC"
	LabelOrderByInputIscsiLunSnapshotNumASC LabelOrderByInput = "iscsi_lun_snapshot_num_ASC"

	// LabelOrderByInputIscsiLunSnapshotNumDESC captures enum value "iscsi_lun_snapshot_num_DESC"
	LabelOrderByInputIscsiLunSnapshotNumDESC LabelOrderByInput = "iscsi_lun_snapshot_num_DESC"

	// LabelOrderByInputIscsiTargetNumASC captures enum value "iscsi_target_num_ASC"
	LabelOrderByInputIscsiTargetNumASC LabelOrderByInput = "iscsi_target_num_ASC"

	// LabelOrderByInputIscsiTargetNumDESC captures enum value "iscsi_target_num_DESC"
	LabelOrderByInputIscsiTargetNumDESC LabelOrderByInput = "iscsi_target_num_DESC"

	// LabelOrderByInputKeyASC captures enum value "key_ASC"
	LabelOrderByInputKeyASC LabelOrderByInput = "key_ASC"

	// LabelOrderByInputKeyDESC captures enum value "key_DESC"
	LabelOrderByInputKeyDESC LabelOrderByInput = "key_DESC"

	// LabelOrderByInputNicNumASC captures enum value "nic_num_ASC"
	LabelOrderByInputNicNumASC LabelOrderByInput = "nic_num_ASC"

	// LabelOrderByInputNicNumDESC captures enum value "nic_num_DESC"
	LabelOrderByInputNicNumDESC LabelOrderByInput = "nic_num_DESC"

	// LabelOrderByInputSystemVlanNumASC captures enum value "system_vlan_num_ASC"
	LabelOrderByInputSystemVlanNumASC LabelOrderByInput = "system_vlan_num_ASC"

	// LabelOrderByInputSystemVlanNumDESC captures enum value "system_vlan_num_DESC"
	LabelOrderByInputSystemVlanNumDESC LabelOrderByInput = "system_vlan_num_DESC"

	// LabelOrderByInputTotalNumASC captures enum value "total_num_ASC"
	LabelOrderByInputTotalNumASC LabelOrderByInput = "total_num_ASC"

	// LabelOrderByInputTotalNumDESC captures enum value "total_num_DESC"
	LabelOrderByInputTotalNumDESC LabelOrderByInput = "total_num_DESC"

	// LabelOrderByInputValueASC captures enum value "value_ASC"
	LabelOrderByInputValueASC LabelOrderByInput = "value_ASC"

	// LabelOrderByInputValueDESC captures enum value "value_DESC"
	LabelOrderByInputValueDESC LabelOrderByInput = "value_DESC"

	// LabelOrderByInputVdsNumASC captures enum value "vds_num_ASC"
	LabelOrderByInputVdsNumASC LabelOrderByInput = "vds_num_ASC"

	// LabelOrderByInputVdsNumDESC captures enum value "vds_num_DESC"
	LabelOrderByInputVdsNumDESC LabelOrderByInput = "vds_num_DESC"

	// LabelOrderByInputVMNumASC captures enum value "vm_num_ASC"
	LabelOrderByInputVMNumASC LabelOrderByInput = "vm_num_ASC"

	// LabelOrderByInputVMNumDESC captures enum value "vm_num_DESC"
	LabelOrderByInputVMNumDESC LabelOrderByInput = "vm_num_DESC"

	// LabelOrderByInputVMSnapshotNumASC captures enum value "vm_snapshot_num_ASC"
	LabelOrderByInputVMSnapshotNumASC LabelOrderByInput = "vm_snapshot_num_ASC"

	// LabelOrderByInputVMSnapshotNumDESC captures enum value "vm_snapshot_num_DESC"
	LabelOrderByInputVMSnapshotNumDESC LabelOrderByInput = "vm_snapshot_num_DESC"

	// LabelOrderByInputVMTemplateNumASC captures enum value "vm_template_num_ASC"
	LabelOrderByInputVMTemplateNumASC LabelOrderByInput = "vm_template_num_ASC"

	// LabelOrderByInputVMTemplateNumDESC captures enum value "vm_template_num_DESC"
	LabelOrderByInputVMTemplateNumDESC LabelOrderByInput = "vm_template_num_DESC"

	// LabelOrderByInputVMVlanNumASC captures enum value "vm_vlan_num_ASC"
	LabelOrderByInputVMVlanNumASC LabelOrderByInput = "vm_vlan_num_ASC"

	// LabelOrderByInputVMVlanNumDESC captures enum value "vm_vlan_num_DESC"
	LabelOrderByInputVMVlanNumDESC LabelOrderByInput = "vm_vlan_num_DESC"

	// LabelOrderByInputVMVolumeNumASC captures enum value "vm_volume_num_ASC"
	LabelOrderByInputVMVolumeNumASC LabelOrderByInput = "vm_volume_num_ASC"

	// LabelOrderByInputVMVolumeNumDESC captures enum value "vm_volume_num_DESC"
	LabelOrderByInputVMVolumeNumDESC LabelOrderByInput = "vm_volume_num_DESC"

	// LabelOrderByInputVMVolumeSnapshotNumASC captures enum value "vm_volume_snapshot_num_ASC"
	LabelOrderByInputVMVolumeSnapshotNumASC LabelOrderByInput = "vm_volume_snapshot_num_ASC"

	// LabelOrderByInputVMVolumeSnapshotNumDESC captures enum value "vm_volume_snapshot_num_DESC"
	LabelOrderByInputVMVolumeSnapshotNumDESC LabelOrderByInput = "vm_volume_snapshot_num_DESC"
)

// for schema
var labelOrderByInputEnum []interface{}

func init() {
	var res []LabelOrderByInput
	if err := json.Unmarshal([]byte(`["cluster_num_ASC","cluster_num_DESC","content_library_image_num_ASC","content_library_image_num_DESC","content_library_vm_template_num_ASC","content_library_vm_template_num_DESC","createdAt_ASC","createdAt_DESC","datacenter_num_ASC","datacenter_num_DESC","disk_num_ASC","disk_num_DESC","elf_image_num_ASC","elf_image_num_DESC","gpu_device_num_ASC","gpu_device_num_DESC","host_num_ASC","host_num_DESC","id_ASC","id_DESC","iscsi_lun_num_ASC","iscsi_lun_num_DESC","iscsi_lun_snapshot_num_ASC","iscsi_lun_snapshot_num_DESC","iscsi_target_num_ASC","iscsi_target_num_DESC","key_ASC","key_DESC","nic_num_ASC","nic_num_DESC","system_vlan_num_ASC","system_vlan_num_DESC","total_num_ASC","total_num_DESC","value_ASC","value_DESC","vds_num_ASC","vds_num_DESC","vm_num_ASC","vm_num_DESC","vm_snapshot_num_ASC","vm_snapshot_num_DESC","vm_template_num_ASC","vm_template_num_DESC","vm_vlan_num_ASC","vm_vlan_num_DESC","vm_volume_num_ASC","vm_volume_num_DESC","vm_volume_snapshot_num_ASC","vm_volume_snapshot_num_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		labelOrderByInputEnum = append(labelOrderByInputEnum, v)
	}
}

func (m LabelOrderByInput) validateLabelOrderByInputEnum(path, location string, value LabelOrderByInput) error {
	if err := validate.EnumCase(path, location, value, labelOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this label order by input
func (m LabelOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLabelOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this label order by input based on context it is used
func (m LabelOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

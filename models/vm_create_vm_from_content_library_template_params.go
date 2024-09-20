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

// VMCreateVMFromContentLibraryTemplateParams Vm create Vm from content library template params
//
// swagger:model VmCreateVmFromContentLibraryTemplateParams
type VMCreateVMFromContentLibraryTemplateParams struct {

	// cloud init
	CloudInit *TemplateCloudInit `json:"cloud_init,omitempty"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// cpu cores
	CPUCores *int32 `json:"cpu_cores,omitempty"`

	// cpu sockets
	CPUSockets *int32 `json:"cpu_sockets,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// disk operate
	DiskOperate *VMDiskOperate `json:"disk_operate,omitempty"`

	// firmware
	Firmware *VMFirmware `json:"firmware,omitempty"`

	// folder id
	FolderID *string `json:"folder_id,omitempty"`

	// gpu devices
	GpuDevices []*VMGpuOperationParams `json:"gpu_devices,omitempty"`

	// guest os type
	GuestOsType *VMGuestsOperationSystem `json:"guest_os_type,omitempty"`

	// ha
	Ha *bool `json:"ha,omitempty"`

	// host id
	HostID *string `json:"host_id,omitempty"`

	// io policy
	IoPolicy *VMDiskIoPolicy `json:"io_policy,omitempty"`

	// is full copy
	// Required: true
	IsFullCopy *bool `json:"is_full_copy"`

	// max bandwidth
	MaxBandwidth *int64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy *VMDiskIoRestrictType `json:"max_bandwidth_policy,omitempty"`

	// max bandwidth unit
	MaxBandwidthUnit *BPSUnit `json:"max_bandwidth_unit,omitempty"`

	// max iops
	MaxIops *int64 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy *VMDiskIoRestrictType `json:"max_iops_policy,omitempty"`

	// memory
	Memory *int64 `json:"memory,omitempty"`

	// memory unit
	MemoryUnit *ByteUnit `json:"memory_unit,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// owner
	Owner *VMOwnerParams `json:"owner,omitempty"`

	// pci nics
	PciNics *NicWhereInput `json:"pci_nics,omitempty"`

	// status
	Status *VMStatus `json:"status,omitempty"`

	// template id
	// Required: true
	TemplateID *string `json:"template_id"`

	// vcpu
	Vcpu *int32 `json:"vcpu,omitempty"`

	// vm nics
	VMNics []*VMNicParams `json:"vm_nics,omitempty"`

	// vm placement group
	VMPlacementGroup *VMPlacementGroupWhereInput `json:"vm_placement_group,omitempty"`
}

// Validate validates this Vm create Vm from content library template params
func (m *VMCreateVMFromContentLibraryTemplateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudInit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskOperate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpuDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsFullCopy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePciNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMPlacementGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateCloudInit(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudInit) { // not required
		return nil
	}

	if m.CloudInit != nil {
		if err := m.CloudInit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_init")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_init")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateDiskOperate(formats strfmt.Registry) error {
	if swag.IsZero(m.DiskOperate) { // not required
		return nil
	}

	if m.DiskOperate != nil {
		if err := m.DiskOperate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_operate")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateFirmware(formats strfmt.Registry) error {
	if swag.IsZero(m.Firmware) { // not required
		return nil
	}

	if m.Firmware != nil {
		if err := m.Firmware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateGpuDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.GpuDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.GpuDevices); i++ {
		if swag.IsZero(m.GpuDevices[i]) { // not required
			continue
		}

		if m.GpuDevices[i] != nil {
			if err := m.GpuDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpu_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpu_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateGuestOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.GuestOsType) { // not required
		return nil
	}

	if m.GuestOsType != nil {
		if err := m.GuestOsType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guest_os_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("guest_os_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateIoPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IoPolicy) { // not required
		return nil
	}

	if m.IoPolicy != nil {
		if err := m.IoPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateIsFullCopy(formats strfmt.Registry) error {

	if err := validate.Required("is_full_copy", "body", m.IsFullCopy); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthPolicy) { // not required
		return nil
	}

	if m.MaxBandwidthPolicy != nil {
		if err := m.MaxBandwidthPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateMaxBandwidthUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthUnit) { // not required
		return nil
	}

	if m.MaxBandwidthUnit != nil {
		if err := m.MaxBandwidthUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_unit")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateMaxIopsPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxIopsPolicy) { // not required
		return nil
	}

	if m.MaxIopsPolicy != nil {
		if err := m.MaxIopsPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_iops_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateMemoryUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.MemoryUnit) { // not required
		return nil
	}

	if m.MemoryUnit != nil {
		if err := m.MemoryUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory_unit")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validatePciNics(formats strfmt.Registry) error {
	if swag.IsZero(m.PciNics) { // not required
		return nil
	}

	if m.PciNics != nil {
		if err := m.PciNics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pci_nics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pci_nics")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("template_id", "body", m.TemplateID); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateVMNics(formats strfmt.Registry) error {
	if swag.IsZero(m.VMNics) { // not required
		return nil
	}

	for i := 0; i < len(m.VMNics); i++ {
		if swag.IsZero(m.VMNics[i]) { // not required
			continue
		}

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) validateVMPlacementGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.VMPlacementGroup) { // not required
		return nil
	}

	if m.VMPlacementGroup != nil {
		if err := m.VMPlacementGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_placement_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_placement_group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm create Vm from content library template params based on the context it is used
func (m *VMCreateVMFromContentLibraryTemplateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudInit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiskOperate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpuDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGuestOsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemoryUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePciNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMPlacementGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateCloudInit(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudInit != nil {
		if err := m.CloudInit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_init")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_init")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateDiskOperate(ctx context.Context, formats strfmt.Registry) error {

	if m.DiskOperate != nil {
		if err := m.DiskOperate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_operate")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

	if m.Firmware != nil {
		if err := m.Firmware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateGpuDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GpuDevices); i++ {

		if m.GpuDevices[i] != nil {
			if err := m.GpuDevices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpu_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpu_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateGuestOsType(ctx context.Context, formats strfmt.Registry) error {

	if m.GuestOsType != nil {
		if err := m.GuestOsType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guest_os_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("guest_os_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateIoPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IoPolicy != nil {
		if err := m.IoPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxBandwidthPolicy != nil {
		if err := m.MaxBandwidthPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateMaxBandwidthUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxBandwidthUnit != nil {
		if err := m.MaxBandwidthUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_unit")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxIopsPolicy != nil {
		if err := m.MaxIopsPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_iops_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateMemoryUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.MemoryUnit != nil {
		if err := m.MemoryUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory_unit")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidatePciNics(ctx context.Context, formats strfmt.Registry) error {

	if m.PciNics != nil {
		if err := m.PciNics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pci_nics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pci_nics")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMNics); i++ {

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateParams) contextValidateVMPlacementGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.VMPlacementGroup != nil {
		if err := m.VMPlacementGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_placement_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_placement_group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMCreateVMFromContentLibraryTemplateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromContentLibraryTemplateParams) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromContentLibraryTemplateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

package model

import "time"

type CreateVmSnapshotRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	ServerId string                 `json:"server_id"`
	Params   CreateVmSnapshotParams `json:"params"`
}

type CreateVmSnapshotParams struct {
	CreateVmSnapshot CreateVmSnapshot `json:"createImage"`
}

type CreateVmSnapshot struct {
	Name             string `json:"name"`
	InstanceSnapshot string `json:"instance_snapshot"` //true,false
}

type CreateVmSnapshotResponse struct {
	ImageId string `json:"imageId"`
}

type RevertVmSnapshotRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	ServerId string                 `json:"server_id"`
	Params   RevertVmSnapshotParams `json:"params"`
}

type RevertVmSnapshotParams struct {
	RevertVmSnapshot RevertVmSnapshot `json:"rebuild"`
}

type RevertVmSnapshot struct {
	Description string `json:"description,omitempty"`
	ImageRef    string `json:"imageRef"` //快照id
}

type DeleteVmSnapshotRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	Params   DeleteVmSnapshotParams `json:"params"`
}

type DeleteVmSnapshotParams struct {
	Images          []string `json:"images"`
	IsSnapshotImage string   `json:"is_snapshot_image"`
	AvailableZone   string   `json:"available_zone"`
	Region          string   `json:"region"`
}
type QueryVmSnapshotResponse struct {
	SupportLiveResize        string        `json:"__support_live_resize"`
	ImgSignatureHashMethod   string        `json:"img_signature_hash_method"`
	MinDisk                  int           `json:"min_disk"`
	ImageSourceType          string        `json:"__image_source_type"`
	SupportKvmVgpu           string        `json:"__support_kvm_vgpu"`
	ContainerFormat          interface{}   `json:"container_format"`
	ImgSignatureKeyType      string        `json:"img_signature_key_type"`
	SupportKvmNpu            string        `json:"__support_kvm_npu"`
	Protected                bool          `json:"protected"`
	Id                       string        `json:"id"`
	Isregistered             string        `json:"__isregistered"`
	ImgSignature             string        `json:"img_signature"`
	VirtualSize              string        `json:"__virtual_size"`
	HwFirmwareType           string        `json:"hw_firmware_type"`
	OsType                   string        `json:"__os_type"`
	Visibility               string        `json:"visibility"`
	QuickStart               string        `json:"__quick_start"`
	VirtualEnvType           string        `json:"virtual_env_type"`
	SupportKvmGpu            string        `json:"__support_kvm_gpu"`
	SupportKvm               string        `json:"__support_kvm"`
	SupportKvmNvmeSsd        string        `json:"__support_kvm_nvme_ssd"`
	HwDiskBus                string        `json:"hw_disk_bus"`
	ImgDigestValue           string        `json:"img_digest_value"`
	Tags                     []interface{} `json:"tags"`
	Platform                 string        `json:"__platform"`
	OriDiskFormat            string        `json:"ori_disk_format"`
	Size                     int           `json:"size"`
	OsVersion                string        `json:"__os_version"`
	Name                     string        `json:"name"`
	AdminEncrypted           string        `json:"__admin_encrypted"`
	BdmV2                    string        `json:"bdm_v2"`
	Status                   string        `json:"status"`
	Schema                   string        `json:"schema"`
	BaseImageRef             string        `json:"base_image_ref"`
	SystemEncrypted          string        `json:"__system_encrypted"`
	BlockDeviceMapping       string        `json:"block_device_mapping"`
	CreatedAt                time.Time     `json:"created_at"`
	SnapshotFromInstanceName string        `json:"__snapshot_from_instance_name"`
	SnapshotFromInstance     string        `json:"__snapshot_from_instance"`
	File                     string        `json:"file"`
	UpdatedAt                time.Time     `json:"updated_at"`
	Checksum                 interface{}   `json:"checksum"`
	SupportStaticIp          string        `json:"__support_static_ip"`
	FileFormat               string        `json:"file_format"`
	MinRam                   int           `json:"min_ram"`
	Architecture             string        `json:"architecture"`
	Owner                    string        `json:"owner"`
	Imagetype                string        `json:"__imagetype"`
	RootDeviceName           string        `json:"root_device_name"`
	Cloudinit                string        `json:"cloudinit"`
	FileName                 string        `json:"file_name"`
	AccountCode              string        `json:"__account_code"`
	SupportKvmHana           string        `json:"__support_kvm_hana"`
	IsAutoConfig             string        `json:"is_auto_config"`
	OsBit                    string        `json:"__os_bit"`
	DiskFormat               string        `json:"disk_format"`
	Self                     string        `json:"self"`
	Describe                 string        `json:"describe"`
	VirtualSize1             interface{}   `json:"virtual_size"`
}

type CreateVolumeSnapshotRequest struct {
	Domain   string                     `json:"domain"`
	Token    string                     `json:"token"`
	TenantId string                     `json:"tenantId"`
	Params   CreateVolumeSnapshotParams `json:"params"`
}

type CreateVolumeSnapshotParams struct {
	CreateVolumeSnapshot CreateVolumeSnapshot `json:"snapshot"`
}

type CreateVolumeSnapshot struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	VolumeId    string `json:"volume_id"`
	Force       bool   `json:"force"`
}

type CreateVolumeSnapshotResponse struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VolumeId    string `json:"volume_id"`
	Status      string `json:"status"`
	Size        string `json:"size"`
	Metadata    struct {
		SnapshotFromInstance string `json:"__snapshot_from_instance"`
		SnapshotWwn          string `json:"snapshot_wwn"`
		SystemEnableActive   string `json:"__system__enableActive"`
	} `json:"metadata"`
	OsExtendedSnapshotAttributesProviderLocation string `json:"os-extended-snapshot-attributes:provider_location"`
	OsVendorExtAttrVendorDriverMetadata          struct {
	} `json:"os-vendor-ext-attr:vendor_driver_metadata"`
	OsExtendedSnapshotAttributesProjectId string `json:"os-extended-snapshot-attributes:project_id"`
	OsExtendedSnapshotAttributesProgress  string `json:"os-extended-snapshot-attributes:progress"`
}

type RevertVolumeSnapshotRequest struct {
	Domain     string                     `json:"domain"`
	Token      string                     `json:"token"`
	TenantId   string                     `json:"tenantId"`
	SnapshotId string                     `json:"snapshotId"`
	Params     RevertVolumeSnapshotParams `json:"params"`
}

type RevertVolumeSnapshotParams struct {
	RevertVolumeSnapshot RevertVolumeSnapshot `json:"rollback"`
}

type RevertVolumeSnapshot struct {
	Name     string `json:"name,omitempty"`
	VolumeId string `json:"volume_id"`
}

type DeleteVolumeSnapshotRequest struct {
	Domain     string `json:"domain"`
	Token      string `json:"token"`
	TenantId   string `json:"tenantId"`
	SnapshotId string `json:"snapshotId"`
}

type QuerySnapshotRequest struct {
	Domain     string `json:"domain"`
	Token      string `json:"token"`
	TenantId   string `json:"tenantId"`
	Id         string `json:"id,omitempty"`
	InstanceId string `json:"__snapshot_from_instance,omitempty"`
	VolumeId   string `json:"volume_id,omitempty"`
}

type QueryVolumeSnapshotResponse struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VolumeId    string `json:"volume_id"`
	Status      string `json:"status"`
	Size        int    `json:"size"`
	Metadata    struct {
		SnapshotFromInstance string `json:"__snapshot_from_instance"`
		SnapshotWwn          string `json:"snapshot_wwn"`
		SystemEnableActive   string `json:"__system__enableActive"`
	} `json:"metadata"`
	OsExtendedSnapshotAttributesProviderLocation string `json:"os-extended-snapshot-attributes:provider_location"`
	OsVendorExtAttrVendorDriverMetadata          struct {
	} `json:"os-vendor-ext-attr:vendor_driver_metadata"`
	OsExtendedSnapshotAttributesProjectId string `json:"os-extended-snapshot-attributes:project_id"`
	OsExtendedSnapshotAttributesProgress  string `json:"os-extended-snapshot-attributes:progress"`
}

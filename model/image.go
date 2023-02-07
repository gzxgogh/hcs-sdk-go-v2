package model

import "time"

type CreateImageRequest struct {
	Domain   string      `json:"domain"`
	Token    string      `json:"token"`
	TenantId string      `json:"tenantId"`
	Params   CreateImage `json:"params"`
}

type CreateImage struct {
	Name        string `json:"name"`
	InstanceId  string `json:"instance_id"`
	Description string `json:"description,omitempty"`
}

type DeleteImageRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryImageRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id,omitempty"`
}

type QueryImageResponse struct {
	Architecture           string    `json:"architecture"`
	File                   string    `json:"file"`
	Owner                  string    `json:"owner"`
	Id                     string    `json:"id"`
	Size                   int64     `json:"size"`
	Self                   string    `json:"self"`
	Status                 string    `json:"status"`
	Tags                   []string  `json:"tags"`
	Visibility             string    `json:"visibility"`
	Name                   string    `json:"name"`
	Checksum               string    `json:"checksum"`
	Protected              bool      `json:"protected"`
	ContainerFormat        string    `json:"container_format"`
	MinRam                 int       `json:"min_ram"`
	UpdatedAt              time.Time `json:"updated_at"`
	OsBit                  string    `json:"__os_bit"`
	OsVersion              string    `json:"__os_version"`
	Description            string    `json:"__description"`
	DiskFormat             string    `json:"disk_format"`
	Isregistered           string    `json:"__isregistered"`
	Platform               string    `json:"__platform"`
	OsType                 string    `json:"__os_type"`
	MinDisk                int       `json:"min_disk"`
	VirtualEnvType         string    `json:"virtual_env_type"`
	ImageSourceType        string    `json:"__image_source_type"`
	Imagetype              string    `json:"__imagetype"`
	CreatedAt              time.Time `json:"created_at"`
	VirtualSize1           int       `json:"virtual_size"`
	ImageSize              int       `json:"__image_size"`
	DataOrigin             string    `json:"__data_origin"`
	SupportKvm             string    `json:"__support_kvm"`
	SupportXen             string    `json:"__support_xen"`
	SupportLargememory     string    `json:"__support_largememory"`
	SupportDiskintensive   string    `json:"__support_diskintensive"`
	SupportHighperformance string    `json:"__support_highperformance"`
	SupportKvmGpuType      string    `json:"__support_kvm_gpu_type"`
	SupportXenHana         string    `json:"__support_xen_hana"`
	SupportKvmInfiniband   string    `json:"__support_kvm_infiniband"`
	IsConfigInit           string    `json:"__is_config_init"`
	SystemSupportMarket    bool      `json:"__system_support_market"`
}

type ImageMetadata struct {
	SupportLiveResize      string `json:"__support_live_resize"`
	SystemEncrypted        string `json:"__system_encrypted"`
	ImgSignatureHashMethod string `json:"img_signature_hash_method"`
	ImageSourceType        string `json:"__image_source_type"`
	SupportKvmVgpu         string `json:"__support_kvm_vgpu"`
	ImgSignatureKeyType    string `json:"img_signature_key_type"`
	SupportKvmNpu          string `json:"__support_kvm_npu"`
	SupportStaticIp        string `json:"__support_static_ip"`
	Isregistered           string `json:"__isregistered"`
	FileFormat             string `json:"file_format"`
	Architecture           string `json:"architecture"`
	ImgSignature           string `json:"img_signature"`
	VirtualSize            string `json:"__virtual_size"`
	HwFirmwareType         string `json:"hw_firmware_type"`
	OsType                 string `json:"__os_type"`
	Imagetype              string `json:"__imagetype"`
	FileName               string `json:"file_name"`
	Cloudinit              string `json:"cloudinit"`
	QuickStart             string `json:"__quick_start"`
	VirtualEnvType         string `json:"virtual_env_type"`
	SupportKvmGpu          string `json:"__support_kvm_gpu"`
	SupportKvm             string `json:"__support_kvm"`
	AccountCode            string `json:"__account_code"`
	SupportKvmNvmeSsd      string `json:"__support_kvm_nvme_ssd"`
	HwDiskBus              string `json:"hw_disk_bus"`
	ImgDigestValue         string `json:"img_digest_value"`
	Platform               string `json:"__platform"`
	SupportKvmHana         string `json:"__support_kvm_hana"`
	OriDiskFormat          string `json:"ori_disk_format"`
	IsAutoConfig           string `json:"is_auto_config"`
	OsBit                  string `json:"__os_bit"`
	OsVersion              string `json:"__os_version"`
	AdminEncrypted         string `json:"__admin_encrypted"`
	Describe               string `json:"describe"`
}

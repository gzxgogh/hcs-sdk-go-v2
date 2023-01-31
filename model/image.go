package model

import "time"

type QueryImageRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id,omitempty"`
}

type QueryImageResponse struct {
	OSEXTIMGSIZESize int64         `json:"OS-EXT-IMG-SIZE:size"`
	Metadata         ImageMetadata `json:"metadata"`
	Created          time.Time     `json:"created"`
	MinRam           int           `json:"minRam"`
	Name             string        `json:"name"`
	Progress         int           `json:"progress"`
	Id               string        `json:"id"`
	MinDisk          int           `json:"minDisk"`
	Updated          time.Time     `json:"updated"`
	Status           string        `json:"status"`
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

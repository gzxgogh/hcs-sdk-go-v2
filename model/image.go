package model

import "time"

type CreateVmImageRequest struct {
	Domain    string        `json:"domain"`
	Token     string        `json:"token"`
	TenantId  string        `json:"tenantId"`
	EcsDomain string        `json:"ecsDomain"`
	Params    CreateVmImage `json:"params"`
}

type CreateVmImage struct {
	Name        string `json:"name"`
	InstanceId  string `json:"instance_id"`
	Description string `json:"description,omitempty"`
}

type CreateVolumeImageRequest struct {
	Domain    string                  `json:"domain"`
	Token     string                  `json:"token"`
	TenantId  string                  `json:"tenantId"`
	EcsDomain string                  `json:"ecsDomain"`
	Params    CreateVolumeImageParams `json:"params"`
}

type CreateVolumeImageParams struct {
	Name       string      `json:"name,omitempty"`
	InstanceId string      `json:"instance_id"`
	DataImages []DataImage `json:"data_images"`
}

type DataImage struct {
	VolumeId string `json:"volume_id"`
}

type CreateVolumeImageResponse struct {
	ImageId string `json:"imageId"`
}

type DeleteImageRequest struct {
	Domain   string      `json:"domain"`
	Token    string      `json:"token"`
	TenantId string      `json:"tenantId"`
	Params   DeleteImage `json:"params"`
}

type DeleteImage struct {
	Images []string `json:"images"`
}

type UpdateImageRequest struct {
	Domain   string        `json:"domain"`
	Token    string        `json:"token"`
	TenantId string        `json:"tenantId"`
	ImageId  string        `json:"imageId"`
	Params   []UpdateImage `json:"params"`
}

type UpdateImage struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

type QueryImageRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
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
	VirtualSize            string    `json:"__virtual_size"`
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
	WholeImage             string    `json:"__whole_image"`
}

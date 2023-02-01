package model

type QueryVolumeTypeRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
}

type QueryVolumeTypeResponse struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	IsPublic   bool       `json:"is_public"`
	ExtraSpecs ExtraSpecs `json:"extra_specs"`
}

type ExtraSpecs struct {
	VolumeBackendName string `json:"volume_backend_name"`
	AvailabilityZone  string `json:"availability-zone"`
}

type CreateVolumeRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenantId"`
	Params   CreateVolumeParams `json:"params"`
}

type CreateVolumeParams struct {
	Volume Volume `json:"volume"`
}

type Volume struct {
	Size               int    `json:"size"`              //大小
	AvailabilityZone   string `json:"availability_zone"` //可用区域
	Name               string `json:"name"`              //名称
	Description        string `json:"description,omitempty"`
	Multiattach        bool   `json:"multiattach"`        //是指是否允许该系统卷同时挂载在不同的VM上
	ImageRef           string `json:"imageRef,omitempty"` //镜像id,必须指定volumeType为business_type_01
	VolumeType         string `json:"volume_type"`        //SAS,business_type_01,hhw
	ConsistencygroupId string `json:"consistencygroup_id,omitempty"`
	SourceVolid        string `json:"source_volid,omitempty"`
	SnapshotId         string `json:"snapshot_id,omitempty"` //快照id
}

type CreateVolumeResponse struct {
	Id               string        `json:"id"`
	Name             string        `json:"name"`
	AvailabilityZone string        `json:"availabilityZone"`
	Status           string        `json:"status"`
	Attachments      []interface{} `json:"attachments"`
}

type DeleteVolumeRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateVolumeRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenantId"`
	Id       string             `json:"id"`
	Params   UpdateVolumeParams `json:"params"`
}

type UpdateVolumeParams struct {
	UpdateVolume UpdateVolume `json:"volume"`
}

type UpdateVolume struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type QueryVolumeRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id,omitempty"`
}

type QueryVolumeResponse struct {
	Id                 string        `json:"id"`
	Name               string        `json:"name"`
	Status             string        `json:"status"`
	Attachments        []interface{} `json:"attachments"`
	AvailabilityZone   string        `json:"availability_zone"`
	SourceVolid        string        `json:"source_volid"`
	SnapshotId         string        `json:"snapshot_id"`
	Description        string        `json:"description"`
	CreatedAt          string        `json:"created_at"`
	VolumeType         string        `json:"volume_type"`
	Size               int           `json:"size"`
	ReplicationStatus  string        `json:"replication_status"`
	Encrypted          bool          `json:"encrypted"`
	UserId             string        `json:"user_id"`
	ConsistencygroupId string        `json:"consistencygroup_id"`
	Bootable           string        `json:"bootable"`
	Shareable          bool          `json:"shareable"`
	Multiattach        bool          `json:"multiattach"`
	Metadata           struct {
		LunWwn         string `json:"lun_wwn"`
		StorageType    string `json:"StorageType"`
		TakeOverLunWwn string `json:"take_over_lun_wwn"`
		Readonly       string `json:"readonly"`
	} `json:"metadata"`
}

type UpgradeVolumeRequest struct {
	Domain   string              `json:"domain"`
	Token    string              `json:"token"`
	TenantId string              `json:"tenantId"`
	Id       string              `json:"id"`
	Params   UpgradeVolumeParams `json:"params"`
}

type UpgradeVolumeParams struct {
	UpgradeVolume UpgradeVolume `json:"os-extend"`
}

type UpgradeVolume struct {
	Size int `json:"new_size"`
}

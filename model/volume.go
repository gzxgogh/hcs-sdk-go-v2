package model

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
	SnapshotId         string `json:"snapshot_id,omitempty"`
}

type CreateVolumeResponse struct {
	Id               string        `json:"id"`
	Name             string        `json:"name"`
	AvailabilityZone string        `json:"availabilityZone"`
	Status           string        `json:"status"`
	Attachments      []interface{} `json:"attachments"`
}

type T struct {
	Id               string        `json:"id"`
	Name             string        `json:"name"`
	Status           string        `json:"status"`
	Attachments      []interface{} `json:"attachments"`
	AvailabilityZone string        `json:"availability_zone"`
	SourceVolid      interface{}   `json:"source_volid"`
	SnapshotId       interface{}   `json:"snapshot_id"`
	Description      interface{}   `json:"description"`
	CreatedAt        string        `json:"created_at"`
	VolumeType       string        `json:"volume_type"`
	Size             int           `json:"size"`
	Metadata         struct {
	} `json:"metadata"`
	ReplicationStatus  string      `json:"replication_status"`
	Encrypted          bool        `json:"encrypted"`
	UserId             string      `json:"user_id"`
	ConsistencygroupId interface{} `json:"consistencygroup_id"`
	Bootable           string      `json:"bootable"`
	UpdatedAt          interface{} `json:"updated_at"`
	Shareable          bool        `json:"shareable"`
	Multiattach        bool        `json:"multiattach"`
}

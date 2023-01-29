package model

type CreateVolumeRequest struct {
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Volume   Volume `json:"volume"`
}

type Volume struct {
	Size             int    `json:"size"`              //大小
	AvailabilityZone string `json:"availability_zone"` //可用区域
	Name             string `json:"name"`              //名称
	Multiattach      bool   `json:"multiattach"`       //是指是否允许该系统卷同时挂载在不同的VM上
	ImageRef         string `json:"imageRef"`          //镜像id
}

type CreateVolumeResponse struct {
	Id               string        `json:"id"`
	Name             string        `json:"name"`
	AvailabilityZone string        `json:"availabilityZone"`
	Status           string        `json:"status"`
	Attachments      []interface{} `json:"attachments"`
}

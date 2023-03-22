package model

import "time"

var ObjTypeList = []ObjType{
	{
		Id:   "562958543355904",
		Name: "VM",
	},
	{
		Id:   "562967133290496",
		Name: "VOLUME",
	},
	{
		Id:   "562971428257792",
		Name: "EIP",
	},
	{
		Id:   "563001493028864",
		Name: "NIC",
	},
	{
		Id:   "562975723225088",
		Name: "BANDWIDTH",
	},
}

var VmIndicatorList = []Indicator{
	{
		Id:   562958543421441,
		Name: "CPU使用率",
		Unit: "%",
	},
	{
		Id:   562958543486979,
		Name: "内存使用率",
		Unit: "%",
	},
	{
		Id:   562958543552537,
		Name: "网络流入速率",
		Unit: "KB/S",
	},
	{
		Id:   562958543552538,
		Name: "网络流出速率",
		Unit: "KB/S",
	},
	{
		Id:   562958543618052,
		Name: "云硬盘使用率",
		Unit: "%",
	},
	{
		Id:   562958543618061,
		Name: "云硬盘IO写入",
		Unit: "%",
	},
	{
		Id:   562958543618062,
		Name: "云硬盘IO读出",
		Unit: "%",
	},
	{
		Id:   562958543618072,
		Name: "磁盘读操作速率",
		Unit: "requests/s",
	},
	{
		Id:   562958543618073,
		Name: "磁盘写操作速率",
		Unit: "requests/s",
	},
}

var VolumeIndicatorList = []Indicator{
	{
		Id:   562967133290497,
		Name: "磁盘读速率",
		Unit: "KB/s",
	},
	{
		Id:   562967133290498,
		Name: "磁盘写速率",
		Unit: "KB/s",
	},
	{
		Id:   562967133290499,
		Name: "磁盘读操作速率",
		Unit: "requests/s",
	},
	{
		Id:   562967133290500,
		Name: "磁盘写操作速率",
		Unit: "requests/s",
	},
	{
		Id:   562967133290501,
		Name: "云硬盘使用率",
		Unit: "%",
	},
}

var EipIndicatorList = []Indicator{
	{
		Id:   562971428257793,
		Name: "上行流量",
		Unit: "KB",
	},
	{
		Id:   562971428257794,
		Name: "上行带宽",
		Unit: "KB/s",
	},
	{
		Id:   562971428257795,
		Name: "下行流量",
		Unit: "KB",
	},
	{
		Id:   562971428257796,
		Name: "下行带宽",
		Unit: "KB/s",
	},
	{
		Id:   562971428257797,
		Name: "出网带宽使用率",
		Unit: "%",
	},
}

var NicIndicatorList = []Indicator{
	{
		Id:   563001493094401,
		Name: "网卡流入字节速率",
		Unit: "B/S",
	},
	{
		Id:   563001493094402,
		Name: "网卡流出字节速率",
		Unit: "B/s",
	},
	{
		Id:   563001493094403,
		Name: "端口收丢包率",
		Unit: "%",
	},
	{
		Id:   563001493094404,
		Name: "端口发丢包率",
		Unit: "%",
	},
}

var BandwidthIndicatorList = []Indicator{
	{
		Id:   562975723225089,
		Name: "上行流量",
		Unit: "KB",
	},
	{
		Id:   562975723225090,
		Name: "上行带宽",
		Unit: "KB/s",
	},
	{
		Id:   562975723225091,
		Name: "下行流量",
		Unit: "KB",
	},
	{
		Id:   562975723225092,
		Name: "下行带宽",
		Unit: "KB/s",
	},
}

type ObjType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Indicator struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

type GetIndicatorRequest struct {
	ObjTypeName string `json:"objTypeName"`
}

type GetInstanceRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
}

type GetCloudVmResponse struct {
	OwnerType             string    `json:"ownerType"`
	ExtraSpecs            string    `json:"extraSpecs"`
	TenantType            string    `json:"tenantType"`
	ResourcePoolNativeId  string    `json:"resourcePoolNativeId"`
	PrivateIps            string    `json:"privateIps"`
	OwnerId               string    `json:"ownerId"`
	CreatedAt             time.Time `json:"createdAt"`
	PowerState            string    `json:"powerState"`
	TaskState             string    `json:"taskState"`
	Id                    string    `json:"id"`
	PodId                 string    `json:"podId"`
	BizRegionId           string    `json:"bizRegionId"`
	ImageId               string    `json:"imageId"`
	FloatingIp            string    `json:"floatingIp"`
	IpAddress             string    `json:"ipAddress"`
	HostId                string    `json:"hostId"`
	HypervisorType        string    `json:"hypervisorType"`
	ResId                 string    `json:"resId"`
	Tags                  string    `json:"tags"`
	IsLocal               bool      `json:"is_Local"`
	ClassName             string    `json:"class_Name"`
	AzoneId               string    `json:"azoneId"`
	PhysicalHostId        string    `json:"physicalHostId"`
	ResourcePoolId        string    `json:"resourcePoolId"`
	FlavorPerformanceType string    `json:"flavorPerformanceType"`
	Name                  string    `json:"name"`
	ProjectName           string    `json:"projectName"`
	ProjectId             string    `json:"projectId"`
	BizRegionName         string    `json:"bizRegionName"`
	Status                string    `json:"status"`
	ResourcePoolType      string    `json:"resourcePoolType"`
	LogicalRegionName     string    `json:"logicalRegionName"`
	AzoneName             string    `json:"azoneName"`
	RegionName            string    `json:"regionName"`
	FlavorRamSize         string    `json:"flavorRamSize"`
	FlavorId              string    `json:"flavorId"`
	FlavorDiskSize        string    `json:"flavorDiskSize"`
	ClusterId             string    `json:"clusterId"`
	TenantName            string    `json:"tenantName"`
	OsVersion             string    `json:"osVersion"`
	ClusterName           string    `json:"clusterName"`
	OsType                string    `json:"osType"`
	LastModified          int64     `json:"last_Modified"`
	FlavorVcpu            string    `json:"flavorVcpu"`
	VdcId                 string    `json:"vdcId"`
	UnHexId               string    `json:"unHexId"`
	UserId                string    `json:"userId"`
	ResourcePoolName      string    `json:"resourcePoolName"`
	BizRegionNativeId     string    `json:"bizRegionNativeId"`
	LogicalRegionId       string    `json:"logicalRegionId"`
	RegionId              string    `json:"regionId"`
	Management            bool      `json:"management"`
	TenantId              string    `json:"tenantId"`
	VdcName               string    `json:"vdcName"`
	NativeId              string    `json:"nativeId"`
	LaunchedAt            string    `json:"launchedAt"`
	VmState               string    `json:"vmState"`
	CpuCoreNum            string    `json:"cpuCoreNum"`
}

type GetPerformanceRequest struct {
	Domain string                 `json:"domain"`
	Token  string                 `json:"token"`
	Params GetVmPerformanceParams `json:"params"`
}

type GetVmPerformanceParams struct {
	ObjTypeId    string   `json:"obj_type_id"`
	IndicatorIds []int64  `json:"indicator_ids"`
	ObjIds       []string `json:"obj_ids"`  //相应的resId
	Interval     string   `json:"interval"` //MINUTE（分钟）、HOUR（小时）DAY（天）、WEEK（星期）、MONTH（月）。
	Range        string   `json:"range"`    //LAST_5_MINUTE（最近5分钟，返回数据粒度为分钟,LAST_1_HOUR（最近1小时，返回数据粒度为分钟）
	// LAST_1_DAY（最近1天，返回数据粒度支持分钟、小时，默认按小时返回）,LAST_1_WEEK（最近1星期，返回数据支持分钟、小时、天，默认按天返回）
	//LAST_1_MONTH（最近1个月，返回数据支持分钟、小时、天、星期默认按星期返回）,LAST_1_QUARTER（最近一个季度，返回数据支持分钟、小时、天、星期、月，默认按月返回）
	//HALF_1_YEAR（最近半年，返回数据支持分钟、小时、天、星期、月、季度，默认按季度返回）,LAST_1_YEAR（最近1年，返回数据支持分钟、小时、天、星期、月、季度、半年，默认按半年返回）
}

type GetCloudVolumeResponse struct {
	OwnerType         string `json:"ownerType"`
	TenantType        string `json:"tenantType"`
	SourceVolId       string `json:"sourceVolId"`
	LogicalRegionName string `json:"logicalRegionName"`
	AzoneName         string `json:"azoneName"`
	RegionName        string `json:"regionName"`
	Description       string `json:"description"`
	Remark            string `json:"remark"`
	OwnerId           string `json:"ownerId"`
	ShareType         string `json:"shareType"`
	Bootable          string `json:"bootable"`
	CreatedAt         string `json:"createdAt"`
	DataPlane         string `json:"dataPlane"`
	OwnerName         string `json:"ownerName"`
	TenantName        string `json:"tenantName"`
	VolumeTypeId      string `json:"volumeTypeId"`
	ConfirmStatus     string `json:"confirmStatus"`
	Id                string `json:"id"`
	LastModified      int64  `json:"last_Modified"`
	VdcId             string `json:"vdcId"`
	BizRegionId       string `json:"bizRegionId"`
	LunWwn            string `json:"lunWwn"`
	ClassId           int    `json:"class_Id"`
	OriginalState     string `json:"originalState"`
	AttachBackendName string `json:"attachBackendName"`
	ResId             string `json:"resId"`
	ResourcePoolName  string `json:"resourcePoolName"`
	UserId            string `json:"userId"`
	KeystoneId        string `json:"keystoneId"`
	LogicalRegionId   string `json:"logicalRegionId"`
	IsLocal           bool   `json:"is_Local"`
	ClassName         string `json:"class_Name"`
	AzoneId           string `json:"azoneId"`
	Size              string `json:"size"`
	Encrypted         string `json:"encrypted"`
	RegionId          string `json:"regionId"`
	Management        bool   `json:"management"`
	ResourcePoolId    string `json:"resourcePoolId"`
	Name              string `json:"name"`
	TenantId          string `json:"tenantId"`
	VdcName           string `json:"vdcName"`
	AttachBackendId   string `json:"attachBackendId"`
	NativeId          string `json:"nativeId"`
	ProjectId         string `json:"projectId"`
	BizRegionName     string `json:"bizRegionName"`
	ResourcePoolType  string `json:"resourcePoolType"`
	Status            string `json:"status"`
}

type GetCloudEipResponse struct {
	LogicalRegionName string `json:"logicalRegionName"`
	RegionName        string `json:"regionName"`
	NetworkName       string `json:"networkName"`
	OwnerId           string `json:"ownerId"`
	TenantName        string `json:"tenantName"`
	FixedIpAddress    string `json:"fixedIpAddress"`
	NetworkId         string `json:"networkId"`
	Id                string `json:"id"`
	LastModified      int64  `json:"last_Modified"`
	VdcId             string `json:"vdcId"`
	BizRegionId       string `json:"bizRegionId"`
	HwsBandwidthId    string `json:"hwsBandwidthId"`
	ResId             string `json:"resId"`
	ResourcePoolName  string `json:"resourcePoolName"`
	FloatingIpAddress string `json:"floatingIpAddress"`
	LogicalRegionId   string `json:"logicalRegionId"`
	IsLocal           bool   `json:"is_Local"`
	ClassName         string `json:"class_Name"`
	RegionId          string `json:"regionId"`
	CreateTime        string `json:"createTime"`
	HwsBandwidthSize  string `json:"hwsBandwidthSize"`
	ResourcePoolId    string `json:"resourcePoolId"`
	TenantId          string `json:"tenantId"`
	VdcName           string `json:"vdcName"`
	NativeId          string `json:"nativeId"`
	ProjectId         string `json:"projectId"`
	BizRegionName     string `json:"bizRegionName"`
	Status            string `json:"status"`
	ResourcePoolType  string `json:"resourcePoolType"`
}

type GetCloudNicResponse struct {
	OwnerType         string `json:"ownerType"`
	LogicalRegionName string `json:"logicalRegionName"`
	RegionName        string `json:"regionName"`
	OwnerId           string `json:"ownerId"`
	OwnerName         string `json:"ownerName"`
	ConfirmStatus     string `json:"confirmStatus"`
	Id                string `json:"id"`
	LastModified      int64  `json:"last_Modified"`
	VdcId             string `json:"vdcId"`
	PodNativeId       string `json:"podNativeId"`
	BizRegionId       string `json:"bizRegionId"`
	DevId             string `json:"devId"`
	DevNativeId       string `json:"devNativeId"`
	VnicType          string `json:"vnicType"`
	ResId             string `json:"resId"`
	ResourcePoolName  string `json:"resourcePoolName"`
	KeystoneId        string `json:"keystoneId"`
	LogicalRegionId   string `json:"logicalRegionId"`
	IsLocal           bool   `json:"is_Local"`
	ClassName         string `json:"class_Name"`
	RegionId          string `json:"regionId"`
	ResourcePoolId    string `json:"resourcePoolId"`
	Name              string `json:"name"`
	NativeId          string `json:"nativeId"`
	Status            string `json:"status"`
	MacAddr           string `json:"macAddr"`
	ResourcePoolType  string `json:"resourcePoolType"`
}

type GetCloudBandwidthResponse struct {
	OwnerType         string `json:"ownerType"`
	TenantType        string `json:"tenantType"`
	LogicalRegionName string `json:"logicalRegionName"`
	RegionName        string `json:"regionName"`
	OwnerId           string `json:"ownerId"`
	ShareType         string `json:"shareType"`
	DataPlane         string `json:"dataPlane"`
	OwnerName         string `json:"ownerName"`
	TenantName        string `json:"tenantName"`
	Id                string `json:"id"`
	LastModified      int64  `json:"last_Modified"`
	VdcId             string `json:"vdcId"`
	BizRegionId       string `json:"bizRegionId"`
	FloatingIp        string `json:"floatingIp"`
	ClassId           int    `json:"class_Id"`
	ResId             string `json:"resId"`
	ResourcePoolName  string `json:"resourcePoolName"`
	LogicalRegionId   string `json:"logicalRegionId"`
	IsLocal           bool   `json:"is_Local"`
	ClassName         string `json:"class_Name"`
	Size              string `json:"size"`
	RegionId          string `json:"regionId"`
	ResourcePoolId    string `json:"resourcePoolId"`
	Name              string `json:"name"`
	TenantId          string `json:"tenantId"`
	VdcName           string `json:"vdcName"`
	NativeId          string `json:"nativeId"`
	ProjectId         string `json:"projectId"`
	BizRegionName     string `json:"bizRegionName"`
	ResourcePoolType  string `json:"resourcePoolType"`
}

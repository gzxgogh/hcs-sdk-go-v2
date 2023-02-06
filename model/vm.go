package model

import "time"

type CreateVmRequest struct {
	Domain   string         `json:"domain"`
	Token    string         `json:"token"`
	TenantId string         `json:"tenantId"`
	Params   CreateVmParams `json:"params"`
}

type CreateVmParams struct {
	Server CreateVm `json:"server"`
}

type CreateVm struct {
	ImageRef         string           `json:"imageRef"`   //镜像id
	FlavorRef        string           ` json:"flavorRef"` //规格id
	Description      string           `json:"description,omitempty"`
	Name             string           `json:"name"` //名称
	Personality      []Personality    `json:"personality,omitempty"`
	UserData         string           `json:"user_data,omitempty"`
	AdminPass        string           `json:"adminPass,omitempty"` //虚拟机密码，当前使用cloud-init的方式实现密码注入，该参数无效
	KeyName          string           `json:"key_name,omitempty"`  //如果需要使用SSH密钥方式登录 云服务器，请指定已创建密钥的 名称。约束：当key_name与 user_data同时指定时，user_data只做用户数据注入。
	VpcId            string           `json:"vpcid"`               //待创建云服务器所属虚拟私有云（简称VPC），需要指定已创建VPC的ID，UUID格式。
	Nics             []Nic            `json:"nics"`
	PublicIp         PublicIp         `json:"publicip,omitempty"`           //配置云服务器的弹性IP信息， TYPE3场景不支持，弹性IP有三 种配置方式。 不使用（无该字 段）。 自动分配，需要指定新 创建弹性IP的信息。 使用已 有，需要指定已创建弹性IP的信 息。
	Count            int              `json:"count,omitempty"`              //创建数量最大1000
	IsAutoRename     bool             `json:"isAutoRename,omitempty"`       //是否允许重名
	RootVolume       RootVolume       `json:"root_volume"`                  //系统盘信息
	DataVolumes      []DataVolume     `json:"data_volumes,omitempty"`       //数据盘信息
	SecurityGroups   []SecurityGroups `json:"security_groups,omitempty"`    //安全组信息
	AvailabilityZone string           `json:"availability_zone"`            //可用区域
	Extendparam      Extendparam      `json:"extendparam,omitempty"`        //创建云服务器附加信息
	Metadata         Metadata         `json:"metadata,omitempty"`           //创建云服务器元数据。 创建密 码方式鉴权的Windows弹性云 服务器时，为必填字段。 说明 该字段暂不支持用户写入数据， 但是当使用Windows镜像创建 弹性云服务器时，该字段为必选 字段。
	OsSchedulerHints OsSchedulerHints `json:"os:scheduler_hints,omitempty"` //云服务器调度信息。
	Tags             []string         `json:"tags,omitempty"`               //标签的格式为“key.value”。 其中，key的长度不超过36个字符，value的长度不超过43个字符。
	PowerOn          bool             `json:"power_on"`                     //创建完后云主机的状态，true:开机,false:关机
	Extra            Extra            `json:"extra,omitempty"`
}

type Personality struct {
	Contents string `json:"contents"` //注入文件的内容。 该值应指定 为注入文件的内容进行base64 格式编码后的信息。
	Path     string `json:"path"`     //注入文件路径信息。 .Linux系统 请输入注入文件保存路径，例如 “/etc/foo.txt”。
}
type Nic struct {
	IpAddress       string                   `json:"ip_address,omitempty"`       //待创建云服务器网卡的IP地址， IPv4格式。
	SubnetId        string                   `json:"subnet_id"`                  //待创建云服务器的网卡信息。需 要指定vpcid对应VPC下已创建 的网络（network）的ID，UUID格式。
	NicType         string                   `json:"nictype,omitempty"`          //网卡类型名称。
	IpAddressV6     string                   `json:"ip_address_v6,omitempty"`    //待创建云服务器网卡的IP地址， IPv6格式。
	PhysicalNetwork string                   `json:"physical_network,omitempty"` //物理网络
	BindingProfile  map[string]string        `json:"binding:profile,omitempty"`  //提供用户设置自定义信息。异构 计算时，属性external_vlan为接 入外部网络的vlan，当采用 access方式时设置为0。格式、 用法参考网络接口中创建端口_ 社区兼容接口中binding:profile 用法。
	ExtraDhcpOpts   []map[string]interface{} `json:"extra_dhcp_opts,omitempty"`  //DHCP的扩展Option，格式、用法参考网络接口中创建端口_社区兼容接口中extra_dhcp_opts对象说明。
}
type PublicIp struct {
	Eip Eip    `json:"eip,omitempty"` //配置云服务器自动分配弹性IP 时，创建弹性IP的配置参数。
	Id  string `json:"id,omitempty"`  //为待创建云服务器分配已有弹性IP时，分配的弹性IP的ID，UUID格式。约束：只能分配状态（status）为DOWN的弹性 IP。
}
type Eip struct {
	Bandwidth Bandwidth `json:"bandwidth"` //弹性IP地址带宽参数。
	IpType    string    `json:"iptype"`    //弹性IP外部网络名称。
}
type Bandwidth struct {
	Chargemode string `json:"chargemode,omitempty"` //带宽的计费类型。未传该字段，表示按带宽计费。字段值为空，表示按带宽计费。字段值为 “traffic”，表示按流量计费。字段为其它值，会导致创建云服务器失败。说明：如果share_type是WHOLE并且id有值，该参数会忽略。
	Id         string `json:"id,omitempty"`         //带宽ID，创建WHOLE类型带宽的弹性IP时可以指定之前的共享带宽创建。取值范围：WHOLE类型的带宽ID。说明当创建WHOLE类型的带宽时，该字段必选。
	ShareType  string `json:"sharetype"`            //带宽的共享类型。共享类型枚举：PER，表示独享。WHOLE，表示共享。
	Size       int    `json:"size,omitempty"`       //带宽（Mbit/s），取值范围为[1,300]。如果 share_type是PER，该参数必选 项；如果share_type是WHOLE并且id有值,该参数会忽略
}
type RootVolume struct {
	Size       int    `json:"size,omitempty"` //系统盘大小，容量单位为GB。 约束： 1.系统盘大小取值应不 小于镜像支持的系统盘的最小值 (镜像的min_disk属性)。 2.若该 参数没有指定或者指定为0，系 统盘大小默认取值为镜像中系统 盘的最小值(镜像的min_disk属 性)。
	VolumeType string `json:"volumetype"`     //磁盘类型
}
type DataVolume struct {
	Passthrough bool   `json:"hw:passthrough,omitempty"` //数据卷是否使用SCSI锁。如果使 用，请将该字段值配置为 “true”，否则，请勿填写该字段。
	Multiattach bool   `json:"multiattach,omitempty"`    //是否指定为共享磁盘
	Size        int    `json:"size"`                     //数据盘大小，容量单位为GB，输入大小范围为[1,65536]。
	VolumeType  string `json:"volumetype"`               //磁盘类型
	VolumeId    string `json:"volume_id,omitempty"`      //1.数据盘 ID参数不为空，创建虚拟机时，不再创建空白的数据盘，使用已有的数据盘；2.数据盘ID参数为空，创建虚拟机时，默认创建空白的数据盘
}
type SecurityGroups struct {
	Id string `json:"id"` //待创建云服务器的安全组，会对 创建云服务器中配置的网卡生效。需要指定已有安全组的ID，UUID格式。
}
type Extendparam struct {
	Postfix string `json:"postfix,omitempty"` //虚拟机名称的后缀，默认为空。 取值范围[0001-9999]，限制：1、该值与参数count值之和不能大于9999；2、该值长度与虚拟机名称长度之和不能超过参数name的最大长度限制。
}
type Metadata struct {
	AdminPass        string `json:"admin_pass,omitempty"` //以Windows镜像创建的弹性云服务器Administrator用户的密码。密码复杂度要求：1.长度为8-26位。
	HwWatchdogAction string `json:"hw_watchdog_action,omitempty"`
	HaPolicyType     string `json:"_ha_policy_type,omitempty"`
}
type OsSchedulerHints struct {
	Group string `json:"group,omitempty"` //云服务器组ID，UUID格式。
}
type Extra struct {
	Device Device `json:"device"`
}
type Device struct {
	Cdrom string `json:"cdrom,omitempty"`
}

type CreateVmFromVolumeRequest struct {
	Domain   string                   `json:"domain"`
	Token    string                   `json:"token"`
	TenantId string                   `json:"tenantId"`
	Params   CreateVmFromVolumeParams `json:"params"`
}

type CreateVmFromVolumeParams struct {
	Server CreateVmFromVolume `json:"server"`
}

type CreateVmFromVolume struct {
	FlavorRef          string               ` json:"flavorRef"`              //规格id
	Name               string               `json:"name"`                    //名称
	Networks           []NetWork            `json:"networks"`                //网络的uuid
	BlockDeviceMapping []BlockDeviceMapping `json:"block_device_mapping_v2"` //系统卷信息
	AvailabilityZone   string               `json:"availability_zone"`       //可用区域
}

type NetWork struct {
	Uuid string `json:"uuid"`
}

type BlockDeviceMapping struct {
	Uuid                string `json:"uuid"`                  //系统卷id
	SourceType          string `json:"source_type"`           //指卷设备的源头类型,当前只支持volume、image、snapshot类型
	DestinationType     string `json:"destination_type"`      //指卷设备的当前类型， 当前只支持volume类型
	DeleteOnTermination string `json:"delete_on_termination"` //指删除弹性云服务器时，是否删除卷，默认值False,true
	BootIndex           int    `json:"boot_index"`            //boot_index是指启动标识，“0”代表启动盘，“-1“代表非启动盘。
}

type CreateVmFromVolumeResponse struct {
	Id        string `json:"id"`
	AdminPass string `json:"adminPass"`
}

type DeleteVmRequest struct {
	Domain   string         `json:"domain"`
	Token    string         `json:"token"`
	TenantId string         `json:"tenantId"`
	Params   DeleteVmParams `json:"params"`
}

type DeleteVmParams struct {
	DeleteVm       []DeleteVm `json:"servers"`
	DeletePublicIp bool       `json:"delete_publicip"`
	DeleteVolume   bool       `json:"delete_volume"`
}

type DeleteVm struct {
	Id string `json:"id"`
}

type UpdateVmRequest struct {
	Domain   string         `json:"domain"`
	Token    string         `json:"token"`
	TenantId string         `json:"tenantId"`
	ServerId string         `json:"server_id"`
	Params   UpdateVmParams `json:"params"`
}

type UpdateVmParams struct {
	UpdateVm UpdateVm `json:"server"`
}

type UpdateVm struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type QueryVmRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	ServerId string `json:"server_id,omitempty"`
}

type QueryVmResponse struct {
	TenantId                         string                           `json:"tenant_id"`
	Metadata                         VmMetadata                       `json:"metadata"`
	Addresses                        map[string]interface{}           `json:"addresses"` //"addresses": { "subnet-74e5-gogh": [ { "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:fd:98:59", "OS-EXT-IPS:type": "fixed", "addr": "192.168.123.140", "version": 4} ]},
	OSEXTSTSTaskState                interface{}                      `json:"OS-EXT-STS:task_state"`
	OSDCFDiskConfig                  string                           `json:"OS-DCF:diskConfig"`
	OSEXTAZAvailabilityZone          string                           `json:"OS-EXT-AZ:availability_zone"`
	OSEXTSTSPowerState               int                              `json:"OS-EXT-STS:power_state"`
	Id                               string                           `json:"id"`
	OsExtendedVolumesVolumesAttached OsExtendedVolumesVolumesAttached `json:"os-extended-volumes:volumes_attached"`
	OSEXTSRVATTRHost                 string                           `json:"OS-EXT-SRV-ATTR:host"`
	Image                            VmImage                          `json:"image,omitempty"`
	OSSRVUSGTerminatedAt             interface{}                      `json:"OS-SRV-USG:terminated_at"`
	AccessIPv4                       string                           `json:"accessIPv4"`
	AccessIPv6                       string                           `json:"accessIPv6"`
	Created                          time.Time                        `json:"created"`
	HostId                           string                           `json:"hostId"`
	OSEXTSRVATTRHypervisorHostname   string                           `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	Flavor                           VmFlavor                         `json:"flavor"`
	KeyName                          string                           `json:"key_name"`
	SecurityGroups                   VmSecurityGroups                 `json:"security_groups"`
	ConfigDrive                      string                           `json:"config_drive"`
	OSEXTSTSVmState                  string                           `json:"OS-EXT-STS:vm_state"`
	UserId                           string                           `json:"user_id"`
	OSEXTSRVATTRInstanceName         string                           `json:"OS-EXT-SRV-ATTR:instance_name"`
	Name                             string                           `json:"name"`
	Progress                         int                              `json:"progress"`
	OSSRVUSGLaunchedAt               string                           `json:"OS-SRV-USG:launched_at"`
	Updated                          time.Time                        `json:"updated"`
	Status                           string                           `json:"status"`
}

type ItemNotFound struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type VmMetadata struct {
	ProductId                 string `json:"productId"`
	InstanceVwatchdog         string `json:"__instance_vwatchdog"`
	ServerExpiry              string `json:"server_expiry"`
	HaPolicyType              string `json:"_ha_policy_type"`
	CascadedInstanceExtrainfo string `json:"cascaded.instance_extrainfo"`
}

type OsExtendedVolumesVolumesAttached []struct {
	Id string `json:"id"`
}

type VmImage struct {
	Id string `json:"id"`
}

type VmFlavor struct {
	Id string `json:"id"`
}

type VmSecurityGroups []struct {
	Name string `json:"name"`
}

type StartVmRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	ServerId string `json:"server_id"`
	Action   Start  `json:"action"`
}

type Start struct {
	Start map[string]interface{} `json:"os-start"`
}

type StopVmRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	ServerId string `json:"server_id"`
	Action   Stop   `json:"action"`
}

type Stop struct {
	Stop map[string]interface{} `json:"os-stop"`
}

type RebootVmRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	ServerId string `json:"server_id"`
	Action   Reboot `json:"action"`
}

type Reboot struct {
	Type RebootType `json:"reboot"`
}

type RebootType struct {
	Type string `json:"type"` //普通重启（“SOFT”）和强制重启（“HARD”）
}

type VmRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	ServerId string `json:"server_id"`
}

type QueryVmNicResponse []struct {
	NetId     string   `json:"net_id"`
	PortId    string   `json:"port_id"`
	MacAddr   string   `json:"mac_addr"`
	PortState string   `json:"port_state"`
	FixedIps  FixedIps `json:"fixed_ips"`
}

type FixedIps []struct {
	SubnetId  string `json:"subnet_id,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
}

type AttachNicRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	ServerId string          `json:"server_id"`
	Params   AttachNicParams `json:"params"`
}

type AttachNicParams struct {
	InterfaceAttachment InterfaceAttachment `json:"interfaceAttachment"`
}

type InterfaceAttachment struct {
	NetId string `json:"net_id"`
}

type BatchAttachNicRequest struct {
	Domain   string               `json:"domain"`
	Token    string               `json:"token"`
	TenantId string               `json:"tenantId"`
	ServerId string               `json:"server_id"`
	Params   BatchAttachNicParams `json:"params"`
}

type BatchAttachNicParams struct {
	Nics []BatchNic `json:"nics"`
}

type BatchNic struct {
	SubnetId         string           `json:"subnet_id"`                   //待创建云服务器的网卡信息。需 要指定vpcid对应VPC下已创建 的网络（network）的ID，UUID格式。
	IpAddress        string           `json:"ip_address,omitempty"`        //待创建云服务器网卡的IP地址， IPv4格式。
	IpAddressV6      string           `json:"ip_address_v6,omitempty"`     //待创建云服务器网卡的IP地址， IPv6格式。
	PhysicalNetwork  string           `json:"physical_network,omitempty"`  //物理网络
	AvailabilityZone string           `json:"availability_zone,omitempty"` //可用区域
	SecurityGroups   []SecurityGroups `json:"security_groups"`
}

type DetachNicRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	ServerId string `json:"server_id"`
	NicId    string `json:"nic_id"`
}

type AttachVolumeRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenantId"`
	ServerId string             `json:"server_id"`
	Params   AttachVolumeParams `json:"params"`
}

type AttachVolumeParams struct {
	VolumeAttachment VolumeAttachment `json:"volumeAttachment"`
}

type VolumeAttachment struct {
	VolumeId string `json:"volumeId"`
	Device   string `json:"device"`
}

type AttachVolumeResponse struct {
	VolumeId string `json:"volumeId"`
	Id       string `json:"id"`
	Device   string `json:"device"`
	ServerId string `json:"serverId"`
}

type DetachVolumeRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	ServerId string `json:"server_id"`
	VolumeId string `json:"volume_id"`
}

type UpgradeVmRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	ServerId string          `json:"server_id"`
	Params   UpgradeVmParams `json:"params"`
}

type UpgradeVmParams struct {
	Resize Resize `json:"resize"`
}

type Resize struct {
	FlavorRef       string `json:"flavorRef"`
	DedicatedHostId string `json:"dedicated_host_id,omitempty"`
}

type ChangeVmImageRequest struct {
	Domain   string              `json:"domain"`
	Token    string              `json:"token"`
	TenantId string              `json:"tenantId"`
	ServerId string              `json:"server_id"`
	Params   ChangeVmImageParams `json:"params"`
}

type ChangeVmImageParams struct {
	ChangeVmImage ChangeVmImage `json:"os-change"`
}

type ChangeVmImage struct {
	AdminPass string            `json:"adminPass,omitempty"` //虚拟机密码，当前使用cloud-init的方式实现密码注入，该参数无效
	KeyName   string            `json:"key_name,omitempty"`
	UserId    string            `json:"user_id,omitempty"`
	ImageId   string            `json:"imageid"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type OnlineUpgradeVmRequest struct {
	Domain   string                `json:"domain"`
	Token    string                `json:"token"`
	TenantId string                `json:"tenantId"`
	ServerId string                `json:"server_id"`
	Params   OnlineUpgradeVmParams `json:"params"`
}

type OnlineUpgradeVmParams struct {
	Resize Resize `json:"live-resize"`
}

type CloneVmRequest struct {
	Domain   string        `json:"domain"`
	Token    string        `json:"token"`
	TenantId string        `json:"tenantId"`
	ServerId string        `json:"server_id"`
	Params   CloneVmParams `json:"params"`
}

type CloneVmParams struct {
	CloneVm CloneVm `json:"os-clone"`
}

type CloneVm struct {
	Name      string   `json:"name,omitempty"`
	CloneNum  int      `json:"clone_num,omitempty"`
	CloneType string   `json:"clone_type,omitempty"`
	KeyName   string   `json:"key_name,omitempty"`
	UserData  string   `json:"user_data,omitempty"`
	Metadata  string   `json:"metadata,omitempty"`
	PowerOn   bool     `json:"power_on"`
	Postfix   string   `json:"postfix,omitempty"`
	VpcId     string   `json:"vpc_id,omitempty"`
	Nics      NicClone `json:"nics,omitempty"`
}

type NicClone struct {
	SubNetId       string           `json:"subnet_id,omitempty"`
	SecurityGroups []SecurityGroups `json:"security_groups,omitempty"`
}

type QueryConsoleAddRequest struct {
	Domain   string                `json:"domain"`
	Token    string                `json:"token"`
	TenantId string                `json:"tenantId"`
	ServerId string                `json:"server_id"`
	Params   QueryConsoleAddParams `json:"params"`
}

type QueryConsoleAddParams struct {
	ConsoleAdd ConsoleAdd `json:"remote_console"`
}

type ConsoleAdd struct {
	Type     string `json:"type"`
	Protocol string `json:"protocol"`
}

type QueryConsoleAddResponse struct {
	Url      string `json:"url"`
	Type     string `json:"type"`
	Protocol string `json:"protocol"`
}

type ChangeVmPwdRequest struct {
	Domain   string            `json:"domain"`
	Token    string            `json:"token"`
	TenantId string            `json:"tenantId"`
	ServerId string            `json:"server_id"`
	Params   ChangeVmPwdParams `json:"params"`
}

type ChangeVmPwdParams struct {
	ResetPassword ResetPassword `json:"reset-password"`
}

type ResetPassword struct {
	Password string `json:"new_password"`
}

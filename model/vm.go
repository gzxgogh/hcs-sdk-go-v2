package model

import "time"

type CreateVmRequest struct {
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Server   Server `json:"server"`
}

type Server struct {
	FlavorRef          string             ` json:"flavorRef"`              //规格id
	Name               string             `json:"name"`                    //名称
	Networks           []NetWork          `json:"networks"`                //网络的uuid
	BlockDeviceMapping BlockDeviceMapping `json:"block_device_mapping_v2"` //系统卷信息
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

type CreateVmResponse struct {
	Id        string `json:"id"`
	AdminPass string `json:"adminPass"`
}

type QueryVmRequest struct {
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
	Image                            VmImage                          `json:"image"`
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
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	ServerId string                 `json:"server_id"`
	Start    map[string]interface{} `json:"os-start"`
}

type StopVmRequest struct {
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	ServerId string                 `json:"server_id"`
	Stop     map[string]interface{} `json:"os-stop"`
}

type RebootVmRequest struct {
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	ServerId string `json:"server_id"`
	Reboot   Reboot `json:"reboot"`
}

type Reboot struct {
	Type string `json:"type"` //普通重启（“SOFT”）和强制重启（“HARD”）
}

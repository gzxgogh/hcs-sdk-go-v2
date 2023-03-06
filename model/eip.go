package model

type CreateFloatIpRequest struct {
	Domain   string              `json:"domain"`
	Token    string              `json:"token"`
	TenantId string              `json:"tenantId"`
	Params   CreateFloatIpParams `json:"params"`
}

type CreateFloatIpParams struct {
	FloatingIp FloatingIp `json:"floatingIp"`
}

type FloatingIp struct {
	FloatingNetworkId string `json:"floating_network_id"`
	PortId            string `json:"port_id"`
}

type CreateFloatIpResponse struct {
	Id                string `json:"id"`
	RouterId          string `json:"router_id"`
	TenantId          string `json:"tenant_id"`
	FloatingNetworkId string `json:"floating_network_id"`
	FixedIpAddress    string `json:"fixed_ip_address"`
	FloatingIpAddress string `json:"floating_ip_address"`
	PortId            string `json:"port_id"`
	Status            string `json:"status"`
	QosPolicyId       string `json:"qos_policy_id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	ProjectId         string `json:"project_id"`
}

type CreateEipRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	Params   CreateEipParams `json:"params"`
}

type CreateEipParams struct {
	PublicIp  EipPublicIp  `json:"publicip"`
	Bandwidth EipBandwidth `json:"bandwidth,omitempty"`
}

type EipPublicIp struct {
	Type      string `json:"type"` //网络名称
	SubnetId  string `json:"subnet_id,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
}

type EipBandwidth struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	ShareType string `json:"share_type"` //带宽共享类型。独占带宽为PER，共享带宽为WHOLE。
}

type DeleteEipRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateEipRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	Id       string          `json:"id"`
	Params   UpdateEipParams `json:"params"`
}

type UpdateEipParams struct {
	UpdateEip UpdateEip `json:"publicip"`
}

type UpdateEip struct {
	PortId      string `json:"port_id,omitempty"`
	QosPolicyId string `json:"qos_policy_id,omitempty"`
}

type QueryEipRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id,omitempty"`
	PortId   string `json:"port_id"`
}

type QueryEipResponse struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	ExternalNetId      string `json:"external_net_id"`
	PortId             string `json:"port_id"`
	PublicIpAddress    string `json:"public_ip_address"`
	PrivateIpAddress   string `json:"private_ip_address"`
	Status             string `json:"status"`
	TenantId           string `json:"tenant_id"`
	ProjectId          string `json:"project_id"`
	CreateTime         string `json:"create_time"`
	BandwidthId        string `json:"bandwidth_id"`
	BandwidthName      string `json:"bandwidth_name"`
	BandwidthShareType string `json:"bandwidth_share_type"`
	BandwidthSize      int    `json:"bandwidth_size"`
	Profile            struct {
		ProductId string `json:"product_id"`
		RegionId  string `json:"region_id"`
		OrderId   string `json:"order_id"`
	} `json:"profile"`
	RouterId    string `json:"router_id"`
	OpStatus    int    `json:"op_status"`
	PortDetails struct {
		Name         string `json:"name"`
		NetworkId    string `json:"network_id"`
		MacAddress   string `json:"mac_address"`
		AdminStateUp string `json:"admin_state_up"`
		Status       string `json:"status"`
		DeviceId     string `json:"device_id"`
		DeviceOwner  string `json:"device_owner"`
	} `json:"port_details"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

type UpdateBandwidthRequest struct {
	Domain      string                `json:"domain"`
	Token       string                `json:"token"`
	TenantId    string                `json:"tenantId"`
	BandwidthId string                `json:"bandwidthId"`
	Params      UpdateBandwidthParams `json:"params"`
}

type UpdateBandwidthParams struct {
	UpdateBandwidth UpdateBandwidth `json:"bandwidth"`
}

type UpdateBandwidth struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Size        int    `json:"size,omitempty"`
}

type CreateShareBandwidthRequest struct {
	Domain   string                     `json:"domain"`
	Token    string                     `json:"token"`
	TenantId string                     `json:"tenantId"`
	Params   CreateShareBandwidthParams `json:"params"`
}

type CreateShareBandwidthParams struct {
	Bandwidth ShareBandwidth `json:"bandwidth"`
}

type ShareBandwidth struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

type DeleteShareBandwidthRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateShareBandwidthRequest struct {
	Domain   string                     `json:"domain"`
	Token    string                     `json:"token"`
	TenantId string                     `json:"tenantId"`
	Id       string                     `json:"id,omitempty"`
	Params   UpdateShareBandwidthParams `json:"params"`
}

type UpdateShareBandwidthParams struct {
	UpdateShareBandwidth UpdateShareBandwidth `json:"bandwidth"`
}

type UpdateShareBandwidth struct {
	Name        string `json:"name,omitempty"`
	Size        int    `json:"size,omitempty"`
	Description string `json:"description,omitempty"`
}

type QueryShareBandwidthRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id,omitempty"`
}

type QueryShareBandwidthResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
	Description  string `json:"description"`
	ShareType    string `json:"share_type"`
	TenantId     string `json:"tenant_id"`
	PublicipInfo []struct {
		PublicipId      string `json:"publicip_id"`
		PublicipAddress string `json:"publicip_address"`
		IpVersion       int    `json:"ip_version"`
		PublicipType    string `json:"publicip_type"`
	} `json:"publicip_info"`
	BandwidthType string `json:"bandwidth_type"`
	Shared        bool   `json:"shared"`
	RuleType      string `json:"rule_type"`
	Status        string `json:"status"`
}

type ShareBandwidthAttachEipRequest struct {
	Domain      string                        `json:"domain"`
	Token       string                        `json:"token"`
	TenantId    string                        `json:"tenantId"`
	BandwidthId string                        `json:"bandwidthId"`
	Params      ShareBandwidthAttachEipParams `json:"params"`
}

type ShareBandwidthAttachEipParams struct {
	ShareBandwidthAttachEip ShareBandwidthAttachEip `json:"bandwidth"`
}

type ShareBandwidthAttachEip struct {
	Bandwidth PublicIpInfo `json:"publicip_info"`
	Size      int          `json:"size"`
}
type PublicIpInfo []struct {
	PublicIpId string `json:"publicip_id"`
}

type ShareBandwidthDetachEipRequest struct {
	Domain      string                        `json:"domain"`
	Token       string                        `json:"token"`
	TenantId    string                        `json:"tenantId"`
	BandwidthId string                        `json:"bandwidthId"`
	Params      ShareBandwidthDetachEipParams `json:"params"`
}

type ShareBandwidthDetachEipParams struct {
	ShareBandwidthDetachEip ShareBandwidthDetachEip `json:"bandwidth"`
}

type ShareBandwidthDetachEip struct {
	Bandwidth PublicIpInfo `json:"publicip_info"`
	Size      int          `json:"size"`
}

type CreateFipRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	Params   CreateFipParams `json:"params"`
}

type CreateFipParams struct {
	CreateFip CreateFip `json:"floatingip_create_dict"`
}

type CreateFip struct {
	FloatingNetworkId string `json:"floating_network_id"`
	FixedIpAddress    string `json:"fixed_ip_address,omitempty"`
	FloatingIpAddress string `json:"floating_ip_address,omitempty"`
	SubnetId          string `json:"subnet_id,omitempty"`
	PortId            string `json:"port_id,omitempty"`
}

type DeleteFipRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateFipRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	Id       string          `json:"id"`
	Params   UpdateFipParams `json:"params"`
}

type UpdateFipParams struct {
	UpdateFip UpdateFip `json:"floatingip_update_dict"`
}

type UpdateFip struct {
	FixedIpAddress string `json:"fixed_ip_address,omitempty"`
	PortId         string `json:"port_id,omitempty"`
}

type QueryFipRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryFipResponse struct {
	Id                string `json:"id"`
	Status            string `json:"status"`
	RouterId          string `json:"router_id"`
	TenantId          string `json:"tenant_id"`
	FloatingNetworkId string `json:"floating_network_id"`
	FixedIpAddress    string `json:"fixed_ip_address"`
	QosPolicyId       string `json:"qos_policy_id"`
	FloatingIpAddress string `json:"floating_ip_address"`
	PortId            string `json:"port_id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

package model

type CreateVpcRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	Params   CreateVpcParams `json:"params"`
}

type CreateVpcParams struct {
	CreateVpc CreateVpc `json:"vpc"`
}

type CreateVpc struct {
	Name                string `json:"name"`
	ExternalGatewayInfo struct {
		NetworkId string `json:"network_id"`
	} `json:"external_gateway_info"`
	Ntp   []string `json:"ntp"`
	NtpV6 []string `json:"ntp_v6"`
}

type DeleteVpcRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateVpcRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	Id       string          `json:"id"`
	Params   UpdateVpcParams `json:"params"`
}

type UpdateVpcParams struct {
	UpdateVpc CreateVpc `json:"vpc"`
}

type QueryVpcRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryVpcResponse struct {
	Id                  string      `json:"id"`
	Name                string      `json:"name"`
	Cidr                interface{} `json:"cidr"`
	Status              string      `json:"status"`
	ExternalGatewayInfo struct {
		NetworkId        string `json:"network_id"`
		EnableSnat       bool   `json:"enable_snat"`
		ExternalFixedIps []struct {
			SubnetId  string `json:"subnet_id"`
			IpAddress string `json:"ip_address"`
		} `json:"external_fixed_ips"`
	} `json:"external_gateway_info"`
	Routes []struct {
		Destination string `json:"destination"`
		Nexthop     string `json:"nexthop"`
	} `json:"routes"`
	EnableSharedSnat         bool     `json:"enable_shared_snat"`
	ServiceAvailabilityZones []string `json:"service_availability_zones"`
	Ntp                      []string `json:"ntp"`
	Expiry                   string   `json:"expiry"`
	OpStatus                 int      `json:"op_status"`
	ProjectId                string   `json:"project_id"`
	NtpV6                    []string `json:"ntp_v6"`
}

type CreateSubnetRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenantId"`
	Params   CreateSubnetParams `json:"params"`
}

type CreateSubnetParams struct {
	CreateSubnet CreateSubnet `json:"subnet"`
}

type CreateSubnet struct {
	Ipv4EnableMulticast bool              `json:"ipv4_enable_multicast,omitempty"` //IPv4组播是否启用
	Name                string            `json:"name"`
	Cidr                string            `json:"cidr"`                          //子网CIDR。
	GatewayIp           string            `json:"gateway_ip,omitempty"`          //子网网关。
	DhcpEnable          bool              `json:"dhcp_enable,omitempty"`         //是否使能DHCP。
	PrimaryDns          string            `json:"primary_dns,omitempty"`         //子网DNS服务器地址，IP格式
	SecondaryDns        string            `json:"secondary_dns,omitempty"`       //子网DNS服务器地址，IP格式。
	DnsList             []string          `json:"dnsList,omitempty"`             //子网DNS服务器地址列表。
	NtpList             []string          `json:"ntpList,omitempty"`             //子网NTP服务器地址列表。
	VpcId               string            `json:"vpc_id"`                        //子网所在VPC ID。
	AvailabilityZone    string            `json:"availability_zone,omitempty"`   //子网所在的可用分区标识。
	PhysicalNetwork     string            `json:"physical_network,omitempty"`    //物理网络信息。
	AllocationPools     []AllocationPools `json:"allocation_pools,omitempty"`    //动态分配IP地址的地址池。
	SegmentationId      int               `json:"segmentation_id,omitempty"`     //虚拟网络的VLAN ID。
	External            bool              `json:"external,omitempty"`            //true标识直连网。
	Routed              bool              `json:"routed,omitempty"`              //指示是否是VPC路由网络
	Ipv4Enable          bool              `json:"ipv4_enable,omitempty"`         //是否使能IPv4，只支持True，默 认为True
	Ipv6Enable          bool              `json:"ipv6_enable,omitempty"`         //是否使能IPv6。
	CidrV6              string            `json:"cidr_v6,omitempty"`             //IPv6子网CIDR。
	GatewayIpV6         string            `json:"gateway_ip_v6,omitempty"`       //IPv6网关。
	DnsListV6           []string          `json:"dns_list_v6,omitempty"`         //IPv6 DNS服务器地址列表。
	AllocationPoolsV6   []AllocationPools `json:"allocation_pools_v6,omitempty"` //IPv6动态分配IP的地址池。
	Ipv6AddressMode     string            `json:"ipv6_address_mode,omitempty"`   //IPv6地址控制模式，Type1支持：dhcpv6-stateful模式，Type2和Type3支持：dhcpv6-stateful,dhcpv6-stateless,slaac三种模式。约束： ipv6_address_mode和ipv6_ra_mode取值必须相同
	Ipv6RaMode          string            `json:"ipv6_ra_mode,omitempty"`        //IPv6路由通告模式，Type1支持：dhcpv6-stateful模式，Type2和Type3支持：dhcpv6-stateful,dhcpv6-stateless,slaac三种模式。约束：ipv6_address_mode和ipv6_ra_mode取值必须相同
	Description         string            `json:"description,omitempty"`
	ProviderNetworkType string            `json:"provider:network_type,omitempty"` //网络类型，vlan或vxlan。
}

type AllocationPools struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type DeleteSubnetRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateSubnetRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenantId"`
	VpcId    string             `json:"vpcId"`
	SubnetId string             `json:"subnetId"`
	Params   UpdateSubnetParams `json:"params"`
}

type UpdateSubnetParams struct {
	UpdateSubnet UpdateSubnet `json:"subnet"`
}

type UpdateSubnet struct {
	Ipv4EnableMulticast bool     `json:"ipv4_enable_multicast,omitempty"` //IPv4组播是否启用
	Name                string   `json:"name,omitempty"`
	DhcpEnable          bool     `json:"dhcp_enable,omitempty"`   //是否使能DHCP。
	PrimaryDns          string   `json:"primary_dns,omitempty"`   //子网DNS服务器地址，IP格式
	SecondaryDns        string   `json:"secondary_dns,omitempty"` //子网DNS服务器地址，IP格式。
	DnsList             []string `json:"dnsList,omitempty"`       //子网DNS服务器地址列表。
	NtpList             []string `json:"ntpList,omitempty"`       //子网NTP服务器地址列表。
	Routed              bool     `json:"routed,omitempty"`        //指示是否是VPC路由网络
	Ipv6Enable          bool     `json:"ipv6_enable,omitempty"`   //是否使能IPv6。
	Description         string   `json:"description,omitempty"`
}

type QuerySubnetRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QuerySubnetResponse struct {
	Name            string   `json:"name"`
	Id              string   `json:"id"`
	Cidr            string   `json:"cidr"`
	DnsList         []string `json:"dnsList"`
	Status          string   `json:"status"`
	VpcId           string   `json:"vpc_id"`
	CidrV6          string   `json:"cidr_v6"`
	GatewayIp       string   `json:"gateway_ip"`
	GatewayIpV6     string   `json:"gateway_ip_v6"`
	DhcpEnable      bool     `json:"dhcp_enable"`
	PrimaryDns      string   `json:"primary_dns"`
	SecondaryDns    string   `json:"secondary_dns"`
	DnsListV6       []string `json:"dns_list_v6"`
	AllocationPools []struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"allocation_pools"`
	AllocationPoolsV6 []struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"allocation_pools_v6"`
	SegmentationId      int    `json:"segmentation_id"`
	NeutronSubnetId     string `json:"neutron_subnet_id"`
	NeutronIpv6SubnetId string `json:"neutron_ipv6_subnet_id"`
	NeutronNetworkId    string `json:"neutron_network_id"`
	Description         string `json:"description"`
	Ipv6Enable          bool   `json:"ipv6_enable"`
}

type CreatePrivateIpRequest struct {
	Domain   string                `json:"domain"`
	Token    string                `json:"token"`
	TenantId string                `json:"tenantId"`
	Params   CreatePrivateIpParams `json:"params"`
}

type CreatePrivateIpParams struct {
	CreatePrivateIp CreatePrivateIp `json:"privateips"`
}

type CreatePrivateIp []struct {
	SubnetId  string `json:"subnet_id"`
	IpAddress string `json:"ip_address,omitempty"`
}

type DeletePrivateIpRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryPrivateIpRequest struct {
	Domain      string `json:"domain"`
	Token       string `json:"token"`
	TenantId    string `json:"tenantId"`
	SubnetId    string `json:"subnetId"`
	PrivateIpId string `json:"privateIpId"`
}

type QueryPrivateIpResponse struct {
	Status      string `json:"status"`
	Id          string `json:"id"`
	SubnetId    string `json:"subnet_id"`
	TenantId    string `json:"tenant_id"`
	ProjectId   string `json:"project_id"`
	DeviceOwner string `json:"device_owner"`
	IpAddress   string `json:"ip_address"`
}

type BandVipRequest struct {
	Domain      string    `json:"domain"`
	Token       string    `json:"token"`
	TenantId    string    `json:"tenantId"`
	PrivateIpId string    `json:"privateIpId"`
	Params      VipParams `json:"params"`
}

type VipParams struct {
	Id string `json:"id"`
}

type UnBandVipRequest struct {
	Domain      string    `json:"domain"`
	Token       string    `json:"token"`
	TenantId    string    `json:"tenantId"`
	PrivateIpId string    `json:"privateIpId"`
	Params      VipParams `json:"params"`
}

type CreateQosPolicyRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	Params   QosPolicyParams `json:"params"`
}

type QosPolicyParams struct {
	QosPolicy QosPolicy `json:"policy"`
}

type QosPolicy struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Shared      bool   `json:"shared,omitempty"`
	TenantId    string `json:"tenant_id,omitempty"`
}

type DeleteQosPolicyRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	PolicyId string `json:"policyId"`
}

type UpdateQosPolicyRequest struct {
	Domain   string          `json:"domain"`
	Token    string          `json:"token"`
	TenantId string          `json:"tenantId"`
	PolicyId string          `json:"policyId"`
	Params   QosPolicyParams `json:"params"`
}

type QueryQosPolicyRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryQosPolicyResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Shared      bool   `json:"shared"`
	TenantId    string `json:"tenant_id"`
	Rules       []struct {
		Id           string `json:"id"`
		QosPolicyId  string `json:"qos_policy_id"`
		MaxKbps      int    `json:"max_kbps"`
		MaxBurstKbps int    `json:"max_burst_kbps"`
		Direction    string `json:"direction"`
	} `json:"rules"`
}

type QueryPolicyRuleRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	PolicyId string `json:"policyId"`
	RuleId   string `json:"ruleId"`
}

type QueryExternalNetworksRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
}

type QueryExternalNetworksResponse struct {
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	VpcNetworkType    string   `json:"vpc_network_type"`
	Group             string   `json:"group"`
	AvailabilityZones []string `json:"availability_zones"`
}

type CreatePeerRequest struct {
	Domain   string     `json:"domain"`
	Token    string     `json:"token"`
	TenantId string     `json:"tenantId"`
	Params   CreatePeer `json:"params"`
}

type CreatePeer struct {
	Name           string `json:"name"`
	LocalVpcId     string `json:"local_vpc_id"`
	PeerVpcId      string `json:"peer_vpc_id"`
	PeerDomainName string `json:"peer_domain_name,omitempty"`
	PeerProjectId  string `json:"peer_project_id,omitempty"`
}

type VpcInfo struct {
	VpcId    string `json:"vpc_id"`
	TenantId string `json:"tenant_id,omitempty"`
}

type DeletePeerRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdatePeerRequest struct {
	Domain   string       `json:"domain"`
	Token    string       `json:"token"`
	TenantId string       `json:"tenantId"`
	Id       string       `json:"id"`
	Params   UpdateParams `json:"params"`
}

type UpdateParams struct {
	Name string `json:"name"`
}

type QueryPeerRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryPeerResponse struct {
	Id               string      `json:"id"`
	Name             string      `json:"name"`
	Identity         string      `json:"identity"`
	RequesterVpcInfo PeerVpcInfo `json:"requesterVpcInfo"`
	AccepterVpcInfo  PeerVpcInfo `json:"accepterVpcInfo"`
	Status           string      `json:"status"`
	ApplyTime        string      `json:"apply_time"`
	DeleteTime       string      `json:"delete_time"`
	AcceptTime       string      `json:"accept_time"`
	ExpireTime       string      `json:"expire_time"`
	CurrentStatus    string      `json:"current_status"`
}

type PeerVpcInfo struct {
	DomainName          string `json:"domain_name"`
	VpcName             string `json:"vpc_name"`
	VpcId               string `json:"vpc_id"`
	Cidr                string `json:"cidr"`
	PeeringOptions      string `json:"peering_options"`
	ExternalGatewayInfo struct {
		NetworkId        string `json:"network_id"`
		EnableSnat       bool   `json:"enable_snat"`
		ExternalFixedIps []struct {
			SubnetId  string `json:"subnet_id"`
			IpAddress string `json:"ip_address"`
		} `json:"external_fixed_ips"`
	} `json:"external_gateway_info"`
	TenantId string `json:"tenant_id"`
}

type ActionPeerRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	PeerId   string `json:"peerId"`
}

type ActionPeerRouterRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	VpcId    string                 `json:"vpcId"`
	Params   ActionPeerRouterParams `json:"params"`
}

type ActionPeerRouterParams struct {
	ActionPeerRouter ActionPeerRouter `json:"routes"`
}

type ActionPeerRouter []struct {
	Destination string `json:"destination"` //对应的vpcId下子网的cidr
	Nexthop     string `json:"nexthop"`     //对等连接的id
}

type QueryPeerRouterRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	VpcId    string `json:"vpcId"`
}

type QueryPeerRouterResponse struct {
	Destination string `json:"destination"`
	Nexthop     string `json:"nexthop"`
}

type QueryPortsRequest struct {
	Domain         string `json:"domain"`
	Token          string `json:"token"`
	TenantId       string `json:"tenantId"`
	Id             string `json:"id,omitempty"`
	Status         string `json:"status,omitempty"`    //端口状态。取值范围： ACTIVE、BUILD、DOWN
	FixedIps       string `json:"fixed_ips,omitempty"` //端口IP地址。
	NetworkId      string `json:"network_id,omitempty"`
	DeviceId       string `json:"device_id,omitempty"` //端口所属设备id，如虚拟机 uuid。
	SecurityGroups string `json:"security_groups"`     //安全组id
}

type QueryPortsResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	NetworkId    string `json:"network_id"`
	AdminStateUp bool   `json:"admin_state_up"`
	MacAddress   string `json:"mac_address"`
	FixedIps     []struct {
		SubnetId  string `json:"subnet_id"`
		IpAddress string `json:"ip_address"`
	} `json:"fixed_ips"`
	DeviceId            string   `json:"device_id"`
	DeviceOwner         string   `json:"device_owner"`
	TenantId            string   `json:"tenant_id"`
	Status              string   `json:"status"`
	SecurityGroups      []string `json:"security_groups"`
	AllowedAddressPairs []struct {
		Ipaddress  string `json:"ipaddress"`
		MacAddress string `json:"macAddress"`
	} `json:"allowed_address_pairs"`
	ExtraDhcpOpts []struct {
		OptName  string `json:"opt_name"`
		OptValue string `json:"opt_value"`
	} `json:"extra_dhcp_opts"`
	BindingVifType    string                 `json:"binding:vif_type"`
	BindingVifDetails map[string]interface{} `json:"binding:vif_details"`
	BindingHostId     string                 `json:"binding:host_id"`
	BindingProfile    map[string]interface{} `json:"binding:profile"`
	BindingVnicType   string                 `json:"binding:vnic_type"`
	Description       string                 `json:"description"`
	CreatedAt         string                 `json:"created_at"`
	UpdatedAt         string                 `json:"updated_at"`
	ProjectId         string                 `json:"project_id"`
}

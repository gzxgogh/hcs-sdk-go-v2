package model

type QueryPublicNetRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
}

type QueryPublicNetResponse struct {
	TenantId                string   `json:"tenant_id"`
	ProviderPhysicalNetwork string   `json:"provider:physical_network"`
	Shared                  bool     `json:"shared"`
	PortSecurityEnabled     bool     `json:"port_security_enabled"`
	NetScope                string   `json:"net_scope"`
	ProviderSegmentationId  int      `json:"provider:segmentation_id"`
	Description             string   `json:"description"`
	ProviderNetworkType     string   `json:"provider:network_type"`
	ExtraDhcpOpts           []string `json:"extra_dhcp_opts"`
	CreatedAt               string   `json:"created_at"`
	AvailabilityZoneHints   []string `json:"availability_zone_hints"`
	UpdatedAt               string   `json:"updated_at"`
	ProjectId               string   `json:"project_id"`
	AdminStateUp            bool     `json:"admin_state_up"`
	Subnets                 []string `json:"subnets"`
	VlanTransparent         string   `json:"vlan_transparent"`
	Id                      string   `json:"id"`
	Ipv4AddressScope        string   `json:"ipv4_address_scope"`
	AvailabilityZones       []string `json:"availability_zones"`
	IsDefault               bool     `json:"is_default"`
	Mtu                     int      `json:"mtu"`
	Tags                    []string `json:"tags"`
	Ipv6AddressScope        string   `json:"ipv6_address_scope"`
	RouterExternal          bool     `json:"router:external"`
	Name                    string   `json:"name"`
	QosPolicyId             string   `json:"qos_policy_id"`
	Status                  string   `json:"status"`
}

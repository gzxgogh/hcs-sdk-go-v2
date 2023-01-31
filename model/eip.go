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

package model

type CreateSgRequest struct {
	Domain   string         `json:"domain"`
	Token    string         `json:"token"`
	TenantId string         `json:"tenantId"`
	Params   CreateSgParams `json:"params"`
}

type CreateSgParams struct {
	CreateSg CreateSg `json:"security_group"`
}

type CreateSg struct {
	Name string `json:"name"`
}

type DeleteSgRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateSgRequest struct {
	Domain   string         `json:"domain"`
	Token    string         `json:"token"`
	TenantId string         `json:"tenantId"`
	Id       string         `json:"id"`
	Params   UpdateSgParams `json:"params"`
}

type UpdateSgParams struct {
	UpdateSg UpdateSg `json:"security_group"`
}

type UpdateSg struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type QuerySgRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id,omitempty"`
}

type QuerySgResponse struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	VpcId       string   `json:"vpc_id"`
	ProjectId   string   `json:"project_id"`
	Description string   `json:"description"`
	Rules       []SgRule `json:"security_group_rules"`
}

type SgRule struct {
	Id              string `json:"id"`
	Direction       string `json:"direction"` //规则方向，取值范围ingress/egress
	Ethertype       string `json:"ethertype"` //网络类型，取值范围IPv4/IPv6
	Description     string `json:"description"`
	PortRangeMax    int    `json:"port_range_max"`   //最大端口，当协议类型为ICMP时，该值表示 ICMP的code。范围：1-65535（当表示code时是0-255）
	PortRangeMin    int    `json:"port_range_min"`   //最小端口，当协议类型为ICMP时，该值表示ICMP的type。protocol为tcp和udp时，port_range_max和port_range_min必须同时输入，且port_range_max应大于等于 port_range_min。protocol为icmp时，指定ICMPcode（port_range_max）时，必须同时指定ICMP type（port_range_min）。范围：1-65535（当表示code时是0-255）
	Protocol        string `json:"protocol"`         //协议类型或直接指定IP协议号
	RemoteGroupId   string `json:"remote_group_id"`  //所属安全组的对端id。与remote_ip_prefix参数二选一
	RemoteIpPrefix  string `json:"remote_ip_prefix"` //对端ip网段。与remote_group_id参数二选一最大长度：255
	SecurityGroupId string `json:"security_group_id"`
	TenantId        string `json:"tenant_Id"`
	CreatedAt       string `json:"created_at"`
}

type CreateSgRuleRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenantId"`
	Params   CreateSgRuleParams `json:"params"`
}

type CreateSgRuleParams struct {
	CreateSgRule CreateSgRule `json:"security_group_rule"`
}

type CreateSgRule struct {
	Description     string `json:"description,omitempty"`
	Direction       string `json:"direction"`                  //规则方向，取值范围ingress/egress
	Ethertype       string `json:"ethertype,omitempty"`        //网络类型，取值范围IPv4/IPv6
	PortRangeMax    int    `json:"port_range_max,omitempty"`   //最大端口，当协议类型为ICMP时，该值表示 ICMP的code。范围：1-65535（当表示code时是0-255）
	PortRangeMin    int    `json:"port_range_min,omitempty"`   //最小端口，当协议类型为ICMP时，该值表示ICMP的type。protocol为tcp和udp时，port_range_max和port_range_min必须同时输入，且port_range_max应大于等于 port_range_min。protocol为icmp时，指定ICMPcode（port_range_max）时，必须同时指定ICMP type（port_range_min）。范围：1-65535（当表示code时是0-255）
	Protocol        string `json:"protocol,omitempty"`         //协议类型或直接指定IP协议号,tcp', 'udp', 'icmp', 'icmpv6
	RemoteGroupId   string `json:"remote_group_id,omitempty"`  //所属安全组的对端id。与remote_ip_prefix参数二选一
	RemoteIpPrefix  string `json:"remote_ip_prefix,omitempty"` //对端ip网段。与remote_group_id参数二选一最大长度：255
	SecurityGroupId string `json:"security_group_id"`
	TenantId        string `json:"tenant_id,omitempty"`
}

type DeleteSgRuleRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	RuleId   string `json:"ruleId"`
}

type QuerySgRuleRequest struct {
	Domain          string `json:"domain"`
	Token           string `json:"token"`
	TenantId        string `json:"tenantId"`
	SecurityGroupId string `json:"security_group_id,omitempty"`
	Id              string `json:"id,omitempty"`
}

type BandNicToSgRequest struct {
	Domain   string            `json:"domain"`
	Token    string            `json:"token"`
	TenantId string            `json:"tenantId"`
	PortId   string            `json:"portId"`
	Params   BandNicToSgParams `json:"params"`
}

type BandNicToSgParams struct {
	BandNicToSg BandNicToSg `json:"port"`
}

type BandNicToSg struct {
	SecurityGroups []string `json:"security_groups"`
}

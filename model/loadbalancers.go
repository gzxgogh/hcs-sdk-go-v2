package model

type CreateLoadBalancersRequest struct {
	Domain   string                    `json:"domain"`
	Token    string                    `json:"token"`
	TenantId string                    `json:"tenantId"`
	Params   CreateLoadBalancersParams `json:"params"`
}

type CreateLoadBalancersParams struct {
	CreateLoadBalancers CreateLoadBalancers `json:"loadbalancer"`
}

type CreateLoadBalancers struct {
	Name              string      `json:"name"`
	Description       string      `json:"description,omitempty"`
	AdminStateUp      bool        `json:"admin_state_up,omitempty"`       //管理状态：true/false。 使用说
	Provider          string      `json:"provider,omitempty"`             //供应者。 使用说明：RegionType I只支持vlb。
	TenantId          string      `json:"tenant_id,omitempty"`            //租户id
	VipSubnetId       string      `json:"vip_subnet_id"`                  //分配vip的子网ID。 使用说明： 仅支持内网类型。
	VipAddress        string      `json:"vip_address,omitempty"`          //VIP的IP地址。
	FlavorId          string      `json:"flavor_id,omitempty"`            //flavor的ID。 使用说明：暂不支持。
	L4FlavorId        string      `json:"l4_flavor_id,omitempty"`         //四层规格ID，默认值为null
	L7FlavorId        string      `json:"l7_flavor_id,omitempty"`         //七层规格ID，默认值为null
	IpTargetEnable    bool        `json:"ip_target_enable,omitempty"`     //是否开启跨vpc后端:true/false。默认为false
	ElbVirsubnetIds   []string    `json:"elb_virsubnet_ids,omitempty"`    //elb后端子网id列表,最大64个,需要和lb同属一个vpc
	ElbVirsubnetL4Ips []VirSubnet `json:"elb_virsubnet_l4_ips,omitempty"` //elb TCP监听器后端子网ip列表，最大64个
	ElbVirsubnetL7Ips []VirSubnet `json:"elb_virsubnet_l7_ips,omitempty"` //elb HTTP/HTTPS监听器后端子网ip列表，最大64个
}

type VirSubnet struct {
	NetworkId string   `json:"network_id"` //后端子网id
	Ips       []string `json:"ips"`        //elb后端子网ip列表
}

type DeleteLoadBalancersRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateLoadBalancersRequest struct {
	Domain   string                    `json:"domain"`
	Token    string                    `json:"token"`
	TenantId string                    `json:"tenantId"`
	Id       string                    `json:"id"`
	Params   UpdateLoadBalancersParams `json:"params"`
}

type UpdateLoadBalancersParams struct {
	UpdateLoadBalancers UpdateLoadBalancers `json:"loadbalancer"`
}

type UpdateLoadBalancers struct {
	Name              string      `json:"name,omitempty"`
	Description       string      `json:"description,omitempty"`
	L4FlavorId        string      `json:"l4_flavor_id,omitempty"`         //四层规格ID，默认值为null
	L7FlavorId        string      `json:"l7_flavor_id,omitempty"`         //七层规格ID，默认值为null
	IpTargetEnable    bool        `json:"ip_target_enable,omitempty"`     //是否开启跨vpc后端:true/false。默认为false
	ElbVirsubnetIds   []string    `json:"elb_virsubnet_ids,omitempty"`    //elb后端子网id列表,最大64个,需要和lb同属一个vpc
	ElbVirsubnetL4Ips []VirSubnet `json:"elb_virsubnet_l4_ips,omitempty"` //elb TCP监听器后端子网ip列表，最大64个
	ElbVirsubnetL7Ips []VirSubnet `json:"elb_virsubnet_l7_ips,omitempty"` //elb HTTP/HTTPS监听器后端子网ip列表，最大64个
}

type QueryLoadBalancersRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryLoadBalancersResponse struct {
	AdminStateUp bool   `json:"admin_state_up"`
	Description  string `json:"description"`
	FlavorId     string `json:"flavor_id"`
	Id           string `json:"id"`
	L4FlavorId   string `json:"l4_flavor_id"`
	L7FlavorId   string `json:"l7_flavor_id"`
	Listeners    []struct {
		Id string `json:"id"`
	} `json:"listeners"`
	TenantId string `json:"tenant_id"`
	Name     string `json:"name"`
	Pools    []struct {
		Id string `json:"id"`
	} `json:"pools"`
	Provider           string `json:"provider"`
	OperatingStatus    string `json:"operating_status"`
	ProvisioningStatus string `json:"provisioning_status"`
	Tags               []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"tags"`
	VipAddress        string      `json:"vip_address"`
	VipSubnetId       string      `json:"vip_subnet_id"`
	VipPortId         string      `json:"vip_port_id"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
	IpTargetEnable    bool        `json:"ip_target_enable"`
	ElbVirsubnetIds   []string    `json:"elb_virsubnet_ids"`
	ElbVirsubnetL4Ips []VirSubnet `json:"elb_virsubnet_l4_ips"`
	ElbVirsubnetL7Ips []VirSubnet `json:"elb_virsubnet_l7_ips"`
}

type CreateListenersRequest struct {
	Domain   string                `json:"domain"`
	Token    string                `json:"token"`
	TenantId string                `json:"tenantId"`
	Params   CreateListenersParams `json:"params"`
}

type CreateListenersParams struct {
	CreateListeners CreateListeners `json:"listener"`
}

type CreateListeners struct {
	AdminStateUp           bool     `json:"admin_state_up,omitempty"`            //管理状态 true/false。
	ConnectionLimit        int      `json:"connection_limit,omitempty"`          //最大连接数，取值范围[-1,2147483647]。
	DefaultPoolId          string   `json:"default_pool_id,omitempty"`           //关联的pool id
	DefaultTlsContainerRef string   `json:"default_tls_container_ref,omitempty"` //TLS secrets容器的引用。 使用说明：Type1和Type2融合ELB填写证书ID。
	Description            string   `json:"description,omitempty"`
	Http2Enable            bool     `json:"http2_enable,omitempty"`            //HTTP2功能的开启状态。取值范围： true/false。 true：开启。false：关闭。默认为false。仅针对监听器的协议为TERMINATED_HTTPS有意义
	EnhanceL7policyEnable  bool     `json:"enhance_l7policy_enable,omitempty"` //高级路由功能的开启状态。取值范围： true/false。 true：开启。false：关闭。开启后无法关闭默认为false。
	KeepaliveTimeout       int      `json:"keepalive_timeout,omitempty"`       //空闲超时时间。 使用说明：监听器协议为UDP时不能设置该参数，显示为null； 协议为TCP时可设置该参数，取值范围[10, 4000]，单位为秒，默认值为300； 协议为HTTP或TERMINATED_HTTPS时可设置该参数，取值范围[0, 4000]，单位为秒，默认值为60。
	LoadbalancerId         string   `json:"loadbalancer_id"`                   //关联的负载均衡器 ID
	Name                   string   `json:"name,omitempty"`
	Protocol               string   `json:"protocol"`                     //监听协议号。 使用说明：支持TCP、UDP、HTTP和TERMINATED_HTTPS。
	ProtocolPort           int      `json:"protocol_port"`                //监听端口，取值范围[1,65535]。使用说明：监听协议为TCP时，端口号可以取值为0。
	SniContainerRefs       []string `json:"sni_container_refs,omitempty"` //TLS secrets容器引用列表。 使 用说明：暂不支持。
	TenantId               string   `json:"tenant_id,omitempty"`
}

type DeleteListenersRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateListenersRequest struct {
	Domain   string                `json:"domain"`
	Token    string                `json:"token"`
	TenantId string                `json:"tenant_id"`
	Id       string                `json:"id"`
	Params   UpdateListenersParams `json:"params"`
}

type UpdateListenersParams struct {
	UpdateListeners UpdateListeners `json:"listener"`
}

type UpdateListeners struct {
	ConnectionLimit       int    `json:"connection_limit,omitempty"` //最大连接数，取值范围[-1,2147483647]。
	Description           string `json:"description,omitempty"`
	KeepaliveTimeout      int    `json:"keepalive_timeout,omitempty"` //空闲超时时间。 使用说明：监听器协议为UDP时不能设置该参数，显示为null； 协议为TCP时可设置该参数，取值范围[10, 4000]，单位为秒，默认值为300； 协议为HTTP或TERMINATED_HTTPS时可设置该参数，取值范围[0, 4000]，单位为秒，默认值为60。
	Name                  string `json:"name,omitempty"`
	TimeoutClientData     int    `json:"timeout_client_data,omitempty"`     //请求超时时间。 使用说明：监听器协议为UDP或TCP时不能设置该参数，显示为null； 协议为HTTP或TERMINATED_HTTPS时可设置该参数，取值范围[1, 300]，单位为秒，默认值为60。
	TimeoutMemberData     int    `json:"timeout_member_data,omitempty"`     //响应超时时间。
	EnhanceL7policyEnable bool   `json:"enhance_l7policy_enable,omitempty"` //高级路由功能的开启状态。取值范围： true/false。 true：开启。false：关闭。开启后无法关闭默认为false。
}

type QueryListenersRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryListenersResponse struct {
	AdminStateUp            bool        `json:"admin_state_up"`
	ClientCaTlsContainerRef interface{} `json:"client_ca_tls_container_ref"`
	ConnectionLimit         int         `json:"connection_limit"`
	CreatedAt               string      `json:"created_at"`
	UpdatedAt               string      `json:"updated_at"`
	DefaultPoolId           string      `json:"default_pool_id,omitempty"` //关联的pool id
	DefaultTlsContainerRef  string      `json:"default_tls_container_ref"`
	Description             string      `json:"description"`
	Http2Enable             bool        `json:"http2_enable,omitempty"`
	EnhanceL7PolicyEnable   bool        `json:"enhance_l7policy_enable"`
	Id                      string      `json:"id"`
	KeepaliveTimeout        int         `json:"keepalive_timeout"`
	Loadbalancers           []struct {
		Id string `json:"id"`
	} `json:"loadbalancers"`
	Name             string   `json:"name"`
	ProtocolPort     int      `json:"protocol_port"`
	Protocol         string   `json:"protocol"`
	SniContainerRefs []string `json:"sni_container_refs"`
	Tags             []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"tags"`
	TenantId          string      `json:"tenant_id"`
	TimeoutClientData interface{} `json:"timeout_client_data"`
	TimeoutMemberData interface{} `json:"timeout_member_data"`
}

type CreatePoolRequest struct {
	Domain   string           `json:"domain"`
	Token    string           `json:"token"`
	TenantId string           `json:"tenantId"`
	Params   CreatePoolParams `json:"params"`
}

type CreatePoolParams struct {
	CreatePool CreatePool `json:"pool"`
}

type CreatePool struct {
	AdminStateUp       bool               `json:"admin_state_up,omitempty"` //管理状态 true/false。
	Description        string             `json:"description,omitempty"`
	LbAlgorithm        string             `json:"lb_algorithm"`              //负载均衡算法，可以为ROUND_ROBIN、LEAST_CONNECTIONS、SOURCE_IP。
	ListenerId         string             `json:"listener_id,omitempty"`     //pool关联的listener ID。使用说明：listener_id和loadbalancer_id中至少指定一个。
	LoadbalancerId     string             `json:"loadbalancer_id,omitempty"` //pool关联的listener ID。使用说明：listener_id和loadbalancer_id中至少指定一个。
	Name               string             `json:"name,omitempty"`
	Protocol           string             `json:"protocol"` //后端协议。 使用说明：支持TCP、UDP和HTTP。listener的protocol为TCP时pool的protocol必须为TCP，listener的protocol为UDP时pool的protocol必须为UDP，listener的protocol为HTTP或HTTPS时pool的protocol必须为HTTP。
	SessionPersistence SessionPersistence `json:"session_persistence,omitempty"`
	TenantId           string             `json:"tenant_id,omitempty"`
}

type SessionPersistence struct {
	CookName           string `json:"cook_name,omitempty"`           //cookie名字。 使用说明：只有当type为APP_COOKIE时才支持。
	Type               string `json:"type,omitempty"`                //类型，可以为SOURCE_IP、HTTP_COOKIE、APP_COOKIE。
	PersistenceTimeout string `json:"persistence_timeout,omitempty"` //会话保持的超时时间。当type为APP_COOKIE时不生效。 取值范围：● [1,60]（分钟）：当后端云服务器组的protocol为TCP、UDP时。● [1,1440]（分钟）：当后端云服务器组的protocol为HTTP时。
}

type DeletePoolRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdatePoolRequest struct {
	Domain   string           `json:"domain"`
	Token    string           `json:"token"`
	TenantId string           `json:"tenant_id"`
	Id       string           `json:"id"`
	Params   UpdatePoolParams `json:"params"`
}

type UpdatePoolParams struct {
	UpdatePool UpdatePool `json:"pool"`
}

type UpdatePool struct {
	Description        string             `json:"description,omitempty"`
	LbAlgorithm        string             `json:"lb_algorithm,omitempty"` //负载均衡算法，可以为ROUND_ROBIN、LEAST_CONNECTIONS、SOURCE_IP。
	Name               string             `json:"name,omitempty"`
	SessionPersistence SessionPersistence `json:"session_persistence,omitempty"`
}

type QueryPoolRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenant_id"`
	Id       string `json:"id"`
}

type QueryPoolResponse struct {
	AdminStateUp    bool   `json:"admin_state_up"`
	Description     string `json:"description"`
	HealthmonitorId string `json:"healthmonitor_id"`
	Id              string `json:"id"`
	LbAlgorithm     string `json:"lb_algorithm"`
	ListenerId      string `json:"listener_id"`
	Listeners       []struct {
		Id string `json:"id"`
	} `json:"listeners"`
	Loadbalancers []struct {
		Id string `json:"id"`
	} `json:"loadbalancers"`
	Members []struct {
		Id string `json:"id"`
	} `json:"members"`
	Name               string             `json:"name"`
	Protocol           string             `json:"protocol"`
	TenantId           string             `json:"tenant_id"`
	SessionPersistence SessionPersistence `json:"session_persistence"`
}

type CreateFlavorRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenant_id"`
	Params   CreateFlavorParams `json:"params"`
}

type CreateFlavorParams struct {
	CreateFlavor CreateFlavor `json:"flavor"`
}

type CreateFlavor struct {
	Name      string       `json:"name,omitempty"`
	Shared    string       `json:"shared,omitempty"`     //暂未使用
	ProjectId string       `json:"project_id,omitempty"` //项目id
	Type      string       `json:"type"`                 //类型，只能填l4或者l7
	TenantId  string       `json:"tenant_id,omitempty"`
	Info      []FlavorInfo `json:"info"`
}

type FlavorInfo struct {
	FlavorType string `json:"flavor_type"` //flavor指标。 使用说明：当flavor类型为l4时，flavorType可填cps（每秒新建连接数）、connection（并发）、bandwidth（带宽）； 当flavorType可填cps类型为l7时，flavorType可填cps（每秒新建连接数）、connection（并发）、bandwidth（带宽）、qps（每秒查询数）。
	Value      int    `json:"value"`       //指标大小。 使用说明：所有指标取值范围为[0，设定最大值]，当输入超过设定最大值时，自动取设定的最大值。其中最大值与网络节点数有关，下面列出的为单个网络节点上的指标最大值，实际使用时可乘以网络节点数量。 当flavorType类型为l4时，cps最大值为87000，单位为个；connection最大值为2200000，单位为个/秒；bandwidth最大值为16384，单位为Mbit/s； 当flavorType类型为l7时，cps最大值为20000，单位为个；connection最大值为320000，单位为个/秒；bandwidth最大值为10384，单位为Mbit/s；qps最大值为300000，单位为个。
}

type DeleteFlavorRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateFlavorRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenantId"`
	Id       string             `json:"id"`
	Params   UpdateFlavorParams `json:"params"`
}

type UpdateFlavorParams struct {
	UpdateFlavor UpdateFlavor `json:"flavor"`
}

type UpdateFlavor struct {
	Name string       `json:"name,omitempty"`
	Info []FlavorInfo `json:"info,omitempty"`
}

type QueryFlavorRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryFlavorResponse struct {
	Id            string                 `json:"id"`
	Name          string                 `json:"name"`
	Shared        bool                   `json:"shared"`     //是否共享
	ProjectId     string                 `json:"project_id"` //项目id
	Type          string                 `json:"type"`       //类型，只能填l4或者l7
	FlavorSoldOut bool                   `json:"flavor_sold_out"`
	TenantId      string                 `json:"tenant_id"`
	Status        string                 `json:"status"`
	Info          map[string]interface{} `json:"info"`
}

type CreatePolicyRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenant_id"`
	Params   CreatePolicyParams `json:"params"`
}

type CreatePolicyParams struct {
	CreatePolicy CreatePolicy `json:"l7policy"`
}

type CreatePolicy struct {
	Action             string            `json:"action"`                   //REDIRECT_TO_POOL：将匹配的流量转发到redirect_pool_id指定的后端云服务器组上；REDIRECT_TO_LISTENER：将listener_id指定的HTTP监听器的流量重定向到redirect_listener_id指定的TERMINATED_HTTPS监听器上；REDIRECT_TO_URL：重定向到URL；
	AdminStateUp       bool              `json:"admin_state_up,omitempty"` //管理状态 true/false。
	Description        string            `json:"description,omitempty"`
	ListenerId         string            `json:"listener_id"` //转发策略对应的监听器ID
	Name               string            `json:"name,omitempty"`
	Position           int               `json:"position,omitempty"`             //转发优先级，从1递增，最高 100。
	Priority           int               `json:"priority,omitempty"`             //转发策略的优先级。 当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。 数字越小表示优先级越高，同一监听器下不允许重复。 当action为REDIRECT_TO_LISTENER时，仅支持指定为0，优先级最高。 当关联的listener没有开启enhance_l7policy_enable，按原有policy的排序逻辑，自动排序。各域名之间优先级独立，相同域名下，按path的compare_type排序，精确>前缀>正则，匹配类型相同时，path的长度越长优先级越高。若policy下只有域名rule，没有路径rule，默认path为前缀匹配/。 当关联的listener开启了enhance_l7policy_enable，且不传该字段，则新创建的转发策略的优先级的值为：同一监听器下已有转发策略的优先级的最大值+1。因此，若当前已有转发策略的优先级的最大值是10000，新创建会因超出取值范围10000而失败。此时可通过传入指定priority，或调整原有policy的优先级来避免错误。若监听器下没有转发策略，则新建的转发策略的优先级为1。 不支持该字段，请勿使用。 最小值： 0 最大值： 10000
	RedirectPoolId     string            `json:"redirect_pool_id,omitempty"`     //转发到pool的ID。● 当action为REDIRECT_TO_POOL时为必选字段；● 当action为REDIRECT_TO_LISTENER时，不可指定该字段。
	RedirectListenerId string            `json:"redirect_listener_id,omitempty"` //流量匹配后转发到的监听器的ID。● 当action为REDIRECT_TO_LISTENER时为必选字段；● 当action为REDIRECT_TO_POOL时，不可指定该字段。
	RedirectUrlConfig  RedirectUrlConfig `json:"redirect_url_config,omitempty"`  //转发到的url配置。 当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。● 当action为REDIRECT_TO_URL时生效，且为必选字段，其他action不可指定，否则报错。 格式：protocol://host:port/ path?query protocol、host、 port、 path不允许同时不传或同时传${xxx}（${xxx}表示原值，如${host}表示被转发的请求URL的host部分）。 protocol和port传入的值不能与l7policy关联的监听器一致且host、 path同时不传或同时传${xxx}。● 不支持该字段，请勿使用。
	RedirectUrl        string            `json:"redirect_url,omitempty"`         //转发策略重定向到的url，默认值：null。 该字段为预留字段，暂未启用。
	TenantId           string            `json:"tenant_id,omitempty"`
}

type RedirectUrlConfig struct {
	StatusCode int    `json:"status_code,omitempty"` //重定向后的返回码。 取值范围： ● 301 ● 302 ● 303 ●307 ● 308 最小长度： 1 最大长度： 16
	Protocol   string `json:"protocol,omitempty"`    //重定向的协议。 默认值${protocol}表示继承原值（即与被转发请求保持一致）。 取值范围： ● HTTP ● HTTPS ● ${protocol} 缺省值： ${protocol} 最小长度： 1 最大长度： 36
	Host       string `json:"host,omitempty"`        //重定向的主机名。
	Port       string `json:"port,omitempty"`        //重定向到的端口。
	Path       string `json:"path,omitempty"`        //重定向的路径。 默认值${path}表示继承原值（即与被转发请求保持一致）。
	Query      string `json:"query,omitempty"`       //重定向的查询字符串。 默认$ {query}表示继承原值
}

type DeletePolicyRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdatePolicyRequest struct {
	Domain   string             `json:"domain"`
	Token    string             `json:"token"`
	TenantId string             `json:"tenantId"`
	Id       string             `json:"id"`
	Params   UpdatePolicyParams `json:"params"`
}

type UpdatePolicyParams struct {
	UpdatePolicy UpdatePolicy `json:"l7policy"`
}

type UpdatePolicy struct {
	AdminStateUp       bool              `json:"admin_state_up,omitempty"` //管理状态 true/false。
	Description        string            `json:"description,omitempty"`
	Name               string            `json:"name,omitempty"`
	RedirectPoolId     string            `json:"redirect_pool_id,omitempty"`     //转发到pool的ID。● 当action为REDIRECT_TO_POOL时为必选字段；● 当action为REDIRECT_TO_LISTENER时，不可指定该字段。
	RedirectListenerId string            `json:"redirect_listener_id,omitempty"` //流量匹配后转发到的监听器的ID。● 当action为REDIRECT_TO_LISTENER时为必选字段；● 当action为REDIRECT_TO_POOL时，不可指定该字段。
	RedirectUrlConfig  RedirectUrlConfig `json:"redirect_url_config"`            //转发到的url配置。 当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。● 当action为REDIRECT_TO_URL时生效，且为必选字段，其他action不可指定，否则报错。 格式：protocol://host:port/ path?query protocol、host、 port、 path不允许同时不传或同时传${xxx}（${xxx}表示原值，如${host}表示被转发的请求URL的host部分）。 protocol和port传入的值不能与l7policy关联的监听器一致且host、 path同时不传或同时传${xxx}。● 不支持该字段，请勿使用。
	Priority           int               `json:"priority,omitempty"`             //转发策略的优先级。 当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。 数字越小表示优先级越高，同一监听器下不允许重复。 当action为REDIRECT_TO_LISTENER时，仅支持指定为0，优先级最高。 当关联的listener没有开启enhance_l7policy_enable，按原有policy的排序逻辑，自动排序。各域名之间优先级独立，相同域名下，按path的compare_type排序，精确>前缀>正则，匹配类型相同时，path的长度越长优先级越高。若policy下只有域名rule，没有路径rule，默认path为前缀匹配/。 当关联的listener开启了enhance_l7policy_enable，且不传该字段，则新创建的转发策略的优先级的值为：同一监听器下已有转发策略的优先级的最大值+1。因此，若当前已有转发策略的优先级的最大值是10000，新创建会因超出取值范围10000而失败。此时可通过传入指定priority，或调整原有policy的优先级来避免错误。若监听器下没有转发策略，则新建的转发策略的优先级为1。 不支持该字段，请勿使用。 最小值： 0 最大值： 10000
}

type QueryPolicyRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryPolicyResponse struct {
	Action             string            `json:"action"`
	AdminStateUp       bool              `json:"admin_state_up"`
	Description        string            `json:"description"`
	Id                 string            `json:"id"`
	Name               string            `json:"name"`
	ListenerId         string            `json:"listener_id"`
	Position           int               `json:"position"`
	Priority           int               `json:"priority"`
	ProvisioningStatus string            `json:"provisioning_status"`
	RedirectPoolId     string            `json:"redirect_pool_id"`
	RedirectListenerId string            `json:"redirect_listener_id"`
	RedirectUrl        string            `json:"redirect_url"`
	RedirectUrlConfig  RedirectUrlConfig `json:"redirect_url_config"`
	Rules              []struct {
		Id string `json:"id"`
	} `json:"rules"`
	TenantId string `json:"tenant_id"`
}

type CreatePolicyRuleRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenant_id"`
	PolicyId string                 `json:"policyId"`
	Params   CreatePolicyRuleParams `json:"params"`
}

type CreatePolicyRuleParams struct {
	CreatePolicyRule CreatePolicyRule `json:"l7rule"`
}

type CreatePolicyRule struct {
	AdminStateUp bool           `json:"admin_state_up,omitempty"` //管理状态 true/false。
	CompareType  string         `json:"compare_type"`             //匹配方式：type为HOST_NAME时可以为EQUAL_TO；type为PATH时可以为REGEX，STARTS_WITH，EQUAL_TO
	Invert       bool           `json:"invert,omitempty"`         //是否反向匹配，true/false。使用说明：固定为false。该字段能更新但不会生效。
	Key          string         `json:"key,omitempty"`            //匹配内容的键值。 使用说明：目前匹配内容为HOST_NAME和PATH时，该字段不生效。该字段能更新但不会生效。
	TenantId     string         `json:"tenant_id,omitempty"`
	Type         string         `json:"type"`                 //匹配内容：可以为 HOST_NAME，PATH。
	Value        string         `json:"value"`                //匹配内容的值。其值不能包含空格。 使用说明：● 当type为HOST_NAME时，取值范围：String(100)，字符串只能包含英文字母、数字、“-”或“.”，必须以字母或数字开头。● 当type为PATH时，取值范围：String (128)。当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:| /()[]{}，且必须以"/"开头。
	Conditions   RuleConditions `json:"conditions,omitempty"` //转发规则的匹配条件。 当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效。配置了conditions后，字段key、字段value的值无意义。若指定了conditons，该rule的条件数为conditons列表长度。列表中key必须相同，value不允许重复。不支持该字段，请勿使用。
}

type RuleConditions struct {
	Key   string `json:"key,omitempty"`   //匹配项的名称。 当type为HOST_NAME、PATH时，该字段固定为空字符串。key的长度限制[1,40]，只允许包含字母、数字和-。 最小长度： 1 最大长度： 128
	Value string `json:"value,omitempty"` //匹配项的值。● 当type为HOST_NAME时，key固定为空字符串，value表示域名的值。value长度[1， 128]，字符串只能包含英文字母、数字、“-”或“.”，必须以字母或数字开头。● 当type为PATH时， key固定为空字符串，value表示请求路径的值。value长度[1，128]。当转发规则的compare_type为STARTS_WITH、 EQUAL_TO时，字符串只能包含英文字母、数字、 _~';@^-%#&$.*+?,=!:|/()[]{}，且必须以"/"开头。 最小长度： 1 最大长度： 128
}

type DeletePolicyRuleRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	RuleId   string `json:"ruleId"`
	PolicyId string `json:"policyId"`
}

type UpdatePolicyRuleRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	RuleId   string                 `json:"ruleId"`
	PolicyId string                 `json:"policyId"`
	Params   UpdatePolicyRuleParams `json:"params"`
}

type UpdatePolicyRuleParams struct {
	UpdatePolicyRule UpdatePolicyRule `json:"l7rule"`
}

type UpdatePolicyRule struct {
	CompareType string         `json:"compare_type"`         //匹配方式：type为HOST_NAME时可以为EQUAL_TO；type为PATH时可以为REGEX，STARTS_WITH，EQUAL_TO
	Invert      bool           `json:"invert,omitempty"`     //是否反向匹配，true/false。使用说明：固定为false。该字段能更新但不会生效。
	Key         string         `json:"key,omitempty"`        //匹配内容的键值。 使用说明：目前匹配内容为HOST_NAME和PATH时，该字段不生效。该字段能更新但不会生效。
	Type        string         `json:"type"`                 //匹配内容：可以为 HOST_NAME，PATH。
	Value       string         `json:"value"`                //匹配内容的值。其值不能包含空格。 使用说明：● 当type为HOST_NAME时，取值范围：String(100)，字符串只能包含英文字母、数字、“-”或“.”，必须以字母或数字开头。● 当type为PATH时，取值范围：String (128)。当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:| /()[]{}，且必须以"/"开头。
	Conditions  RuleConditions `json:"conditions,omitempty"` //转发规则的匹配条件。 当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效。配置了conditions后，字段key、字段value的值无意义。若指定了conditons，该rule的条件数为conditons列表长度。列表中key必须相同，value不允许重复。不支持该字段，请勿使用。
}

type QueryPolicyRuleRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	RuleId   string `json:"ruleId,omitempty"`
	PolicyId string `json:"policyId"`
}

type QueryPolicyRuleResponse struct {
	CompareType        string           `json:"compare_type"`
	AdminStateUp       bool             `json:"admin_state_up"`
	Id                 string           `json:"id"`
	Invert             bool             `json:"invert"`
	TenantId           string           `json:"tenant_id"`
	Key                string           `json:"key"`
	ProvisioningStatus string           `json:"provisioning_status"`
	Value              string           `json:"value"`
	Type               string           `json:"type"`
	Conditions         []RuleConditions `json:"conditions"`
}

type CreateWhitelistsRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenant_id"`
	Params   CreateWhitelistsParams `json:"params"`
}

type CreateWhitelistsParams struct {
	CreateWhitelists CreateWhitelists `json:"whitelist"`
}

type CreateWhitelists struct {
	EnableWhitelist bool   `json:"enable_whitelist,omitempty"` //是否开启访问控制开关。
	TenantId        string `json:"tenant_id,omitempty"`
	ListenerId      string `json:"listener_id"`
	Whitelist       string `json:"whitelist,omitempty"` //白名单IP列表用,分割
}

type DeleteWhitelistsRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateWhitelistsRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	Id       string                 `json:"id"`
	Params   UpdateWhitelistsParams `json:"params"`
}

type UpdateWhitelistsParams struct {
	UpdateWhitelists UpdateWhitelists `json:"whitelist"`
}

type UpdateWhitelists struct {
	EnableWhitelist bool   `json:"enable_whitelist"`    //是否开启访问控制开关。
	Whitelist       string `json:"whitelist,omitempty"` //白名单IP列表用,分割
}

type QueryWhitelistsRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryWhitelistsResponse struct {
	Id              string `json:"id"`
	EnableWhitelist bool   `json:"enable_whitelist"` //是否开启访问控制开关。
	TenantId        string `json:"tenant_id,omitempty"`
	ListenerId      string `json:"listener_id"`
	Whitelist       string `json:"whitelist"` //白名单IP列表用,分割
}

type CreateCertificateRequest struct {
	Domain   string                  `json:"domain"`
	Token    string                  `json:"token"`
	TenantId string                  `json:"tenant_id"`
	Params   CreateCertificateParams `json:"params"`
}

type CreateCertificateParams struct {
	Certificate    string `json:"certificate"` //服务端公有密钥证书。
	Description    string `json:"description,omitempty"`
	Domain         string `json:"domain,omitempty"`          //服务端证书所签域名。
	EncCertificate string `json:"enc_certificate,omitempty"` //国密加密证书内容。只支持国密环境配置，其他环境默认为空。
	EncPrivateKey  string `json:"enc_private_key,omitempty"` //国密加密证书私钥。只支持国密环境配置，其他环境默认为空。
	Name           string `json:"name,omitempty"`
	PrivateKey     string `json:"private_key"`    //服务端私有密钥。
	Type           string `json:"type,omitempty"` //证书类型，可以为client，server，server_sm。
	TenantId       string `json:"tenant_id,omitempty"`
}

type DeleteCertificateRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type UpdateCertificateRequest struct {
	Domain   string                 `json:"domain"`
	Token    string                 `json:"token"`
	TenantId string                 `json:"tenantId"`
	Id       string                 `json:"id"`
	Params   UpdateWhitelistsParams `json:"params"`
}

type UpdateCertificateParams struct {
	Certificate    string `json:"certificate,omitempty"` //服务端公有密钥证书。
	Description    string `json:"description,omitempty"`
	Domain         string `json:"domain,omitempty"`          //服务端证书所签域名。
	EncCertificate string `json:"enc_certificate,omitempty"` //国密加密证书内容。只支持国密环境配置，其他环境默认为空。
	EncPrivateKey  string `json:"enc_private_key,omitempty"` //国密加密证书私钥。只支持国密环境配置，其他环境默认为空。
	Name           string `json:"name,omitempty"`
	PrivateKey     string `json:"private_key,omitempty"` //服务端私有密钥。
}

type QueryCertificateRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type QueryCertificateResponse struct {
	Id             string `json:"id"`
	Certificate    string `json:"certificate"` //服务端公有密钥证书。
	CreateTime     string `json:"create_time"`
	Description    string `json:"description,"`
	Domain         string `json:"domain"` //服务端证书所签域名。
	Name           string `json:"name"`
	PrivateKey     string `json:"private_key"` //服务端私有密钥。
	Type           string `json:"type"`        //证书类型，可以为client，server，server_sm。
	UpdateTime     string `json:"update_time"`
	ExpireTime     string `json:"expire_time"`
	TenantId       string `json:"tenant_id,omitempty"`
	EncCertificate string `json:"enc_certificate"` //国密加密证书内容。只支持国密环境配置，其他环境默认为空。
	EncPrivateKey  string `json:"enc_private_key"` //国密加密证书私钥。只支持国密环境配置，其他环境默认为空。
	AdminStateUp   bool   `json:"admin_state_up"`
}

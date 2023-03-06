package hcs

import (
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

// vpc
func CreateVpc(params model.CreateVpcRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcs`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryVpcResponse
	utils.FromJSON(utils.ToJSON(res["vpc"]), &obj)

	return models.Success[any](obj)
}

func DeleteVpc(params model.DeleteVpcRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcs/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateVpc(params model.UpdateVpcRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcs/%s`, params.Domain, params.TenantId, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryVpcResponse
	utils.FromJSON(utils.ToJSON(res["vpc"]), &obj)

	return models.Success[any](obj)
}

func QueryVpc(params model.QueryVpcRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v1/%s/vpcs`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryVpcResponse
		utils.FromJSON(utils.ToJSON(res["vpcs"]), &finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v1/%s/vpcs/%s`, params.Domain, params.TenantId, params.Id)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		if res["code"] != nil {
			return models.Error(-1, fmt.Sprintf(`Router %s could not be found`, params.Id))
		}
		var obj model.QueryVpcResponse
		utils.FromJSON(utils.ToJSON(res["vpc"]), &obj)

		return models.Success[any](obj)
	}
}

// 子网
func CreateSubnet(params model.CreateSubnetRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/subnets`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QuerySubnetResponse
	utils.FromJSON(utils.ToJSON(res["subnet"]), &obj)

	return models.Success[any](obj)
}

func DeleteSubnet(params model.DeleteSubnetRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/subnets/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateSubnet(params model.UpdateSubnetRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcs/%s/subnets/%s`, params.Domain, params.TenantId, params.VpcId, params.SubnetId)
	_, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	url = fmt.Sprintf(`%s/v1/%s/subnets/%s`, params.Domain, params.TenantId, params.SubnetId)
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QuerySubnetResponse
	utils.FromJSON(utils.ToJSON(res["subnet"]), &obj)

	return models.Success[any](obj)
}

func QuerySubnet(params model.QuerySubnetRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v1/%s/subnets`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QuerySubnetResponse
		utils.FromJSON(utils.ToJSON(res["subnets"]), &finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v1/%s/subnets/%s`, params.Domain, params.TenantId, params.Id)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		if res["code"] != nil {
			return models.Error(-1, fmt.Sprintf(`Router %s could not be found`, params.Id))
		}
		var obj model.QuerySubnetResponse
		utils.FromJSON(utils.ToJSON(res["subnet"]), &obj)

		return models.Success[any](obj)
	}
}

// 私有ip
func CreatePrivateIp(params model.CreatePrivateIpRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/privateips`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list []model.QueryPrivateIpResponse
	utils.FromJSON(utils.ToJSON(res["privateips"]), &list)

	return models.Success[any](list)
}

func DeletePrivateIp(params model.DeletePrivateIpRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/privateips/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func QueryPrivateIp(params model.QueryPrivateIpRequest) models.Result[any] {
	if params.PrivateIpId == "" {
		url := fmt.Sprintf(`%s/v1/%s/subnets/%s/privateips`, params.Domain, params.TenantId, params.SubnetId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryPrivateIpResponse
		utils.FromJSON(utils.ToJSON(res["privateips"]), &finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v1/%s/privateips/%s`, params.Domain, params.TenantId, params.PrivateIpId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		if res["code"] != nil {
			return models.Error(-1, fmt.Sprintf(`Router %s could not be found`, params.PrivateIpId))
		}
		var obj model.QueryPrivateIpResponse
		utils.FromJSON(utils.ToJSON(res["privateip"]), &obj)

		return models.Success[any](obj)
	}
}

// 虚拟ip
func QueryVip(params model.QueryVipRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/ports?device_owner=neutron:VIP_PORT`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryPrivateIpResponse
		utils.FromJSON(utils.ToJSON(res["ports"]), &list)

		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2.0/ports/%s`, params.Domain, params.Id)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		obj := model.QueryPrivateIpResponse{}
		utils.FromJSON(utils.ToJSON(res["port"]), &obj)

		return models.Success[any](obj)
	}

}

func BandVip(params model.BandVipRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vport/%s`, params.Domain, params.TenantId, params.PrivateIpId)
	_, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UnBandVip(params model.UnBandVipRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vport/%s`, params.Domain, params.TenantId, params.PrivateIpId)
	_, err := request.Delete(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

// qos策略
func CreateQosPolicy(params model.CreateQosPolicyRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/qos/policies`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryQosPolicyResponse
	utils.FromJSON(utils.ToJSON(res["policy"]), &obj)

	return models.Success[any](obj)
}

func DeleteQosPolicy(params model.DeleteQosPolicyRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/qos/policies/%s`, params.Domain, params.PolicyId)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateQosPolicy(params model.UpdateQosPolicyRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/qos/policies/%s`, params.Domain, params.PolicyId)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryQosPolicyResponse
	utils.FromJSON(utils.ToJSON(res["policy"]), &obj)

	return models.Success[any](obj)
}

func QueryQosPolicy(params model.QueryQosPolicyRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/qos/policies`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryQosPolicyResponse
		utils.FromJSON(utils.ToJSON(res["policies"]), &finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v2.0/qos/policies/%s`, params.Domain, params.Id)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		if res["NeutronError"] != nil {
			var errObj model.ItemNotFound
			utils.FromJSON(utils.ToJSON(res["NeutronError"]), &errObj)
			return models.Error(-1, errObj.Message)
		}
		var obj model.QueryQosPolicyResponse
		utils.FromJSON(utils.ToJSON(res["policy"]), &obj)

		return models.Success[any](obj)
	}
}

func CreateQosPolicyRule(params model.CreateQosPolicyRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/qos/policies/%s/bandwidth_limit_rules`, params.Domain, params.PolicyId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryQosPolicyResponse
	utils.FromJSON(utils.ToJSON(res["policy"]), &obj)

	return models.Success[any](obj)
}

// todo	端口限速规则
func QueryPolicyBandwidthRule(params model.QueryPolicyBandwidthRuleRequest) models.Result[any] {
	if params.RuleId == "" {
		url := fmt.Sprintf(`%s/v2.0/qos/policies/%s/bandwidth_limit_rules`, params.Domain, params.PolicyId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryQosPolicyResponse
		utils.FromJSON(utils.ToJSON(res["policies"]), &finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v2.0/qos/policies/%s/bandwidth_limit_rules/%s`, params.Domain, params.PolicyId, params.RuleId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		if res["NeutronError"] != nil {
			var errObj model.ItemNotFound
			utils.FromJSON(utils.ToJSON(res["NeutronError"]), &errObj)
			return models.Error(-1, errObj.Message)
		}
		var obj model.QueryQosPolicyResponse
		utils.FromJSON(utils.ToJSON(res["policy"]), &obj)

		return models.Success[any](obj)
	}
}

// 查询外部网络
func QueryExternalNetworks(params model.QueryExternalNetworksRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/external_networks`, params.Domain, params.TenantId)
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var finalList []model.QueryExternalNetworksResponse
	utils.FromJSON(utils.ToJSON(res["vpc_external_networks"]), &finalList)

	return models.Success[any](finalList)
}

// 查询公有网络
func QueryPublicNet(params model.QueryPublicNetRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/networks?router:external=True&tags=service_type=Internet`, params.Domain)
	if params.Id != "" {
		url = fmt.Sprintf(`%s/v2.0/networks?router:external=True&tags=service_type=Internet&id=%s`, params.Domain, params.Id)
	}
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)

	var list []model.QueryPublicNetResponse
	utils.FromJSON(utils.ToJSON(res["networks"]), &list)

	return models.Success[any](list)
}

func GetIpCapacity(params model.GetIpCapacityRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/network-ip-availabilities/%s`, params.Domain, params.PublicNetId)
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.GetIpCapacityResponse
	utils.FromJSON(utils.ToJSON(res["network_ip_availability"]), &result)
	return models.Success[any](result)
}

// 对等连接
func CreatePeer(params model.CreatePeerRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcpeering`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryPeerResponse
	utils.FromJSON(utils.ToJSON(res["vpc_peering_connection"]), &obj)

	return models.Success[any](obj)
}

func DeletePeer(params model.DeletePeerRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcpeering/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdatePeer(params model.UpdatePeerRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcpeering/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func QueryPeer(params model.QueryPeerRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v1/%s/vpcpeering`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryPeerResponse
		utils.FromJSON(utils.ToJSON(res["vpc_peering_connections"]), &finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v1/%s/vpcpeering`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryPeerResponse
		utils.FromJSON(utils.ToJSON(res["vpc_peering_connections"]), &finalList)
		if len(finalList) == 0 {
			return models.Error(-1, fmt.Sprintf(`Router %s could not be found`, params.Id))
		}
		return models.Success[any](finalList[0])
	}
}

func AcceptPeer(params model.ActionPeerRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcpeering/%s/accept`, params.Domain, params.TenantId, params.PeerId)
	dataStr, err := request.Put(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryPeerResponse
	utils.FromJSON(utils.ToJSON(res["vpc_peering_connection"]), &obj)

	return models.Success[any](obj)
}

func RejectPeer(params model.ActionPeerRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcpeering/%s/reject`, params.Domain, params.TenantId, params.PeerId)
	dataStr, err := request.Post(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryPeerResponse
	utils.FromJSON(utils.ToJSON(res["vpc_peering_connection"]), &obj)

	return models.Success[any](obj)
}

// 添加本端路由或者添加对端路由取决于vpcId的值
func CreatePeerRouter(params model.ActionPeerRouterRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcpeering/router/%s/addroutes`, params.Domain, params.TenantId, params.VpcId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var finalList []model.QueryPeerRouterResponse
	utils.FromJSON(utils.ToJSON(res["routes"]), &finalList)

	return models.Success[any](finalList)
}

func DeletePeerRouter(params model.ActionPeerRouterRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcpeering/router/%s/removeroutes`, params.Domain, params.TenantId, params.VpcId)
	dataStr, err := request.Delete(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var finalList []model.QueryPeerRouterResponse
	utils.FromJSON(utils.ToJSON(res["routes"]), &finalList)

	return models.Success[any](finalList)
}

func QueryPeerRouter(params model.QueryPeerRouterRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/vpcpeering/router/%s/queryroutes`, params.Domain, params.TenantId, params.VpcId)
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var finalList []model.QueryPeerRouterResponse
	utils.FromJSON(utils.ToJSON(res["routes"]), &finalList)

	return models.Success[any](finalList)
}

// 端口
func CreatePorts(params model.CreatePortsRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/ports`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryPortsResponse
	utils.FromJSON(utils.ToJSON(res["port"]), &obj)

	return models.Success[any](obj)
}

func DeletePorts(params model.DeletePortsRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/ports/%s`, params.Domain, params.TenantId, params.PortId)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdatePorts(params model.UpdatePortsRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/ports/%s`, params.Domain, params.TenantId, params.PortId)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryPortsResponse
	utils.FromJSON(utils.ToJSON(res["port"]), &obj)

	return models.Success[any](obj)
}

func QueryPorts(params model.QueryPortsRequest) models.Result[any] {
	if params.PortId == "" {
		url := fmt.Sprintf(`%s/v1/%s/ports`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryPortsResponse
		utils.FromJSON(utils.ToJSON(res["ports"]), &finalList)
		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v1/%s/ports/%s`, params.Domain, params.TenantId, params.PortId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var obj model.QueryPortsResponse
		utils.FromJSON(utils.ToJSON(res["port"]), &obj)
		return models.Success[any](obj)
	}
}

// 路由
// todo No VPC peering exist with id ff808082859c57d101861162b3fa0053"
func CreateRoutes(params model.CreateRoutesRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/vpc/routes`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryRoutesResponse
	utils.FromJSON(utils.ToJSON(res["route"]), &obj)

	return models.Success[any](obj)
}

func DeleteRoutes(params model.DeleteRoutesRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/vpc/routes/%s`, params.Domain, params.RouteId)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return models.Success[any](nil)
}

func QueryRoutes(params model.QueryRoutesRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/vpc/routes`, params.Domain)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryRoutesResponse
		utils.FromJSON(utils.ToJSON(res["routes"]), &finalList)
		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v2.0/vpc/routes/%s`, params.Domain, params.Id)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var obj model.QueryRoutesResponse
		utils.FromJSON(utils.ToJSON(res["route"]), &obj)
		return models.Success[any](obj)
	}
}

// NAT网关
// todo the API does not exist or has not been published in the environment
func CreateNatGateways(params model.CreateNatGatewaysRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/nat_gateways`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryNatGatewaysResponse
	utils.FromJSON(utils.ToJSON(res["nat_gateway"]), &obj)

	return models.Success[any](obj)
}

func DeleteNatGateways(params model.DeleteNatGatewaysRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/nat_gateways/%s`, params.Domain, params.NatGatewayId)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateNatGateways(params model.UpdateNatGatewaysRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/nat_gateways/%s`, params.Domain, params.NatGatewayId)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryNatGatewaysResponse
	utils.FromJSON(utils.ToJSON(res["nat_gateway"]), &obj)

	return models.Success[any](obj)
}

func QueryNatGateways(params model.QueryNatGatewaysRequest) models.Result[any] {
	if params.NatGatewayId == "" {
		url := fmt.Sprintf(`%s/v2.0/%s/nat_gateways`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryNatGatewaysResponse
		utils.FromJSON(utils.ToJSON(res["nat_gateways"]), &finalList)
		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v2.0/nat_gateways/%s`, params.Domain, params.NatGatewayId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var obj model.QueryNatGatewaysResponse
		utils.FromJSON(utils.ToJSON(res["nat_gateway"]), &obj)
		return models.Success[any](obj)
	}
}

// SNAT规则	the API does not exist or has not been published in the environment
func CreateSNatRule(params model.CreateSNatRulesRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/snat_rules`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QuerySNatRulesResponse
	utils.FromJSON(utils.ToJSON(res["snat_rule"]), &obj)

	return models.Success[any](obj)
}

func DeleteSNatRules(params model.DeleteSNatRulesRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/snat_rules/%s`, params.Domain, params.SNatRuleId)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func QuerySNatRules(params model.QuerySNatRulesRequest) models.Result[any] {
	if params.SNatRuleId == "" {
		url := fmt.Sprintf(`%s/v2.0/snat_rules`, params.Domain)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QuerySNatRulesResponse
		utils.FromJSON(utils.ToJSON(res["snat_rules"]), &finalList)
		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v2.0/snat_rules/%s`, params.Domain, params.NatGatewayId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var obj model.QuerySNatRulesResponse
		utils.FromJSON(utils.ToJSON(res["snat_rule"]), &obj)
		return models.Success[any](obj)
	}
}

// DNAT规则	the API does not exist or has not been published in the environment
func CreateDNatRule(params model.CreateDNatRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/dnat_rules`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryDNatRulesResponse
	utils.FromJSON(utils.ToJSON(res["nat_gateway"]), &obj)

	return models.Success[any](obj)
}

func DeleteDNatRule(params model.DeleteDNatRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/dnat_rules/%s`, params.Domain, params.DNatRuleId)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateDNatRule(params model.UpdateDNatRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/dnat_rules/%s`, params.Domain, params.DNatRuleId)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryDNatRulesResponse
	utils.FromJSON(utils.ToJSON(res["nat_gateway"]), &obj)

	return models.Success[any](obj)
}

func QueryDNatRules(params model.QueryDNatRulesRequest) models.Result[any] {
	if params.DNatRuleId == "" {
		url := fmt.Sprintf(`%s/v2.0/dnat_rules`, params.Domain)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryDNatRulesResponse
		utils.FromJSON(utils.ToJSON(res["dnat_rules"]), &finalList)
		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v2.0/dnat_rules/%s`, params.Domain, params.DNatRuleId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var obj model.QueryDNatRulesResponse
		utils.FromJSON(utils.ToJSON(res["dnat_rule"]), &obj)
		return models.Success[any](obj)
	}
}

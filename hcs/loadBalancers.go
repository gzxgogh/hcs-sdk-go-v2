package hcs

import (
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

// 负载均衡
func CreateLoadBalancer(params model.CreateLoadBalancersRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/loadbalancers`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryLoadBalancersResponse
	utils.FromJSON(utils.ToJSON(res["loadbalancer"]), &result)
	return models.Success[any](result)
}

func DeleteLoadBalancer(params model.DeleteLoadBalancersRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/loadbalancers/%s`, params.Domain, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateLoadBalancer(params model.UpdateLoadBalancersRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/loadbalancers/%s`, params.Domain, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryLoadBalancersResponse
	utils.FromJSON(utils.ToJSON(res["loadbalancer"]), &result)
	return models.Success[any](result)
}

func QueryLoadBalancer(params model.QueryLoadBalancersRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/lbaas/loadbalancers`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryLoadBalancersResponse
		utils.FromJSON(utils.ToJSON(res["loadbalancers"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2.0/lbaas/loadbalancers/%s`, params.Domain, params.Id)
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
		var obj model.QueryLoadBalancersResponse
		utils.FromJSON(utils.ToJSON(res["loadbalancer"]), &obj)
		return models.Success[any](obj)
	}
}

// 侦听器
func CreateListener(params model.CreateListenersRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/listeners`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryListenersResponse
	utils.FromJSON(utils.ToJSON(res["listener"]), &result)
	return models.Success[any](result)
}

func DeleteListener(params model.DeleteListenersRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/listeners/%s`, params.Domain, params.Id)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateListener(params model.UpdateListenersRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/listeners/%s`, params.Domain, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryListenersResponse
	utils.FromJSON(utils.ToJSON(res["listener"]), &result)
	return models.Success[any](result)
}

func QueryListener(params model.QueryListenersRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/lbaas/listeners`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryListenersResponse
		utils.FromJSON(utils.ToJSON(res["listeners"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2.0/lbaas/listeners/%s`, params.Domain, params.Id)
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
		var obj model.QueryListenersResponse
		utils.FromJSON(utils.ToJSON(res["listener"]), &obj)
		return models.Success[any](obj)
	}
}

// 监听池
func CreatePool(params model.CreatePoolRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/pools`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryPoolResponse
	utils.FromJSON(utils.ToJSON(res["pool"]), &result)
	return models.Success[any](result)
}

func DeletePool(params model.DeletePoolRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/pools/%s`, params.Domain, params.Id)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdatePool(params model.UpdatePoolRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/pools/%s`, params.Domain, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryPoolResponse
	utils.FromJSON(utils.ToJSON(res["pool"]), &result)
	return models.Success[any](result)
}

func QueryPool(params model.QueryPoolRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/lbaas/pools`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryPoolResponse
		utils.FromJSON(utils.ToJSON(res["pools"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2.0/lbaas/pools/%s`, params.Domain, params.Id)
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
		var obj model.QueryPoolResponse
		utils.FromJSON(utils.ToJSON(res["pool"]), &obj)
		return models.Success[any](obj)
	}
}

// 指标参数
func CreateFlavor(params model.CreateFlavorRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/flavors`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryFlavorResponse
	utils.FromJSON(utils.ToJSON(res["flavor"]), &result)
	return models.Success[any](result)
}

func DeleteFlavor(params model.DeleteFlavorRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/flavors/%s`, params.Domain, params.Id)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateFlavor(params model.UpdateFlavorRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/flavors/%s`, params.Domain, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryFlavorResponse
	utils.FromJSON(utils.ToJSON(res["flavor"]), &result)
	return models.Success[any](result)
}

func QueryFlavor(params model.QueryFlavorRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/lbaas/flavors`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryFlavorResponse
		utils.FromJSON(utils.ToJSON(res["flavors"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2.0/lbaas/flavors/%s`, params.Domain, params.Id)
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
		var obj model.QueryFlavorResponse
		utils.FromJSON(utils.ToJSON(res["flavor"]), &obj)
		return models.Success[any](obj)
	}
}

// 转发策略
func CreatePolicy(params model.CreatePolicyRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryPolicyResponse
	utils.FromJSON(utils.ToJSON(res["l7policy"]), &result)
	return models.Success[any](result)
}

func DeletePolicy(params model.DeletePolicyRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies/%s`, params.Domain, params.Id)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdatePolicy(params model.UpdatePolicyRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies/%s`, params.Domain, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryPolicyResponse
	utils.FromJSON(utils.ToJSON(res["l7policy"]), &result)
	return models.Success[any](result)
}

func QueryPolicy(params model.QueryPolicyRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryPolicyResponse
		utils.FromJSON(utils.ToJSON(res["l7policies"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies/%s`, params.Domain, params.Id)
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
		var obj model.QueryPolicyResponse
		utils.FromJSON(utils.ToJSON(res["l7policy"]), &obj)
		return models.Success[any](obj)
	}
}

// 转发规则
func CreatePolicyRule(params model.CreatePolicyRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies/%s/rules`, params.Domain, params.PolicyId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryPolicyRuleResponse
	utils.FromJSON(utils.ToJSON(res["rule"]), &result)
	return models.Success[any](result)
}

func DeletePolicyRule(params model.DeletePolicyRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies/%s/rules/%s`, params.Domain, params.PolicyId, params.RuleId)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdatePolicyRule(params model.UpdatePolicyRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies/%s/rules/%s`, params.Domain, params.PolicyId, params.RuleId)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryPolicyRuleResponse
	utils.FromJSON(utils.ToJSON(res["rule"]), &result)
	return models.Success[any](result)
}

func QueryPolicyRule(params model.QueryPolicyRuleRequest) models.Result[any] {
	if params.RuleId == "" {
		url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies/%s/rules`, params.Domain, params.PolicyId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryPolicyRuleResponse
		utils.FromJSON(utils.ToJSON(res["rules"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2.0/lbaas/l7policies/%s/rules/%s`, params.Domain, params.PolicyId, params.RuleId)
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
		var obj model.QueryPolicyRuleResponse
		utils.FromJSON(utils.ToJSON(res["rule"]), &obj)
		return models.Success[any](obj)
	}
}

// 白名单
func CreateWhitelist(params model.CreateWhitelistsRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/whitelists`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryWhitelistsResponse
	utils.FromJSON(utils.ToJSON(res["whitelist"]), &result)
	return models.Success[any](result)
}

func DeleteWhitelist(params model.DeleteWhitelistsRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/whitelists/%s`, params.Domain, params.Id)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateWhitelist(params model.UpdateWhitelistsRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/whitelists/%s`, params.Domain, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryWhitelistsResponse
	utils.FromJSON(utils.ToJSON(res["whitelist"]), &result)
	return models.Success[any](result)
}

func QueryWhitelist(params model.QueryWhitelistsRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/lbaas/whitelists`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryWhitelistsResponse
		utils.FromJSON(utils.ToJSON(res["whitelists"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2.0/lbaas/whitelists/%s`, params.Domain, params.Id)
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
		var obj model.QueryWhitelistsResponse
		utils.FromJSON(utils.ToJSON(res["whitelist"]), &obj)
		return models.Success[any](obj)
	}
}

// ssl证书
func CreateCertificate(params model.CreateCertificateRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/certificates`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryCertificateResponse
	utils.FromJSON(utils.ToJSON(res["certificate"]), &result)
	return models.Success[any](result)
}

func DeleteCertificate(params model.DeleteCertificateRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/certificates/%s`, params.Domain, params.Id)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateCertificate(params model.UpdateCertificateRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/lbaas/certificates/%s`, params.Domain, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryCertificateResponse
	utils.FromJSON(utils.ToJSON(res["certificate"]), &result)
	return models.Success[any](result)
}

func QueryCertificate(params model.QueryCertificateRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.0/lbaas/certificates`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryCertificateResponse
		utils.FromJSON(utils.ToJSON(res["certificates"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2.0/lbaas/certificates/%s`, params.Domain, params.Id)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		if res["code"] != nil {
			return models.Error(-1, fmt.Sprint(res["message"]))
		}
		var obj model.QueryCertificateResponse
		utils.FromJSON(utils.ToJSON(res["certificate"]), &obj)
		return models.Success[any](obj)
	}
}

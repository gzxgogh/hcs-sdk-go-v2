package hcs

import (
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

func CreateSg(params model.CreateSgRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/security-groups`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QuerySgResponse
	utils.FromJSON(utils.ToJSON(res["security_group"]), &obj)
	return models.Success[any](obj)
}

func DeleteSg(params model.DeleteSgRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/security-groups/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateSg(params model.UpdateSgRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/security-groups/%s`, params.Domain, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QuerySgResponse
	utils.FromJSON(utils.ToJSON(res["security_group"]), &obj)
	return models.Success[any](obj)
}

func QuerySg(params model.QuerySgRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v1/%s/security-groups`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QuerySgResponse
		utils.FromJSON(utils.ToJSON(res["security_groups"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v1/%s/security-groups/%s`, params.Domain, params.TenantId, params.Id)
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
		var obj model.QuerySgResponse
		utils.FromJSON(utils.ToJSON(res["security_group"]), &obj)
		return models.Success[any](obj)
	}
}

func CreateSgRule(params model.CreateSgRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/security-group-rules`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.SgRule
	utils.FromJSON(utils.ToJSON(res["security_group_rule"]), &obj)
	return models.Success[any](obj)
}

func DeleteSgRule(params model.DeleteSgRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/security-group-rules/%s`, params.Domain, params.TenantId, params.RuleId)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func QuerySgRule(params model.QuerySgRuleRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/security-group-rules`, params.Domain, params.TenantId)
	dataStr, err := request.Get(url, params.Token, params)
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
	var list []model.SgRule
	utils.FromJSON(utils.ToJSON(res["security_group_rules"]), &list)
	return models.Success[any](list)
}

func BandNicToSg(params model.BandNicToSgRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/ports/%s`, params.Domain, params.PortId)
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

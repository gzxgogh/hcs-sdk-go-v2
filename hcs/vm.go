package hcs

import (
	"fmt"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
	"hcs-sdk-go-v2/model"
	"hcs-sdk-go-v2/request"
	"time"
)

func CreateVm(params model.CreateVmRequest) models.Result[any] {
	url := fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2/%s/servers`, params.TenantId)
	dataStr, err := request.Post(url, params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var servers model.CreateVmResponse
	utils.FromJSON(utils.ToJSON(res["servers"]), &servers)

	return models.Success[any](servers)
}

func QueryVm(params model.QueryVmRequest) models.Result[any] {
	if params.ServerId == "" {
		url := fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2/%s/servers`, params.TenantId)
		dataStr, err := request.Get(url, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var flavors []map[string]interface{}
		utils.FromJSON(utils.ToJSON(res["servers"]), &flavors)

		var finalList []model.QueryVmResponse
		for _, item := range flavors {
			newUrl := url + "/" + fmt.Sprint(item["id"])
			dataStr, err := request.Get(newUrl, params)
			if err != nil {
				return models.Error(-1, err.Error())
			}
			res := make(map[string]interface{})
			utils.FromJSON(dataStr, &res)

			var result model.QueryVmResponse
			utils.FromJSON(utils.ToJSON(res["server"]), &result)
			finalList = append(finalList, result)
		}
		logs.Debug("最终的结果:{}", finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2/%s/servers/%s`, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)

		return models.Success[any](vmInstance)
	}
}

func StartVm(params model.StartVmRequest) models.Result[any] {
	url := fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2/%s/servers/%s/action`, params.TenantId, params.ServerId)
	_, err := request.VmActionPost(url, params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	for {
		url = fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2/%s/servers/%s`, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)
		if vmInstance.Status == "ACTIVE" {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return models.Success[any](nil)
}

func StopVm(params model.StopVmRequest) models.Result[any] {
	url := fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2/%s/servers/%s/action`, params.TenantId, params.ServerId)
	_, err := request.VmActionPost(url, params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	for {
		url = fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2/%s/servers/%s`, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)
		if vmInstance.Status == "SHUTOFF" {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return models.Success[any](nil)
}

func RebootVm(params model.RebootVmRequest) models.Result[any] {
	url := fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2/%s/servers/%s/action`, params.TenantId, params.ServerId)
	_, err := request.VmActionPost(url, params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	for {
		url = fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2/%s/servers/%s`, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)
		if vmInstance.Status == "ACTIVE" {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return models.Success[any](nil)
}

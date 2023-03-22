package hcs

import (
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

func GetObjType() models.Result[any] {
	list := model.ObjTypeList
	return models.Success[any](list)
}

func GetIndicator(params model.GetIndicatorRequest) models.Result[any] {
	var list []model.Indicator
	switch params.ObjTypeName {
	case "VM":
		list = model.VmIndicatorList
	case "VOLUME":
		list = model.VolumeIndicatorList
	case "EIP":
		list = model.EipIndicatorList
	case "NIC":
		list = model.NicIndicatorList
	case "BANDWIDTH":
		list = model.BandwidthIndicatorList
	}
	return models.Success[any](list)
}

func GetCloudVm(params model.GetInstanceRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_VM?pageNo=%d&pageSize=%d`, params.Domain, params.PageNo, params.PageSize)
	if params.NativeId != "" {
		url = fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_VM?condition={"constraint":[{"simple":{"name":"nativeId","value":"%s","operator":"equal"}}]}`, params.Domain, params.NativeId)
	}
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list []model.GetCloudVmResponse
	utils.FromJSON(utils.ToJSON(res["objList"]), &list)

	return models.Success[any](list)
}

func GetCloudVolume(params model.GetInstanceRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_VOLUME?pageNo=%d&pageSize=%d`, params.Domain, params.PageNo, params.PageSize)
	if params.NativeId != "" {
		url = fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_VOLUME?condition={"constraint":[{"simple":{"name":"nativeId","value":"%s","operator":"equal"}}]}`, params.Domain, params.NativeId)
	}
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list []model.GetCloudVolumeResponse
	utils.FromJSON(utils.ToJSON(res["objList"]), &list)

	return models.Success[any](list)
}

func GetCloudEip(params model.GetInstanceRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_EIP?pageNo=%d&pageSize=%d`, params.Domain, params.PageNo, params.PageSize)
	if params.NativeId != "" {
		url = fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_EIP?condition={"constraint":[{"simple":{"name":"nativeId","value":"%s","operator":"equal"}}]}`, params.Domain, params.NativeId)
	}
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list []model.GetCloudEipResponse
	utils.FromJSON(utils.ToJSON(res["objList"]), &list)

	return models.Success[any](list)
}

func GetCloudNic(params model.GetInstanceRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_VM_NIC?pageNo=%d&pageSize=%d`, params.Domain, params.PageNo, params.PageSize)
	if params.NativeId != "" {
		url = fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_VM_NIC?condition={"constraint":[{"simple":{"name":"nativeId","value":"%s","operator":"equal"}}]}`, params.Domain, params.NativeId)
	}
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list []model.GetCloudNicResponse
	utils.FromJSON(utils.ToJSON(res["objList"]), &list)

	return models.Success[any](list)
}

func GetCloudBandwidth(params model.GetInstanceRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_BANDWIDTHS?pageNo=%d&pageSize=%d`, params.Domain, params.PageNo, params.PageSize)
	if params.NativeId != "" {
		url = fmt.Sprintf(`%s/rest/tenant-resource/v1/instances/CLOUD_BANDWIDTHS?condition={"constraint":[{"simple":{"name":"nativeId","value":"%s","operator":"equal"}}]}`, params.Domain, params.NativeId)
	}
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list []model.GetCloudBandwidthResponse
	utils.FromJSON(utils.ToJSON(res["objList"]), &list)

	return models.Success[any](list)
}

func GetPerformance(params model.GetPerformanceRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/rest/performance/v1/data-svc/history-data/action/query`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var data map[string]interface{}
	utils.FromJSON(utils.ToJSON(res["data"]), &data)

	return models.Success[any](data)
}

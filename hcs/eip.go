package hcs

import (
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

// 创建并绑定弹性ip
func CreateEipAndBand(params model.CreateFloatIpRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/floatingips`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.CreateFloatIpResponse
	utils.FromJSON(utils.ToJSON(res["floatingip"]), &result)
	return models.Success[any](result)
}

func CreateEip(params model.CreateEipRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/publicips`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryEipResponse
	utils.FromJSON(utils.ToJSON(res["publicip"]), &obj)
	return models.Success[any](obj)
}

func DeleteEip(params model.DeleteEipRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/publicips/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return models.Success[any](nil)
}

// 可以绑定和解绑弹性ip
func UpdateEip(params model.UpdateEipRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/publicips/%s`, params.Domain, params.TenantId, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryEipResponse
	utils.FromJSON(utils.ToJSON(res["publicip"]), &obj)
	return models.Success[any](obj)
}

func QueryEip(params model.QueryEipRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v1/%s/publicips`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryEipResponse
		utils.FromJSON(utils.ToJSON(res["publicips"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v1/%s/publicips/%s`, params.Domain, params.TenantId, params.Id)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}

		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		if res["itemNotFound"] != nil {
			var errObj model.ItemNotFound
			utils.FromJSON(utils.ToJSON(res["itemNotFound"]), &errObj)
			return models.Error(-1, errObj.Message)
		}
		var obj model.QueryEipResponse
		utils.FromJSON(utils.ToJSON(res["publicip"]), &obj)
		return models.Success[any](obj)
	}

}

func CreateShareBandwidth(params model.CreateShareBandwidthRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/%s/bandwidths`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.CreateShareBandwidthResponse
	utils.FromJSON(utils.ToJSON(res["bandwidth"]), &obj)
	return models.Success[any](obj)
}

func DeleteShareBandwidth(params model.DeleteShareBandwidthRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/%s/bandwidths/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return models.Success[any](nil)
}

func UpdateShareBandwidth(params model.UpdateShareBandwidthRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/%s/bandwidths/%s`, params.Domain, params.TenantId, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.CreateShareBandwidthResponse
	utils.FromJSON(utils.ToJSON(res["bandwidth"]), &obj)
	return models.Success[any](obj)
}

func QueryShareBandwidth(params model.QueryShareBandwidthRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v1/%s/bandwidths`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryEipResponse
		utils.FromJSON(utils.ToJSON(res["bandwidths"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v1/%s/bandwidths/%s`, params.Domain, params.TenantId, params.Id)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}

		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		if res["code"] != nil {
			return models.Error(-1, fmt.Sprintf(`QoS policy %s could not be found.`, params.Id))
		}
		var obj model.QueryEipResponse
		utils.FromJSON(utils.ToJSON(res["bandwidth"]), &obj)
		return models.Success[any](obj)
	}

}

func ShareBandwidthAttachEip(params model.ShareBandwidthAttachEipRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/%s/bandwidths/%s/insert`, params.Domain, params.TenantId, params.BandwidthId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.ShareBandwidthAttachEipResponse
	utils.FromJSON(utils.ToJSON(res["bandwidth"]), &obj)
	return models.Success[any](obj)
}

func ShareBandwidthDetachEip(params model.ShareBandwidthDetachEipRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.0/%s/bandwidths/%s/remove`, params.Domain, params.TenantId, params.BandwidthId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.ShareBandwidthAttachEipResponse
	utils.FromJSON(utils.ToJSON(res["bandwidth"]), &obj)
	return models.Success[any](obj)
}

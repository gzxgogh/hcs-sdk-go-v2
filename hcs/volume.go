package hcs

import (
	"fmt"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
	"hcs-sdk-go-v2/model"
	"hcs-sdk-go-v2/request"
)

func QueryVolumeType(params model.QueryVolumeTypeRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/types`, params.Domain, params.TenantId)
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list []model.QueryVolumeTypeResponse
	utils.FromJSON(utils.ToJSON(res["volume_types"]), &list)

	return models.Success[any](list)
}

func CreateVolume(params model.CreateVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/volumes`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.CreateVolumeResponse
	utils.FromJSON(utils.ToJSON(res["volume"]), &result)

	return models.Success[any](result)
}

func DeleteVolume(params model.DeleteVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/volumes/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateVolume(params model.UpdateVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/volumes/%s`, params.Domain, params.TenantId, params.Id)
	dataStr, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.CreateVolumeResponse
	utils.FromJSON(utils.ToJSON(res["volume"]), &result)
	return models.Success[any](result)
}

func QueryVolume(params model.QueryVolumeRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`https://%s/v2/%s/volumes/detail`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}

		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryVolumeResponse
		utils.FromJSON(utils.ToJSON(res["volumes"]), &finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`https://%s/v2/%s/volumes/%s`, params.Domain, params.TenantId, params.Id)
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
		var obj model.QueryVolumeResponse
		utils.FromJSON(utils.ToJSON(res["volume"]), &obj)
		return models.Success[any](obj)
	}
}

func UpgradeVolume(params model.UpgradeVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/volumes/%s/action`, params.Domain, params.TenantId, params.Id)
	_, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return models.Success[any](nil)
}

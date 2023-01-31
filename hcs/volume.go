package hcs

import (
	"fmt"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
	"hcs-sdk-go-v2/model"
	"hcs-sdk-go-v2/request"
)

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

func QueryVolume(params model.CreateVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/volumes`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var servers model.CreateVolumeResponse
	utils.FromJSON(utils.ToJSON(res["volume"]), &servers)

	return models.Success[any](servers)
}

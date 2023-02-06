package hcs

import (
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

func CreateImage(params model.CreateImageRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/cloudimages/action`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)

	var result model.QueryImageResponse
	utils.FromJSON(utils.ToJSON(res["image"]), &result)
	return models.Success[any](result)
}

func DeleteImage(params model.DeleteImageRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/images/%s`, params.Domain, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func QueryImage(params model.QueryImageRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2.1/%s/images`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var images []map[string]interface{}
		utils.FromJSON(utils.ToJSON(res["images"]), &images)

		var finalList []model.QueryImageResponse
		for _, item := range images {
			//name := fmt.Sprint(item["name"])
			newUrl := url + "/" + fmt.Sprint(item["id"])
			dataStr, err := request.Get(newUrl, params.Token, params)
			if err != nil {
				return models.Error(-1, err.Error())
			}
			res := make(map[string]interface{})
			utils.FromJSON(dataStr, &res)

			var result model.QueryImageResponse
			utils.FromJSON(utils.ToJSON(res["image"]), &result)
			finalList = append(finalList, result)
		}
		logs.Debug("最终的结果:{}", finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v2.1/%s/images/%s`, params.Domain, params.TenantId, params.Id)
		dataStr, err := request.Get(url, params.Token, params)
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
		var result model.QueryImageResponse
		utils.FromJSON(utils.ToJSON(res["image"]), &result)
		return models.Success[any](result)
	}
}

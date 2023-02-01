package hcs

import (
	"fmt"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
	"hcs-sdk-go-v2/model"
	"hcs-sdk-go-v2/request"
)

func QuerySpec(params model.QuerySpecRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`https://%s/v2.1/%s/flavors`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var flavors []map[string]interface{}
		utils.FromJSON(utils.ToJSON(res["flavors"]), &flavors)

		var finalList []model.QuerySpecResponse
		for _, item := range flavors {
			//name := fmt.Sprint(item["name"])
			newUrl := url + "/" + fmt.Sprint(item["id"])
			dataStr, err := request.Get(newUrl, params.Token, params)
			if err != nil {
				return models.Error(-1, err.Error())
			}
			res := make(map[string]interface{})
			utils.FromJSON(dataStr, &res)

			var result model.QuerySpecResponse
			utils.FromJSON(utils.ToJSON(res["flavor"]), &result)
			finalList = append(finalList, result)
		}
		logs.Debug("最终的结果:{}", finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`https://%s/v2.1/%s/flavors/%s`, params.Domain, params.TenantId, params.Id)
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
		var result model.QuerySpecResponse
		utils.FromJSON(utils.ToJSON(res["flavor"]), &result)
		return models.Success[any](result)
	}

}

func DeleteSpec(params model.DeleteSpecRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2.1/%s/flavors/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return models.Success[any](nil)
}

func CreateSpec(params model.DeleteSpecRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2.1/%s/flavors/%s`, params.Domain, params.TenantId, params.Id)
	_, err := request.Post(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return models.Success[any](nil)
}

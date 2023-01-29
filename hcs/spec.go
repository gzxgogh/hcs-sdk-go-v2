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
	url := fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2.1/%s/flavors`, params.TenantId)
	dataStr, err := request.Get(url, params)
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
		dataStr, err := request.Get(newUrl, params)
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
}

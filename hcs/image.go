package hcs

import (
	"fmt"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
	"hcs-sdk-go-v2/model"
	"hcs-sdk-go-v2/request"
)

func QueryImage(params model.QueryImageRequest) models.Result[any] {
	url := fmt.Sprintf(`https://ecs.cn-global-1.hyy.com:5443/v2.1/%s/images`, params.TenantId)
	dataStr, err := request.Get(url, params)
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
		dataStr, err := request.Get(newUrl, params)
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
}

package hcs

import (
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

func CreateImage(params model.CreateImageRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/cloudimages/wholeimages/action`, params.Domain)
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
		url := fmt.Sprintf(`%s/v2/cloudimages`, params.Domain)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryImageResponse
		utils.FromJSON(utils.ToJSON(res["images"]), &list)

		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2/images/%s`, params.Domain, params.Id)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var result model.QueryImageResponse
		utils.FromJSON(utils.ToJSON(res), &result)
		if result.Id == "" {
			return models.Error(-1, fmt.Sprintf(`Router %s could not be found`, params.Id))
		}
		return models.Success[any](result)
	}
}

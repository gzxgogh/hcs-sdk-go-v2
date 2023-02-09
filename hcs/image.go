package hcs

import (
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

func CreateVmImage(params model.CreateVmImageRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/cloudimages/wholeimages/action`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	var res model.JobResponse
	utils.FromJSON(dataStr, &res)
	if res.Error.Message != "" {
		return models.Error(-1, res.Error.Message)
	}
	_, err = ExecCreateJob(params.EcsDomain, params.TenantId, params.Token, res.JobId)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	//不会告知uuid,通过名称查询镜像找到该镜像。
	return QueryImage(model.QueryImageRequest{
		Domain:   params.Domain,
		TenantId: params.TenantId,
		Token:    params.Token,
		Name:     params.Params.Name,
	})
}

func CreateVolumeImage(params model.CreateVolumeImageRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/cloudimages/action`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	var res model.JobResponse
	utils.FromJSON(dataStr, &res)
	if res.Error.Message != "" {
		return models.Error(-1, res.Error.Message)
	}
	result, err := ExecCreateJob(params.EcsDomain, params.TenantId, params.Token, res.JobId)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	var subJobInfo model.SubJobInfo
	utils.FromJSON(utils.ToJSON(result), &subJobInfo)
	obj := model.CreateVolumeImageResponse{
		ImageId: subJobInfo.Entities["image_id"].(string),
	}
	return models.Success[any](obj)
}

func DeleteImage(params model.DeleteImageRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/cloudimages`, params.Domain)
	_, err := request.Delete(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpdateImage(params model.UpdateImageRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/cloudimages/%s`, params.Domain, params.ImageId)
	dataStr, err := request.PATCH(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.QueryImageResponse
	utils.FromJSON(utils.ToJSON(res), &result)

	return models.Success[any](result)
}

func QueryImage(params model.QueryImageRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2/cloudimages`, params.Domain)
		dataStr, err := request.Get(url, params.Token, params)
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

package hcs

import (
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
)

func CreateVmSnapshot(params model.CreateVmSnapshotRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.1/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	var res model.JobResponse
	utils.FromJSON(dataStr, &res)
	if res.Error.Message != "" {
		return models.Error(-1, res.Error.Message)
	}
	result, err := ExecCreateJob(params.Domain, params.TenantId, params.Token, res.JobId)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	var subJobInfo model.SubJobInfo
	utils.FromJSON(utils.ToJSON(result), &subJobInfo)
	obj := model.CreateVmSnapshotResponse{
		ImageId: subJobInfo.Entities["image_id"].(string),
	}
	return models.Success[any](obj)
}

func RevertVmSnapshot(params model.RevertVmSnapshotRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.1/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	var res model.JobResponse
	utils.FromJSON(dataStr, &res)
	if res.Error.Message != "" {
		return models.Error(-1, res.Error.Message)
	}
	err = ExecJob(params.Domain, params.TenantId, params.Token, res.JobId)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func DeleteVmSnapshot(params model.DeleteVmSnapshotRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/cloudimages`, params.Domain)
	params.Params.IsSnapshotImage = "true"
	_, err := request.Delete(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func QueryVmSnapshot(params model.QuerySnapshotRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/images`, params.Domain)
	dataStr, err := request.Get(url, params.Token, params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list []model.QueryVmSnapshotResponse
	utils.FromJSON(utils.ToJSON(res["images"]), &list)
	return models.Success[any](list)
}

func CreateVolumeSnapshot(params model.CreateVolumeSnapshotRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/snapshots`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.CreateVolumeSnapshotResponse
	utils.FromJSON(utils.ToJSON(res["snapshot"]), &obj)
	return models.Success[any](obj)
}

func RevertVolumeSnapshot(params model.RevertVolumeSnapshotRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/os-vendor-snapshots/%s/rollback`, params.Domain, params.TenantId, params.SnapshotId)
	_, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return models.Success[any](nil)
}

func DeleteVolumeSnapshot(params model.DeleteVolumeSnapshotRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/snapshots/%s`, params.Domain, params.TenantId, params.SnapshotId)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return models.Success[any](nil)
}

func QuerySnapshot(params model.QuerySnapshotRequest) models.Result[any] {
	if params.Id == "" {
		url := fmt.Sprintf(`%s/v2/%s/snapshots/detail`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var list []model.QueryVolumeSnapshotResponse
		utils.FromJSON(utils.ToJSON(res["snapshots"]), &list)
		return models.Success[any](list)
	} else {
		url := fmt.Sprintf(`%s/v2/%s/snapshots/%s`, params.Domain, params.TenantId, params.Id)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var obj model.QueryVolumeSnapshotResponse
		utils.FromJSON(utils.ToJSON(res["snapshot"]), &obj)
		return models.Success[any](obj)
	}
}

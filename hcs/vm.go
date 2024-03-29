package hcs

import (
	"errors"
	"fmt"
	"github.com/gzxgogh/hcs-sdk-go-v2/model"
	"github.com/gzxgogh/hcs-sdk-go-v2/request"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
	"time"
)

func CreateVm(params model.CreateVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1.1/%s/cloudservers`, params.Domain, params.TenantId)
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
	obj := model.CreateVmResponse{
		ServerId: subJobInfo.Entities["server_id"].(string),
	}
	return models.Success[any](obj)
}

func CreateVmFromVolume(params model.CreateVmFromVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var servers model.CreateVmFromVolumeResponse
	utils.FromJSON(utils.ToJSON(res["servers"]), &servers)

	return models.Success[any](servers)
}

func DeleteVm(params model.DeleteVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/cloudservers/delete`, params.Domain, params.TenantId)
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

func UpdateVm(params model.UpdateVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2.1/%s/servers/%s`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return QueryVm(model.QueryVmRequest{
		Domain:   params.Domain,
		TenantId: params.TenantId,
		Token:    params.Token,
		ServerId: params.ServerId,
	})
}

func QueryVm(params model.QueryVmRequest) models.Result[any] {
	if params.ServerId == "" {
		url := fmt.Sprintf(`%s/v1/%s/cloudservers/detail`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var finalList []model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["servers"]), &finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`%s/v1/%s/cloudservers/%s`, params.Domain, params.TenantId, params.ServerId)
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
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)

		return models.Success[any](vmInstance)
	}
}

func StartVm(params model.StartVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Action)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	for i := 0; i < 60; i++ {
		url = fmt.Sprintf(`%s/v2/%s/servers/%s`, params.Domain, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)
		if vmInstance.Status == "ACTIVE" {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return models.Success[any](nil)
}

func StopVm(params model.StopVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Action)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	for i := 0; i < 60; i++ {
		url = fmt.Sprintf(`%s/v2/%s/servers/%s`, params.Domain, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)
		if vmInstance.Status == "SHUTOFF" {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return models.Success[any](nil)
}

func RebootVm(params model.RebootVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Action)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	for i := 0; i < 60; i++ {
		url = fmt.Sprintf(`%s/v2/%s/servers/%s`, params.Domain, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)
		if vmInstance.Status == "ACTIVE" {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return models.Success[any](nil)
}

func ResetVm(params model.ResetVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/cloudservers/%s/reinstallos`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	for i := 0; i < 60; i++ {
		url = fmt.Sprintf(`%s/v2/%s/servers/%s`, params.Domain, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)
		if vmInstance.Status == "ACTIVE" {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return models.Success[any](nil)
}

func QueryVmNic(params model.VmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/os-interface`, params.Domain, params.TenantId, params.ServerId)
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
	var list model.QueryVmNicResponse
	utils.FromJSON(utils.ToJSON(res["interfaceAttachments"]), &list)

	return models.Success[any](list)
}

func AttachVmNic(params model.AttachNicRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/os-interface`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.QueryVmNicResponse
	utils.FromJSON(utils.ToJSON(res["interfaceAttachments"]), &obj)

	return models.Success[any](obj)
}

func BatchAttachVmNic(params model.BatchAttachNicRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/cloudservers/%s/nics`, params.Domain, params.TenantId, params.ServerId)
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

func DetachVmNic(params model.DetachNicRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/os-interface/%s`, params.Domain, params.TenantId, params.ServerId, params.NicId)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func AttachVmVolume(params model.AttachVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/os-volume_attachments`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var obj model.AttachVolumeResponse
	utils.FromJSON(utils.ToJSON(res["volumeAttachment"]), &obj)

	return models.Success[any](obj)
}

func DetachVmVolume(params model.DetachVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/os-volume_attachments/%s`, params.Domain, params.TenantId, params.ServerId, params.VolumeId)
	_, err := request.DeleteWithoutBody(url, params.Token)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func QueryCanUpgradeSpec(params model.QueryCanUpgradeSpecRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/cloudservers/resize_flavors`, params.Domain, params.TenantId)
	dataStr, err := request.Get(url, params.Token, params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list []model.QuerySpecResponse
	utils.FromJSON(utils.ToJSON(res["flavors"]), &list)

	return models.Success[any](list)
}

func UpgradeVm(params model.UpgradeVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1.1/%s/cloudservers/%s/resize`, params.Domain, params.TenantId, params.ServerId)
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

func ChangeVmImage(params model.ChangeVmImageRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/cloudservers/%s/changeos`, params.Domain, params.TenantId, params.ServerId)
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

func LockVm(params model.LockVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UnlockVm(params model.UnlockVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v2/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

// 在线变更规格目前调用失败，原因：check_live_resize fail!
func OnlineUpgradeVm(params model.OnlineUpgradeVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/cloudservers/%s/live-resize`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	//var res model.JobResponse
	//utils.FromJSON(dataStr, &res)
	//
	//err =ExecJob(params.Domain,params.TenantId,params.Token,res.JobId)
	//if err != nil {
	//	return models.Error(-1, err.Error())
	//}

	return models.Success[any](nil)
}

// 克隆云主机调用失败，原因：convert request to action vm fail!
func CloneVm(params model.CloneVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/cloudservers/%s/action`, params.Domain, params.TenantId, params.ServerId)
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

func QueryConsoleAddress(params model.QueryConsoleAddRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/cloudservers/%s/remote_console`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
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
	var result model.QueryConsoleAddResponse
	utils.FromJSON(utils.ToJSON(res["remote_console"]), &result)

	return models.Success[any](result)
}

func GetPassword(params model.UnlockVmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/cloudservers/%s/os-server-password`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

// 使用此API，需预先安装重置密码插件
func ChangeVmPwd(params model.ChangeVmPwdRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/cloudservers/%s/os-reset-password`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Put(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	return models.Success[any](nil)
}

func CheckCanChangePwd(params model.VmRequest) models.Result[any] {
	url := fmt.Sprintf(`%s/v1/%s/cloudservers/%s/os-resetpwd-flag`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	return models.Success[any](res)
}

func ExecJob(domain, tenantId, token, jobId string) error {
	for {
		time.Sleep(10 * time.Second)
		url := fmt.Sprintf(`%s/v1/%s/jobs/%s`, domain, tenantId, jobId)
		dataStr, err := request.Get(url, token, nil)
		if err != nil {
			return err
		}
		var jobRes model.JobInfo
		utils.FromJSON(dataStr, &jobRes)
		if jobRes.Status == "SUCCESS" {
			return nil
		}
		if jobRes.Status == "FAIL" {
			if len(jobRes.Entities.SubJobs) > 0 {
				return errors.New(jobRes.Entities.SubJobs[0].FailReason)
			} else {
				return errors.New(jobRes.FailReason)
			}
		}
		if jobRes.JobId == "" {
			var jobErr model.JobErr
			utils.FromJSON(dataStr, &jobErr)
			return errors.New(jobErr.ErrorMsg)
		}
	}
}

func ExecCreateJob(domain, tenantId, token, jobId string) (interface{}, error) {
	for {
		time.Sleep(10 * time.Second)
		url := fmt.Sprintf(`%s/v1/%s/jobs/%s`, domain, tenantId, jobId)
		dataStr, err := request.Get(url, token, nil)
		if err != nil {
			return model.SubJobInfo{}, err
		}
		var jobRes model.JobInfo
		utils.FromJSON(dataStr, &jobRes)
		if jobRes.Status == "SUCCESS" {
			if len(jobRes.Entities.SubJobs) > 0 {
				return jobRes.Entities.SubJobs[0], nil
			} else {
				var subJobRes model.SubJobInfo
				utils.FromJSON(dataStr, &subJobRes)
				return subJobRes, err
			}
		}
		if jobRes.Status == "FAIL" {
			if len(jobRes.Entities.SubJobs) > 0 {
				return nil, errors.New(jobRes.Entities.SubJobs[0].FailReason)
			} else {
				return nil, errors.New(jobRes.FailReason)
			}
		}
		if jobRes.JobId == "" {
			var jobErr model.JobErr
			utils.FromJSON(dataStr, &jobErr)
			return nil, errors.New(jobErr.ErrorMsg)
		}
	}
}

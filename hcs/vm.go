package hcs

import (
	"fmt"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
	"hcs-sdk-go-v2/model"
	"hcs-sdk-go-v2/request"
	"time"
)

func CreateVm(params model.CreateVmRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/servers`, params.Domain, params.TenantId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var servers model.CreateVmResponse
	utils.FromJSON(utils.ToJSON(res["servers"]), &servers)

	return models.Success[any](servers)
}

func QueryVm(params model.QueryVmRequest) models.Result[any] {
	if params.ServerId == "" {
		url := fmt.Sprintf(`https://%s/v2/%s/servers`, params.Domain, params.TenantId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var flavors []map[string]interface{}
		utils.FromJSON(utils.ToJSON(res["servers"]), &flavors)

		var finalList []model.QueryVmResponse
		for _, item := range flavors {
			newUrl := url + "/" + fmt.Sprint(item["id"])
			dataStr, err := request.Get(newUrl, params.Token, params)
			if err != nil {
				return models.Error(-1, err.Error())
			}
			res := make(map[string]interface{})
			utils.FromJSON(dataStr, &res)

			var result model.QueryVmResponse
			utils.FromJSON(utils.ToJSON(res["server"]), &result)
			finalList = append(finalList, result)
		}
		logs.Debug("最终的结果:{}", finalList)

		return models.Success[any](finalList)
	} else {
		url := fmt.Sprintf(`https://%s/v2/%s/servers/%s`, params.Domain, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params.Token, params)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		res := make(map[string]interface{})
		utils.FromJSON(dataStr, &res)
		var vmInstance model.QueryVmResponse
		utils.FromJSON(utils.ToJSON(res["server"]), &vmInstance)

		return models.Success[any](vmInstance)
	}
}

func StartVm(params model.StartVmRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Action)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	for i := 0; i < 60; i++ {
		url = fmt.Sprintf(`https://%s/v2/%s/servers/%s`, params.Domain, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params.Token, params)
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
	url := fmt.Sprintf(`https://%s/v2/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Action)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	for i := 0; i < 60; i++ {
		url = fmt.Sprintf(`https://%s/v2/%s/servers/%s`, params.Domain, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params.Token, params)
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
	url := fmt.Sprintf(`https://%s/v2/%s/servers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Action)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	for i := 0; i < 60; i++ {
		url = fmt.Sprintf(`https://%s/v2/%s/servers/%s`, params.Domain, params.TenantId, params.ServerId)
		dataStr, err := request.Get(url, params.Token, params)
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
	url := fmt.Sprintf(`https://%s/v2/%s/servers/%s/os-interface`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Get(url, params.Token, params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list model.QueryVmNicResponse
	utils.FromJSON(utils.ToJSON(res["interfaceAttachments"]), &list)

	return models.Success[any](list)
}

func AttachVmNic(params model.AttachNicRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/servers/%s/os-interface`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list model.QueryVmNicResponse
	utils.FromJSON(utils.ToJSON(res["interfaceAttachments"]), &list)

	return models.Success[any](list)
}

func DetachVmNic(params model.DetachNicRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/servers/%s/os-interface/%s`, params.Domain, params.TenantId, params.ServerId, params.NicId)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func AttachVmVolume(params model.AttachVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/servers/%s/os-volume_attachments`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var list model.AttachVolumeResponse
	utils.FromJSON(utils.ToJSON(res["volumeAttachment"]), &list)

	return models.Success[any](list)
}

func DetachVmVolume(params model.DetachVolumeRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2/%s/servers/%s/os-volume_attachments/%s`, params.Domain, params.TenantId, params.ServerId, params.VolumeId)
	_, err := request.Delete(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}

	return models.Success[any](nil)
}

func UpgradeVm(params model.UpgradeVmRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v1.1/%s/cloudservers/%s/resize`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	var res model.JobResponse
	utils.FromJSON(dataStr, &res)

	for i := 0; i < 60; i++ {
		url = fmt.Sprintf(`https://%s/v1/%s/jobs/%s`, params.Domain, params.TenantId, res.JobId)
		dataStr, err = request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		var jobRes model.JobInfo
		utils.FromJSON(dataStr, &jobRes)
		if jobRes.Status == "SUCCESS" {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return models.Success[any](nil)
}

// 在线变更规格目前调用失败，原因：check_live_resize fail!
func OnlineUpgradeVm(params model.OnlineUpgradeVmRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v1/%s/cloudservers/%s/live-resize`, params.Domain, params.TenantId, params.ServerId)
	_, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	//var res model.JobResponse
	//utils.FromJSON(dataStr, &res)
	//
	//for {
	//	url = fmt.Sprintf(`https://%s/v1/%s/jobs/%s`, params.Domain, params.TenantId, res.JobId)
	//	dataStr, err = request.Get(url, params.Token, nil)
	//	if err != nil {
	//		return models.Error(-1, err.Error())
	//	}
	//	var jobRes model.JobInfo
	//	utils.FromJSON(dataStr, &jobRes)
	//	if jobRes.Status == "SUCCESS" {
	//		break
	//	}
	//	time.Sleep(1 * time.Second)
	//}

	return models.Success[any](nil)
}

// 克隆云主机调用失败，原因：convert request to action vm fail!
func CloneVm(params model.CloneVmRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v1/%s/cloudservers/%s/action`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	var res model.JobResponse
	utils.FromJSON(dataStr, &res)

	for i := 0; i < 60; i++ {
		url = fmt.Sprintf(`https://%s/v1/%s/jobs/%s`, params.Domain, params.TenantId, res.JobId)
		dataStr, err = request.Get(url, params.Token, nil)
		if err != nil {
			return models.Error(-1, err.Error())
		}
		var jobRes model.JobInfo
		utils.FromJSON(dataStr, &jobRes)
		if jobRes.Status == "SUCCESS" {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return models.Success[any](nil)
}

// 创建并绑定弹性ip
func CreateEipAndBand(params model.CreateFloatIpRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2.0/floatingips`, params.Domain)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)
	var result model.CreateFloatIpResponse
	utils.FromJSON(utils.ToJSON(res["floatingip"]), &result)
	return models.Success[any](result)
}

func QueryConsoleAddress(params model.QueryConsoleAddRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v1/%s/cloudservers/%s/remote_console`, params.Domain, params.TenantId, params.ServerId)
	dataStr, err := request.Post(url, params.Token, params.Params)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)

	var result model.QueryConsoleAddResponse
	utils.FromJSON(utils.ToJSON(res["remote_console"]), &result)

	return models.Success[any](result)
}

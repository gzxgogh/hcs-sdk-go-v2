package hcs

import (
	"fmt"
	"github.com/maczh/mgin/models"
	"github.com/maczh/mgin/utils"
	"hcs-sdk-go-v2/model"
	"hcs-sdk-go-v2/request"
)

func QueryPublicNet(params model.QueryPublicNetRequest) models.Result[any] {
	url := fmt.Sprintf(`https://%s/v2.0/networks?router:external=True&tags=service_type=Internet`, params.Domain)
	dataStr, err := request.Get(url, params.Token, nil)
	if err != nil {
		return models.Error(-1, err.Error())
	}
	res := make(map[string]interface{})
	utils.FromJSON(dataStr, &res)

	var list []model.QueryPublicNetResponse
	utils.FromJSON(utils.ToJSON(res["networks"]), &list)

	return models.Success[any](list)
}

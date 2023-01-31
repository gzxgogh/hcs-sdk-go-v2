package model

type QuerySpecRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id,omitempty"`
}

type QuerySpecResponse struct {
	Id                     string  `json:"id"`
	Name                   string  `json:"name"`
	Ram                    int     `json:"ram"`
	Disk                   int     `json:"disk"`
	Vcpus                  int     `json:"vcpus"`
	Swap                   string  `json:"swap"`
	OSFLVEXTDATAEphemeral  int     `json:"OS-FLV-EXT-DATA:ephemeral"`
	OSFLVDISABLEDDisabled  bool    `json:"OS-FLV-DISABLED:disabled"`
	RxtxFactor             float64 `json:"rxtx_factor"`
	OsFlavorAccessIsPublic bool    `json:"os-flavor-access:is_public"`
}

type DeleteSpecRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

type CreateSpecRequest struct {
	Domain   string `json:"domain"`
	Token    string `json:"token"`
	TenantId string `json:"tenantId"`
}

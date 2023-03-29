package model

type GetPushDeviceInfoInput struct {
	Id uint `json:"id" ` // 推送设备表主键id
}

type AddPushDeviceInput struct {
	DeviceName string `json:"device_name"      ` // 设备名称
	BaseUrl    string `json:"base_url"         ` // 推送基础URL
}

type SetPushDeviceServiceInput struct {
	Id            uint   `json:"id"               ` // 推送设备表主键id
	ServiceIdList string `json:"service_id_list"  ` // 推送服务id列表
}

type SetPushDeviceServiceOutput struct {
	SuccessList []int `json:"success_list" ` // 成功列表
	FailList    []int `json:"fail_list"    ` // 失败列表
}

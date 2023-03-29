package model

type PushBasicData struct {
	ServiceSign string `json:"service_sign" ` // 推送服务标识
	PushData    string `json:"push_data"    ` // 推送数据
}

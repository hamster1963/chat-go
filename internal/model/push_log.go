package model

import "github.com/gogf/gf/v2/os/gtime"

type AddPushLogInput struct {
	PushInfo    interface{} `json:"push_info"    ` // 推送内容
	PushService string      `json:"push_service" ` // 推送服务
	PushTime    *gtime.Time `json:"push_time"    ` // 推送时间
	PushStatus  bool        `json:"push_status"  ` // 推送状态
	ErrInfo     string      `json:"err_info"     ` // 错误日志
}

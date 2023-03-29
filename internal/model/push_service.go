package model

import (
	"push/internal/model/entity"
)

type GetPushServiceListInput struct {
	Page int `json:"page" dc:"分页码"`
	Size int `json:"size" dc:"分页数量"`
}

type GetPushServiceCount struct {
	Id uint `json:"id" dc:"推送服务表主键id"`
}

type GetPushServiceListOutput struct {
	List  []*entity.PushService `json:"list" description:"列表"`
	Page  int                   `json:"page" description:"分页码"`
	Size  int                   `json:"size" description:"分页数量"`
	Total int                   `json:"total" description:"数据总数"`
}

type AddPushServiceInput struct {
	ServiceName string `json:"service_name"      ` // 推送服务名称
	ServiceSign string `json:"service_sign"      ` // 推送服务标识
}

type SetPushServiceInput struct {
	Id            uint `json:"id"                ` // 推送服务表主键id
	ServiceStatus bool `json:"service_status"    ` // 服务状态
}

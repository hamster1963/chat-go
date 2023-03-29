package api_push_service

import (
	"github.com/gogf/gf/v2/frame/g"
	"push/internal/consts"
	"push/internal/model"
)

// GetServiceListReq 获取服务列表 Req请求
type GetServiceListReq struct {
	g.Meta `method:"get" tags:"推送服务" summary:"获取服务列表" dc:"获取服务列表"`
	Page   int `json:"page" dc:"分页码"`
	Size   int `json:"size" dc:"分页数量"`
}

// GetServiceListRes 获取服务列表 Res返回
type GetServiceListRes struct {
	*model.GetPushServiceListOutput
}

// AddPushServiceReq 新增推送服务 Req请求
type AddPushServiceReq struct {
	g.Meta      `method:"post" tags:"推送服务" summary:"新增推送服务" dc:"新增推送服务"`
	ServiceName string `json:"service_name" dc:"推送服务名称" v:"required#请输入推送服务名称" `
	ServiceSign string `json:"service_sign" dc:"推送服务标识" v:"required#请输入推送服务标识"`
}

// AddPushServiceRes 新增推送服务 Res返回
type AddPushServiceRes struct {
	*consts.DefaultActionMessage
}

// SetPushServiceStatusReq 设置推送服务状态 Req请求
type SetPushServiceStatusReq struct {
	g.Meta        `method:"put" tags:"推送服务" summary:"设置推送服务" dc:"设置推送服务"`
	Id            uint `json:"id" dc:"推送服务ID" v:"required#请输入推送服务ID"`
	ServiceStatus bool `json:"service_status" dc:"服务状态" v:"required#请输入服务状态"`
}

// SetPushServiceStatusRes 设置推送服务 Res返回
type SetPushServiceStatusRes struct {
	*consts.DefaultActionMessage
}

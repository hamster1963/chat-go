package api_push_device

import (
	"github.com/gogf/gf/v2/frame/g"
	"push/internal/consts"
	"push/internal/model"
)

// AddPushDeviceReq 新增推送设备 Req请求
type AddPushDeviceReq struct {
	g.Meta     `method:"post" tags:"推送设备" summary:"新增推送设备" dc:"新增推送设备"`
	DeviceName string `json:"device_name"   v:"required#请输入设备名称"   `      // 设备名称
	BaseUrl    string `json:"base_url"     v:"required#请输入设备推送基础URL"    ` // 推送基础URL
}

// AddPushDeviceRes 新增推送设备 Res返回
type AddPushDeviceRes struct {
	*consts.DefaultActionMessage
}

// SetPushDeviceServiceReq 设置推送设备服务 Req请求
type SetPushDeviceServiceReq struct {
	g.Meta        `method:"put" tags:"推送设备" summary:"设置推送设备服务" dc:"设置推送设备服务"`
	Id            uint   `json:"id" dc:"推送设备ID" v:"required#请输入推送设备ID"`
	ServiceIdList string `json:"service_id_list" dc:"服务ID列表" v:"required#请输入服务ID列表"`
}

// SetPushDeviceServiceRes 设置推送设备服务 Res返回
type SetPushDeviceServiceRes struct {
	*model.SetPushDeviceServiceOutput
}

// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PushService is the golang structure of table push_service for DAO operations like Where/Data.
type PushService struct {
	g.Meta          `orm:"table:push_service, do:true"`
	Id              interface{} // 推送服务表主键id
	ServiceName     interface{} // 推送服务名称
	ServiceSign     interface{} // 推送服务标识
	ServiceStatus   interface{} // 服务状态
	ServicePushTime *gtime.Time // 服务最新推送时间
}
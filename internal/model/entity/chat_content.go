// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatContent is the golang structure for table chat_content.
type ChatContent struct {
	Id         uint        `json:"id"          ` // 聊天信息ID
	Content    string      `json:"content"     ` // 聊天内容
	Read       int         `json:"read"        ` // 是否已读
	SendId     int         `json:"send_id"     ` // 发送者ID
	ReceiveId  int         `json:"receive_id"  ` // 接收者ID
	CreateTime *gtime.Time `json:"create_time" ` // 创建时间
	ReadTime   *gtime.Time `json:"read_time"   ` // 已读时间
}

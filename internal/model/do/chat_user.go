// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatUser is the golang structure of table chat_user for DAO operations like Where/Data.
type ChatUser struct {
	g.Meta       `orm:"table:chat_user, do:true"`
	Id           interface{} // 用户ID
	Username     interface{} // 用户名
	LastChatTime *gtime.Time // 上次聊天时间
}

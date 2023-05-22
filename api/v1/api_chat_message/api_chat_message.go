package api_chat_message

import (
	"chat-go/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// AddMessageReq 新增聊天记录 Req请求
type AddMessageReq struct {
	g.Meta    `method:"post" tags:"聊天数据" summary:"新增聊天记录" dc:"新增聊天记录"`
	Content   *string `json:"content"     description:"聊天内容" v:"required #请输入 content"`     // 聊天内容
	ReceiveId *int    `json:"receive_id"  description:"接收者ID" v:"required #请输入 receive_id"` // 接收者ID
}

// GetMessageByIdReq 获取聊天信息(Id) Req请求
type GetMessageByIdReq struct {
	g.Meta `method:"get" tags:"聊天数据" summary:"获取聊天信息(Id)" dc:"获取聊天信息(Id)"`
	Id     *string `json:"id" description:"聊天记录ID" v:"required #请输入 id"` // 聊天记录ID
}

// GetMessageByIdRes 获取聊天信息(Id) Res返回
type GetMessageByIdRes struct {
	*entity.ChatContent
}

package c_chat_message

import (
	"chat-go/api/v1/api_chat_message"
	"chat-go/internal/global/g_consts"
	"chat-go/internal/global/g_functions"
	"chat-go/internal/model/m_chat_content"
	"chat-go/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type cChatMessage struct{}

var ChatMessage = &cChatMessage{}

// AddMessage
//
//	@dc: 添加消息
//	@author: laixin   @date:2023/5/22 19:33:26
func (c *cChatMessage) AddMessage(ctx context.Context, req *api_chat_message.AddMessageReq) (res g_consts.DefaultStringRes, err error) {
	// 获取用户ID
	userId := gconv.PtrInt(ctx.Value("userId"))
	if userId == nil || *userId == 0 {
		err := g_functions.ResErr(400, "cChatMessage AddMessage 获取用户ID错误！", err)
		return nil, err
	}
	// 调用logic层
	res, err = service.ChatMessage().AddMessage(ctx, &m_chat_content.AddChatContent{
		Content:   req.Content,
		SendId:    userId,
		ReceiveId: req.ReceiveId,
	})
	if err != nil {
		err := g_functions.ResErr(400, "cChatMessage AddMessage logic层错误！", err)
		return nil, err
	}
	return
}

// GetMessageById
//
//	@dc: 根据消息Id查询表信息
//	@author: laixin   @date:2023/5/23 00:33:04
func (c *cChatMessage) GetMessageById(ctx context.Context, req *api_chat_message.GetMessageByIdReq) (res *api_chat_message.GetMessageByIdRes, err error) {
	// 获取用户ID
	userId := gconv.PtrInt(ctx.Value("userId"))
	if userId == nil || *userId == 0 {
		err := g_functions.ResErr(400, "cChatMessage AddMessage 获取用户ID错误！", err)
		return nil, err
	}
	// 调用logic层
	message, err := service.ChatMessage().GetMessageById(ctx, req.Id)
	if err != nil {
		err := g_functions.ResErr(400, "cChatMessage AddMessage logic层错误！", err)
		return nil, err
	}
	if message.ReceiveId != *userId {
		err := g_functions.ResErr(400, "cChatMessage AddMessage 用户无此消息权限", err)
		return nil, err
	}
	res = &api_chat_message.GetMessageByIdRes{
		message,
	}
	if err != nil {
		err := g_functions.ResErr(400, "cChatMessage AddMessage 数据转换错误", err)
		return nil, err
	}
	return
}

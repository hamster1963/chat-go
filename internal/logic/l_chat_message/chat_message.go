package l_chat_message

import (
	"chat-go/internal/dao"
	"chat-go/internal/global/g_functions"
	"chat-go/internal/model/entity"
	"chat-go/internal/model/m_chat_content"
	"chat-go/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

// ==========================================================================
// logic 初始化
// ==========================================================================

type sChatMessage struct {
}

func init() {
	service.RegisterChatMessage(New())
}

func New() service.IChatMessage {
	return &sChatMessage{}
}

// GetMessageById
//
//	@dc:主键查询表信息
//	@params:表主键id-in
//	@response:表原结构信息entity
//	@author:auto @date:2023/5/21 16:57:37
func (s *sChatMessage) GetMessageById(ctx context.Context, in *string) (out *entity.ChatContent, err error) {
	db := dao.ChatContent.Ctx(ctx)
	err = db.Where("id", in).Scan(&out)
	if err != nil {
		err = g_functions.ResErr(500, "sChatMessage SelectById 数据库查询错误！", err)
		return
	}
	return
}

// AddMessage
//
//	@dc: 添加消息
//	@params: 聊天内容-in
//	@response: 聊天内容ID-out, 错误信息-err
//	@author:laixin @date:2023/5/21 17:00:12
func (s *sChatMessage) AddMessage(ctx context.Context, in *m_chat_content.AddChatContent) (out *string, err error) {
	db := dao.ChatContent.Ctx(ctx)
	data, err := db.OmitEmpty().InsertAndGetId(in)
	if err != nil {
		return nil, fmt.Errorf("sChatMessage AddMessage 数据库新增错误 %v", err)
	} else {
		out = gconv.PtrString(data)
	}
	return
}

package c_chat_login

import (
	v1 "chat-go/api/v1"
	"chat-go/internal/global/g_functions"
	"chat-go/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"
)

type cChatLogin struct{}

var ChatLogin = &cChatLogin{}

// Login
//
//	@dc: 用户登入
//	@author: laixin   @date:2023/5/19 08:16:00
func (c *cChatLogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	chatUser, err := service.ChatLogin().SelectByUsername(ctx, &req.Username)
	if err != nil {
		return nil, err
	}
	if chatUser == nil {
		return nil, g_functions.ResErr(400, "用户不存在", nil)
	}
	jwt, err := service.ChatLogin().GenerateUserJWT(ctx, chatUser)
	if err != nil {
		return nil, g_functions.ResErr(400, "生成JWT失败", err)
	}
	glog.Info(ctx, "用户登入成功，user_id为：", chatUser.Id)
	return &v1.LoginRes{
		JWT: jwt,
	}, nil
}

// TokenAuthTest
//
//	@dc: token鉴权测试接口
//	@author: laixin   @date:2023/5/19 20:09:55
func (c *cChatLogin) TokenAuthTest(ctx context.Context, _ *v1.EmptyReq) (res *v1.EmptyRes, err error) {
	glog.Info(ctx, "token鉴权测试接口完成")
	userID := ctx.Value("user_id")
	if userID == nil {
		glog.Warning(ctx, "token鉴权测试接口完成，但是user_id为nil")
		return nil, nil
	}
	glog.Info(ctx, "token鉴权测试接口完成，user_id为：", userID)
	return
}

package r_hamster_router

import (
	"chat-go/internal/controller/c_chat_login"
	"chat-go/internal/controller/c_chat_message"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		BindChatUser(group)
		BindChatMessage(group)
	})
}

// BindChatUser 注册用户路由
func BindChatUser(group *ghttp.RouterGroup) {
	group.Group("/user", func(group *ghttp.RouterGroup) {
		// 自定中间件设置
		// group.Middleware(middleware.JWTAuth)
		// Bind注册路由
		group.Bind(c_chat_login.ChatLogin)
	})
}

// BindChatMessage 注册消息路由
func BindChatMessage(group *ghttp.RouterGroup) {
	group.Group("/message", func(group *ghttp.RouterGroup) {
		// 自定中间件设置
		// group.Middleware(g_middleware.SMiddlewares.JWTAuth)
		// Bind注册路由
		group.Bind(c_chat_message.ChatMessage)
	})
}

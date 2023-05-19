package r_hamster_router

import (
	"chat-go/internal/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		BindChatUser(group)
	})
}

// BindChatUser 注册用户路由
func BindChatUser(group *ghttp.RouterGroup) {
	group.Group("/user", func(group *ghttp.RouterGroup) {
		// 自定中间件设置
		// group.Middleware(middleware.JWTAuth)
		// Bind注册路由
		group.Bind(controller.ChatLogin)
	})
}

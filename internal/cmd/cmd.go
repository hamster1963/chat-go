package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"push/internal/consts"
	"push/internal/controller"
	"push/internal/logic/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetNameToUriType(1) // 将路由中的下划线转换为驼峰命名
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.ALL("/", func(r *ghttp.Request) {
					r.Response.Write(consts.Ui)
				})

				group.Middleware(middleware.HandlerResponse)
				group.Bind(
					controller.PushService, // 推送服务
					controller.PushDevice,  // 推送设备
					controller.PushCore,    // 推送核心
				)
			})
			s.Run()
			return nil
		},
	}
)

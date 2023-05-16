package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"push/boot"
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
			s.SetIndexFolder(true)
			s.AddSearchPath("resource/public/html")
			s.SetIndexFiles([]string{"index.html"})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.MiddlewareCORS)
				group.Middleware(middleware.HandlerResponse)
				group.Bind(
					controller.DataCore,
				)
			})

			if err := boot.Boot(); err != nil {
				glog.Fatal(ctx, "初始化任务失败: ", err)
			}
			s.Run()
			return nil
		},
	}
)

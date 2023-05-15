package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"push/boot"
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
				group.Middleware(middleware.MiddlewareCORS)
				group.Middleware(middleware.HandlerResponse)
				group.Bind(
					controller.DataCore,
				)
			})

			if err := boot.Boot(); err != nil {
				glog.Fatal(context.Background(), "boot failed: ", err)
			}
			glog.Debug(context.Background(), "定时任务启动成功")

			s.Run()
			return nil
		},
	}
)

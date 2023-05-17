package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"home-network-watcher/boot"
	"home-network-watcher/internal/consts"
	"home-network-watcher/internal/controller"
	"home-network-watcher/internal/logic/middleware"
	binInfo "home-network-watcher/utility/bin_utils"
)

var (
	VersionString = "GitTag:" + binInfo.GitTag + "\n" +
		"GitCommitLog:" + binInfo.GitCommitLog + "\n" +
		"GitStatus:" + binInfo.GitStatus + "\n" +
		"BuildTime:" + binInfo.BuildTime + "\n" +
		"BuildGoVersion:" + binInfo.BuildGoVersion + "\n"
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetNameToUriType(1) // 将路由中的下划线转换为驼峰命名

			s.Group("/", func(group *ghttp.RouterGroup) {
				// 首页HTML
				group.ALL("/", func(r *ghttp.Request) {
					r.Response.Write(consts.IndexHTML)
				})
				group.ALL("/version", func(r *ghttp.Request) {
					r.Response.Write(VersionString)
				})
				// 中间件
				group.Middleware(middleware.MiddlewareCORS)
				group.Middleware(middleware.HandlerResponse)
				// 接口绑定
				group.Bind(
					controller.DataCore,
				)
			})
			// 初始化
			if err := boot.Boot(); err != nil {
				glog.Fatal(ctx, "初始化任务失败: ", err)
			}

			s.Run()
			return nil
		},
	}
)

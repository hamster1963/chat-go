package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"push/internal/consts"
	"push/internal/controller"
	"push/internal/logic/middleware"
	"push/utility/network_utils"
	"push/utility/push_utils"
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

			g.Dump("开始获取科学上网网速")
			_, err = gcron.AddSingleton(ctx, "@every 1s", func(ctx context.Context) {
				err = network_utils.ProxyNetwork.GetProxyNetwork()
				if err != nil {
					g.Dump(err)
				}
			}, "获取代理速度")
			if err != nil {
				panic(err)
			}

			g.Dump("开始获取家庭路由器网速")
			_, err = gcron.AddSingleton(ctx, "@every 1s", func(ctx context.Context) {
				err = network_utils.NetworkUtils.GetHomeNetwork()
				if err != nil {
					g.Dump(err)
				}
			}, "获取家庭路由器速度")
			if err != nil {
				panic(err)
			}

			g.Dump("开始获取当前代理节点信息")
			_, err = gcron.AddSingleton(ctx, "@every 5s", func(ctx context.Context) {
				err = network_utils.NodeUtils.GetNodeInfo()
				if err != nil {
					g.Dump(err)
				}
			}, "获取当前代理节点信息")
			if err != nil {
				panic(err)
			}

			g.Dump("开始推送科学上网网速")
			_, err = gcron.AddSingleton(ctx, "@every 1s", func(ctx context.Context) {
				err = push_utils.PushUtils.ProxyPushCore(ctx)
				if err != nil {
					g.Dump(err)
				}
			}, "推送科学上网速度")
			if err != nil {
				panic(err)
			}

			g.Dump("开始存储出站流量")
			_, err = gcron.AddSingleton(ctx, "@midnight", func(ctx context.Context) {
				err = push_utils.PushUtils.StoreOutbound()
				if err != nil {
					g.Dump(err)
				}
			}, "存储出站流量")
			if err != nil {
				panic(err)
			}
			err = push_utils.PushUtils.StoreOutbound()
			if err != nil {
				panic(err)
			}
			g.Dump("开始推送出站流量")
			_, err = gcron.AddSingleton(ctx, "@every 6h", func(ctx context.Context) {
				err = push_utils.PushUtils.GetUsedOutboundAndPush()
				if err != nil {
					g.Dump(err)
				}
			}, "推送出站流量")
			if err != nil {
				panic(err)
			}
			s.Run()
			return nil
		},
	}
)

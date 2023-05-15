package boot

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
	"push/utility/network_utils"
	"push/utility/push_utils"
)

func Boot() (err error) {
	ctx := context.TODO()
	glog.Notice(ctx, "开始获取科学上网网速")
	_, err = gcron.AddSingleton(ctx, "@every 1s", func(ctx context.Context) {
		err = network_utils.ProxyNetwork.GetProxyNetwork()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "获取代理速度")
	if err != nil {
		panic(err)
	}

	glog.Notice(ctx, "开始获取家庭路由器网速")
	_, err = gcron.AddSingleton(ctx, "@every 1s", func(ctx context.Context) {
		err = network_utils.NetworkUtils.GetHomeNetwork()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "获取家庭路由器速度")
	if err != nil {
		panic(err)
	}

	glog.Notice(ctx, "开始获取当前代理节点信息")
	_, err = gcron.AddSingleton(ctx, "@every 5s", func(ctx context.Context) {
		err = network_utils.NodeUtils.GetNodeInfo()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "获取当前代理节点信息")
	if err != nil {
		panic(err)
	}

	glog.Notice(ctx, "开始推送科学上网网速")
	_, err = gcron.AddSingleton(ctx, "@every 1s", func(ctx context.Context) {
		err = push_utils.PushUtils.ProxyPushCore(ctx)
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "推送科学上网速度")
	if err != nil {
		panic(err)
	}

	glog.Notice(ctx, "开始存储出站流量")
	_, err = gcron.AddSingleton(ctx, "@midnight", func(ctx context.Context) {
		err = push_utils.PushUtils.StoreOutbound()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "存储出站流量")
	if err != nil {
		panic(err)
	}
	err = push_utils.PushUtils.StoreOutbound()
	if err != nil {
		panic(err)
	}

	glog.Notice(ctx, "开始推送出站流量")
	_, err = gcron.AddSingleton(ctx, "@every 6h", func(ctx context.Context) {
		err = push_utils.PushUtils.GetUsedOutboundAndPush()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "推送出站流量")
	if err != nil {
		panic(err)
	}
	return nil
}

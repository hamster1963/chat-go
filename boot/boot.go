package boot

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
	"home-network-watcher/utility/network_utils"
	"home-network-watcher/utility/push_utils"
)

func Boot() (err error) {
	_, err = gcron.AddOnce(context.TODO(), "@every 2s", func(ctx context.Context) {
		glog.Debug(context.Background(), "定时任务启动中...")
		if err := BootMethod(); err != nil {
			glog.Fatal(context.Background(), "定时任务启动失败: ", err)
		}
		glog.Debug(context.Background(), "定时任务启动成功")
	}, "开始启动定时任务")
	if err != nil {
		return err
	}
	return nil
}

func BootMethod() (err error) {
	var ctx = context.TODO()

	glog.Notice(ctx, "开始获取科学上网网速")
	_, err = gcron.AddSingleton(ctx, "@every 1s", func(ctx context.Context) {
		err = network_utils.ProxyNetwork.GetProxyNetwork()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "获取代理速度")
	if err != nil {
		return err
	}

	glog.Notice(ctx, "开始获取家庭路由器网速")
	_, err = gcron.AddSingleton(ctx, "@every 1s", func(ctx context.Context) {
		err = network_utils.NetworkUtils.GetHomeNetwork()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "获取家庭路由器速度")
	if err != nil {
		return err
	}

	glog.Notice(ctx, "开始获取当前代理节点信息")
	_, err = gcron.AddSingleton(ctx, "@every 5s", func(ctx context.Context) {
		err = network_utils.NodeUtils.GetNodeInfo()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "获取当前代理节点信息")
	if err != nil {
		return err
	}

	glog.Notice(ctx, "开始推送科学上网网速")
	_, err = gcron.AddSingleton(ctx, "@every 1s", func(ctx context.Context) {
		err = push_utils.PushUtils.ProxyPushCore(ctx)
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "推送科学上网速度")
	if err != nil {
		return err
	}

	glog.Notice(ctx, "开始存储出站流量")
	_, err = gcron.AddSingleton(ctx, "@midnight", func(ctx context.Context) {
		err = push_utils.PushUtils.StoreOutbound()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "存储出站流量")
	if err != nil {
		return err
	}
	// 进行第一次流量数据缓存
	err = push_utils.PushUtils.StoreOutbound()
	if err != nil {
		return err
	}

	glog.Notice(ctx, "开始推送出站流量")
	_, err = gcron.AddSingleton(ctx, "@every 6h", func(ctx context.Context) {
		err = push_utils.PushUtils.GetUsedOutboundAndPush()
		if err != nil {
			glog.Warning(ctx, err)
		}
	}, "推送出站流量")
	if err != nil {
		return err
	}
	return nil
}

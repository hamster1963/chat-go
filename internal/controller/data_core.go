package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	v1 "home-network-watcher/api/v1"
)

type cDataCore struct{}

var DataCore = &cDataCore{}

// GetNetWorkData
//
//	@dc: 获取网络信息
//	@author: laixin   @date:2023/4/2 20:37:34
func (c *cDataCore) GetNetWorkData(_ context.Context, _ *v1.GetNetworkDataReq) (res *v1.GetNetworkDataRes, err error) {
	// 从缓存中获取数据
	nodeInfo, err := gcache.Get(context.Background(), "nodeInfo")
	if err != nil {
		return nil, err
	}
	homeNetwork, err := gcache.Get(context.Background(), "homeNetwork")
	if err != nil {
		return nil, err
	}
	proxyNetwork, err := gcache.Get(context.Background(), "proxyNetwork")
	if err != nil {
		return nil, err
	}
	res = &v1.GetNetworkDataRes{
		NodeInfo:     nodeInfo,
		HomeNetwork:  homeNetwork,
		ProxyNetwork: proxyNetwork,
	}
	return
}

package network_utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type uNetworkUtils struct{}

var NetworkUtils = &uNetworkUtils{}

// GetHomeNetwork
//
//	@dc: 获取家庭路由器网速
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/2 19:43:13
func (u uNetworkUtils) GetHomeNetwork() (err error) {
	homeNetwork := g.Map{
		"time":        "",
		"rxSpeedKbps": 0,
		"txSpeedKbps": 0,
		"rxSpeedMbps": 0,
		"txSpeedMbps": 0,
	}

	url := "http://120.24.211.49:35600/json/stats.json"
	response, err := g.Client().Get(context.Background(), url)
	defer func(response *gclient.Response) {
		err := response.Close()
		if err != nil {
			g.Dump(err)
		}
	}(response)
	if err != nil {
		g.Dump(err)
		return err
	}
	jsonData := gjson.New(response.ReadAllString())
	rxSpeed := jsonData.Get("servers.0.network_rx") // 下载速度
	txSpeed := jsonData.Get("servers.0.network_tx") // 上传速度

	// 速度单位转换
	rxSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024))
	txSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024))
	homeNetwork["rxSpeedKbps"] = rxSpeedKbps
	homeNetwork["txSpeedKbps"] = txSpeedKbps

	// 转换成MB
	rxSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024/1024))
	txSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024/1024))
	homeNetwork["rxSpeedMbps"] = rxSpeedMbps
	homeNetwork["txSpeedMbps"] = txSpeedMbps

	homeNetwork["time"] = gtime.Now().String()
	err = gcache.Set(context.Background(), "homeNetwork", homeNetwork, 0)
	if err != nil {
		return err
	}
	return nil
}

package push_utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"push/internal/consts"
)

// GetTotalOutbound
//
//	@dc: 获取总出口流量
//	@params:
//	@response:
//	@author: laixin   @date:2023/5/11 00:20:58
func (u *uPushUtils) GetTotalOutbound() (outBoundStr string, err error) {
	var (
		totalOutbound float64
	)
	proxyUserList, err := PushUtils.GetProxyUser()
	if err != nil {
		return
	}
	// g.Dump(proxyUserList)
	// 计算全部用户的出口流量
	for _, user := range proxyUserList {
		userMap := gconv.Map(user)
		totalOutbound += gconv.Float64(userMap["up"]) + gconv.Float64(userMap["down"])
	}
	// byte -> G
	outBoundG := totalOutbound / consts.ByteToG
	outBoundStr = fmt.Sprintf("%.2f\n", outBoundG)
	return
}

// StoreOutbound
//
//	@dc: 存储出口流量
//	@params:
//	@response:
//	@author: laixin   @date:2023/5/11 02:20:10
func (u *uPushUtils) StoreOutbound() (err error) {
	outBoundStr, err := PushUtils.GetTotalOutbound()
	if err != nil {
		return
	}
	nowTime := gtime.Now().String()
	outBoundInfo := g.Map{
		"time":     nowTime,
		"outBound": outBoundStr,
	}
	_ = gcache.Set(context.Background(), "outBoundInfo", outBoundInfo, 0)
	return
}

// GetUsedOutboundAndPush
//
//	@dc: 获取已使用的出口流量
//	@params:
//	@response:
//	@author: laixin   @date:2023/5/11 02:22:34
func (u *uPushUtils) GetUsedOutboundAndPush() (err error) {
	if outBoundData, err := gcache.Get(context.Background(), "outBoundInfo"); !outBoundData.IsMap() || err != nil {
		if err != nil {
			g.Dump(err)
			return err
		} else {
			err = fmt.Errorf("outBoundData is not map")
			return err
		}
	} else {
		proxyUserList, err := PushUtils.GetProxyUser()
		if err != nil {
			return err
		}
		// g.Dump(proxyUserList)
		var totalOutbound float64
		// 计算全部用户的出口流量
		for _, user := range proxyUserList {
			userMap := gconv.Map(user)
			totalOutbound += gconv.Float64(userMap["up"]) + gconv.Float64(userMap["down"])
		}
		// byte -> G
		outBoundG := totalOutbound / consts.ByteToG
		outBoundStr := fmt.Sprintf("%.2f\n", outBoundG)
		outBoundMap := outBoundData.MapStrStr()
		duringTime := gtime.NewFromStr(gtime.Now().String()).Sub(gtime.NewFromStr(outBoundMap["time"]))
		usedOutBound := gconv.String(gconv.Float64(outBoundStr) - gconv.Float64(outBoundMap["outBound"]))
		g.Dump("过去" + duringTime.String() + "的出口流量为：" + usedOutBound + "GB")
		// 推送到Bark
		err = PushUtils.PushOutboundToBark(usedOutBound, duringTime.String())
		if err != nil {
			return err
		}
	}
	return
}

// PushOutboundToBark
//
//	@dc: 推送出口流量到Bark
//	@params:
//	@response:
//	@author: laixin   @date:2023/5/11 02:37:19
func (u *uPushUtils) PushOutboundToBark(usedOutBound, duringTime string) (err error) {
	url := "http://120.24.211.49:10399/PushCore"
	pushClient := g.Client()
	pushClient.SetHeader("Push-Sign", "ProxyOutbound_Push")
	// 设置为json
	pushClient.SetHeader("Content-Type", "application/json")
	response, err := pushClient.ContentJson().Post(context.Background(), url, g.Map{
		"duringTime":   duringTime,
		"usedOutBound": usedOutBound,
	})
	defer func(response *gclient.Response) {
		err := response.Close()
		if err != nil {
			g.Dump(err)
		}
	}(response)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		g.Dump(response.ReadAllString())
		err := gerror.New("推送失败")
		return err
	}
	return
}

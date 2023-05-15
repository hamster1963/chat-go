package network_utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type uProxyNetwork struct{}

var ProxyNetwork = &uProxyNetwork{}

// GetProxyNetwork
//
//	@dc:
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/2 20:06:21
func (u uProxyNetwork) GetProxyNetwork() (err error) {
	proxyNetwork := g.Map{
		"time":        "",
		"rxSpeedKbps": 0,
		"txSpeedKbps": 0,
		"rxSpeedMbps": 0,
		"txSpeedMbps": 0,
	}
	session, err := u.GetSession()
	// 通过xui进行网速的获取
	url := "http://xui.xinyu.today:580/server/status"
	post, err := g.Client().SetCookieMap(session).Post(context.Background(), url)
	defer func(post *gclient.Response) {
		err := post.Close()
		if err != nil {
			glog.Warning(context.Background(), err)
		}
	}(post)
	if err != nil {
		return err
	}
	if post.StatusCode != 200 {
		glog.Warning(context.Background(), "获取网速失败")
		return err
	}
	jsonData := gjson.New(post.ReadAllString())
	rxSpeed := jsonData.Get("obj.netIO.down") // 下载速度
	txSpeed := jsonData.Get("obj.netIO.up")   // 上传速度

	// 速度单位转换
	rxSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024))
	txSpeedKbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024))
	proxyNetwork["rxSpeedKbps"] = rxSpeedKbps
	proxyNetwork["txSpeedKbps"] = txSpeedKbps

	// 转换成MB
	rxSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(rxSpeed)/1024/1024))
	txSpeedMbps := gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(txSpeed)/1024/1024))
	proxyNetwork["rxSpeedMbps"] = rxSpeedMbps
	proxyNetwork["txSpeedMbps"] = txSpeedMbps

	proxyNetwork["time"] = gtime.Now().String()
	err = gcache.Set(context.Background(), "proxyNetwork", proxyNetwork, 0)
	if err != nil {
		return err
	}

	return err
}

func (u uProxyNetwork) GetSession() (sessionMap map[string]string, err error) {
	url := "http://xui.xinyu.today:580/login"
	post, err := g.Client().Post(context.Background(), url, g.Map{
		"username": "hamster",
		"password": "Deny1963!",
	})
	defer func(post *gclient.Response) {
		err := post.Close()
		if err != nil {
			glog.Warning(context.Background(), err)
		}
	}(post)
	if err != nil {
		return nil, err
	}
	if post.StatusCode != 200 {
		return nil, fmt.Errorf("登录失败")
	}
	if post.Header.Get("Set-Cookie") == "" {
		return nil, fmt.Errorf("获取Cookie失败")
	}
	// 将session存入缓存
	err = gcache.Set(context.Background(), "proxySession", post.GetCookieMap(), 15*time.Minute)
	if err != nil {
		return nil, err
	}
	return post.GetCookieMap(), nil
}

package push_utils

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"push/utility/network_utils"
	"time"
)

type uPushUtils struct{}

var PushUtils = &uPushUtils{}

// GetProxyNetwork
//
//	@dc: 获取科学上网速度
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/19 18:47:05
func (u *uPushUtils) GetProxyNetwork() (proxyNetworkUp string, err error) {
	proxyData, err := gcache.Get(context.Background(), "proxyNetwork")
	if err != nil {
		g.Dump(err)
		return "", err
	}
	proxyNetworkUp = gjson.New(proxyData.Map()).Get("txSpeedMbps").String()
	return
}

// PushToBark
//
//	@dc:
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/19 18:47:47
func (u *uPushUtils) PushToBark(proxyNetworkUp, maxFlowUser string, maxFlow int) (err error) {
	url := "http://120.24.211.49:30001/PushCore"
	pushClient := g.Client()
	pushClient.SetHeader("Push-Sign", "ProxyNetwork_Push")
	// 设置为json
	pushClient.SetHeader("Content-Type", "application/json")
	response, err := pushClient.ContentJson().Post(context.Background(), url, g.Map{
		"proxyNetworkUp": proxyNetworkUp,
		"maxFlow":        maxFlow,
		"maxFlowUser":    maxFlowUser,
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

// GetProxyUser
//
//	@dc: 获取代理占用用户
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/23 17:57:52
func (u *uPushUtils) GetProxyUser() (userList []interface{}, err error) {
	var (
		session map[string]string
	)

	// 尝试获取缓存中的session
	sessionData, err := gcache.Get(context.Background(), "proxySession")
	if err != nil {
		return nil, err
	}

	if sessionData == nil {
		// 获取session
		session, err = network_utils.ProxyNetwork.GetSession()
		if err != nil {
			return nil, err
		}
	} else {
		session = sessionData.MapStrStr()
	}

	// 获取代理占用用户
	url := "http://xui.xinyu.today:580/xui/inbound/list"
	post, err := g.Client().SetCookieMap(session).Post(context.Background(), url)
	defer func(post *gclient.Response) {
		err := post.Close()
		if err != nil {
			g.Dump(err)
		}
	}(post)
	if err != nil {
		return nil, err
	}
	if post.StatusCode != 200 {
		g.Dump(post.ReadAllString())
		return
	}
	jsonData := gjson.New(post.ReadAllString())
	userList = jsonData.Get("obj").Array()
	if len(userList) == 0 {
		err := gerror.New("代理用户为空")
		return nil, err
	}
	return userList, nil
}

// ProxyPushCore
//
//	@dc:
//	@params:
//	@response:
//	@author: laixin   @date:2023/4/19 18:48:13
func (u *uPushUtils) ProxyPushCore(ctx context.Context) (err error) {
	// 获取科学上网速度
	proxyNetworkUp, err := u.GetProxyNetwork()
	if err != nil {
		return
	}
	proxyNetworkUpSpeed := gconv.Float64(proxyNetworkUp)
	// 进行10s内超过速率限制次数判断
	// 速率限制
	speedLimit := "6"
	limitTime := 10
	if proxyNetworkUpSpeed > gconv.Float64(speedLimit) {
		g.Dump("速率超过限制" + proxyNetworkUp)
		// 速率超过限制
		// 获取缓存中的速率超过限制次数
		count, err := gcache.Get(ctx, "proxyNetworkUpSpeedCount")
		if err != nil {
			return err
		}
		if count == nil {
			err := gcache.Set(ctx, "proxyNetworkUpSpeedCount", 1, 20*time.Second)
			if err != nil {
				return err
			}
			// 获取占用用户
			userList, err := u.GetProxyUser()
			if err != nil {
				g.Dump(err)
				return err
			}
			// 获取用户当前流量存入缓存
			err = gcache.Set(ctx, "proxyUserFlow", userList, 0)
			if err != nil {
				return err
			}
			return nil
		}
		countInt := count.Int()
		g.Dump("当前超出次数" + gconv.String(countInt))
		if countInt > limitTime {
			// 清空缓存
			_, _ = gcache.Remove(ctx, "proxyNetworkUpSpeedCount")
			_, _ = gcache.Remove(ctx, "proxyUserFlow")
		}
		if countInt == limitTime {
			var maxFLow int
			var maxFLowUser string
			// 计算用户流量变化
			// 获取缓存中的用户流量
			userList, _ := gcache.Get(ctx, "proxyUserFlow")
			if userList == nil {
				g.Dump("用户流量为空")
			}
			if userList != nil {
				// 获取当前用户流量
				userListNow, err := u.GetProxyUser()
				if err != nil {
					g.Dump(err)
				}
				// 计算用户流量变化
				for _, user := range userList.Array() {
					userMap := gconv.Map(user)
					for _, userNow := range userListNow {
						userNowMap := gconv.Map(userNow)
						if userMap["id"] == userNowMap["id"] {
							// 计算用户流量变化
							totalFlow := gconv.Int(userNowMap["down"])/1024/1024 - gconv.Int(userMap["down"])/1024/1024
							if totalFlow >= maxFLow {
								maxFLow = totalFlow
								maxFLowUser = gconv.String(userNowMap["remark"])
							}
						}
					}
				}
			}
			// 进行推送
			err = u.PushToBark(proxyNetworkUp, maxFLowUser, maxFLow)
			if err != nil {
				g.Dump(err)
			}
			g.Dump("------推送成功------")
			// 清空缓存
			_, err = gcache.Remove(ctx, "proxyNetworkUpSpeedCount")
			_, err = gcache.Remove(ctx, "proxyUserFlow")
			if err != nil {
				g.Dump(err)
			}
		} else {
			// 速率超过限制次数+1
			err = gcache.Set(ctx, "proxyNetworkUpSpeedCount", countInt+1, gcache.MustGetExpire(ctx, "proxyNetworkUpSpeedCount"))
			if err != nil {
				return err
			}
		}

	} else {
		// 速率未超过限制
	}

	return
}

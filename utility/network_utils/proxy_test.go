package network_utils

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

// 测试方法注释
func Test18_01_04(t *testing.T) {
	url := "http://xui.xinyu.today:580/login"
	post, err := g.Client().Post(context.Background(), url, g.Map{
		"username": "hamster",
		"password": "Deny1963!",
	})
	if err != nil {
		g.Dump(err)
	}
	if post.StatusCode != 200 {
		g.Dump("登录失败")
	}
	if post.Header.Get("Set-Cookie") == "" {
		g.Dump("获取Cookie失败")
	}
	g.Dump(post.Header.Get("Set-Cookie"))
	g.Dump(post.GetCookieMap())
}

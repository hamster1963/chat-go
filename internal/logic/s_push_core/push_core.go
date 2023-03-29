package s_push_core

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"push/internal/model"
	"push/internal/service"
)

type sPushCore struct{}

func init() {
	service.RegisterPushCore(New())
}

func New() *sPushCore {
	return &sPushCore{}
}

// PushCore
//
//	@dc: 推送核心
//	@params:
//	@response:
//	@author:laixin @date:2023/3/29 23:20:16
func (s *sPushCore) PushCore(ctx context.Context, in *model.PushBasicData) (err error) {
	g.Dump("start")
	serviceInfo, err := service.PushService().SearchPushServiceBySign(ctx, in)
	g.Dump(serviceInfo)
	if err != nil {
		err := service.PushLog().AddPushLog(ctx, &model.AddPushLogInput{
			PushInfo:    in.PushData,
			PushService: "error",
			PushTime:    gtime.Now(),
			PushStatus:  false,
			ErrInfo:     err.Error(),
		})
		if err != nil {
			return err
		}
	}
	g.Dump(serviceInfo.ServiceName)
	nowTime := gtime.Now()
	err = service.PushService().UpdatePushServicePushTime(ctx, serviceInfo.Id)
	if err != nil {
		err = service.PushLog().AddPushLog(ctx, &model.AddPushLogInput{
			PushInfo:    "in.PushData",
			PushService: serviceInfo.ServiceName,
			PushTime:    nowTime,
			PushStatus:  false,
			ErrInfo:     err.Error(),
		})
		if err != nil {
			return err
		}
		return err
	}
	baseList, err := service.PushDevice().GetDeviceBaseListByServiceId(ctx, serviceInfo.Id)
	if err != nil {
		err = service.PushLog().AddPushLog(ctx, &model.AddPushLogInput{
			PushInfo:    "in.PushData",
			PushService: serviceInfo.ServiceName,
			PushTime:    nowTime,
			PushStatus:  false,
			ErrInfo:     err.Error(),
		})
		if err != nil {
			return err
		}
		return err
	}
	g.Dump(baseList)
	// 进行信息组装推送
	// 获取推送信息
	decode, err := gjson.Decode(in.PushData)
	if err != nil {
		err = service.PushLog().AddPushLog(ctx, &model.AddPushLogInput{
			PushInfo:    "in.PushData",
			PushService: serviceInfo.ServiceName,
			PushTime:    nowTime,
			PushStatus:  false,
			ErrInfo:     err.Error(),
		})
		if err != nil {
			return err
		}
		return err
	}
	decodeMap := gconv.Map(decode)
	switch serviceInfo.ServiceSign {
	case "UptimeKuma_Push":
		msg := decodeMap["msg"]
		// 组装推送信息
		postData := g.Map{
			"icon":  "http://120.24.211.49:3999/upload/logo1.png",
			"body":  msg,
			"title": "Uptime",
			"sound": "calypso",
			"group": "UptimeKuma",
		}
		// 获取推送设备信息
		for _, v := range baseList {
			err := globalPost(v.BaseUrl, postData)
			if err != nil {
				err = service.PushLog().AddPushLog(ctx, &model.AddPushLogInput{
					PushInfo:    "in.PushData",
					PushService: serviceInfo.ServiceName,
					PushTime:    nowTime,
					PushStatus:  false,
					ErrInfo:     err.Error(),
				})
				if err != nil {
					return err
				}
				return err
			}
		}
	case "DDNS_Push":
		g.Dump("DDNS")
		msg := decodeMap["title"]
		// 组装推送信息
		postData := g.Map{
			"icon":  "https://120.24.211.49/favicon.ico",
			"body":  msg,
			"title": "DDNS",
			"sound": "shake",
			"group": "DDNS-GO",
		}
		// 获取推送设备信息
		for _, v := range baseList {
			err := globalPost(v.BaseUrl, postData)
			if err != nil {
				err = service.PushLog().AddPushLog(ctx, &model.AddPushLogInput{
					PushInfo:    "in.PushData",
					PushService: serviceInfo.ServiceName,
					PushTime:    nowTime,
					PushStatus:  false,
					ErrInfo:     err.Error(),
				})
				if err != nil {
					return err
				}
				return err
			}
		}
	default:
		err = gerror.New("未知推送服务")
		return err
	}

	err = service.PushLog().AddPushLog(ctx, &model.AddPushLogInput{
		PushInfo:    "in.PushData",
		PushService: serviceInfo.ServiceName,
		PushTime:    nowTime,
		PushStatus:  true,
		ErrInfo:     "",
	})
	if err != nil {
		return err
	}

	return
}

func globalPost(url string, data g.Map) (err error) {
	postRes, err := g.Client().Post(context.Background(), url, data)
	if err != nil {
		return err
	}
	if postRes.StatusCode != 200 {
		return gerror.New("推送失败")
	}
	return
}

package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"push/api/v1/api_push_core"
	"push/internal/model"
	"push/internal/service"
)

type cPushCore struct{}

var PushCore = &cPushCore{}

// PushCore
//
//	@dc: 推送核心
//	@author: laixin   @date:2023/3/30 00:11:02
func (c *cPushCore) PushCore(ctx context.Context, req *api_push_core.PushCoreReq) (res *api_push_core.PushCoreRes, err error) {
	// 获取推送header或者body
	request := g.RequestFromCtx(ctx)
	sign := request.Header.Get("Push-Sign")
	if sign == "" {
		sign = gconv.String(gconv.Map(request.GetBodyString())["Push-Sign"])
		if sign == "" {
			return nil, gerror.New("推送签名不能为空")
		}
	}
	g.Dump(sign)
	if request.GetBodyString() == "" {
		return nil, gerror.New("推送内容不能为空")
	}
	err = service.PushCore().PushCore(ctx, &model.PushBasicData{
		ServiceSign: sign,
		PushData:    request.GetBodyString(),
	})
	if err != nil {
		return nil, err
	}
	return
}

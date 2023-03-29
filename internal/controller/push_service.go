package controller

import (
	"context"
	"push/api/v1/api_push_service"
	"push/internal/consts"
	"push/internal/model"
	"push/internal/service"
)

type cPushService struct{}

var PushService = &cPushService{}

// GetPushServiceList
//
//	@dc: 获取服务列表
//	@author: laixin   @date:2023/3/29 18:18:54
func (c *cPushService) GetPushServiceList(ctx context.Context, req *api_push_service.GetServiceListReq) (res *api_push_service.GetServiceListRes, err error) {
	res = &api_push_service.GetServiceListRes{}
	res.GetPushServiceListOutput, err = service.PushService().GetPushServiceList(ctx, model.GetPushServiceListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return
}

// AddPushService
//
//	@dc: 新增推送服务
//	@author: laixin   @date:2023/3/29 19:35:38
func (c *cPushService) AddPushService(ctx context.Context, req *api_push_service.AddPushServiceReq) (res *api_push_service.AddPushServiceRes, err error) {
	// 检测服务名称或者服务标识是否存在
	err = service.PushService().AddPushService(ctx, model.AddPushServiceInput{
		ServiceName: req.ServiceName,
		ServiceSign: req.ServiceSign,
	})
	if err != nil {
		return nil, err
	}
	msg := &consts.DefaultActionMessage{Message: consts.DefaultSuccessMessage}
	res = &api_push_service.AddPushServiceRes{
		DefaultActionMessage: msg,
	}
	return
}

// SetPushServiceStatus
//
//	@dc: 设置推送服务状态
//	@author: laixin   @date:2023/3/29 20:53:18
func (c *cPushService) SetPushServiceStatus(ctx context.Context, req *api_push_service.SetPushServiceStatusReq) (res *api_push_service.SetPushServiceStatusRes, err error) {
	err = service.PushService().SetPushServiceStatus(ctx, model.SetPushServiceInput{
		Id:            req.Id,
		ServiceStatus: req.ServiceStatus,
	})
	if err != nil {
		return nil, err
	}
	msg := &consts.DefaultActionMessage{Message: consts.DefaultSuccessMessage}
	res = &api_push_service.SetPushServiceStatusRes{DefaultActionMessage: msg}
	return
}

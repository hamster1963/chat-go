package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"push/api/v1/api_push_device"
	"push/internal/consts"
	"push/internal/model"
	"push/internal/service"
)

type cPushDevice struct{}

var PushDevice = &cPushDevice{}

// AddPushDevice
//
//	@dc: 添加推送设备
//	@author: laixin   @date:2023/3/29 22:12:54
func (c *cPushDevice) AddPushDevice(ctx context.Context, req *api_push_device.AddPushDeviceReq) (res *api_push_device.AddPushDeviceRes, err error) {
	res = new(api_push_device.AddPushDeviceRes)
	err = service.PushDevice().AddPushDevice(ctx, &model.AddPushDeviceInput{
		DeviceName: req.DeviceName,
		BaseUrl:    req.BaseUrl,
	})
	if err != nil {
		return nil, err
	}
	msg := &consts.DefaultActionMessage{Message: consts.DefaultSuccessMessage}
	res = &api_push_device.AddPushDeviceRes{
		DefaultActionMessage: msg,
	}
	return
}

// SetPushDeviceService
//
//	@dc:
//	@author: laixin   @date:2023/3/29 22:58:06
func (c *cPushDevice) SetPushDeviceService(ctx context.Context, req *api_push_device.SetPushDeviceServiceReq) (res *api_push_device.SetPushDeviceServiceRes, err error) {
	out, err := service.PushDevice().SetPushDeviceService(ctx, &model.SetPushDeviceServiceInput{
		Id:            req.Id,
		ServiceIdList: req.ServiceIdList,
	})
	if err != nil {
		return nil, err
	}
	g.Dump(out)
	res = &api_push_device.SetPushDeviceServiceRes{
		SetPushDeviceServiceOutput: out,
	}
	return
}

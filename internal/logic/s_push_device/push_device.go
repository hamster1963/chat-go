package s_push_device

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"push/internal/dao"
	"push/internal/model"
	"push/internal/model/entity"
	"push/internal/service"
)

type sPushDevice struct {
}

func init() {
	service.RegisterPushDevice(New())
}

func New() *sPushDevice {
	return &sPushDevice{}
}

// AddPushDevice
//
//	@dc: 添加推送设备
//	@params:
//	@response:
//	@author:laixin @date:2023/3/29 22:06:18
func (s *sPushDevice) AddPushDevice(ctx context.Context, in *model.AddPushDeviceInput) (err error) {
	var m = dao.PushDevice.Ctx(ctx)
	// 判断设备名称是否存在
	count, err := m.Where("device_name = ?", in.DeviceName).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("设备名称已存在")
	}
	// 判断设备推送基础URL是否存在
	count, err = m.Where("base_url = ?", in.BaseUrl).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("设备推送基础URL已存在")
	}
	_, err = m.Insert(&entity.PushDevice{
		DeviceName:   in.DeviceName,
		BaseUrl:      in.BaseUrl,
		DeviceStatus: true,
		CreateTime:   gtime.Now(),
	})
	return
}

// SetPushDeviceService
//
//	@dc: 设置设备推送服务
//	@params:
//	@response:
//	@author:laixin @date:2023/3/29 22:40:08
func (s *sPushDevice) SetPushDeviceService(ctx context.Context, in *model.SetPushDeviceServiceInput) (out *model.SetPushDeviceServiceOutput, err error) {
	var m = dao.PushDevice.Ctx(ctx)
	out = &model.SetPushDeviceServiceOutput{}
	// 判断服务是否存在
	g.Dump("ok")
	serviceList := gconv.SliceInt(in.ServiceIdList)
	g.Dump(serviceList)
	for _, v := range serviceList {
		count, err := service.PushService().GetPushServiceCount(ctx, &model.GetPushDeviceInfoInput{
			Id: gconv.Uint(v),
		})
		if err != nil {
			return nil, err
		}
		g.Dump(count)
		if count == 0 {
			out.FailList = append(out.FailList, v)
		}
		if count > 0 {
			out.SuccessList = append(out.SuccessList, v)
		}
	}
	g.Dump(out)
	// 更新设备服务
	in.ServiceIdList = gconv.String(out.SuccessList)
	_, err = m.Where("id = ?", in.Id).Data("service_id_list", in.ServiceIdList).Update()
	if err != nil {
		return nil, err
	}
	return
}

// GetDeviceInfo
//
//	@dc: 获取设备信息
//	@params:
//	@response:
//	@author:laixin @date:2023/3/29 21:21:22
func (s *sPushDevice) GetDeviceInfo(ctx context.Context, in *model.GetPushDeviceInfoInput) (out *entity.PushDevice, err error) {
	var m = dao.PushDevice.Ctx(ctx)
	out = new(entity.PushDevice)
	err = m.Where("id = ?", in.Id).Scan(out)
	return
}

// GetDeviceCount
//
//	@dc: 获取设备数量
//	@params:
//	@response:
//	@author:laixin @date:2023/3/29 21:51:46
func (s *sPushDevice) GetDeviceCount(ctx context.Context, in *model.GetPushDeviceInfoInput) (count int, err error) {
	var m = dao.PushDevice.Ctx(ctx)
	count, err = m.Where("id = ?", in.Id).Count()
	return
}

// GetDeviceBaseListByServiceId
//
//	@dc:
//	@params:
//	@response:
//	@author:laixin @date:2023/3/30 00:02:52
func (s *sPushDevice) GetDeviceBaseListByServiceId(ctx context.Context, in uint) (out []*entity.PushDevice, err error) {
	var m = dao.PushDevice.Ctx(ctx)
	var deviceList []*entity.PushDevice
	err = m.Where("id>0").Scan(&deviceList)
	if err != nil {
		return nil, err
	}
	for _, v := range deviceList {
		// 判断设备是否启用且服务是否启用
		if !v.DeviceStatus {
			continue
		}
		serviceList := gconv.SliceInt(v.ServiceIdList)
		for _, v2 := range serviceList {
			if v2 == gconv.Int(in) {
				out = append(out, v)
			}
		}
	}
	return
}

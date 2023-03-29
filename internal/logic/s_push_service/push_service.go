package s_push_service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"push/internal/dao"
	"push/internal/model"
	"push/internal/model/entity"
	"push/internal/service"
)

type sPushService struct {
}

func init() {
	service.RegisterPushService(New())
}

func New() *sPushService {
	return &sPushService{}
}

// SearchPushServiceBySign
//
//	@dc:
//	@params:
//	@response:
//	@author:laixin @date:2023/3/29 23:21:44
func (s *sPushService) SearchPushServiceBySign(ctx context.Context, in *model.PushBasicData) (out *entity.PushService, err error) {
	var m = dao.PushService.Ctx(ctx)
	out = new(entity.PushService)
	err = m.Where("service_sign", in.ServiceSign).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// GetPushServiceCount
//
//	@dc: 获取推送服务数量
//	@params:
//	@response:
//	@author:laixin @date:2023/3/29 22:19:24
func (s *sPushService) GetPushServiceCount(ctx context.Context, in *model.GetPushDeviceInfoInput) (out int, err error) {
	var m = dao.PushService.Ctx(ctx)
	out, err = m.Where("id = ?", in.Id).Count()
	return
}

func (s *sPushService) GetPushServiceList(ctx context.Context, input model.GetPushServiceListInput) (output *model.GetPushServiceListOutput, err error) {
	var m = dao.PushService.Ctx(ctx)
	output = &model.GetPushServiceListOutput{
		Page: input.Page,
		Size: input.Size,
	}
	listModel := m.Page(input.Page, input.Size)
	var list []*entity.PushService
	if err := listModel.Scan(&list); err != nil {
		return output, err
	}
	if len(list) == 0 {
		return output, nil
	}
	output.Total, err = m.Count()
	if err != nil {
		return output, err
	}
	output.List = list
	return
}

func (s *sPushService) AddPushService(ctx context.Context, input model.AddPushServiceInput) (err error) {
	var m = dao.PushService.Ctx(ctx)
	// 判断服务名称是否存在
	count, err := m.Where("service_name", input.ServiceName).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("服务名称已存在")
	}
	// 判断服务标识是否存在
	count, err = m.Where("service_sign", input.ServiceSign).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("服务标识已存在")
	}

	_, err = m.Insert(&entity.PushService{
		ServiceName:     input.ServiceName,
		ServiceSign:     input.ServiceSign,
		ServiceStatus:   true,
		ServicePushTime: gtime.Now(),
	})
	if err != nil {
		return err
	}
	return
}

func (s *sPushService) SetPushServiceStatus(ctx context.Context, input model.SetPushServiceInput) (err error) {
	var m = dao.PushService.Ctx(ctx)
	if input.Id == 0 {
		return gerror.New("id不能为空")
	}
	if input.ServiceStatus != true && input.ServiceStatus != false {
		return gerror.New("服务状态不正确")
	}
	_, err = m.Where("id", input.Id).UpdateAndGetAffected(input)
	if err != nil {
		return err
	}
	return
}

// UpdatePushServicePushTime
//
//	@dc: 更新推送服务推送时间
//	@params:
//	@response:
//	@author:laixin @date:2023/3/30 00:33:18
func (s *sPushService) UpdatePushServicePushTime(ctx context.Context, in uint) (err error) {
	var m = dao.PushService.Ctx(ctx)
	_, err = m.Where("id", in).UpdateAndGetAffected(map[string]interface{}{
		"service_push_time": gtime.Now(),
	})
	return
}

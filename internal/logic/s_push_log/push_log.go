package s_push_log

import (
	"context"
	"push/internal/dao"
	"push/internal/model"
	"push/internal/service"
)

type sPushLog struct{}

func init() {
	service.RegisterPushLog(New())
}

func New() *sPushLog {
	return &sPushLog{}
}

// AddPushLog
//
//	@dc: 添加推送日志
//	@params:
//	@response:
//	@author:laixin @date:2023/3/29 23:25:52
func (s *sPushLog) AddPushLog(ctx context.Context, in *model.AddPushLogInput) (err error) {
	var m = dao.PushLog.Ctx(ctx)
	_, err = m.Insert(in)
	return
}

package api_push_core

import (
	"github.com/gogf/gf/v2/frame/g"
	"push/internal/consts"
)

// PushCoreReq 推送核心 Req请求
type PushCoreReq struct {
	g.Meta `method:"post" tags:"推送核心" summary:"推送核心" dc:"推送核心"`
}

// PushCoreRes 推送核心 Res返回
type PushCoreRes struct {
	*consts.DefaultActionMessage
}

package controller

import (
	"context"
	v1 "home-network-watcher/api/v1"
	binInfo "home-network-watcher/utility/bin_utils"
	"runtime"
)

type cVersion struct{}

var Version = &cVersion{}

// GetVersion
//
//	@dc: 获取版本信息
//	@author: laixin   @date:2023/5/17 02:27:52
func (c *cVersion) GetVersion(_ context.Context, _ *v1.GetVersionReq) (res *v1.GetVersionRes, err error) {
	res = &v1.GetVersionRes{
		GitTag:         binInfo.GitTag,
		GitCommitLog:   binInfo.GitCommitLog,
		GitStatus:      binInfo.GitStatus,
		BuildTime:      binInfo.BuildTime,
		BuildGoVersion: binInfo.BuildGoVersion,
		GoVersion:      runtime.GOOS + "/" + runtime.GOARCH,
	}
	return
}

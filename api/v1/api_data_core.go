package v1

import "github.com/gogf/gf/v2/frame/g"

// GetNetworkDataReq 获取网络信息 Req请求
type GetNetworkDataReq struct {
	g.Meta `method:"get" tags:"家庭网络" summary:"获取网络信息" dc:"获取网络信息"`
}

// GetNetworkDataRes 获取网络信息 Res返回
type GetNetworkDataRes struct {
	NodeInfo     interface{} `json:"nodeInfo" dc:"节点信息"`
	HomeNetwork  interface{} `json:"homeNetwork" dc:"家庭网络"`
	ProxyNetwork interface{} `json:"proxyNetwork" dc:"科学上网"`
}

// GetVersionReq 获取版本信息 Req请求
type GetVersionReq struct {
	g.Meta `method:"get" tags:"服务信息" summary:"获取版本信息" dc:"获取版本信息"`
}

// GetVersionRes 获取版本信息 Res返回
type GetVersionRes struct {
	GitTag         string `json:"git_tag" dc:"git_tag"`
	GitCommitLog   string `json:"git_commit_log" dc:"git_commit_log"`
	GitStatus      string `json:"git_status" dc:"git_status"`
	BuildTime      string `json:"build_time" dc:"build_time"`
	BuildGoVersion string `json:"build_go_version" dc:"build_go_version"`
	GoVersion      string `json:"go_version" dc:"go_version"`
	Runtime        string `json:"runtime" dc:"runtime"`
}

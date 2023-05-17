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

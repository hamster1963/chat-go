package api_chat_login

import "github.com/gogf/gf/v2/frame/g"

// LoginReq 用户登入 Req请求
type LoginReq struct {
	g.Meta   `method:"post" tags:"用户管理" summary:"用户登入" dc:"用户登入"`
	Username string `json:"username"     v:"required#请输入用户名"        ` // 用户ID
}

// LoginRes 用户登入 Res返回
type LoginRes struct {
	JWT *string `json:"jwt" dc:"JWT"`
}

// EmptyReq 空请求 Req请求
type EmptyReq struct {
	g.Meta `method:"get" tags:"空请求" summary:"空请求" dc:"空请求"`
}

// EmptyRes 空请求 Res返回
type EmptyRes struct {
}

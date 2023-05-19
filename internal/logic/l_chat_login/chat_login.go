package l_chat_login

import (
	"chat-go/internal/dao"
	"chat-go/internal/global/g_consts"
	"chat-go/internal/model/entity"
	"chat-go/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// ==========================================================================
// logic 初始化
// ==========================================================================

type sChatLogin struct {
}

func init() {
	service.RegisterChatLogin(New())
}

func New() service.IChatLogin {
	return &sChatLogin{}
}

// SelectByUsername
//
//	@author:auto @date:2023/5/19 08:19:54
func (s *sChatLogin) SelectByUsername(ctx context.Context, in *string) (out *entity.ChatUser, err error) {
	db := dao.ChatUser.Ctx(ctx)
	err = db.Where("username", in).Scan(&out)
	if err != nil {
		err = gerror.NewCode(gcode.New(400, "数据库错误", nil), "查询失败")
		return
	}
	return
}

// GenerateUserJWT
//
//	@dc: 生成用户JWT
//	@params: 用户信息-in
//	@response: JWT-out, err-out
//	@author:laixin @date:2023/5/19 11:22:18
func (s *sChatLogin) GenerateUserJWT(ctx context.Context, in *entity.ChatUser) (out *string, err error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "chat-go",
		Audience:  []string{gconv.String(in.Id)},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(g_consts.JWTKey)
	return &ss, err
}

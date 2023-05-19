package g_functions

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func ResErr(code int, errMsg ...interface{}) error {
	resMsg := ""
	errDetail := ""
	switch len(errMsg) {
	case 0:
		resMsg = "Unknown error reason"
	case 1:
		resMsg = gconv.String(errMsg[0])
	default:
		for _, v := range gconv.SliceStr(errMsg[:len(errMsg)-1]) {
			resMsg += v + ","
		}
		resMsg = resMsg[:len(resMsg)-1]
		errDetail = gconv.String(errMsg[len(errMsg)-1])
	}
	return gerror.NewCode(gcode.New(code, resMsg, nil), errDetail)
}

package consts

const (
// DefaultSuccessMessage = "操作成功"
// DefaultFailMessage    = "操作失败"
)

type DefaultActionMessage struct {
	Message string `json:"message" dc:"返回信息"`
}

const (
	ByteToG = 1024 * 1024 * 1024
)

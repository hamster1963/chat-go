package main

import (
	_ "chat-go/internal/logic"
	_ "chat-go/internal/packed"
	"chat-go/internal/router"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	router.Main.Run(gctx.New())
}

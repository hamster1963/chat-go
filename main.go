package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "push/internal/logic"
	_ "push/internal/packed"

	"push/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}

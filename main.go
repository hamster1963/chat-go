package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"home-network-watcher/internal/cmd"
	_ "home-network-watcher/internal/logic"
	_ "home-network-watcher/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}

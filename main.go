package main

import (
	"flag"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"os"
	"push/internal/cmd"
	_ "push/internal/logic"
	_ "push/internal/packed"
	binInfo "push/utility/bin_utils"
)

func main() {
	v := flag.Bool("v", false, "Show bin info.")
	flag.Parse()
	if *v {
		_, _ = fmt.Fprint(os.Stderr, binInfo.StringifyMultiLine())
		os.Exit(1)
	}
	cmd.Main.Run(gctx.New())
}

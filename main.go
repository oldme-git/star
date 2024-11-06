package main

import (
	_ "star/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"star/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

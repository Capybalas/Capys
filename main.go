package main

import (
	_ "Capys/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"Capys/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

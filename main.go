package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "fanzhou-go/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"fanzhou-go/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}

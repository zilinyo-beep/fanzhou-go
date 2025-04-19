package main

import (
	_ "fanzhou-go/internal/app/boot"
	_ "fanzhou-go/internal/app/system/packed"
	"fanzhou-go/internal/cmd"
	_ "fanzhou-go/task"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}

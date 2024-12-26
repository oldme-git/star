package main

import (
	"errors"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"star/internal/cmd"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var err error

	g.I18n().SetLanguage("zh-CN")

	err = connDb()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}

func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("连接到数据库失败")
	}
	return nil
}

package main

import (
	"errors"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "star/internal/logic"
	_ "star/internal/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"star/internal/cmd"
)

func main() {
	var err error

	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")

	// 检查数据库是否能连接
	err = connData()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}

// connData 检查数据库连接是否正常
func connData() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("连接到数据库失败")
	}
	return nil
}

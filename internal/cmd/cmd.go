package cmd

import (
	"context"

	"star/internal/controller/account"
	"star/internal/controller/users"
	"star/internal/controller/words"
	"star/internal/logic/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
    Main = gcmd.Command{
        Name:  "main",
        Usage: "main",
        Brief: "start http server",
        Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
            s := g.Server()
            s.Group("/", func(group *ghttp.RouterGroup) {
                group.Middleware(ghttp.MiddlewareHandlerResponse)
                group.Group("/v1", func(group *ghttp.RouterGroup) {
                    group.Bind(
                        users.NewV1(),
                    )
                    group.Group("/", func(group *ghttp.RouterGroup) {
                        group.Middleware(middleware.Auth)
                        group.Bind(
                            account.NewV1(),
                            words.NewV1(),
                        )
                    })
                })
            })
            s.Run()
            return nil
        },
    }
)

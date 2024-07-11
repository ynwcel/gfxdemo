package ghttpx

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/ynwcel/gfxdemo/internal/ghttpx/apiv1"
	"github.com/ynwcel/gfxdemo/internal/logic/bizx"
	"github.com/ynwcel/gfxdemo/internal/svcx"
)

func Start() error {
	var (
		ctx    = context.Background()
		server = ghttp.GetServer()
	)
	if err := server.SetConfigWithMap(svcx.Cfg().MustGet(ctx, "ghttpx").Map()); err != nil {
		return err
	}
	server.SetSessionStorage(gsession.NewStorageMemory())
	server.Group("/", func(group *ghttp.RouterGroup) {
		server.Group("/api", func(apiGroup *ghttp.RouterGroup) {
			apiGroup.Middleware(bizx.ServerMiddleware())
			apiGroup.GET("/ping", apiv1.PingHandler())
		})
	})

	server.Run()
	return nil
}

package ghttpx

import (
	"io/fs"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/ynwcel/gfxdemo/internal/svcx"
	"github.com/ynwcel/gfxdemo/internal/util"
	"github.com/ynwcel/gfxdemo/pkg/gfx"
	"github.com/ynwcel/gfxdemo/public"
)

func Start() error {
	httpx := gfx.NewHttpServerx(svcx.Cfg().MustGet(gctx.GetInitCtx(), "server").Map())

	bind_assets_router(httpx)
	bind_api_router(httpx)

	return httpx.ListenAndRun()
}

func bind_assets_router(httpx *gfx.HttpServerx) {
	var (
		assets_folder_name = "assets"
		assets_fs          fs.FS
		err                error
	)
	if assets_fs, err = fs.Sub(public.EmbedFS(), assets_folder_name); err != nil {
		panic(gerror.Wrapf(err, "get-embed-subfs-error:(%s)", assets_folder_name))
	}
	httpx.StaticsFS("/assets", assets_fs)
}

func bind_api_router(httpx *gfx.HttpServerx) {
	httpx.Group("/api", func(gApi *ghttp.RouterGroup) {
		gApi.Middleware(ghttp.MiddlewareCORS)
		gApi.GET("/ping", func(r *ghttp.Request) {
			svcx.Log().Info(r.Context(), r.RequestURI)
			data := svcx.Cfg().MustData(r.Context())
			r.Response.WriteJson(util.JsonOk(data, "pong"))
		})
	})
}

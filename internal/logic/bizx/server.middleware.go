package bizx

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/ynwcel/gfxdemo/internal/svcx"
)

func ServerMiddleware() ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		r.Middleware.Next()
		r.Response.Header().Set("gfxversion", svcx.AppVersion())
	}
}

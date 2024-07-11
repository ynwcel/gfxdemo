package bizx

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func ServerMiddleware() ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		r.Middleware.Next()
		r.Response.Header().Set("Server222", "gfxdemo1")
	}
}

package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/ynwcel/gfxdemo/internal/svcx"
)

func PingHandler() ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		svcx.Log().Info(r.Context(), "apiv1/PingHandler")
		r.Response.WriteJson(g.Map{
			"code": 200,
			"data": nil,
			"msg":  "pong",
		})
	}
}

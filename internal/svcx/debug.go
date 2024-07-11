package svcx

import "github.com/gogf/gf/v2/os/gctx"

func AppIsDebug() bool {
	if v, err := Cfg().Get(gctx.GetInitCtx(), "app_debug"); err == nil {
		return v.Bool()
	}
	return false
}

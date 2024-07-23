package svcx

import (
	"os"

	"github.com/gogf/gf/v2/os/gctx"
)

func AppIsDebug() bool {
	if v, err := Cfg().Get(gctx.GetInitCtx(), app_debug_key); err == nil {
		return v.Bool()
	}
	return false
}

func AppVersion() string {
	if v := os.Getenv(app_version_key); len(v) > 0 {
		return v
	}
	return "0.0.1"
}

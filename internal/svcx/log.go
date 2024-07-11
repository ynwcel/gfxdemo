package svcx

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/ynwcel/gfxdemo/internal/core"
)

func Log(group ...string) glog.ILogger {
	var (
		groupName = glog.DefaultName
		cache_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cache_key = fmt.Sprintf("svcx.glog.%s", groupName)
	if !svcxcaches.Contains(cache_key) {
		ctx := gctx.GetInitCtx()
		cfg := Cfg()
		cfg_map := cfg.MustGet(ctx, fmt.Sprintf("logger.%s", groupName)).Map()
		vlog := new_gfglog(cfg_map)
		svcxcaches.Set(cache_key, vlog)
	}
	return svcxcaches.Get(cache_key).(glog.ILogger)
}

func new_gfglog(cfg_map map[string]any) glog.ILogger {
	log := glog.New()
	log.SetConfig(gfglog_default_cfg())
	if len(cfg_map) > 0 {
		if err := log.SetConfigWithMap(cfg_map); err != nil {
			panic(err)
		}
	}
	log = log.Line(true)
	log.SetHandlers(glog.HandlerJson)
	return log
}

func gfglog_default_cfg() glog.Config {
	def := glog.DefaultConfig()
	def.Path = core.APP_RUNTIME_DIR
	def.File = "log.{Y-m-d}.log"
	def.TimeFormat = "2006-01-02 15:04:05.999"
	def.StdoutPrint = false
	def.Level = glog.LEVEL_ALL
	def.Flags = glog.F_FILE_SHORT | glog.F_TIME_DATE | glog.F_TIME_MILLI
	return def
}

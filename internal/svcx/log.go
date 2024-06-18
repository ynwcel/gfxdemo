package svcx

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

func Log(group ...string) glog.ILogger {
	var (
		groupName    = glog.DefaultName
		cfg_key      = ""
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cfg_key = fmt.Sprintf("logger.%s", groupName)
	instance_key = fmt.Sprintf("svcx.%s", cfg_key)
	log := svcx_maps.Get(instance_key)
	if log == nil {
		ctx := gctx.GetInitCtx()
		cfg_map := Cfg().MustGet(ctx, cfg_key).Map()
		vlog := new_gf_glog(cfg_map)
		svcx_maps.Set(instance_key, vlog)
		return vlog
	}
	return log.(glog.ILogger)
}

func new_gf_glog(cfg_map map[string]any) glog.ILogger {
	log := glog.New()
	log.SetConfig(gf_glog_default_cfg())
	if len(cfg_map) > 0 {
		if err := log.SetConfigWithMap(cfg_map); err != nil {
			panic(err)
		}
	}
	log = log.Line(true)
	log.SetHandlers(glog.HandlerJson)
	return log
}

func gf_glog_default_cfg() glog.Config {
	def := glog.DefaultConfig()
	def.Path = RUNTIME_PATH
	def.File = "log.{Y-m-d}.log"
	def.TimeFormat = "2006-01-02 15:04:05.999"
	def.StdoutPrint = false
	def.Level = glog.LEVEL_ALL
	def.Flags = glog.F_FILE_LONG | glog.F_TIME_DATE | glog.F_TIME_MILLI
	return def
}

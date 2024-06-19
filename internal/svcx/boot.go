package svcx

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

const (
	svcx_runmode_mapkey = "svcx.run_mode"
	RUNTIME_PATH        = "./runtimes"
	STORAGES_PATH       = "./storages"
)

var (
	svcx_maps = gmap.NewStrAnyMap(true)
)

func Bootstrap(run_mode_is_debug bool) error {
	if !svcx_maps.SetIfNotExist(svcx_runmode_mapkey, run_mode_is_debug) {
		return gerror.New("Repeat-Set-RunMode")
	}
	if err := setup_dirs(); err != nil {
		return err
	}
	if err := setup_gf_glog(run_mode_is_debug); err != nil {
		return err
	}
	return nil
}

func setup_dirs() error {
	var (
		dirs = []string{
			RUNTIME_PATH,
			STORAGES_PATH,
		}
	)
	for _, d := range dirs {
		if !gfile.IsDir(d) {
			if err := gfile.Mkdir(d); err != nil {
				return err
			}
		}
	}
	return nil
}

func setup_gf_glog(debug bool) error {
	var (
		def_glog    = glog.DefaultLogger()
		def_log_cfg = gf_glog_default_cfg()
	)
	def_log_cfg.StdoutPrint = true
	def_log_cfg.File = "app.{Y-m-d}.log"

	if !debug {
		def_log_cfg.StdoutPrint = false
	}
	def_glog.SetConfig(def_log_cfg)
	def_glog.SetHandlers(glog.HandlerJson)
	glog.SetDefaultLogger(def_glog)
	return nil
}

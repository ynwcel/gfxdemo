package svcx

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/ynwcel/gfxdemo/internal/core"
	"github.com/ynwcel/gfxdemo/internal/util"
)

var svcxcaches = gmap.NewStrAnyMap(true)

func Bootstrap() error {
	var (
		cache_key = "svcx.bootstraped"
	)
	if svcxcaches.Contains(cache_key) {
		return gerror.New("repeat run svcx.bootstrap")
	}
	return util.Try(func() error {
		if err := bootstrap_mkdirs(); err != nil {
			return err
		}
		if err := bootstrap_goframe(); err != nil {
			return err
		}
		_ = Cfg()
		return nil
	})
}

func bootstrap_mkdirs() error {
	dirs := []string{
		core.APP_RUNTIME_DIR,
		core.APP_STORAGE_DIR,
	}
	for _, d := range dirs {
		if !gfile.IsDir(d) {
			if err := gfile.Mkdir(d); err != nil {
				return err
			}
		}
	}
	return nil
}

func bootstrap_goframe() error {
	return nil
}

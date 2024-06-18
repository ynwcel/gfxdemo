package svcx

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/ynwcel/gfxdemo/public"
)

func Cfg(names ...string) *gcfg.Config {
	var (
		cfg_name     = gcfg.DefaultConfigFileName
		svcx_cfg_key = ""
	)
	if len(names) > 0 && len(names[0]) > 0 {
		cfg_name = names[0]
	}
	svcx_cfg_key = fmt.Sprintf("svcx.cfg.%s", cfg_name)
	cfg := svcx_maps.Get(svcx_cfg_key)
	if cfg == nil {
		var (
			cfg_fname   = fmt.Sprintf("%s.yaml", strings.TrimRight(cfg_name, filepath.Ext(cfg_name)))
			cfg_adapt   *gcfg.AdapterContent
			cfg_content = gfile.GetBytes(fmt.Sprintf("./%s", cfg_fname))
			err         error
		)
		if len(cfg_content) <= 0 {
			cfg_content, err = public.ReadFile(cfg_fname)
		}
		if err != nil {
			panic(gerror.Wrapf(err, "read-cfg-file-error:%s", cfg_fname))
		}
		if cfg_adapt, err = gcfg.NewAdapterContent(string(cfg_content)); err != nil {
			panic(gerror.Wrap(err, "new-cfg-adapter-content-error"))
		}
		cfg = gcfg.NewWithAdapter(cfg_adapt)
		svcx_maps.Set(svcx_cfg_key, cfg)
	}
	return cfg.(*gcfg.Config)
}

package svcx

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/ynwcel/gfxdemo/internal/core"
	"github.com/ynwcel/gfxdemo/public"
)

func Cfg() *gcfg.Config {
	var (
		cache_key = "svcx.gcfg"
	)
	if !svcxcaches.Contains(cache_key) {
		var (
			cur_cfg_file = fmt.Sprintf("./%s", core.APP_CFG_FILENAME)
			pub_cfg_file = fmt.Sprintf("./public/%s", core.APP_CFG_FILENAME)
			cfg_adapter  *gcfg.AdapterContent
			cfg_content  []byte
			err          error
		)
		if gfile.Exists(cur_cfg_file) {
			if cfg_content, err = os.ReadFile(cur_cfg_file); err != nil {
				panic(gerror.Wrapf(err, "read-cfg-file-error:%v", cur_cfg_file))
			}
		} else if gfile.Exists(pub_cfg_file) {
			if cfg_content, err = os.ReadFile(pub_cfg_file); err != nil {
				panic(gerror.Wrapf(err, "read-cfg-file-error:%v", pub_cfg_file))
			}
		} else {
			if cfg_content, err = fs.ReadFile(public.FS(true), core.APP_CFG_FILENAME); err != nil {
				panic(gerror.Wrapf(err, "read-embed-cfg-file-error:%v", core.APP_CFG_FILENAME))
			}
		}
		if cfg_adapter, err = gcfg.NewAdapterContent(string(cfg_content)); err != nil {
			panic(gerror.Wrap(err, "new-cfg-content-adapter-error"))
		}
		svcxcaches.Set(cache_key, gcfg.NewWithAdapter(cfg_adapter))
	}
	return svcxcaches.Get(cache_key).(*gcfg.Config)
}

package svcx

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func Cache(group ...string) *gcache.Cache {
	var (
		groupName = gredis.DefaultGroupName
		cache_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cache_key = fmt.Sprintf("svcx.gcache.%s", groupName)
	if !svcxcaches.Contains(cache_key) {
		var (
			ctx              = gctx.GetInitCtx()
			vcache           *gcache.Cache
			cfg_sub_type_key string = fmt.Sprintf("cache.%s.type", groupName)
			cfg_type_val     string
		)
		if cfg_val, err := Cfg().Get(ctx, cfg_sub_type_key); err == nil {
			cfg_type_val = cfg_val.String()
		}
		if strings.EqualFold(cfg_type_val, "redis") {
			redis := new_gfgreids(Cfg().MustGet(ctx, fmt.Sprintf("cache.%s", groupName)).Map())
			vcache = gcache.NewWithAdapter(gcache.NewAdapterRedis(redis))
		} else {
			vcache = gcache.NewWithAdapter(gcache.NewAdapterMemory())
		}
		svcxcaches.Set(cache_key, vcache)
	}
	return svcxcaches.Get(cache_key).(*gcache.Cache)
}

package svcx

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func Cache(group ...string) *gcache.Cache {

	var (
		groupName    = "default"
		cfg_key      = ""
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cfg_key = fmt.Sprintf("cache.%s", groupName)
	instance_key = fmt.Sprintf("svcx.%s", cfg_key)
	cache := svcx_maps.Get(instance_key)
	if cache == nil {
		var (
			ctx              = gctx.GetInitCtx()
			adapter          gcache.Adapter
			vcache           *gcache.Cache
			cfg_sub_type_key string = fmt.Sprintf("%s.type", cfg_key)
		)
		if cache_type := Cfg().MustGet(ctx, cfg_sub_type_key).String(); strings.EqualFold(cache_type, "redis") {
			redis := new_gf_greids(Cfg().MustGet(ctx, cfg_key).Map())
			adapter = gcache.NewAdapterRedis(redis)
		}
		if adapter == nil {
			adapter = gcache.NewAdapterMemory()
		}
		vcache = gcache.NewWithAdapter(adapter)
		svcx_maps.Set(instance_key, vcache)
		return vcache
	}
	return cache.(*gcache.Cache)
}

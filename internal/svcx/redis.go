package svcx

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func Redis(group ...string) *gredis.Redis {
	var (
		groupName    = gredis.DefaultGroupName
		cfg_key      = ""
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cfg_key = fmt.Sprintf("redis.%s", groupName)
	instance_key = fmt.Sprintf("svcx.%s", cfg_key)
	redis := svcx_maps.Get(instance_key)
	if redis == nil {
		ctx := gctx.GetInitCtx()
		cfg_map := Cfg().MustGet(ctx, cfg_key).Map()
		vredis := new_gf_greids(cfg_map)
		if _, err := vredis.Do(gctx.GetInitCtx(), "ping"); err != nil {
			panic(err)
		}
		svcx_maps.Set(instance_key, vredis)
		return vredis
	}
	return redis.(*gredis.Redis)
}

func new_gf_greids(cfg map[string]any) *gredis.Redis {
	var (
		redis    *gredis.Redis
		cfg_node = new(gredis.Config)
		err      error
	)
	if err = gconv.Scan(cfg, &cfg_node); err != nil {
		panic(gerror.Wrap(err, "parse-redis-config-failed"))
	}
	if cfg_node.DialTimeout <= 0 {
		cfg_node.DialTimeout = time.Second
	}
	if cfg_node.ReadTimeout <= 0 {
		cfg_node.ReadTimeout = time.Second
	}

	if cfg_node.WriteTimeout <= 0 {
		cfg_node.WriteTimeout = time.Second
	}

	if redis, err = gredis.New(cfg_node); err != nil {
		panic(gerror.Wrapf(err, "new-redis-instance-failed"))
	}
	return redis
}

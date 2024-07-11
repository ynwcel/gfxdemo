package svcx

import (
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func DB(group ...string) gdb.DB {
	var (
		groupName = gdb.DefaultGroupName
		cache_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cache_key = fmt.Sprintf("svcx.gdb.%s", groupName)
	if !svcxcaches.Contains(cache_key) {
		var (
			ctx = gctx.GetInitCtx()
			cfg = Cfg().MustGet(ctx, fmt.Sprintf("database.%s", groupName)).Interface()
			vdb = new_gfgdb(groupName, cfg)
		)
		if AppIsDebug() {
			vdb.SetDebug(true)
			vdb.SetLogger(Log())
		}
		if db_logger, err := Cfg().Get(ctx, "database.logger"); err == nil {
			vdb.SetLogger(new_gfglog(db_logger.Map()))
		}
		svcxcaches.Set(cache_key, vdb)
	}
	return svcxcaches.Get(cache_key).(gdb.DB)
}
func new_gfgdb(group string, cfg any) gdb.DB {

	switch cfg.(type) {
	case []any:
		return new_gdb_by_slices(group, cfg.([]any))
	case map[string]any:
		return new_gdb_by_map(cfg.(map[string]any))
	default:
		panic("invalid-param-of-new-gdb")
	}
}

func new_gdb_by_slices(group string, cfg []any) gdb.DB {
	var (
		cfg_group = gdb.ConfigGroup{}
		db        gdb.DB
		err       error
	)
	if err = gconv.Scan(cfg, &cfg_group); err != nil {
		panic(gerror.Wrap(err, "scan-gdb-config-node-failed"))
	}

	gdb.SetConfigGroup(group, cfg_group)
	if db, err = gdb.NewByGroup(group); err != nil {
		panic(gerror.Wrap(err, "new-gdb-failed"))
	}
	return db
}

func new_gdb_by_map(cfg map[string]any) gdb.DB {
	var (
		dbnode_cfg = gdb.ConfigNode{}
		db         gdb.DB
		err        error
	)
	if err = gconv.Scan(cfg, &dbnode_cfg); err != nil {
		panic(gerror.Wrap(err, "scan-gdb-config-node-failed"))
	}
	if db, err = gdb.New(dbnode_cfg); err != nil {
		panic(gerror.Wrap(err, "new-gdb-failed"))
	}
	return db
}

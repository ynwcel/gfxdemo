package svcx

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func DB(group ...string) gdb.DB {
	var (
		groupName    = gdb.DefaultGroupName
		cfg_key      = ""
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cfg_key = fmt.Sprintf("database.%s", groupName)
	instance_key = fmt.Sprintf("svcx.%s", cfg_key)
	db := svcx_maps.Get(instance_key)
	if db == nil {
		ctx := gctx.GetInitCtx()
		cfg_any := Cfg().MustGet(ctx, cfg_key).Interface()
		vdb := new_gf_gdb(cfg_any)
		if err1, err2 := vdb.PingMaster(), vdb.PingSlave(); err1 != nil || err2 != nil {
			if err1 != nil {
				panic(err1)
			} else {
				panic(err2)
			}
		}
		if RunModeIsDebug() {
			vdb.SetDebug(true)
			vdb.SetLogger(Log())
		}
		if db_logger, err := Cfg().Get(ctx, "database.logger"); err == nil {
			vdb.SetLogger(new_gf_glog(db_logger.Map()))
		}
		svcx_maps.Set(instance_key, vdb)
		return vdb
	}
	return db.(gdb.DB)
}

func new_gf_gdb(cfg any) gdb.DB {

	switch cfg.(type) {
	case []any:
		return new_gdb_by_slices(cfg.([]any))
	case map[string]any:
		return new_gdb_by_map(cfg.(map[string]any))
	default:
		panic("invalid-param-of-new-gdb")
	}
}

func new_gdb_by_slices(cfg []any) gdb.DB {
	var (
		groupName = fmt.Sprintf("pgdb_%d", time.Now().Unix())
		group     = gdb.ConfigGroup{}
		db        gdb.DB
		err       error
	)
	if err = gconv.Scan(cfg, &group); err != nil {
		panic(gerror.Wrap(err, "scan-gdb-config-node-failed"))
	}

	gdb.SetConfigGroup(groupName, group)
	if db, err = gdb.NewByGroup(groupName); err != nil {
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

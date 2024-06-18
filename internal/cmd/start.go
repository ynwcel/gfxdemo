package cmd

import (
	"github.com/ynwcel/gfxdemo/internal/gcronx"
	"github.com/ynwcel/gfxdemo/internal/ghttpx"
	"github.com/ynwcel/gfxdemo/internal/grpcx"
	"github.com/ynwcel/gfxdemo/internal/util"
	"golang.org/x/sync/errgroup"
)

func (cx *cmdx) start_process() error {
	var (
		err_group = new(errgroup.Group)
	)
	if cx.ghttpx {
		err_group.Go(func() error {
			return util.Try(ghttpx.Start)
		})
	}
	if cx.gcronx {
		err_group.Go(func() error {
			return util.Try(gcronx.Start)
		})
	}
	if cx.grpcx {
		err_group.Go(func() error {
			return util.Try(grpcx.Start)
		})
	}
	return err_group.Wait()
}

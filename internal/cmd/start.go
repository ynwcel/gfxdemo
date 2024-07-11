package cmd

import (
	"os"
	"os/signal"

	"github.com/ynwcel/gfxdemo/internal/gcronx"
	"github.com/ynwcel/gfxdemo/internal/ghttpx"
	"github.com/ynwcel/gfxdemo/internal/grpcx"
	"github.com/ynwcel/gfxdemo/internal/svcx"
	"github.com/ynwcel/gfxdemo/internal/util"
	"github.com/ynwcel/gox/gerrgroup"
)

func (cx *cmdx) start_sub_process() error {
	if !cx.ghttpx && !cx.gcronx && !cx.grpcx {
		return nil
	}
	var (
		signChan = make(chan os.Signal, 1)
		errgroup = gerrgroup.New()
		err      error
	)
	if err = svcx.Bootstrap(); err != nil {
		return err
	}
	if cx.ghttpx {
		errgroup.Go(func() error {
			return util.Try(ghttpx.Start)
		})
	}
	if cx.gcronx {
		errgroup.Go(func() error {
			return util.Try(gcronx.Start)
		})
	}
	if cx.gcronx {
		errgroup.Go(func() error {
			return util.Try(grpcx.Start)
		})
	}

	signal.Notify(signChan, os.Interrupt)
	select {
	case err := <-errgroup.Wait():
		return err
	case <-signChan:
		return nil
	}
}

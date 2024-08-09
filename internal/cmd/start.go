package cmd

import (
	"os"
	"os/signal"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/ynwcel/gfxdemo/internal/gcronx"
	"github.com/ynwcel/gfxdemo/internal/ghttpx"
	"github.com/ynwcel/gfxdemo/internal/grpcx"
	"github.com/ynwcel/gfxdemo/internal/svcx"
	"github.com/ynwcel/gfxdemo/internal/util"
)

func (cx *cmdx) start_process() error {
	if !cx.ghttpx && !cx.gcronx && !cx.grpcx {
		return nil
	}
	if err := svcx.Bootstrap(cx.version); err != nil {
		return gerror.Wrap(err, "failed to bootstrap service")
	}
	var (
		signChan = make(chan os.Signal, 2)
		errChan  = make(chan error, 3)
	)
	if cx.ghttpx {
		go func() {
			if err := util.Try(ghttpx.Start); err != nil {
				errChan <- err
			}
		}()
	}
	if cx.gcronx {
		go func() {
			if err := util.Try(gcronx.Start); err != nil {
				errChan <- err
			}
		}()
	}
	if cx.grpcx {
		go func() {
			if err := util.Try(grpcx.Start); err != nil {
				errChan <- err
			}
		}()
	}

	signal.Notify(signChan, os.Interrupt, os.Kill)
	select {
	case err := <-errChan:
		return gerror.Wrap(err, "one or more services failed to start")
	case <-signChan:
		return gerror.New("user signal to stop")
	}
}

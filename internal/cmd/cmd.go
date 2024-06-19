package cmd

import (
	"os"

	"github.com/ynwcel/gox/gflagx"
)

type cmdx struct {
	flag   *gflagx.Flagx
	debug  bool
	gcronx bool
	ghttpx bool
	grpcx  bool
	unpack bool
}

func New(appVersion string) *cmdx {
	cx := cmdx{
		flag: gflagx.NewFlagx().SetVersion(appVersion),
	}
	cx.flag.BoolVar(&cx.debug, "debug", false, "set run mode is debug")
	cx.flag.BoolVar(&cx.ghttpx, "ghttpx", false, "start run ghttpx")
	cx.flag.BoolVar(&cx.gcronx, "gcronx", false, "start run gcronx")
	cx.flag.BoolVar(&cx.grpcx, "grpcx", false, "start run grpcx")
	cx.flag.BoolVar(&cx.unpack, "unpack", false, "unpack public folder")
	return &cx
}

func (cx *cmdx) Run() error {
	var (
		args = os.Args[1:]
		err  error
	)
	if err = cx.flag.Parse(args); err != nil {
		return err
	}
	if len(args) <= 0 || cx.flag.HasSetHelpFlag() {
		cx.flag.Usage()
		return nil
	}
	if err = cx.unpack_public(); err != nil {
		return err
	}
	if err = cx.start_process(); err != nil {
		return err
	}
	return nil
}

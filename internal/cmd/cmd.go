package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

type cmdx struct {
	flagset *pflag.FlagSet
	unpack  bool
	ghttpx  bool
	gcronx  bool
	grpcx   bool
	help    bool
}

func New(appVersion string) *cmdx {
	app := os.Args[0]
	cx := &cmdx{
		flagset: pflag.NewFlagSet(filepath.Base(app), pflag.ContinueOnError),
	}
	cx.flagset.SortFlags = false
	cx.flagset.BoolVar(&cx.ghttpx, "ghttpx", false, "start ghttpx process")
	cx.flagset.BoolVar(&cx.gcronx, "gcronx", false, "start gcronx process")
	cx.flagset.BoolVar(&cx.grpcx, "grpcx", false, "start grpcx process")
	cx.flagset.BoolVar(&cx.unpack, "unpack", false, "unpack public folder")
	cx.flagset.BoolVarP(&cx.help, "help", "h", false, "show help message")
	cx.flagset.Usage = func() {
		fmt.Println("Usage:")
		fmt.Printf("  %s [options]\n", filepath.Base(app))
		fmt.Printf("Options:\n")
		cx.flagset.PrintDefaults()
		fmt.Println("Version:")
		fmt.Printf("  %s", appVersion)
		fmt.Println()
	}
	return cx
}

func (cx *cmdx) Run() error {
	var (
		args = os.Args[1:]
	)
	if err := cx.flagset.Parse(args); err != nil {
		return err
	}
	if len(args) <= 0 || cx.help {
		cx.flagset.Usage()
		return nil
	}
	if err := cx.unpack_public_resouce(); err != nil {
		return err
	}
	if err := cx.start_sub_process(); err != nil {
		return err
	}
	return nil
}

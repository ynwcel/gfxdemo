package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

type cmdx struct {
	flagset *pflag.FlagSet
	version string
	unpack  bool
	ghttpx  bool
	gcronx  bool
	grpcx   bool
	help    bool
}

func NewWithVersion(appVersion string) *cmdx {
	cx := &cmdx{
		version: appVersion,
		flagset: pflag.NewFlagSet(filepath.Base(os.Args[0]), pflag.ContinueOnError),
	}
	cx.flagset.SortFlags = false
	cx.flagset.BoolVar(&cx.ghttpx, "ghttpx", false, "start ghttpx process")
	cx.flagset.BoolVar(&cx.gcronx, "gcronx", false, "start gcronx process")
	cx.flagset.BoolVar(&cx.grpcx, "grpcx", false, "start grpcx process")
	cx.flagset.BoolVar(&cx.unpack, "unpack", false, "unpack public folder")
	cx.flagset.BoolVarP(&cx.help, "help", "h", false, "show help message")
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
		fmt.Println("Usage:")
		fmt.Printf("  %s [options]\n", filepath.Base(os.Args[0]))
		fmt.Printf("Options:\n")
		cx.flagset.PrintDefaults()
		fmt.Println("Version:")
		fmt.Printf("  %s", cx.version)
		fmt.Println()
		return nil
	}
	if err := cx.unpack_resouce(); err != nil {
		return err
	}
	if err := cx.start_process(); err != nil {
		return err
	}
	return nil
}

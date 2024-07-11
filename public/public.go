package public

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	"github.com/ynwcel/gfxdemo/internal/core"
)

var (
	//go:embed all:**
	embed_fs embed.FS
	real_fs  = os.DirFS(core.APP_PUBLIC_DIR)
)

func FS(embed bool) fs.FS {
	if embed {
		return embed_fs
	} else {
		return real_fs
	}
}

func ListEmbedFils() ([]string, error) {
	var (
		rfs    = FS(true)
		listfn func(fs.FS, string) ([]string, error)
	)
	listfn = func(f fs.FS, s string) ([]string, error) {
		var result = make([]string, 0, 50)
		files, err := fs.Glob(f, s)
		if err != nil {
			return result, err
		}
		for _, fstr := range files {
			fstat, err := fs.Stat(f, fstr)
			if err != nil {
				return result, err
			}
			result = append(result, fstr)
			if fstat.IsDir() {
				if subfs, err := listfn(f, fmt.Sprintf("%s/*", fstr)); err != nil {
					return result, err
				} else {
					result = append(result, subfs...)
				}
			}
		}
		return result, nil
	}
	return listfn(rfs, "*")
}

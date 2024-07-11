package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/ynwcel/gfxdemo/internal/core"
	"github.com/ynwcel/gfxdemo/public"
)

func (cx *cmdx) unpack_resouce() error {
	if !cx.unpack {
		return nil
	}
	var (
		base_dir   = core.APP_PUBLIC_DIR
		embed_fs   = public.FS(true)
		files, err = public.ListEmbedFils()
	)
	if err != nil {
		return gerror.Wrap(err, "get-embed-files-error")
	}
	if err = gfile.Mkdir(base_dir); err != nil {
		return gerror.Wrap(err, "create-public-folder-error")
	}
	for _, file := range files {
		if strings.EqualFold(filepath.Ext(file), ".go") {
			continue
		}
		public_file := fmt.Sprintf("%s/%s", base_dir, strings.TrimLeft(file, "/"))
		fstat, err := fs.Stat(embed_fs, file)
		if err != nil {
			return gerror.Wrap(err, "get-embed-file-stat-error")
		}
		if fstat.IsDir() {
			gfile.Mkdir(public_file)
		} else if content, err := fs.ReadFile(embed_fs, file); err != nil {
			return gerror.Wrapf(err, "get-embed-file-content-error:(%s)", file)
		} else if err := gfile.PutBytes(public_file, content); err != nil {
			return gerror.Wrapf(err, "write-file-error:(%s)", public_file)
		}
	}
	return nil
}

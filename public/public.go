package public

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

var (
	real_path = "./public"
	//go:embed all:*
	embed_fs embed.FS
	real_fs  = os.DirFS(real_path)
)

func EmbedFS() fs.FS {
	return embed_fs
}

func RealFS() fs.FS {
	return real_fs
}

func Stat(path string) (fs.FileInfo, error) {
	if rstate, err := fs.Stat(real_fs, path); err == nil {
		return rstate, nil
	} else {
		return fs.Stat(embed_fs, path)
	}
}

func ReadFile(filename string) ([]byte, error) {
	if content, err := fs.ReadFile(real_fs, filename); err == nil {
		return content, nil
	} else {
		return fs.ReadFile(embed_fs, filename)
	}
}

func ReadDir(path string) ([]os.DirEntry, error) {
	if ds, err := fs.ReadDir(real_fs, path); err == nil {
		return ds, nil
	} else {
		return fs.ReadDir(embed_fs, path)
	}
}

func ListFiles() ([]string, error) {
	var (
		rfs    = EmbedFS()
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

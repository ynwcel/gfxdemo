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

// FS 根据 embed 参数决定返回嵌入式文件系统还是实际文件系统。
// 当 embed 为 true 时，返回嵌入式文件系统；否则，返回实际文件系统。
func FS(embed bool) fs.FS {
	if embed {
		return embed_fs
	} else {
		return real_fs
	}
}

// SubFS 根据目录名返回一个文件系统子集。
// 它首先尝试使用现实世界中的文件系统(real_fs)来查找目录。
// 如果目录不存在于现实世界中的文件系统中，它将尝试在嵌入的文件系统(embed_fs)中查找。
func SubFS(dirname string) (fs.FS, error) {
	if _, err := fs.Stat(real_fs, dirname); err != nil {
		return fs.Sub(real_fs, dirname)
	} else {
		return fs.Sub(embed_fs, dirname)
	}
}

// ReadFile 根据文件路径读取文件内容。
// 它首先尝试从实际文件系统中读取文件。如果失败了，它会尝试从嵌入的文件系统中读取。
// 这个函数的存在是为了提供一种从两种不同来源读取文件内容的灵活方式。
func ReadFile(fpath string) ([]byte, error) {
	if content, err := fs.ReadFile(real_fs, fpath); err == nil {
		return content, err
	} else {
		return fs.ReadFile(embed_fs, fpath)
	}
}

// ReadDir 读取指定目录下的文件和子目录信息。
// 它首先尝试从现实文件系统(real_fs)中读取目录内容，如果失败，则尝试从嵌入文件系统(embed_fs)中读取。
// 这种设计允许在不同的文件系统之间灵活切换，例如在本地开发和部署时使用不同的文件系统。
func ReadDir(dirpath string) ([]fs.DirEntry, error) {
	if dirs, err := fs.ReadDir(real_fs, dirpath); err == nil {
		return dirs, err
	} else {
		return fs.ReadDir(embed_fs, dirpath)
	}
}

// Stat 函数用于获取给定文件路径的文件信息。
// 它首先尝试从实际文件系统中获取文件信息。如果失败，则尝试从嵌入的文件系统中获取。
// 这个函数的存在是为了提供一种统一的方式来访问不同来源的文件信息，而不需要调用者关心具体的文件系统实现。
func Stat(fpath string) (fs.FileInfo, error) {
	if info, err := fs.Stat(real_fs, fpath); err == nil {
		return info, err
	} else {
		return fs.Stat(embed_fs, fpath)
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

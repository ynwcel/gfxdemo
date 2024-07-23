package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	var (
		setgomod      = flag.String("setgomod", "", "change go mod name")
		build         = flag.Bool("build", false, "run go build")
		build_version = flag.String("build.version", "0.0.1", "set build version")
		err           error
	)
	flag.Usage = func() {
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	switch {
	case *build:
		err = hack_build(*build_version)
	case len(*setgomod) > 0:
		err = hack_setgomod(*setgomod)
	default:
		flag.Usage()
		err = fmt.Errorf("no command")
	}
	if err != nil {
		fmt.Printf("Error:%v\n", err)
		os.Exit(1)
	}
}
func hack_build(build_version string) error {
	var (
		cur_time = time.Now().Format("060102.1504")
		cmd_args = []string{"build", "-ldflags", fmt.Sprintf("-X main.appVersion=%s.%s", build_version, cur_time)}
	)
	fmt.Printf("> go %s\n", strings.Join(cmd_args, " "))
	if err := exec.Command("go", cmd_args...).Run(); err != nil {
		return fmt.Errorf("run-go-build-error:%W", err)
	}
	return nil
}

func hack_setgomod(new_mod string) error {
	var (
		cur_dir           = filepath.Dir(".")
		old_mod           string
		old_gomod_file    = fmt.Sprintf("%s/go.mod", cur_dir)
		old_gomod_content []byte
		walk              func(path string) []string
		err               error
		mod_regexp        = regexp.MustCompile(`^[\w][\w\d\/\-]+$`)
		gofile_regexp     = regexp.MustCompile(`.*?\.go.*?`)
		cmd_args          = []string{"mod", "tidy"}
	)
	if !mod_regexp.Match([]byte(new_mod)) {
		return errors.New("new-mod-name-error")
	}
	if old_gomod_content, err = os.ReadFile(old_gomod_file); err != nil {
		return fmt.Errorf("read-go.mod-error:%W", err)
	}
	lines := strings.Split(string(old_gomod_content), "\n")
	for _, line := range lines {
		if strings.Index(line, "module ") == 0 {
			old_mod = strings.Split(line, "module ")[1]
			break
		}
	}
	if len(old_mod) == 0 {
		return fmt.Errorf("get-old-mod-name-error")
	}
	walk = func(path string) []string {
		var (
			dirfs         = os.DirFS(path)
			curfiles, err = fs.Glob(dirfs, "*")
			result        = make([]string, 0, len(curfiles)*10)
		)
		if err == nil {
			for _, f := range curfiles {
				fpath := filepath.Join(path, f)
				result = append(result, fpath)
				if finfo, err := fs.Stat(dirfs, f); err == nil {
					if finfo.IsDir() {
						subfs := walk(filepath.Join(path, f))
						result = append(result, subfs...)
					}
				}
			}
		}
		return result
	}
	for _, f := range walk(cur_dir) {
		if gofile_regexp.Match([]byte(f)) || filepath.Base(f) == "go.mod" {
			f_content, err := os.ReadFile(f)
			if err != nil {
				return fmt.Errorf("read-go-file-error:%W", err)
			}
			if bytes.Contains(f_content, []byte(old_mod)) {
				f_content_new := bytes.ReplaceAll(f_content, []byte(old_mod), []byte(new_mod))
				if err = os.WriteFile(f, f_content_new, 0666); err != nil {
					return fmt.Errorf("save-go-file-error:%W", err)
				}
			}
		}
	}
	fmt.Printf("> go %s\n", strings.Join(cmd_args, " "))
	if err = exec.Command("go", cmd_args...).Run(); err != nil {
		return fmt.Errorf("run-go-mod-tidy-error:%W", err)
	}
	return nil
}

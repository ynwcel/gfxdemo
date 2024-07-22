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
		setmode = flag.String("setmode", "", "change package module name")
		build   = flag.String("build", "", "set build version and run go build")
		err     error
	)
	flag.Usage = func() {
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	switch {
	case len(*build) > 0:
		err = hack_build(*build)
	case len(*setmode) > 0:
		err = hack_setmode(*setmode)
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

func hack_setmode(new_mode string) error {
	var (
		cur_dir           = filepath.Dir(".")
		old_mode          string
		old_gomod_file    = fmt.Sprintf("%s/go.mod", cur_dir)
		old_gomod_content []byte
		walk              func(path string) []string
		err               error
		mod_regexp        = regexp.MustCompile(`^[\W\w][\w\W\/\-]+$`)
		gofile_regexp     = regexp.MustCompile(`.*?\.go.*?`)
		cmd_args          = []string{"mod", "tidy"}
	)
	if !mod_regexp.Match([]byte(new_mode)) {
		return errors.New("new-mode-name-error")
	}
	if old_gomod_content, err = os.ReadFile(old_gomod_file); err != nil {
		return fmt.Errorf("read-go.mode-error:%W", err)
	}
	lines := strings.Split(string(old_gomod_content), "\n")
	for _, line := range lines {
		if strings.Index(line, "module ") == 0 {
			old_mode = strings.Split(line, "module ")[1]
			break
		}
	}
	if len(old_mode) == 0 {
		return fmt.Errorf("get-old-mode-name-error")
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
			if bytes.Contains(f_content, []byte(old_mode)) {
				f_content_new := bytes.ReplaceAll(f_content, []byte(old_mode), []byte(new_mode))
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

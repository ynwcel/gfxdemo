package gfx

import (
	"bytes"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

type HttpServerx struct {
	*ghttp.Server
}

func NewHttpServerx(cfg map[string]any, names ...string) *HttpServerx {
	name := ghttp.DefaultServerName
	if len(names) > 0 && len(names[0]) > 0 {
		name = names[0]
	}
	sx := &HttpServerx{
		Server: ghttp.GetServer(name),
	}
	sx_cfg := new(ghttp.ServerConfig)
	gvar.New(cfg).Scan(&sx_cfg)
	sx.SetConfig(*sx_cfg)
	if sx.Logger() == nil {
		sx.SetLogger(glog.DefaultLogger())
	}
	return sx
}

func (sx *HttpServerx) ListenAndRun() error {
	if err := sx.Start(); err != nil {
		return err
	}
	ghttp.Wait()
	return nil
}

func (sx *HttpServerx) StaticsFS(prefix string, staticfs fs.FS, logs ...glog.ILogger) {
	var log glog.ILogger = sx.Server.Logger()
	if len(logs) > 0 && logs[0] != nil {
		log = logs[0]
	}
	sx.BindHandler(fmt.Sprintf("%s/*", prefix), func(r *ghttp.Request) {
		var (
			req_uri  = r.RequestURI
			req_file = strings.TrimLeft(req_uri, prefix)
		)
		if content, err := fs.ReadFile(staticfs, req_file); err != nil {
			log.Errorf(r.GetCtx(), "not-found-file:%s", req_uri)
			r.Response.WriteStatus(http.StatusNotFound, "Not Found")
		} else {
			fstat, _ := fs.Stat(staticfs, req_file)
			http.ServeContent(r.Response.Writer, r.Request, req_file, fstat.ModTime(), bytes.NewReader(content))
		}
	})
}

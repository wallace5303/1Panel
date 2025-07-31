package proxy

import (
	"context"
	"net"
	"net/http"
	"net/http/httputil"
	"path"
	"time"

	"github.com/1Panel-dev/1Panel/core/global"
)

var (
	sockPath = "/etc/1panel/agent.sock"

	LocalAgentProxy *httputil.ReverseProxy
)

func Init() {
	// [gsx] 重新赋值
	sockPath = path.Join(global.CONF.Base.InstallDir, "1panel/agent.sock")
	dialUnix := func(ctx context.Context, network, addr string) (net.Conn, error) {
		return net.Dial("unix", sockPath)
	}
	transport := &http.Transport{
		DialContext:         dialUnix,
		ForceAttemptHTTP2:   false,
		MaxIdleConns:        50,
		MaxIdleConnsPerHost: 50,
		IdleConnTimeout:     30 * time.Second,
	}
	LocalAgentProxy = &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "unix"
		},
		Transport: transport,
		ErrorHandler: func(rw http.ResponseWriter, req *http.Request, err error) {
			rw.WriteHeader(http.StatusBadGateway)
			rw.Write([]byte("Bad Gateway"))
		},
	}
}

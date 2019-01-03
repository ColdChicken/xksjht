package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"runtime"

	"be/common/log"

	"github.com/gorilla/mux"
)

type WWWMux struct {
	r *mux.Router
}

func New() *WWWMux {
	return &WWWMux{r: mux.NewRouter()}
}

func (m *WWWMux) GetRouter() *mux.Router {
	return m.r
}

// 记录日志
func AccessLogHandler(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s - %s", r.Method, r.RequestURI)
		h(w, r)
	}
}

func (m *WWWMux) RegistURLMapping(path string, method string, handle func(http.ResponseWriter, *http.Request)) {
	log.WithFields(log.Fields{
		"path":   path,
		"method": method,
		"handle": runtime.FuncForPC(reflect.ValueOf(handle).Pointer()).Name(),
	}).Info("注册URL映射")
	handle = AccessLogHandler(handle)
	m.r.HandleFunc(path, handle).Methods(method)
}

// 代理的handler
func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debugln("代理请求")
		p.ServeHTTP(w, r)
	}
}

func (m *WWWMux) SetProxy(path string, targetAddress string) {
	remote, err := url.Parse(targetAddress)
	if err != nil {
		log.Errorf("代理注册失败 %s", targetAddress)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	m.r.HandleFunc(path, handler(proxy))
}

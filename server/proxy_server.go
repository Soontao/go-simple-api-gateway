package server

import (
	"github.com/elazarl/goproxy"
	"net/http"
)

type ProxyServer struct {
	*goproxy.ProxyHttpServer
}

func NewProxyServer(ls string) (p *ProxyServer) {
	p = &ProxyServer{goproxy.NewProxyHttpServer()}
	return
}

func (p *ProxyServer) Start(ls string) {
	http.ListenAndServe(ls, p)
}

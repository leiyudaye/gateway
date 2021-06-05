/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-05 18:57:31
 * @LastEditors: lly
 * @LastEditTime: 2021-06-05 23:45:42
 */
package gateway

import (
	"net"
	"net/http"

	"github.com/leiyudaye/gateway/log"
	"github.com/leiyudaye/gateway/reverseproxy"
)

type Gateway struct {
}

func (g *Gateway) ServerHttp() {
	// 创建http反向代理
	proxy := reverseproxy.NewHttpReverseProxy()

	// 开始监听
	log.Info("http proxy. listen=%v", "127.0.0.1:9910")
	http.ListenAndServe(":9910", proxy)
}

func (g *Gateway) ServerGrpc() {
	listen, err := net.Listen("tcp", ":9920")
	if err != nil {
		return
	}

	// 创建grpc反向代理
	log.Info("grpc proxy. listen=%v", "127.0.0.1:9920")
	reverseproxy.NewGrpcReverseProxy(listen)
}

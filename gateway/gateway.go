/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-05 18:57:31
 * @LastEditors: lly
 * @LastEditTime: 2021-06-20 00:34:12
 */
package gateway

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leiyudaye/gateway/lib/comm"
	"github.com/leiyudaye/gateway/log"
	"github.com/leiyudaye/gateway/reverseproxy"
)

const (
	gatewayTypeHttp = 0 // http类型
	gatewayTypeGrpc = 1 // grpc类型
	gatewayTypeGin  = 2 // gin类型
)

type Gateway struct {
	HttpListenAddr string
	GrpcListenAddr string
	GinListenAddr  string
}

func (g *Gateway) CheckAddr(addr string) bool {
	if addr == "" {
		return false
	}
	return true
}

func (g *Gateway) ServerHttp() {
	if !g.CheckAddr(g.HttpListenAddr) {
		panic("no found listen addr")
	}
	// 创建http反向代理
	proxy := reverseproxy.NewHttpReverseProxy()

	// 开始监听
	log.Info("http proxy. listen=%v", g.HttpListenAddr)
	http.ListenAndServe(g.HttpListenAddr, proxy)
}

func (g *Gateway) ServerGrpc() {
	if !g.CheckAddr(g.GrpcListenAddr) {
		panic("no found listen addr")
	}
	listen, err := net.Listen("tcp", g.GrpcListenAddr)
	if err != nil {
		log.Error("listen failed. err=%v", err)
		return
	}

	// 创建grpc反向代理
	log.Info("grpc proxy. listen=%v", g.GrpcListenAddr)
	reverseproxy.NewGrpcReverseProxy(listen)
}

func (g *Gateway) ServerGinForGrpc() {
	if !g.CheckAddr(g.GinListenAddr) {
		panic("no found listen addr")
	}
	// 创建gin路由
	router := gin.Default()
	router.POST("cgi-bin/gateway.cgi", HanleGateway)
	log.Info("gin proxy. listen=%v", g.GinListenAddr)
	router.Run(g.GinListenAddr)
}

func HanleGateway(c *gin.Context) {
	params := comm.GatewayParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp, err := reverseproxy.NewGinForGrpcReverseProxy(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	resp := make(map[string]interface{}, 0)
	result := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(rsp.Data), &resp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result["data"] = resp
	result["code"] = rsp.Code
	result["ts"] = time.Now().Unix()

	c.JSON(http.StatusOK, result) //gin.H{"status": "200", "data": rsp})
}

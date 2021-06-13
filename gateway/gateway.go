/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-05 18:57:31
 * @LastEditors: lly
 * @LastEditTime: 2021-06-14 03:37:01
 */
package gateway

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	lib "github.com/leiyudaye/gateway/lib/comm"
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

func HanleGateway(c *gin.Context) {
	params := lib.GatewayParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(params.Module.Method)
	fmt.Println(params.Module.Param)
	fmt.Println(params)

	rsp, err := reverseproxy.NewGinForGrpcReverseProxy(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	resp := make(map[string]interface{}, 0)
	result := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(rsp), &resp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result["data"] = resp
	result["status"] = "ok"
	result["ts"] = time.Now().Unix()

	c.JSON(http.StatusOK, result) //gin.H{"status": "200", "data": rsp})
}

func (g *Gateway) ServerGinForGrpc() {
	// 创建gin路由
	router := gin.Default()
	router.POST("cgi-bin/gateway.cgi", HanleGateway)
	router.Run("127.0.0.1:9930")
}

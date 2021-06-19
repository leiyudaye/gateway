/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 21:09:55
 * @LastEditors: lly
 * @LastEditTime: 2021-06-20 00:13:05
 */
package main

import (
	"github.com/leiyudaye/gateway/gateway"
)

func main() {
	g := gateway.Gateway{
		HttpListenAddr: "127.0.0.1:9910",
		GrpcListenAddr: "127.0.0.1:9920",
		GinListenAddr:  "127.0.0.1:9930",
	}
	go g.ServerHttp()
	go g.ServerGrpc()
	go g.ServerGinForGrpc()
	for {
	}
}

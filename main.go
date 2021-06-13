/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 21:09:55
 * @LastEditors: lly
 * @LastEditTime: 2021-06-14 01:45:09
 */
package main

import (
	gateway "github.com/leiyudaye/gateway/gateway"
)

func main() {
	g := gateway.Gateway{}
	go g.ServerHttp()
	go g.ServerGrpc()
	go g.ServerGinForGrpc()
	for {
	}
}

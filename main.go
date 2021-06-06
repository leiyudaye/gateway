/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 21:09:55
 * @LastEditors: lly
 * @LastEditTime: 2021-06-06 12:25:37
 */
package main

import (
	gateway "github.com/leiyudaye/gateway/gateway"
)

func main() {
	g := gateway.Gateway{}
	go g.ServerHttp()
	go g.ServerGrpc()
	for {
	}
}

/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 21:09:55
 * @LastEditors: lly
 * @LastEditTime: 2021-06-05 20:04:57
 */
package main

import (
	gateway "github.com/leiyudaye/gateway/gateway"
)

func main() {
	g := gateway.Gateway{}
	go g.ServerHttp()
	g.ServerGrpc()
}

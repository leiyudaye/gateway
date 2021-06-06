/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 00:09:29
 * @LastEditors: lly
 * @LastEditTime: 2021-06-05 18:19:50
 */
package reverseproxy

import (
	"net"
	"testing"

	load "github.com/leiyudaye/gateway/loadbalance"
)

func TestGrpcReverseProxy(t *testing.T) {
	t.Log("TestNewGrpcReverseProxy begin")

	lis, err := net.Listen("tcp", ":8822")
	if err != nil {
		t.Error("err")
	}

	ld := &load.RandomBalance{}
	ld.Add("127.0.0.1:8811")
	NewGrpcReverseProxy(lis, ld)

	t.Log("TestNewGrpcReverseProxy end")
}
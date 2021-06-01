/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 00:09:29
 * @LastEditors: lly
 * @LastEditTime: 2021-06-01 00:12:01
 */
package gateway

import (
	"net"
	"testing"
)

func TestGrpcReverseProxy(t *testing.T) {
	t.Log("TestNewGrpcReverseProxy begin")

	lis, err := net.Listen("tcp", ":8822")
	if err != nil {
		t.Error("err")
	}
	NewGrpcReverseProxy(lis)
	t.Log("TestNewGrpcReverseProxy end")
}

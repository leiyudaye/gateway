/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 00:09:29
 * @LastEditors: lly
 * @LastEditTime: 2021-06-06 22:31:24
 */
package reverseproxy

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

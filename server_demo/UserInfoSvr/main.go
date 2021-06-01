/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 21:45:56
 * @LastEditors: lly
 * @LastEditTime: 2021-06-01 22:00:07
 */
package main

import (
	"flag"
	"gateway/protobuf_go/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 传入参数
	addr := flag.String("addr", "127.0.0.1:8811", "listen addr")
	flag.Parse()

	// 初始化grpc
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("net listen failed, err=%v", err)
	}
	user.RegisterUserInfoServer(s, new(UserInfoImpl))

	log.Printf("listen %s", *addr)
	s.Serve(lis)
}

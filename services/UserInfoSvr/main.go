/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 21:45:56
 * @LastEditors: lly
 * @LastEditTime: 2021-06-09 23:36:37
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/leiyudaye/gateway/discover"
	"github.com/leiyudaye/gateway/protobuf_go/user"

	"google.golang.org/grpc"
)

func main() {
	// 传入参数
	addr := flag.String("addr", "127.0.0.1", "listen addr")
	port := flag.Int("port", 8822, "listen port")
	flag.Parse()

	lisAddr := fmt.Sprintf("%s:%d", *addr, *port)

	// 初始化grpc
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", lisAddr)
	if err != nil {
		log.Fatalf("net listen failed, err=%v", err)
	}
	user.RegisterUserInfoServer(s, new(UserInfoImpl))

	// 注册服务
	disConn, err := discover.NewDiscoverClient("127.0.0.1", 8500)
	if err != nil {
		return
	}
	if !disConn.Register("user.UserInfo", lisAddr, "", *addr, *port, nil) {
		return
	}

	log.Printf("listen %s", *addr)
	s.Serve(lis)
}

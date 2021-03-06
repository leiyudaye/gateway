/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 21:45:56
 * @LastEditors: lly
 * @LastEditTime: 2021-06-21 00:07:00
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/leiyudaye/gateway/discover"
	"github.com/leiyudaye/gateway/protobuf_go/flowcount"

	"google.golang.org/grpc"
)

func main() {
	// 传入参数
	addr := flag.String("addr", "127.0.0.1", "listen addr")
	port := flag.Int("port", 8833, "listen port")
	flag.Parse()

	lisAddr := fmt.Sprintf("%s:%d", *addr, *port)

	// 初始化grpc
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", lisAddr)
	if err != nil {
		log.Fatalf("net listen failed, err=%v", err)
	}
	flowcount.RegisterFlowCountServer(s, new(FlowCountImpl))

	// 注册服务
	disConn, err := discover.NewDiscoverClient("127.0.0.1:8500")
	if err != nil {
		return
	}
	if !disConn.Register("flowcount.FlowCount", lisAddr, "", *addr, *port, nil) {
		return
	}

	log.Printf("listen %s", *addr)
	s.Serve(lis)
}

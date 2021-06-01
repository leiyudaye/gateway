/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 22:24:44
 * @LastEditors: lly
 * @LastEditTime: 2021-06-01 22:39:12
 */
package main

import (
	"context"
	pb "gateway/protobuf_go/user"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8822", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial failed, err=%v", err)
		return
	}
	defer conn.Close()

	cli := pb.NewUserInfoClient(conn)
	rsp, err := cli.GetUserInfo(context.Background(), &pb.GetUserInfoReq{UserID: 1})
	if err != nil {
		log.Fatalf("call GetUserInfo failed, err=%v", err)
		return
	}
	log.Println(rsp)
}

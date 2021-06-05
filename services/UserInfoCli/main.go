/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 22:24:44
 * @LastEditors: lly
 * @LastEditTime: 2021-06-05 23:49:49
 */
package main

import (
	"context"
	"log"

	pb "github.com/leiyudaye/gateway/protobuf_go/user"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9920", grpc.WithInsecure())
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

/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 21:47:43
 * @LastEditors: lly
 * @LastEditTime: 2021-06-01 21:52:54
 */
package main

import (
	"context"

	pb "github.com/leiyudaye/gateway/protobuf_go/user"
)

type UserInfoImpl struct {
}

func (u UserInfoImpl) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoRsp, error) {
	rsp := new(pb.GetUserInfoRsp)
	rsp.Name = "lly"
	return rsp, nil
}

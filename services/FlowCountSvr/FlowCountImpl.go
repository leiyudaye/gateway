/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-01 21:47:43
 * @LastEditors: lly
 * @LastEditTime: 2021-06-20 23:55:01
 */
package main

import (
	"context"

	"github.com/leiyudaye/gateway/log"
	pb "github.com/leiyudaye/gateway/protobuf_go/flowcount"
)

type FlowCountImpl struct {
}

func (u FlowCountImpl) ReportFlowCount(ctx context.Context, req *pb.ReportFlowCountReq) (*pb.ReportFlowCountRsp, error) {
	rsp := new(pb.ReportFlowCountRsp)
	log.Debug("qps=%v", req.QPS)
	return rsp, nil
}

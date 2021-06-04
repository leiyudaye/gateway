/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-31 23:21:56
 * @LastEditors: lly
 * @LastEditTime: 2021-06-01 22:59:26
 */

package gateway

import (
	"context"
	load "gateway/load_balance"
	proxy "gateway/reverse_proxy/proxy_comm"
	"net"

	"google.golang.org/grpc"
)

func NewGrpcReverseProxy(lis net.Listener, ld load.BlanceInterface) {
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		tagAddr, _ := ld.Get("")
		c, err := grpc.DialContext(ctx, tagAddr, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
		return ctx, c, err
	}

	s := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
	)
	if err := s.Serve(lis); err != nil {
		return
	}
}

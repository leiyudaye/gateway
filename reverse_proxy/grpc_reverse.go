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
	proxy "gateway/reverse_proxy/proxy_comm"
	"net"

	"google.golang.org/grpc"
)

func NewGrpcReverseProxy(lis net.Listener) {
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		c, err := grpc.DialContext(ctx, "localhost:8811", grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
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

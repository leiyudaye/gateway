/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-31 23:21:56
 * @LastEditors: lly
 * @LastEditTime: 2021-06-06 22:34:44
 */

package reverseproxy

import (
	"context"
	"errors"
	"net"
	"strings"

	"github.com/leiyudaye/gateway/discover"
	"github.com/leiyudaye/gateway/log"
	middleware "github.com/leiyudaye/gateway/middleware/grpc_middleware"
	proxy "github.com/leiyudaye/gateway/reverseproxy/proxy_comm"

	"google.golang.org/grpc"
)

func NewGrpcReverseProxy(listen net.Listener) {
	disConn, err := discover.NewDiscoverClient("127.0.0.1", 8500)
	if err != nil {
		log.Error("discover connect failed. err=%v", err)
		return
	}

	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {

		list := strings.Split(fullMethodName, "/")
		if len(list) < 2 {
			log.Error("discover failed")
			return ctx, nil, errors.New("discover failed")
		}
		serverName := list[1]
		tagAddr, err := disConn.Discover(serverName)
		if err != nil {
			log.Error(err.Error())
			return ctx, nil, err
		}
		log.Info("discover success, serverName=%v, addr=%v", serverName, tagAddr)

		c, err := grpc.DialContext(ctx, tagAddr, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
		return ctx, c, err
	}

	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(middleware.GrpcFlowCount()),
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
	)
	if err := s.Serve(listen); err != nil {
		return
	}
}

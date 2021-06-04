/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 21:09:55
 * @LastEditors: lly
 * @LastEditTime: 2021-06-01 22:37:56
 */
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	load "gateway/load_balance"
	reverse "gateway/reverse_proxy"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/hashicorp/consul/api"
)

var logger log.Logger

func init() {
	// 创建日志组件
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "Ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "Caller", log.DefaultCaller)
}

func main() {
	lis, err := net.Listen("tcp", ":8822")
	if err != nil {
		return
	}

	ld := &load.RandomBalance{}
	ld.Add("127.0.0.1:8811")
	reverse.NewGrpcReverseProxy(lis, ld)
}

func main1() {
	// 创建环境变量
	var (
		consulHost = flag.String("consul.host", "127.0.0.1", "consul server ip address")
		consulPort = flag.String("consul.port", "8500", "consul server port")
	)
	flag.Parse()

	// 创建consul api 客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://" + *consulHost + ":" + *consulPort
	conCli, err := api.NewClient(consulConfig)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	// 创建反向代理
	proxy := reverse.NewHttpReverseProxy(conCli, logger)

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// 开始监听
	go func() {
		logger.Log("listen...", ":9090")
		errc <- http.ListenAndServe(":9090", proxy)
	}()

	logger.Log("exit", <-errc)
}

/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-05-28 21:09:55
 * @LastEditors: lly
 * @LastEditTime: 2021-05-31 20:47:00
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	reverse "gateway/reverseproxy"
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

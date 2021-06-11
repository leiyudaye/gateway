<!--
 * @Descripttion: 
 * @Author: lly
 * @Date: 2021-05-28 21:09:55
 * @LastEditors: lly
 * @LastEditTime: 2021-06-12 02:28:05
-->


# 1.运行服务
# 运行注册中心
consul agent -dev
redis-server

# 运行Redis


# 运行网关程序
go run main.go 

# 运行userInfo服务端
 go run services/UserInfoSvr/*.go   --port=9911

# 运行userInfo客户端
go run services/UserInfoCli/main.go 
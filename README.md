<!--
 * @Descripttion: 
 * @Author: lly
 * @Date: 2021-05-28 21:09:55
 * @LastEditors: lly
 * @LastEditTime: 2021-06-20 23:55:57
-->


# 1.运行服务
# 运行注册中心
consul agent -dev

# 运行Redis
redis-server

# 运行网关程序
go run main.go 

# 运行userInfo服务端
 go run services/UserInfoSvr/*.go   --port=9911
# 运行FlowCount服务端
 go run services/FlowCountSvr/*.go   --port=9912

# 运行userInfo客户端
go run services/UserInfoCli/main.go 
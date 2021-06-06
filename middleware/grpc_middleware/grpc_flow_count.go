/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-06 22:24:33
 * @LastEditors: lly
 * @LastEditTime: 2021-06-06 22:31:45
 */
package middleware

import (
	"google.golang.org/grpc"
)

func GrpcFlowCount() func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return nil
	}
}
